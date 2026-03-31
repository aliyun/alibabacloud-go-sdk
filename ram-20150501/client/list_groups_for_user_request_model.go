// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *ListGroupsForUserRequest
	GetUserName() *string
}

type ListGroupsForUserRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// Alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListGroupsForUserRequest) SetUserName(v string) *ListGroupsForUserRequest {
	s.UserName = &v
	return s
}

func (s *ListGroupsForUserRequest) Validate() error {
	return dara.Validate(s)
}
