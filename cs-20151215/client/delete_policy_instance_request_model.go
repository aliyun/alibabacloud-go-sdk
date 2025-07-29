// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DeletePolicyInstanceRequest
	GetInstanceName() *string
}

type DeletePolicyInstanceRequest struct {
	// The ID of the policy instance.
	//
	// example:
	//
	// allowed-repos-mqdsf
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
}

func (s DeletePolicyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeletePolicyInstanceRequest) SetInstanceName(v string) *DeletePolicyInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DeletePolicyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
