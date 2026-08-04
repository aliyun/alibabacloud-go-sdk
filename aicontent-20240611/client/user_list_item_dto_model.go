// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserListItemDTO interface {
	dara.Model
	String() string
	GoString() string
	SetDepartments(v []*UserDepartmentDTO) *UserListItemDTO
	GetDepartments() []*UserDepartmentDTO
	SetGmtCreate(v string) *UserListItemDTO
	GetGmtCreate() *string
	SetId(v int64) *UserListItemDTO
	GetId() *int64
	SetLoginName(v string) *UserListItemDTO
	GetLoginName() *string
	SetName(v string) *UserListItemDTO
	GetName() *string
	SetPhone(v string) *UserListItemDTO
	GetPhone() *string
}

type UserListItemDTO struct {
	// example:
	//
	// []
	Departments []*UserDepartmentDTO `json:"departments,omitempty" xml:"departments,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-07-01 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// zhangsan
	LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
	// example:
	//
	// Zhang San
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 13800000000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s UserListItemDTO) String() string {
	return dara.Prettify(s)
}

func (s UserListItemDTO) GoString() string {
	return s.String()
}

func (s *UserListItemDTO) GetDepartments() []*UserDepartmentDTO {
	return s.Departments
}

func (s *UserListItemDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UserListItemDTO) GetId() *int64 {
	return s.Id
}

func (s *UserListItemDTO) GetLoginName() *string {
	return s.LoginName
}

func (s *UserListItemDTO) GetName() *string {
	return s.Name
}

func (s *UserListItemDTO) GetPhone() *string {
	return s.Phone
}

func (s *UserListItemDTO) SetDepartments(v []*UserDepartmentDTO) *UserListItemDTO {
	s.Departments = v
	return s
}

func (s *UserListItemDTO) SetGmtCreate(v string) *UserListItemDTO {
	s.GmtCreate = &v
	return s
}

func (s *UserListItemDTO) SetId(v int64) *UserListItemDTO {
	s.Id = &v
	return s
}

func (s *UserListItemDTO) SetLoginName(v string) *UserListItemDTO {
	s.LoginName = &v
	return s
}

func (s *UserListItemDTO) SetName(v string) *UserListItemDTO {
	s.Name = &v
	return s
}

func (s *UserListItemDTO) SetPhone(v string) *UserListItemDTO {
	s.Phone = &v
	return s
}

func (s *UserListItemDTO) Validate() error {
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
