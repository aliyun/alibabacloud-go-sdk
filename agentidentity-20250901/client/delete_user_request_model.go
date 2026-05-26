// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *DeleteUserRequest
	GetUserName() *string
	SetUserPoolName(v string) *DeleteUserRequest
	GetUserPoolName() *string
}

type DeleteUserRequest struct {
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteUserRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *DeleteUserRequest) SetUserName(v string) *DeleteUserRequest {
	s.UserName = &v
	return s
}

func (s *DeleteUserRequest) SetUserPoolName(v string) *DeleteUserRequest {
	s.UserPoolName = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
