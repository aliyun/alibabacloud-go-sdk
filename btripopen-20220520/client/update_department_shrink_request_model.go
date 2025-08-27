// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDepartmentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeptName(v string) *UpdateDepartmentShrinkRequest
	GetDeptName() *string
	SetManagerEmployeeIdListShrink(v string) *UpdateDepartmentShrinkRequest
	GetManagerEmployeeIdListShrink() *string
	SetOutDeptId(v string) *UpdateDepartmentShrinkRequest
	GetOutDeptId() *string
	SetOutDeptPid(v string) *UpdateDepartmentShrinkRequest
	GetOutDeptPid() *string
}

type UpdateDepartmentShrinkRequest struct {
	// This parameter is required.
	DeptName                    *string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	ManagerEmployeeIdListShrink *string `json:"manager_employee_id_list,omitempty" xml:"manager_employee_id_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dept123
	OutDeptId *string `json:"out_dept_id,omitempty" xml:"out_dept_id,omitempty"`
	// example:
	//
	// dept456
	OutDeptPid *string `json:"out_dept_pid,omitempty" xml:"out_dept_pid,omitempty"`
}

func (s UpdateDepartmentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDepartmentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentShrinkRequest) GetDeptName() *string {
	return s.DeptName
}

func (s *UpdateDepartmentShrinkRequest) GetManagerEmployeeIdListShrink() *string {
	return s.ManagerEmployeeIdListShrink
}

func (s *UpdateDepartmentShrinkRequest) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *UpdateDepartmentShrinkRequest) GetOutDeptPid() *string {
	return s.OutDeptPid
}

func (s *UpdateDepartmentShrinkRequest) SetDeptName(v string) *UpdateDepartmentShrinkRequest {
	s.DeptName = &v
	return s
}

func (s *UpdateDepartmentShrinkRequest) SetManagerEmployeeIdListShrink(v string) *UpdateDepartmentShrinkRequest {
	s.ManagerEmployeeIdListShrink = &v
	return s
}

func (s *UpdateDepartmentShrinkRequest) SetOutDeptId(v string) *UpdateDepartmentShrinkRequest {
	s.OutDeptId = &v
	return s
}

func (s *UpdateDepartmentShrinkRequest) SetOutDeptPid(v string) *UpdateDepartmentShrinkRequest {
	s.OutDeptPid = &v
	return s
}

func (s *UpdateDepartmentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
