// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdpDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentId(v string) *UpdateIdpDepartmentRequest
	GetDepartmentId() *string
	SetDepartmentName(v string) *UpdateIdpDepartmentRequest
	GetDepartmentName() *string
	SetIdpConfigId(v string) *UpdateIdpDepartmentRequest
	GetIdpConfigId() *string
}

type UpdateIdpDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10653
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// This parameter is required.
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 598
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
}

func (s UpdateIdpDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdpDepartmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdpDepartmentRequest) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *UpdateIdpDepartmentRequest) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *UpdateIdpDepartmentRequest) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *UpdateIdpDepartmentRequest) SetDepartmentId(v string) *UpdateIdpDepartmentRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateIdpDepartmentRequest) SetDepartmentName(v string) *UpdateIdpDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateIdpDepartmentRequest) SetIdpConfigId(v string) *UpdateIdpDepartmentRequest {
	s.IdpConfigId = &v
	return s
}

func (s *UpdateIdpDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
