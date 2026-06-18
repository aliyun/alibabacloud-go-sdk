// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAgentStatusDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupAgentStatusDetailsRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetSkillGroupAgentStatusDetailsRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetSkillGroupAgentStatusDetailsRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupAgentStatusDetailsRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupAgentStatusDetailsRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetSkillGroupAgentStatusDetailsRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetSkillGroupAgentStatusDetailsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupAgentStatusDetailsRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupAgentStatusDetailsRequest
	GetStartDate() *int64
}

type GetSkillGroupAgentStatusDetailsRequest struct {
	// Current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// End Datetime Variable as a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1614824972
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether to query by department grouping. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// Specifies whether to query by skill group grouping. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	// A list of skill group IDs.
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// AICCS instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page size. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start date as a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1614824872
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetSkillGroupAgentStatusDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupAgentStatusDetailsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetCurrentPage(v int32) *GetSkillGroupAgentStatusDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetDepIds(v []*int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.DepIds = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetEndDate(v int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAgentStatusDetailsRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAgentStatusDetailsRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetGroupIds(v []*int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.GroupIds = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetInstanceId(v string) *GetSkillGroupAgentStatusDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetPageSize(v int32) *GetSkillGroupAgentStatusDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) SetStartDate(v int64) *GetSkillGroupAgentStatusDetailsRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsRequest) Validate() error {
	return dara.Validate(s)
}
