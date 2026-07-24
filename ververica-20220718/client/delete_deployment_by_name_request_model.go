// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentName(v string) *DeleteDeploymentByNameRequest
	GetDeploymentName() *string
}

type DeleteDeploymentByNameRequest struct {
	// The name of the deployed job, which is typically specified by the user when submitting the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
}

func (s DeleteDeploymentByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentByNameRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentByNameRequest) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *DeleteDeploymentByNameRequest) SetDeploymentName(v string) *DeleteDeploymentByNameRequest {
	s.DeploymentName = &v
	return s
}

func (s *DeleteDeploymentByNameRequest) Validate() error {
	return dara.Validate(s)
}
