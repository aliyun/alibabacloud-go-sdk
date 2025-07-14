// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryUserGroupMemberRequest
	GetKeyword() *string
	SetUserGroupId(v string) *QueryUserGroupMemberRequest
	GetUserGroupId() *string
}

type QueryUserGroupMemberRequest struct {
	// Keyword for the username or nickname of the user group member.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// User group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryUserGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryUserGroupMemberRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *QueryUserGroupMemberRequest) SetKeyword(v string) *QueryUserGroupMemberRequest {
	s.Keyword = &v
	return s
}

func (s *QueryUserGroupMemberRequest) SetUserGroupId(v string) *QueryUserGroupMemberRequest {
	s.UserGroupId = &v
	return s
}

func (s *QueryUserGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
