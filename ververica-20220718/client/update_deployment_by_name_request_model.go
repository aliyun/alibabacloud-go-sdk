// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Deployment) *UpdateDeploymentByNameRequest
	GetBody() *Deployment
	SetDeploymentName(v string) *UpdateDeploymentByNameRequest
	GetDeploymentName() *string
}

type UpdateDeploymentByNameRequest struct {
	// The collection of fields to update. Partial updates are supported.
	//
	// This parameter is required.
	Body *Deployment `json:"body,omitempty" xml:"body,omitempty"`
	// The deployment job name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
}

func (s UpdateDeploymentByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentByNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentByNameRequest) GetBody() *Deployment {
	return s.Body
}

func (s *UpdateDeploymentByNameRequest) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *UpdateDeploymentByNameRequest) SetBody(v *Deployment) *UpdateDeploymentByNameRequest {
	s.Body = v
	return s
}

func (s *UpdateDeploymentByNameRequest) SetDeploymentName(v string) *UpdateDeploymentByNameRequest {
	s.DeploymentName = &v
	return s
}

func (s *UpdateDeploymentByNameRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
