package customrules

import (
	"os"
	"path"
)

func GetDefaultDirectory() string {
	d, _ := os.UserHomeDir()
	return path.Join(d, "whatapp-rules")
}
