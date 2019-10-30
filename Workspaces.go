package workspaces

import (
	"encoding/json"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetWorkspaceNames() []string {
	cmd := exec.Command(
		"gsettings", "get",
		"org.cinnamon.desktop.wm.preferences",
		"workspace-names")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	out_str := string(out)
	out_str = strings.Replace(out_str, "'", "\"", -1)

	var workspaces []string
	json_err := json.Unmarshal([]byte(out_str), &workspaces)
	if json_err != nil {
		log.Fatalln("error:", json_err)
	}
	return workspaces
}

func GetWorkspaceName(desktop_id int) string {
	workspaces := GetWorkspaceNames()
	if desktop_id > len(workspaces)-1 {
		log.Fatalln("Desktop id out of range")
	}
	return workspaces[desktop_id]
}

func GetCurrentWorkspaceName() string {
	workspaces := GetWorkspaceNames()
	desktop_id := GetCurrentDesktopId()
	if desktop_id > len(workspaces)-1 {
		log.Fatalln("Desktop id out of range")
	}
	return workspaces[desktop_id]
}

func GetCurrentDesktopId() int {
	cmd := exec.Command("xdotool", "get_desktop")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	desktop_id, parse_error := strconv.Atoi(
		strings.TrimSuffix(string(out), "\n"))
	if parse_error != nil {
		log.Fatalf(
			"Could not convert output from xdotool to int: %s %s",
			parse_error, string(out))
	}
	return desktop_id
}

func GoToDesktop(id int) {
	cmd2 := exec.Command("xdotool", "set_desktop", string(id))
	err := cmd2.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
