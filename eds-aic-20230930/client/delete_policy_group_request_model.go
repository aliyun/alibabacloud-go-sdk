// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyGroupIds(v []*string) *DeletePolicyGroupRequest
	GetPolicyGroupIds() []*string
}

type DeletePolicyGroupRequest struct {
	// The IDs of the policies.
	//
	// This parameter is required.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
}

func (s DeletePolicyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupRequest) GetPolicyGroupIds() []*string {
	return s.PolicyGroupIds
}

func (s *DeletePolicyGroupRequest) SetPolicyGroupIds(v []*string) *DeletePolicyGroupRequest {
	s.PolicyGroupIds = v
	return s
}

func (s *DeletePolicyGroupRequest) Validate() error {
	return dara.Validate(s)
}
