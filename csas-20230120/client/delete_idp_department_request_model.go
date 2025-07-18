// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdpDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentId(v string) *DeleteIdpDepartmentRequest
	GetDepartmentId() *string
	SetIdpConfigId(v string) *DeleteIdpDepartmentRequest
	GetIdpConfigId() *string
}

type DeleteIdpDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10829
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 507
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
}

func (s DeleteIdpDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdpDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIdpDepartmentRequest) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *DeleteIdpDepartmentRequest) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *DeleteIdpDepartmentRequest) SetDepartmentId(v string) *DeleteIdpDepartmentRequest {
	s.DepartmentId = &v
	return s
}

func (s *DeleteIdpDepartmentRequest) SetIdpConfigId(v string) *DeleteIdpDepartmentRequest {
	s.IdpConfigId = &v
	return s
}

func (s *DeleteIdpDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
