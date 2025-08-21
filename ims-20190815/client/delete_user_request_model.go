// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *DeleteUserRequest
	GetUserId() *string
	SetUserPrincipalName(v string) *DeleteUserRequest
	GetUserPrincipalName() *string
}

type DeleteUserRequest struct {
	// The ID of the RAM user.
	//
	// >  You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// >  You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserRequest) SetUserPrincipalName(v string) *DeleteUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
