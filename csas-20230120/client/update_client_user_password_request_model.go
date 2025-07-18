// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateClientUserPasswordRequest
	GetId() *string
	SetPassword(v string) *UpdateClientUserPasswordRequest
	GetPassword() *string
	SetUsername(v string) *UpdateClientUserPasswordRequest
	GetUsername() *string
}

type UpdateClientUserPasswordRequest struct {
	// example:
	//
	// 1128
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// kehudiyidj
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateClientUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientUserPasswordRequest) GetId() *string {
	return s.Id
}

func (s *UpdateClientUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateClientUserPasswordRequest) GetUsername() *string {
	return s.Username
}

func (s *UpdateClientUserPasswordRequest) SetId(v string) *UpdateClientUserPasswordRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientUserPasswordRequest) SetPassword(v string) *UpdateClientUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateClientUserPasswordRequest) SetUsername(v string) *UpdateClientUserPasswordRequest {
	s.Username = &v
	return s
}

func (s *UpdateClientUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
