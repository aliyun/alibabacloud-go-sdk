// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateServiceCredentialRequest
	GetName() *string
	SetPassword(v string) *CreateServiceCredentialRequest
	GetPassword() *string
	SetScope(v string) *CreateServiceCredentialRequest
	GetScope() *string
	SetType(v string) *CreateServiceCredentialRequest
	GetType() *string
	SetUsername(v string) *CreateServiceCredentialRequest
	GetUsername() *string
}

type CreateServiceCredentialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 张三的Git证书
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zhangsan
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// example:
	//
	// PERSON
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USERNAME_PASSWORD
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zhangsan
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateServiceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceCredentialRequest) GetName() *string {
	return s.Name
}

func (s *CreateServiceCredentialRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateServiceCredentialRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateServiceCredentialRequest) GetType() *string {
	return s.Type
}

func (s *CreateServiceCredentialRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateServiceCredentialRequest) SetName(v string) *CreateServiceCredentialRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceCredentialRequest) SetPassword(v string) *CreateServiceCredentialRequest {
	s.Password = &v
	return s
}

func (s *CreateServiceCredentialRequest) SetScope(v string) *CreateServiceCredentialRequest {
	s.Scope = &v
	return s
}

func (s *CreateServiceCredentialRequest) SetType(v string) *CreateServiceCredentialRequest {
	s.Type = &v
	return s
}

func (s *CreateServiceCredentialRequest) SetUsername(v string) *CreateServiceCredentialRequest {
	s.Username = &v
	return s
}

func (s *CreateServiceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
