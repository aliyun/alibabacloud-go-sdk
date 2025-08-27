// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDepartmentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeptName(v string) *AddDepartmentShrinkRequest
	GetDeptName() *string
	SetManagerEmployeeIdListShrink(v string) *AddDepartmentShrinkRequest
	GetManagerEmployeeIdListShrink() *string
	SetOutDeptId(v string) *AddDepartmentShrinkRequest
	GetOutDeptId() *string
	SetOutDeptPid(v string) *AddDepartmentShrinkRequest
	GetOutDeptPid() *string
}

type AddDepartmentShrinkRequest struct {
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

func (s AddDepartmentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDepartmentShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddDepartmentShrinkRequest) GetDeptName() *string {
	return s.DeptName
}

func (s *AddDepartmentShrinkRequest) GetManagerEmployeeIdListShrink() *string {
	return s.ManagerEmployeeIdListShrink
}

func (s *AddDepartmentShrinkRequest) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *AddDepartmentShrinkRequest) GetOutDeptPid() *string {
	return s.OutDeptPid
}

func (s *AddDepartmentShrinkRequest) SetDeptName(v string) *AddDepartmentShrinkRequest {
	s.DeptName = &v
	return s
}

func (s *AddDepartmentShrinkRequest) SetManagerEmployeeIdListShrink(v string) *AddDepartmentShrinkRequest {
	s.ManagerEmployeeIdListShrink = &v
	return s
}

func (s *AddDepartmentShrinkRequest) SetOutDeptId(v string) *AddDepartmentShrinkRequest {
	s.OutDeptId = &v
	return s
}

func (s *AddDepartmentShrinkRequest) SetOutDeptPid(v string) *AddDepartmentShrinkRequest {
	s.OutDeptPid = &v
	return s
}

func (s *AddDepartmentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
