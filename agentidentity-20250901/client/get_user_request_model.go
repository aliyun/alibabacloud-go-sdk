// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetUserRequest
	GetUserId() *string
	SetUserName(v string) *GetUserRequest
	GetUserName() *string
	SetUserPoolId(v string) *GetUserRequest
	GetUserPoolId() *string
	SetUserPoolName(v string) *GetUserRequest
	GetUserPoolName() *string
}

type GetUserRequest struct {
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPoolId   *string `json:"UserPoolId,omitempty" xml:"UserPoolId,omitempty"`
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetUserRequest) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *GetUserRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

func (s *GetUserRequest) SetUserName(v string) *GetUserRequest {
	s.UserName = &v
	return s
}

func (s *GetUserRequest) SetUserPoolId(v string) *GetUserRequest {
	s.UserPoolId = &v
	return s
}

func (s *GetUserRequest) SetUserPoolName(v string) *GetUserRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
