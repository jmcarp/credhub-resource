package out

import (
	"github.com/cloudfoundry-community/credhub-resource/concourse"
	"github.com/cloudfoundry-community/credhub-resource/credhub"
)

type OutResponse struct {
	Version  concourse.Version    `json:"version"`
	Metadata []concourse.Metadata `json:"metadata"`
}

type OutCommand struct {
	client             *credhub.CredhubClient
	resourcesDirectory string
}

func NewOutCommand(client *credhub.CredhubClient, resourcesDirectory string) OutCommand {
	return OutCommand{
		client:             client,
		resourcesDirectory: resourcesDirectory,
	}
}

func (c OutCommand) Run(outRequest concourse.OutRequest) (OutResponse, error) {
	credentials, err := c.client.FindAllCredentialPaths(outRequest.Source.Path)
	if err != nil {
		return OutResponse{}, err
	}

	version := concourse.NewVersion([]byte(credentials), outRequest.Source.Server)

	concourseOutput := OutResponse{
		Version:  version,
		Metadata: append([]concourse.Metadata{}),
	}

	return concourseOutput, nil
}
