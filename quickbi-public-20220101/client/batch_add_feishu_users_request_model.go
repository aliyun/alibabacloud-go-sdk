// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddFeishuUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeishuUsers(v string) *BatchAddFeishuUsersRequest
	GetFeishuUsers() *string
	SetIsAdmin(v bool) *BatchAddFeishuUsersRequest
	GetIsAdmin() *bool
	SetIsAuthAdmin(v bool) *BatchAddFeishuUsersRequest
	GetIsAuthAdmin() *bool
	SetUserGroupIds(v string) *BatchAddFeishuUsersRequest
	GetUserGroupIds() *string
	SetUserType(v int32) *BatchAddFeishuUsersRequest
	GetUserType() *int32
}

type BatchAddFeishuUsersRequest struct {
	// Information of the users to be added
	//
	// example:
	//
	// {"ad****fd": "TEST", "82****5a": "TEST"}"
	FeishuUsers *string `json:"FeishuUsers,omitempty" xml:"FeishuUsers,omitempty"`
	// Whether the user is an admin user:
	//
	// - true
	//
	// - false
	//
	// Default is false if not provided
	//
	// example:
	//
	// False
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// Whether the user is an authorization administrator
	//
	// - true
	//
	// - false
	//
	// Default is false if not provided
	//
	// example:
	//
	// true
	IsAuthAdmin *bool `json:"IsAuthAdmin,omitempty" xml:"IsAuthAdmin,omitempty"`
	// User group ID(s)
	//
	// example:
	//
	// "0d5fb19b-5555-41f0-99da-1248fc27ca51,0f868dd6_68dd_4d13_8422_c5dca3bd4b61"
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// User type
	//
	// - Developer: 1
	//
	// - Visitor: 2
	//
	// - Analyst: 3
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s BatchAddFeishuUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFeishuUsersRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersRequest) GetFeishuUsers() *string {
	return s.FeishuUsers
}

func (s *BatchAddFeishuUsersRequest) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *BatchAddFeishuUsersRequest) GetIsAuthAdmin() *bool {
	return s.IsAuthAdmin
}

func (s *BatchAddFeishuUsersRequest) GetUserGroupIds() *string {
	return s.UserGroupIds
}

func (s *BatchAddFeishuUsersRequest) GetUserType() *int32 {
	return s.UserType
}

func (s *BatchAddFeishuUsersRequest) SetFeishuUsers(v string) *BatchAddFeishuUsersRequest {
	s.FeishuUsers = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetIsAdmin(v bool) *BatchAddFeishuUsersRequest {
	s.IsAdmin = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetIsAuthAdmin(v bool) *BatchAddFeishuUsersRequest {
	s.IsAuthAdmin = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetUserGroupIds(v string) *BatchAddFeishuUsersRequest {
	s.UserGroupIds = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetUserType(v int32) *BatchAddFeishuUsersRequest {
	s.UserType = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) Validate() error {
	return dara.Validate(s)
}
