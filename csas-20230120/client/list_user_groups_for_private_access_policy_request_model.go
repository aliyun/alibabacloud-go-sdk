// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForPrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *ListUserGroupsForPrivateAccessPolicyRequest
	GetPolicyIds() []*string
}

type ListUserGroupsForPrivateAccessPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForPrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListUserGroupsForPrivateAccessPolicyRequest) SetPolicyIds(v []*string) *ListUserGroupsForPrivateAccessPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
