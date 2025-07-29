// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPolicyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*string) *DeployPolicyInstanceResponseBody
	GetInstances() []*string
}

type DeployPolicyInstanceResponseBody struct {
	// A list of policy instances.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s DeployPolicyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployPolicyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployPolicyInstanceResponseBody) GetInstances() []*string {
	return s.Instances
}

func (s *DeployPolicyInstanceResponseBody) SetInstances(v []*string) *DeployPolicyInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *DeployPolicyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
