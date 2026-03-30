// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *ListGroupsForUserRequest
	GetUserPrincipalName() *string
}

type ListGroupsForUserRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListGroupsForUserRequest) SetUserPrincipalName(v string) *ListGroupsForUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *ListGroupsForUserRequest) Validate() error {
	return dara.Validate(s)
}
