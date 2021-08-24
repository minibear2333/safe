package safe

import "strings"

var (
	filterCommands = []string{"rm -rf /", "halt", "poweroff", "reboot", "shutdown -h now", "shutdown -r now"}
)

func init() {
	tmp := make([]string, 0, len(filterCommands))
	for _, v := range filterCommands {
		tmp = append(tmp, strings.ToLower(strings.ReplaceAll(v, " ", "")))
	}
	filterCommands = tmp
}

// HasHighRiskCmd check has high risk of cmd
func HasHighRiskCmd(cmd string) bool {
	cmd = strings.ToLower(strings.ReplaceAll(cmd, " ", ""))
	for _, v := range filterCommands {
		if strings.Index(cmd, v) > -1 {
			return true
		}
	}
	return false
}
