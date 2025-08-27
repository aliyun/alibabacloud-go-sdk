// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupDepartSaveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeptName(v string) *GroupDepartSaveShrinkRequest
	GetDeptName() *string
	SetManagerIds(v string) *GroupDepartSaveShrinkRequest
	GetManagerIds() *string
	SetOuterDeptId(v string) *GroupDepartSaveShrinkRequest
	GetOuterDeptId() *string
	SetOuterDeptPid(v string) *GroupDepartSaveShrinkRequest
	GetOuterDeptPid() *string
	SetStatus(v int32) *GroupDepartSaveShrinkRequest
	GetStatus() *int32
	SetSubCorpIdListShrink(v string) *GroupDepartSaveShrinkRequest
	GetSubCorpIdListShrink() *string
	SetSyncGroup(v bool) *GroupDepartSaveShrinkRequest
	GetSyncGroup() *bool
}

type GroupDepartSaveShrinkRequest struct {
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
	Status              *int32  `json:"status,omitempty" xml:"status,omitempty"`
	SubCorpIdListShrink *string `json:"sub_corp_id_list,omitempty" xml:"sub_corp_id_list,omitempty"`
	SyncGroup           *bool   `json:"sync_group,omitempty" xml:"sync_group,omitempty"`
}

func (s GroupDepartSaveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GroupDepartSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *GroupDepartSaveShrinkRequest) GetDeptName() *string {
	return s.DeptName
}

func (s *GroupDepartSaveShrinkRequest) GetManagerIds() *string {
	return s.ManagerIds
}

func (s *GroupDepartSaveShrinkRequest) GetOuterDeptId() *string {
	return s.OuterDeptId
}

func (s *GroupDepartSaveShrinkRequest) GetOuterDeptPid() *string {
	return s.OuterDeptPid
}

func (s *GroupDepartSaveShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GroupDepartSaveShrinkRequest) GetSubCorpIdListShrink() *string {
	return s.SubCorpIdListShrink
}

func (s *GroupDepartSaveShrinkRequest) GetSyncGroup() *bool {
	return s.SyncGroup
}

func (s *GroupDepartSaveShrinkRequest) SetDeptName(v string) *GroupDepartSaveShrinkRequest {
	s.DeptName = &v
	return s
}

func (s *GroupDepartSaveShrinkRequest) SetManagerIds(v string) *GroupDepartSaveShrinkRequest {
	s.ManagerIds = &v
	return s
}

func (s *GroupDepartSaveShrinkRequest) SetOuterDeptId(v string) *GroupDepartSaveShrinkRequest {
	s.OuterDeptId = &v
	return s
}

func (s *GroupDepartSaveShrinkRequest) SetOuterDeptPid(v string) *GroupDepartSaveShrinkRequest {
	s.OuterDeptPid = &v
	return s
}

func (s *GroupDepartSaveShrinkRequest) SetStatus(v int32) *GroupDepartSaveShrinkRequest {
	s.Status = &v
	return s
}

func (s *GroupDepartSaveShrinkRequest) SetSubCorpIdListShrink(v string) *GroupDepartSaveShrinkRequest {
	s.SubCorpIdListShrink = &v
	return s
}

func (s *GroupDepartSaveShrinkRequest) SetSyncGroup(v bool) *GroupDepartSaveShrinkRequest {
	s.SyncGroup = &v
	return s
}

func (s *GroupDepartSaveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
