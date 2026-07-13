// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMethod(v string) *CreateUserRequest
	GetAuthMethod() *string
	SetClientToken(v string) *CreateUserRequest
	GetClientToken() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetInstanceId(v string) *CreateUserRequest
	GetInstanceId() *string
	SetName(v string) *CreateUserRequest
	GetName() *string
	SetNote(v string) *CreateUserRequest
	GetNote() *string
	SetPassword(v string) *CreateUserRequest
	GetPassword() *string
}

type CreateUserRequest struct {
	AuthMethod  *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note     *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *CreateUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserRequest) GetName() *string {
	return s.Name
}

func (s *CreateUserRequest) GetNote() *string {
	return s.Note
}

func (s *CreateUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateUserRequest) SetAuthMethod(v string) *CreateUserRequest {
	s.AuthMethod = &v
	return s
}

func (s *CreateUserRequest) SetClientToken(v string) *CreateUserRequest {
	s.ClientToken = &v
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

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetName(v string) *CreateUserRequest {
	s.Name = &v
	return s
}

func (s *CreateUserRequest) SetNote(v string) *CreateUserRequest {
	s.Note = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
