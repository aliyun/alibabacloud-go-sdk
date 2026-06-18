// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAgentStatusDetailsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIdsShrink(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetGroupIdsShrink() *string
	SetInstanceId(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetSkillGroupAgentStatusDetailsShrinkRequest
	GetStartDate() *int64
}

type GetSkillGroupAgentStatusDetailsShrinkRequest struct {
	// Current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
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
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
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

func (s GetSkillGroupAgentStatusDetailsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetCurrentPage(v int32) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetDepIdsShrink(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetEndDate(v int64) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetExistDepartmentGrouping(v bool) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetExistSkillGroupGrouping(v bool) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetGroupIdsShrink(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetInstanceId(v string) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetPageSize(v int32) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) SetStartDate(v int64) *GetSkillGroupAgentStatusDetailsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
