// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetQueueInformationRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetQueueInformationRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetQueueInformationRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetQueueInformationRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetQueueInformationRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetQueueInformationRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetQueueInformationRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetQueueInformationRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetQueueInformationRequest
	GetStartDate() *int64
}

type GetQueueInformationRequest struct {
	// The current page. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// End UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Specifies whether to query by department grouping. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// Specifies whether to query data grouped by skill group. Default value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	// The list of skill group IDs.
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetQueueInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueInformationRequest) GoString() string {
	return s.String()
}

func (s *GetQueueInformationRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetQueueInformationRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetQueueInformationRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetQueueInformationRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetQueueInformationRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetQueueInformationRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetQueueInformationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueueInformationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueueInformationRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetQueueInformationRequest) SetCurrentPage(v int32) *GetQueueInformationRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueInformationRequest) SetDepIds(v []*int64) *GetQueueInformationRequest {
	s.DepIds = v
	return s
}

func (s *GetQueueInformationRequest) SetEndDate(v int64) *GetQueueInformationRequest {
	s.EndDate = &v
	return s
}

func (s *GetQueueInformationRequest) SetExistDepartmentGrouping(v bool) *GetQueueInformationRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetQueueInformationRequest) SetExistSkillGroupGrouping(v bool) *GetQueueInformationRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetQueueInformationRequest) SetGroupIds(v []*int64) *GetQueueInformationRequest {
	s.GroupIds = v
	return s
}

func (s *GetQueueInformationRequest) SetInstanceId(v string) *GetQueueInformationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueueInformationRequest) SetPageSize(v int32) *GetQueueInformationRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueueInformationRequest) SetStartDate(v int64) *GetQueueInformationRequest {
	s.StartDate = &v
	return s
}

func (s *GetQueueInformationRequest) Validate() error {
	return dara.Validate(s)
}
