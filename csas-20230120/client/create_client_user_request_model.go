// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentId(v string) *CreateClientUserRequest
	GetDepartmentId() *string
	SetDescription(v string) *CreateClientUserRequest
	GetDescription() *string
	SetEmail(v string) *CreateClientUserRequest
	GetEmail() *string
	SetIdpConfigId(v string) *CreateClientUserRequest
	GetIdpConfigId() *string
	SetMobileNumber(v string) *CreateClientUserRequest
	GetMobileNumber() *string
	SetPassword(v string) *CreateClientUserRequest
	GetPassword() *string
	SetUsername(v string) *CreateClientUserRequest
	GetUsername() *string
}

type CreateClientUserRequest struct {
	// example:
	//
	// 10797
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 727
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 13641966835
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// kehudiyi
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientUserRequest) GoString() string {
	return s.String()
}

func (s *CreateClientUserRequest) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *CreateClientUserRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateClientUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateClientUserRequest) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *CreateClientUserRequest) GetMobileNumber() *string {
	return s.MobileNumber
}

func (s *CreateClientUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateClientUserRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateClientUserRequest) SetDepartmentId(v string) *CreateClientUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateClientUserRequest) SetDescription(v string) *CreateClientUserRequest {
	s.Description = &v
	return s
}

func (s *CreateClientUserRequest) SetEmail(v string) *CreateClientUserRequest {
	s.Email = &v
	return s
}

func (s *CreateClientUserRequest) SetIdpConfigId(v string) *CreateClientUserRequest {
	s.IdpConfigId = &v
	return s
}

func (s *CreateClientUserRequest) SetMobileNumber(v string) *CreateClientUserRequest {
	s.MobileNumber = &v
	return s
}

func (s *CreateClientUserRequest) SetPassword(v string) *CreateClientUserRequest {
	s.Password = &v
	return s
}

func (s *CreateClientUserRequest) SetUsername(v string) *CreateClientUserRequest {
	s.Username = &v
	return s
}

func (s *CreateClientUserRequest) Validate() error {
	return dara.Validate(s)
}
