// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupDepartSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeptName(v string) *GroupDepartSaveRequest
	GetDeptName() *string
	SetManagerIds(v string) *GroupDepartSaveRequest
	GetManagerIds() *string
	SetOuterDeptId(v string) *GroupDepartSaveRequest
	GetOuterDeptId() *string
	SetOuterDeptPid(v string) *GroupDepartSaveRequest
	GetOuterDeptPid() *string
	SetStatus(v int32) *GroupDepartSaveRequest
	GetStatus() *int32
	SetSubCorpIdList(v []*string) *GroupDepartSaveRequest
	GetSubCorpIdList() []*string
	SetSyncGroup(v bool) *GroupDepartSaveRequest
	GetSyncGroup() *bool
}

type GroupDepartSaveRequest struct {
	// This parameter is required.
	DeptName   *string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	ManagerIds *string `json:"manager_ids,omitempty" xml:"manager_ids,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 001
	OuterDeptId *string `json:"outer_dept_id,omitempty" xml:"outer_dept_id,omitempty"`
	// example:
	//
	// 002
	OuterDeptPid *string `json:"outer_dept_pid,omitempty" xml:"outer_dept_pid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status        *int32    `json:"status,omitempty" xml:"status,omitempty"`
	SubCorpIdList []*string `json:"sub_corp_id_list,omitempty" xml:"sub_corp_id_list,omitempty" type:"Repeated"`
	SyncGroup     *bool     `json:"sync_group,omitempty" xml:"sync_group,omitempty"`
}

func (s GroupDepartSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s GroupDepartSaveRequest) GoString() string {
	return s.String()
}

func (s *GroupDepartSaveRequest) GetDeptName() *string {
	return s.DeptName
}

func (s *GroupDepartSaveRequest) GetManagerIds() *string {
	return s.ManagerIds
}

func (s *GroupDepartSaveRequest) GetOuterDeptId() *string {
	return s.OuterDeptId
}

func (s *GroupDepartSaveRequest) GetOuterDeptPid() *string {
	return s.OuterDeptPid
}

func (s *GroupDepartSaveRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GroupDepartSaveRequest) GetSubCorpIdList() []*string {
	return s.SubCorpIdList
}

func (s *GroupDepartSaveRequest) GetSyncGroup() *bool {
	return s.SyncGroup
}

func (s *GroupDepartSaveRequest) SetDeptName(v string) *GroupDepartSaveRequest {
	s.DeptName = &v
	return s
}

func (s *GroupDepartSaveRequest) SetManagerIds(v string) *GroupDepartSaveRequest {
	s.ManagerIds = &v
	return s
}

func (s *GroupDepartSaveRequest) SetOuterDeptId(v string) *GroupDepartSaveRequest {
	s.OuterDeptId = &v
	return s
}

func (s *GroupDepartSaveRequest) SetOuterDeptPid(v string) *GroupDepartSaveRequest {
	s.OuterDeptPid = &v
	return s
}

func (s *GroupDepartSaveRequest) SetStatus(v int32) *GroupDepartSaveRequest {
	s.Status = &v
	return s
}

func (s *GroupDepartSaveRequest) SetSubCorpIdList(v []*string) *GroupDepartSaveRequest {
	s.SubCorpIdList = v
	return s
}

func (s *GroupDepartSaveRequest) SetSyncGroup(v bool) *GroupDepartSaveRequest {
	s.SyncGroup = &v
	return s
}

func (s *GroupDepartSaveRequest) Validate() error {
	return dara.Validate(s)
}
