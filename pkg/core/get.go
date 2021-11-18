package core

import (
	"github.com/restechnica/semverbot/pkg/version"
)

type GetVersionOptions struct {
	DefaultVersion string
}

// GetVersion gets the current version based on the latest annotated git tag.
// Returns the current version.
func GetVersion(options *GetVersionOptions) string {
	var versionAPI = version.NewAPI()
	return versionAPI.GetVersionOrDefault(options.DefaultVersion)
}
