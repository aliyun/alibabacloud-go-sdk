// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *GetUserRequest
	GetUserName() *string
	SetUserPoolName(v string) *GetUserRequest
	GetUserPoolName() *string
}

type GetUserRequest struct {
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetUserRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserRequest) SetUserName(v string) *GetUserRequest {
	s.UserName = &v
	return s
}

func (s *GetUserRequest) SetUserPoolName(v string) *GetUserRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
