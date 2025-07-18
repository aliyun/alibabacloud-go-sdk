// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupsForRegistrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *ListUserGroupsForRegistrationPolicyRequest
	GetPolicyIds() []*string
}

type ListUserGroupsForRegistrationPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForRegistrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListUserGroupsForRegistrationPolicyRequest) SetPolicyIds(v []*string) *ListUserGroupsForRegistrationPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
