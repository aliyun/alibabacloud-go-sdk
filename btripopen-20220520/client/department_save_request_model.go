// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDepartmentSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepartList(v []*DepartmentSaveRequestDepartList) *DepartmentSaveRequest
	GetDepartList() []*DepartmentSaveRequestDepartList
}

type DepartmentSaveRequest struct {
	DepartList []*DepartmentSaveRequestDepartList `json:"depart_list,omitempty" xml:"depart_list,omitempty" type:"Repeated"`
}

func (s DepartmentSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s DepartmentSaveRequest) GoString() string {
	return s.String()
}

func (s *DepartmentSaveRequest) GetDepartList() []*DepartmentSaveRequestDepartList {
	return s.DepartList
}

func (s *DepartmentSaveRequest) SetDepartList(v []*DepartmentSaveRequestDepartList) *DepartmentSaveRequest {
	s.DepartList = v
	return s
}

func (s *DepartmentSaveRequest) Validate() error {
	return dara.Validate(s)
}

type DepartmentSaveRequestDepartList struct {
	// example:
	//
	// 10
	DepartId *int64 `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// This parameter is required.
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 10
	DepartPid *int64 `json:"depart_pid,omitempty" xml:"depart_pid,omitempty"`
	// example:
	//
	// 001|002|003
	ManagerIds *string `json:"manager_ids,omitempty" xml:"manager_ids,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// testdepartid001
	ThirdDepartId *string `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	// example:
	//
	// testdepartpid001
	ThirdDepartPid *string `json:"third_depart_pid,omitempty" xml:"third_depart_pid,omitempty"`
}

func (s DepartmentSaveRequestDepartList) String() string {
	return dara.Prettify(s)
}

func (s DepartmentSaveRequestDepartList) GoString() string {
	return s.String()
}

func (s *DepartmentSaveRequestDepartList) GetDepartId() *int64 {
	return s.DepartId
}

func (s *DepartmentSaveRequestDepartList) GetDepartName() *string {
	return s.DepartName
}

func (s *DepartmentSaveRequestDepartList) GetDepartPid() *int64 {
	return s.DepartPid
}

func (s *DepartmentSaveRequestDepartList) GetManagerIds() *string {
	return s.ManagerIds
}

func (s *DepartmentSaveRequestDepartList) GetStatus() *int32 {
	return s.Status
}

func (s *DepartmentSaveRequestDepartList) GetThirdDepartId() *string {
	return s.ThirdDepartId
}

func (s *DepartmentSaveRequestDepartList) GetThirdDepartPid() *string {
	return s.ThirdDepartPid
}

func (s *DepartmentSaveRequestDepartList) SetDepartId(v int64) *DepartmentSaveRequestDepartList {
	s.DepartId = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetDepartName(v string) *DepartmentSaveRequestDepartList {
	s.DepartName = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetDepartPid(v int64) *DepartmentSaveRequestDepartList {
	s.DepartPid = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetManagerIds(v string) *DepartmentSaveRequestDepartList {
	s.ManagerIds = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetStatus(v int32) *DepartmentSaveRequestDepartList {
	s.Status = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetThirdDepartId(v string) *DepartmentSaveRequestDepartList {
	s.ThirdDepartId = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) SetThirdDepartPid(v string) *DepartmentSaveRequestDepartList {
	s.ThirdDepartPid = &v
	return s
}

func (s *DepartmentSaveRequestDepartList) Validate() error {
	return dara.Validate(s)
}
