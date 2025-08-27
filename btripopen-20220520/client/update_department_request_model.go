// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeptName(v string) *UpdateDepartmentRequest
	GetDeptName() *string
	SetManagerEmployeeIdList(v []*string) *UpdateDepartmentRequest
	GetManagerEmployeeIdList() []*string
	SetOutDeptId(v string) *UpdateDepartmentRequest
	GetOutDeptId() *string
	SetOutDeptPid(v string) *UpdateDepartmentRequest
	GetOutDeptPid() *string
}

type UpdateDepartmentRequest struct {
	// This parameter is required.
	DeptName              *string   `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	ManagerEmployeeIdList []*string `json:"manager_employee_id_list,omitempty" xml:"manager_employee_id_list,omitempty" type:"Repeated"`
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

func (s UpdateDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentRequest) GetDeptName() *string {
	return s.DeptName
}

func (s *UpdateDepartmentRequest) GetManagerEmployeeIdList() []*string {
	return s.ManagerEmployeeIdList
}

func (s *UpdateDepartmentRequest) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *UpdateDepartmentRequest) GetOutDeptPid() *string {
	return s.OutDeptPid
}

func (s *UpdateDepartmentRequest) SetDeptName(v string) *UpdateDepartmentRequest {
	s.DeptName = &v
	return s
}

func (s *UpdateDepartmentRequest) SetManagerEmployeeIdList(v []*string) *UpdateDepartmentRequest {
	s.ManagerEmployeeIdList = v
	return s
}

func (s *UpdateDepartmentRequest) SetOutDeptId(v string) *UpdateDepartmentRequest {
	s.OutDeptId = &v
	return s
}

func (s *UpdateDepartmentRequest) SetOutDeptPid(v string) *UpdateDepartmentRequest {
	s.OutDeptPid = &v
	return s
}

func (s *UpdateDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
