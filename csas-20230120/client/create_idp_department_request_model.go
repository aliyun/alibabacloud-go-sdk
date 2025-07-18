// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdpDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentName(v string) *CreateIdpDepartmentRequest
	GetDepartmentName() *string
	SetIdpConfigId(v string) *CreateIdpDepartmentRequest
	GetIdpConfigId() *string
}

type CreateIdpDepartmentRequest struct {
	// This parameter is required.
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1222
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
}

func (s CreateIdpDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdpDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateIdpDepartmentRequest) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *CreateIdpDepartmentRequest) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *CreateIdpDepartmentRequest) SetDepartmentName(v string) *CreateIdpDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *CreateIdpDepartmentRequest) SetIdpConfigId(v string) *CreateIdpDepartmentRequest {
	s.IdpConfigId = &v
	return s
}

func (s *CreateIdpDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
