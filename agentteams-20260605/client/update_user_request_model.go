// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMethod(v string) *UpdateUserRequest
	GetAuthMethod() *string
	SetClientToken(v string) *UpdateUserRequest
	GetClientToken() *string
	SetDisplayName(v string) *UpdateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *UpdateUserRequest
	GetEmail() *string
	SetInstanceId(v string) *UpdateUserRequest
	GetInstanceId() *string
	SetName(v string) *UpdateUserRequest
	GetName() *string
	SetNote(v string) *UpdateUserRequest
	GetNote() *string
}

type UpdateUserRequest struct {
	AuthMethod  *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *UpdateUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserRequest) GetName() *string {
	return s.Name
}

func (s *UpdateUserRequest) GetNote() *string {
	return s.Note
}

func (s *UpdateUserRequest) SetAuthMethod(v string) *UpdateUserRequest {
	s.AuthMethod = &v
	return s
}

func (s *UpdateUserRequest) SetClientToken(v string) *UpdateUserRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateUserRequest) SetDisplayName(v string) *UpdateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetInstanceId(v string) *UpdateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserRequest) SetName(v string) *UpdateUserRequest {
	s.Name = &v
	return s
}

func (s *UpdateUserRequest) SetNote(v string) *UpdateUserRequest {
	s.Note = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
