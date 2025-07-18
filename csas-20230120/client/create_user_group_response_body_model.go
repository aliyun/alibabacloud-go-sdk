// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserGroupResponseBody
	GetRequestId() *string
	SetUserGroupId(v string) *CreateUserGroupResponseBody
	GetUserGroupId() *string
}

type CreateUserGroupResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s CreateUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserGroupResponseBody) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetUserGroupId(v string) *CreateUserGroupResponseBody {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
