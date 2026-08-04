// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSetUserRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartmentRoles(v []*DepartmentRoleCmd) *ModelRouterSetUserRolesRequest
	GetDepartmentRoles() []*DepartmentRoleCmd
}

type ModelRouterSetUserRolesRequest struct {
	// The department role list (required, full overwrite).
	//
	// example:
	//
	// [{"clientId":1001,"roleCode":"member"},{"clientId":1002,"roleCode":"member"}]
	DepartmentRoles []*DepartmentRoleCmd `json:"departmentRoles,omitempty" xml:"departmentRoles,omitempty" type:"Repeated"`
}

func (s ModelRouterSetUserRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSetUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterSetUserRolesRequest) GetDepartmentRoles() []*DepartmentRoleCmd {
	return s.DepartmentRoles
}

func (s *ModelRouterSetUserRolesRequest) SetDepartmentRoles(v []*DepartmentRoleCmd) *ModelRouterSetUserRolesRequest {
	s.DepartmentRoles = v
	return s
}

func (s *ModelRouterSetUserRolesRequest) Validate() error {
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
