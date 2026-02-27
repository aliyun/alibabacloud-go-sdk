// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructSaseDepartment interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentId(v string) *OpenStructSaseDepartment
	GetDepartmentId() *string
	SetFullDepartmentIdPath(v string) *OpenStructSaseDepartment
	GetFullDepartmentIdPath() *string
	SetFullDn(v string) *OpenStructSaseDepartment
	GetFullDn() *string
	SetIdpId(v int64) *OpenStructSaseDepartment
	GetIdpId() *int64
	SetName(v string) *OpenStructSaseDepartment
	GetName() *string
	SetParentDepartmentId(v string) *OpenStructSaseDepartment
	GetParentDepartmentId() *string
}

type OpenStructSaseDepartment struct {
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	FullDepartmentIdPath *string `json:"FullDepartmentIdPath,omitempty" xml:"FullDepartmentIdPath,omitempty"`
	FullDn               *string `json:"FullDn,omitempty" xml:"FullDn,omitempty"`
	IdpId                *int64  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentDepartmentId   *string `json:"ParentDepartmentId,omitempty" xml:"ParentDepartmentId,omitempty"`
}

func (s OpenStructSaseDepartment) String() string {
	return dara.Prettify(s)
}

func (s OpenStructSaseDepartment) GoString() string {
	return s.String()
}

func (s *OpenStructSaseDepartment) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *OpenStructSaseDepartment) GetFullDepartmentIdPath() *string {
	return s.FullDepartmentIdPath
}

func (s *OpenStructSaseDepartment) GetFullDn() *string {
	return s.FullDn
}

func (s *OpenStructSaseDepartment) GetIdpId() *int64 {
	return s.IdpId
}

func (s *OpenStructSaseDepartment) GetName() *string {
	return s.Name
}

func (s *OpenStructSaseDepartment) GetParentDepartmentId() *string {
	return s.ParentDepartmentId
}

func (s *OpenStructSaseDepartment) SetDepartmentId(v string) *OpenStructSaseDepartment {
	s.DepartmentId = &v
	return s
}

func (s *OpenStructSaseDepartment) SetFullDepartmentIdPath(v string) *OpenStructSaseDepartment {
	s.FullDepartmentIdPath = &v
	return s
}

func (s *OpenStructSaseDepartment) SetFullDn(v string) *OpenStructSaseDepartment {
	s.FullDn = &v
	return s
}

func (s *OpenStructSaseDepartment) SetIdpId(v int64) *OpenStructSaseDepartment {
	s.IdpId = &v
	return s
}

func (s *OpenStructSaseDepartment) SetName(v string) *OpenStructSaseDepartment {
	s.Name = &v
	return s
}

func (s *OpenStructSaseDepartment) SetParentDepartmentId(v string) *OpenStructSaseDepartment {
	s.ParentDepartmentId = &v
	return s
}

func (s *OpenStructSaseDepartment) Validate() error {
	return dara.Validate(s)
}
