// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStatisticsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetAgentStatisticsShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetAgentStatisticsShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetAgentStatisticsShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetAgentStatisticsShrinkRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetAgentStatisticsShrinkRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetAgentStatisticsShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetAgentStatisticsShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentStatisticsShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentStatisticsShrinkRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetAgentStatisticsShrinkRequest
	GetTimeLatitudeType() *string
}

type GetAgentStatisticsShrinkRequest struct {
	// List of agent IDs.
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
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
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Whether to query by agent group. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// Indicates whether to query by department grouping. Valid values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Time latitude type. Valid values:
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

func (s GetAgentStatisticsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetAgentStatisticsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentStatisticsShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetAgentStatisticsShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentStatisticsShrinkRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetAgentStatisticsShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetAgentStatisticsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentStatisticsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentStatisticsShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentStatisticsShrinkRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetAgentStatisticsShrinkRequest) SetAgentIdsShrink(v string) *GetAgentStatisticsShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetCurrentPage(v int32) *GetAgentStatisticsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetDepIdsShrink(v string) *GetAgentStatisticsShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetEndDate(v int64) *GetAgentStatisticsShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetExistAgentGrouping(v bool) *GetAgentStatisticsShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetExistDepartmentGrouping(v bool) *GetAgentStatisticsShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetInstanceId(v string) *GetAgentStatisticsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetPageSize(v int32) *GetAgentStatisticsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetStartDate(v int64) *GetAgentStatisticsShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) SetTimeLatitudeType(v string) *GetAgentStatisticsShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentStatisticsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
