//go:build !windows

package tray

import (
	"errors"

	"github.com/uppercaveman/ollama-server/app/tray/commontray"
)

func InitPlatformTray(icon, updateIcon []byte) (commontray.OllamaTray, error) {
	return nil, errors.New("not implemented")
}
