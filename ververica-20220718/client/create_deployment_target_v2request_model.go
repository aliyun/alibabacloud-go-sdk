// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Resource) *CreateDeploymentTargetV2Request
	GetBody() *Resource
	SetDeploymentTargetName(v string) *CreateDeploymentTargetV2Request
	GetDeploymentTargetName() *string
}

type CreateDeploymentTargetV2Request struct {
	Body *Resource `json:"body,omitempty" xml:"body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-dt
	DeploymentTargetName *string `json:"deploymentTargetName,omitempty" xml:"deploymentTargetName,omitempty"`
}

func (s CreateDeploymentTargetV2Request) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetV2Request) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetV2Request) GetBody() *Resource {
	return s.Body
}

func (s *CreateDeploymentTargetV2Request) GetDeploymentTargetName() *string {
	return s.DeploymentTargetName
}

func (s *CreateDeploymentTargetV2Request) SetBody(v *Resource) *CreateDeploymentTargetV2Request {
	s.Body = v
	return s
}

func (s *CreateDeploymentTargetV2Request) SetDeploymentTargetName(v string) *CreateDeploymentTargetV2Request {
	s.DeploymentTargetName = &v
	return s
}

func (s *CreateDeploymentTargetV2Request) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
