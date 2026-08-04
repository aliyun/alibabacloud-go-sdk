// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentRoles(v []*DepartmentRoleCmd) *ModelRouterCreateUserRequest
	GetDepartmentRoles() []*DepartmentRoleCmd
	SetLoginName(v string) *ModelRouterCreateUserRequest
	GetLoginName() *string
	SetName(v string) *ModelRouterCreateUserRequest
	GetName() *string
	SetPhone(v string) *ModelRouterCreateUserRequest
	GetPhone() *string
}

type ModelRouterCreateUserRequest struct {
	// The department roles to assign to the user during creation. This parameter is optional.
	//
	// example:
	//
	// []
	DepartmentRoles []*DepartmentRoleCmd `json:"departmentRoles,omitempty" xml:"departmentRoles,omitempty" type:"Repeated"`
	// The logon name. This parameter is required. The logon name can be the same as the phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// zhangsan
	LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
	// The name. This parameter is required. The value must be 2 to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// John Smith
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The phone number. This parameter is required.
	//
	// example:
	//
	// 13800000000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s ModelRouterCreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateUserRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateUserRequest) GetDepartmentRoles() []*DepartmentRoleCmd {
	return s.DepartmentRoles
}

func (s *ModelRouterCreateUserRequest) GetLoginName() *string {
	return s.LoginName
}

func (s *ModelRouterCreateUserRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterCreateUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *ModelRouterCreateUserRequest) SetDepartmentRoles(v []*DepartmentRoleCmd) *ModelRouterCreateUserRequest {
	s.DepartmentRoles = v
	return s
}

func (s *ModelRouterCreateUserRequest) SetLoginName(v string) *ModelRouterCreateUserRequest {
	s.LoginName = &v
	return s
}

func (s *ModelRouterCreateUserRequest) SetName(v string) *ModelRouterCreateUserRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterCreateUserRequest) SetPhone(v string) *ModelRouterCreateUserRequest {
	s.Phone = &v
	return s
}

func (s *ModelRouterCreateUserRequest) Validate() error {
	if s.DepartmentRoles != nil {
		for _, item := range s.DepartmentRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
