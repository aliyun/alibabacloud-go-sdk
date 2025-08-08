// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeployEnvironmentOptions) *DeployEnvironmentRequest
	GetBody() *DeployEnvironmentOptions
}

type DeployEnvironmentRequest struct {
	Body *DeployEnvironmentOptions `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentRequest) GetBody() *DeployEnvironmentOptions {
	return s.Body
}

func (s *DeployEnvironmentRequest) SetBody(v *DeployEnvironmentOptions) *DeployEnvironmentRequest {
	s.Body = v
	return s
}

func (s *DeployEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
