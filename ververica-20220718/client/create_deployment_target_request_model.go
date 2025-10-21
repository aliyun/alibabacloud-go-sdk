// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ResourceSpec) *CreateDeploymentTargetRequest
	GetBody() *ResourceSpec
	SetDeploymentTargetName(v string) *CreateDeploymentTargetRequest
	GetDeploymentTargetName() *string
}

type CreateDeploymentTargetRequest struct {
	Body *ResourceSpec `json:"body,omitempty" xml:"body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-dt
	DeploymentTargetName *string `json:"deploymentTargetName,omitempty" xml:"deploymentTargetName,omitempty"`
}

func (s CreateDeploymentTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetRequest) GetBody() *ResourceSpec {
	return s.Body
}

func (s *CreateDeploymentTargetRequest) GetDeploymentTargetName() *string {
	return s.DeploymentTargetName
}

func (s *CreateDeploymentTargetRequest) SetBody(v *ResourceSpec) *CreateDeploymentTargetRequest {
	s.Body = v
	return s
}

func (s *CreateDeploymentTargetRequest) SetDeploymentTargetName(v string) *CreateDeploymentTargetRequest {
	s.DeploymentTargetName = &v
	return s
}

func (s *CreateDeploymentTargetRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
