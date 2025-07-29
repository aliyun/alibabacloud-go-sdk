// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*string) *DeletePolicyInstanceResponseBody
	GetInstances() []*string
}

type DeletePolicyInstanceResponseBody struct {
	// A list of policy instances.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s DeletePolicyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyInstanceResponseBody) GetInstances() []*string {
	return s.Instances
}

func (s *DeletePolicyInstanceResponseBody) SetInstances(v []*string) *DeletePolicyInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *DeletePolicyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
