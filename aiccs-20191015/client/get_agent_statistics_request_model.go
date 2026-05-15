// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetAgentStatisticsRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetAgentStatisticsRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetAgentStatisticsRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetAgentStatisticsRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetAgentStatisticsRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetAgentStatisticsRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetAgentStatisticsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentStatisticsRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentStatisticsRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetAgentStatisticsRequest
	GetTimeLatitudeType() *string
}

type GetAgentStatisticsRequest struct {
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIds      []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// false
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// example:
	//
	// false
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// minute
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetAgentStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetAgentStatisticsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentStatisticsRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetAgentStatisticsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentStatisticsRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetAgentStatisticsRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetAgentStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentStatisticsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentStatisticsRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetAgentStatisticsRequest) SetAgentIds(v []*int64) *GetAgentStatisticsRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentStatisticsRequest) SetCurrentPage(v int32) *GetAgentStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetDepIds(v []*int64) *GetAgentStatisticsRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentStatisticsRequest) SetEndDate(v int64) *GetAgentStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetExistAgentGrouping(v bool) *GetAgentStatisticsRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetExistDepartmentGrouping(v bool) *GetAgentStatisticsRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetInstanceId(v string) *GetAgentStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetPageSize(v int32) *GetAgentStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetStartDate(v int64) *GetAgentStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentStatisticsRequest) SetTimeLatitudeType(v string) *GetAgentStatisticsRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
