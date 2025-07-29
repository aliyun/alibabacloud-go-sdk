// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*string) *ModifyPolicyInstanceResponseBody
	GetInstances() []*string
}

type ModifyPolicyInstanceResponseBody struct {
	// The list of policy instances that are updated.
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s ModifyPolicyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyInstanceResponseBody) GetInstances() []*string {
	return s.Instances
}

func (s *ModifyPolicyInstanceResponseBody) SetInstances(v []*string) *ModifyPolicyInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *ModifyPolicyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
