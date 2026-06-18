// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineServiceStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetHotlineServiceStatisticsRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetHotlineServiceStatisticsRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetHotlineServiceStatisticsRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetHotlineServiceStatisticsRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetHotlineServiceStatisticsRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetHotlineServiceStatisticsRequest
	GetExistDepartmentGrouping() *bool
	SetExistSkillGroupGrouping(v bool) *GetHotlineServiceStatisticsRequest
	GetExistSkillGroupGrouping() *bool
	SetGroupIds(v []*int64) *GetHotlineServiceStatisticsRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetHotlineServiceStatisticsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetHotlineServiceStatisticsRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetHotlineServiceStatisticsRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetHotlineServiceStatisticsRequest
	GetTimeLatitudeType() *string
}

type GetHotlineServiceStatisticsRequest struct {
	// List of agent IDs.
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// Current page. The value must be greater than **0**. Default value: **1**.
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
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Specifies whether to query by agent group. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// true
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// Specifies whether to query by department group. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// Indicates whether to query by skill group grouping. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ExistSkillGroupGrouping *bool `json:"ExistSkillGroupGrouping,omitempty" xml:"ExistSkillGroupGrouping,omitempty"`
	// List of skill group IDs.
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
	// Page size. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start date UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Time latitude type for grouped queries. Valid values:
	//
	// - **minute**: minute
	//
	// - **hour**: hour
	//
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetHotlineServiceStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineServiceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetHotlineServiceStatisticsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetHotlineServiceStatisticsRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetHotlineServiceStatisticsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetHotlineServiceStatisticsRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetHotlineServiceStatisticsRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetHotlineServiceStatisticsRequest) GetExistSkillGroupGrouping() *bool {
	return s.ExistSkillGroupGrouping
}

func (s *GetHotlineServiceStatisticsRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetHotlineServiceStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineServiceStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHotlineServiceStatisticsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetHotlineServiceStatisticsRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetHotlineServiceStatisticsRequest) SetAgentIds(v []*int64) *GetHotlineServiceStatisticsRequest {
	s.AgentIds = v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetCurrentPage(v int32) *GetHotlineServiceStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetDepIds(v []*int64) *GetHotlineServiceStatisticsRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetEndDate(v int64) *GetHotlineServiceStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetExistAgentGrouping(v bool) *GetHotlineServiceStatisticsRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetExistDepartmentGrouping(v bool) *GetHotlineServiceStatisticsRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetExistSkillGroupGrouping(v bool) *GetHotlineServiceStatisticsRequest {
	s.ExistSkillGroupGrouping = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetGroupIds(v []*int64) *GetHotlineServiceStatisticsRequest {
	s.GroupIds = v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetInstanceId(v string) *GetHotlineServiceStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetPageSize(v int32) *GetHotlineServiceStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetStartDate(v int64) *GetHotlineServiceStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) SetTimeLatitudeType(v string) *GetHotlineServiceStatisticsRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetHotlineServiceStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
