// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistrationPoliciesForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupIds(v []*string) *ListRegistrationPoliciesForUserGroupRequest
	GetUserGroupIds() []*string
}

type ListRegistrationPoliciesForUserGroupRequest struct {
	// This parameter is required.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListRegistrationPoliciesForUserGroupRequest) SetUserGroupIds(v []*string) *ListRegistrationPoliciesForUserGroupRequest {
	s.UserGroupIds = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
