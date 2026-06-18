// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetAgentServiceStatusRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetAgentServiceStatusRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetAgentServiceStatusRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetAgentServiceStatusRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetAgentServiceStatusRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetAgentServiceStatusRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetAgentServiceStatusRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentServiceStatusRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentServiceStatusRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetAgentServiceStatusRequest
	GetTimeLatitudeType() *string
}

type GetAgentServiceStatusRequest struct {
	// A list of agent IDs.
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// The current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of department IDs.
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// End UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Whether to query by agent group. Default Value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// Whether to query by department group. Default Value: **false**. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
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
	// The start UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The time latitude type. Valid values:
	//
	// - **minute**: Minute
	//
	// - **hour**: Hour
	//
	// - **day**: Day
	//
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetAgentServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetAgentServiceStatusRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentServiceStatusRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetAgentServiceStatusRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentServiceStatusRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetAgentServiceStatusRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetAgentServiceStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentServiceStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentServiceStatusRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentServiceStatusRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetAgentServiceStatusRequest) SetAgentIds(v []*int64) *GetAgentServiceStatusRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentServiceStatusRequest) SetCurrentPage(v int32) *GetAgentServiceStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetDepIds(v []*int64) *GetAgentServiceStatusRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentServiceStatusRequest) SetEndDate(v int64) *GetAgentServiceStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetExistAgentGrouping(v bool) *GetAgentServiceStatusRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetExistDepartmentGrouping(v bool) *GetAgentServiceStatusRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetInstanceId(v string) *GetAgentServiceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetPageSize(v int32) *GetAgentServiceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetStartDate(v int64) *GetAgentServiceStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentServiceStatusRequest) SetTimeLatitudeType(v string) *GetAgentServiceStatusRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
