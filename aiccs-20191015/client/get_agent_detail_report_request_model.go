// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDetailReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*int64) *GetAgentDetailReportRequest
	GetAgentIds() []*int64
	SetCurrentPage(v int32) *GetAgentDetailReportRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetAgentDetailReportRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetAgentDetailReportRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetAgentDetailReportRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetAgentDetailReportRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetAgentDetailReportRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentDetailReportRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentDetailReportRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetAgentDetailReportRequest
	GetTimeLatitudeType() *string
}

type GetAgentDetailReportRequest struct {
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
	// true/false
	ExistAgentGrouping *bool `json:"ExistAgentGrouping,omitempty" xml:"ExistAgentGrouping,omitempty"`
	// example:
	//
	// true/false
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
	// day
	TimeLatitudeType *string `json:"TimeLatitudeType,omitempty" xml:"TimeLatitudeType,omitempty"`
}

func (s GetAgentDetailReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportRequest) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *GetAgentDetailReportRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentDetailReportRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetAgentDetailReportRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentDetailReportRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetAgentDetailReportRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetAgentDetailReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentDetailReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentDetailReportRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentDetailReportRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetAgentDetailReportRequest) SetAgentIds(v []*int64) *GetAgentDetailReportRequest {
	s.AgentIds = v
	return s
}

func (s *GetAgentDetailReportRequest) SetCurrentPage(v int32) *GetAgentDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetDepIds(v []*int64) *GetAgentDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetAgentDetailReportRequest) SetEndDate(v int64) *GetAgentDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetExistAgentGrouping(v bool) *GetAgentDetailReportRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetExistDepartmentGrouping(v bool) *GetAgentDetailReportRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetInstanceId(v string) *GetAgentDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetPageSize(v int32) *GetAgentDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetStartDate(v int64) *GetAgentDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentDetailReportRequest) SetTimeLatitudeType(v string) *GetAgentDetailReportRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentDetailReportRequest) Validate() error {
	return dara.Validate(s)
}
