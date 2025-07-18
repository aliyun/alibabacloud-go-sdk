// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentId(v string) *UpdateClientUserRequest
	GetDepartmentId() *string
	SetDescription(v string) *UpdateClientUserRequest
	GetDescription() *string
	SetEmail(v string) *UpdateClientUserRequest
	GetEmail() *string
	SetId(v string) *UpdateClientUserRequest
	GetId() *string
	SetMobileNumber(v string) *UpdateClientUserRequest
	GetMobileNumber() *string
}

type UpdateClientUserRequest struct {
	// example:
	//
	// 10701
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20644
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 13641966835
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
}

func (s UpdateClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientUserRequest) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *UpdateClientUserRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateClientUserRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateClientUserRequest) GetId() *string {
	return s.Id
}

func (s *UpdateClientUserRequest) GetMobileNumber() *string {
	return s.MobileNumber
}

func (s *UpdateClientUserRequest) SetDepartmentId(v string) *UpdateClientUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateClientUserRequest) SetDescription(v string) *UpdateClientUserRequest {
	s.Description = &v
	return s
}

func (s *UpdateClientUserRequest) SetEmail(v string) *UpdateClientUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateClientUserRequest) SetId(v string) *UpdateClientUserRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientUserRequest) SetMobileNumber(v string) *UpdateClientUserRequest {
	s.MobileNumber = &v
	return s
}

func (s *UpdateClientUserRequest) Validate() error {
	return dara.Validate(s)
}
