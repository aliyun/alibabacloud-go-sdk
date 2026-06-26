// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUserRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetStatus(v string) *CreateUserRequest
	GetStatus() *string
	SetUserName(v string) *CreateUserRequest
	GetUserName() *string
	SetUserPoolName(v string) *CreateUserRequest
	GetUserPoolName() *string
}

type CreateUserRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetStatus(v string) *CreateUserRequest {
	s.Status = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) SetUserPoolName(v string) *CreateUserRequest {
	s.UserPoolName = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
