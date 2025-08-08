// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployEnvironmentOptions interface {
	dara.Model
	String() string
	GoString() string
	SetServices(v []*string) *DeployEnvironmentOptions
	GetServices() []*string
}

type DeployEnvironmentOptions struct {
	Services []*string `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s DeployEnvironmentOptions) String() string {
	return dara.Prettify(s)
}

func (s DeployEnvironmentOptions) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentOptions) GetServices() []*string {
	return s.Services
}

func (s *DeployEnvironmentOptions) SetServices(v []*string) *DeployEnvironmentOptions {
	s.Services = v
	return s
}

func (s *DeployEnvironmentOptions) Validate() error {
	return dara.Validate(s)
}
