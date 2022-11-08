package registry

import (
	"context"

	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"
)

// nullRegistry is a special handler that does nothing
type nullRegistry struct{}

// We don't need a constructor for such a simple type
//func newNullHandler() Registry {
//return nullRegistry{}
//}

func (r nullRegistry) PopulateIteration(context.Context) error {
	return nil
}

func (r nullRegistry) BuildStart(context.Context, string) error {
	return nil
}

func (r nullRegistry) BuildDone(
	ctx context.Context,
	buildName string,
	artifacts []sdkpacker.Artifact,
	buildErr error,
) ([]sdkpacker.Artifact, error) {
	return artifacts, nil
}
