// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeptName(v string) *AddDepartmentRequest
	GetDeptName() *string
	SetManagerEmployeeIdList(v []*string) *AddDepartmentRequest
	GetManagerEmployeeIdList() []*string
	SetOutDeptId(v string) *AddDepartmentRequest
	GetOutDeptId() *string
	SetOutDeptPid(v string) *AddDepartmentRequest
	GetOutDeptPid() *string
}

type AddDepartmentRequest struct {
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

func (s AddDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDepartmentRequest) GoString() string {
	return s.String()
}

func (s *AddDepartmentRequest) GetDeptName() *string {
	return s.DeptName
}

func (s *AddDepartmentRequest) GetManagerEmployeeIdList() []*string {
	return s.ManagerEmployeeIdList
}

func (s *AddDepartmentRequest) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *AddDepartmentRequest) GetOutDeptPid() *string {
	return s.OutDeptPid
}

func (s *AddDepartmentRequest) SetDeptName(v string) *AddDepartmentRequest {
	s.DeptName = &v
	return s
}

func (s *AddDepartmentRequest) SetManagerEmployeeIdList(v []*string) *AddDepartmentRequest {
	s.ManagerEmployeeIdList = v
	return s
}

func (s *AddDepartmentRequest) SetOutDeptId(v string) *AddDepartmentRequest {
	s.OutDeptId = &v
	return s
}

func (s *AddDepartmentRequest) SetOutDeptPid(v string) *AddDepartmentRequest {
	s.OutDeptPid = &v
	return s
}

func (s *AddDepartmentRequest) Validate() error {
	return dara.Validate(s)
}
