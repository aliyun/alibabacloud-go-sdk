// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructSaseUserSimple interface {
	dara.Model
	String() string
	GoString() string
	SetDepartments(v []*OpenStructSaseUserSimpleDepartments) *OpenStructSaseUserSimple
	GetDepartments() []*OpenStructSaseUserSimpleDepartments
	SetEmail(v string) *OpenStructSaseUserSimple
	GetEmail() *string
	SetIdpConfigId(v string) *OpenStructSaseUserSimple
	GetIdpConfigId() *string
	SetSaseUserId(v string) *OpenStructSaseUserSimple
	GetSaseUserId() *string
	SetTelephone(v string) *OpenStructSaseUserSimple
	GetTelephone() *string
	SetUsername(v string) *OpenStructSaseUserSimple
	GetUsername() *string
}

type OpenStructSaseUserSimple struct {
	Departments []*OpenStructSaseUserSimpleDepartments `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	Email       *string                                `json:"Email,omitempty" xml:"Email,omitempty"`
	IdpConfigId *string                                `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	SaseUserId  *string                                `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Telephone   *string                                `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	Username    *string                                `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s OpenStructSaseUserSimple) String() string {
	return dara.Prettify(s)
}

func (s OpenStructSaseUserSimple) GoString() string {
	return s.String()
}

func (s *OpenStructSaseUserSimple) GetDepartments() []*OpenStructSaseUserSimpleDepartments {
	return s.Departments
}

func (s *OpenStructSaseUserSimple) GetEmail() *string {
	return s.Email
}

func (s *OpenStructSaseUserSimple) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *OpenStructSaseUserSimple) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *OpenStructSaseUserSimple) GetTelephone() *string {
	return s.Telephone
}

func (s *OpenStructSaseUserSimple) GetUsername() *string {
	return s.Username
}

func (s *OpenStructSaseUserSimple) SetDepartments(v []*OpenStructSaseUserSimpleDepartments) *OpenStructSaseUserSimple {
	s.Departments = v
	return s
}

func (s *OpenStructSaseUserSimple) SetEmail(v string) *OpenStructSaseUserSimple {
	s.Email = &v
	return s
}

func (s *OpenStructSaseUserSimple) SetIdpConfigId(v string) *OpenStructSaseUserSimple {
	s.IdpConfigId = &v
	return s
}

func (s *OpenStructSaseUserSimple) SetSaseUserId(v string) *OpenStructSaseUserSimple {
	s.SaseUserId = &v
	return s
}

func (s *OpenStructSaseUserSimple) SetTelephone(v string) *OpenStructSaseUserSimple {
	s.Telephone = &v
	return s
}

func (s *OpenStructSaseUserSimple) SetUsername(v string) *OpenStructSaseUserSimple {
	s.Username = &v
	return s
}

func (s *OpenStructSaseUserSimple) Validate() error {
	if s.Departments != nil {
		for _, item := range s.Departments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OpenStructSaseUserSimpleDepartments struct {
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	FullDepartmentIdPath *string `json:"FullDepartmentIdPath,omitempty" xml:"FullDepartmentIdPath,omitempty"`
	FullDn               *string `json:"FullDn,omitempty" xml:"FullDn,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentDepartmentId   *string `json:"ParentDepartmentId,omitempty" xml:"ParentDepartmentId,omitempty"`
}

func (s OpenStructSaseUserSimpleDepartments) String() string {
	return dara.Prettify(s)
}

func (s OpenStructSaseUserSimpleDepartments) GoString() string {
	return s.String()
}

func (s *OpenStructSaseUserSimpleDepartments) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *OpenStructSaseUserSimpleDepartments) GetFullDepartmentIdPath() *string {
	return s.FullDepartmentIdPath
}

func (s *OpenStructSaseUserSimpleDepartments) GetFullDn() *string {
	return s.FullDn
}

func (s *OpenStructSaseUserSimpleDepartments) GetName() *string {
	return s.Name
}

func (s *OpenStructSaseUserSimpleDepartments) GetParentDepartmentId() *string {
	return s.ParentDepartmentId
}

func (s *OpenStructSaseUserSimpleDepartments) SetDepartmentId(v string) *OpenStructSaseUserSimpleDepartments {
	s.DepartmentId = &v
	return s
}

func (s *OpenStructSaseUserSimpleDepartments) SetFullDepartmentIdPath(v string) *OpenStructSaseUserSimpleDepartments {
	s.FullDepartmentIdPath = &v
	return s
}

func (s *OpenStructSaseUserSimpleDepartments) SetFullDn(v string) *OpenStructSaseUserSimpleDepartments {
	s.FullDn = &v
	return s
}

func (s *OpenStructSaseUserSimpleDepartments) SetName(v string) *OpenStructSaseUserSimpleDepartments {
	s.Name = &v
	return s
}

func (s *OpenStructSaseUserSimpleDepartments) SetParentDepartmentId(v string) *OpenStructSaseUserSimpleDepartments {
	s.ParentDepartmentId = &v
	return s
}

func (s *OpenStructSaseUserSimpleDepartments) Validate() error {
	return dara.Validate(s)
}
