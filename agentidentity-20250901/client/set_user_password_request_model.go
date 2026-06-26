// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGenerateRandomPassword(v bool) *SetUserPasswordRequest
	GetGenerateRandomPassword() *bool
	SetPassword(v string) *SetUserPasswordRequest
	GetPassword() *string
	SetUserName(v string) *SetUserPasswordRequest
	GetUserName() *string
	SetUserPoolName(v string) *SetUserPasswordRequest
	GetUserPoolName() *string
}

type SetUserPasswordRequest struct {
	GenerateRandomPassword *bool   `json:"GenerateRandomPassword,omitempty" xml:"GenerateRandomPassword,omitempty"`
	Password               *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName               *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPoolName           *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s SetUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *SetUserPasswordRequest) GetGenerateRandomPassword() *bool {
	return s.GenerateRandomPassword
}

func (s *SetUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *SetUserPasswordRequest) GetUserName() *string {
	return s.UserName
}

func (s *SetUserPasswordRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *SetUserPasswordRequest) SetGenerateRandomPassword(v bool) *SetUserPasswordRequest {
	s.GenerateRandomPassword = &v
	return s
}

func (s *SetUserPasswordRequest) SetPassword(v string) *SetUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *SetUserPasswordRequest) SetUserName(v string) *SetUserPasswordRequest {
	s.UserName = &v
	return s
}

func (s *SetUserPasswordRequest) SetUserPoolName(v string) *SetUserPasswordRequest {
	s.UserPoolName = &v
	return s
}

func (s *SetUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
