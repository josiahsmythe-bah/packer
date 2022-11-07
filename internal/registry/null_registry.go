package registry

import (
	"context"

	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"
)

// nullRegistry is a special handler that does nothing
type nullRegistry struct{}

func newNullHandler() Registry {
	return nullRegistry{}
}

func (h nullRegistry) PopulateIteration(context.Context) error {
	return nil
}

func (h nullRegistry) BuildStart(context.Context, string) error {
	return nil
}

func (h nullRegistry) BuildDone(
	ctx context.Context,
	buildName string,
	artifacts []sdkpacker.Artifact,
	buildErr error,
) ([]sdkpacker.Artifact, error) {
	return artifacts, nil
}
