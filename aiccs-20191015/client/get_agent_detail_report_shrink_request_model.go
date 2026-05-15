// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDetailReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdsShrink(v string) *GetAgentDetailReportShrinkRequest
	GetAgentIdsShrink() *string
	SetCurrentPage(v int32) *GetAgentDetailReportShrinkRequest
	GetCurrentPage() *int32
	SetDepIdsShrink(v string) *GetAgentDetailReportShrinkRequest
	GetDepIdsShrink() *string
	SetEndDate(v int64) *GetAgentDetailReportShrinkRequest
	GetEndDate() *int64
	SetExistAgentGrouping(v bool) *GetAgentDetailReportShrinkRequest
	GetExistAgentGrouping() *bool
	SetExistDepartmentGrouping(v bool) *GetAgentDetailReportShrinkRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetAgentDetailReportShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAgentDetailReportShrinkRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetAgentDetailReportShrinkRequest
	GetStartDate() *int64
	SetTimeLatitudeType(v string) *GetAgentDetailReportShrinkRequest
	GetTimeLatitudeType() *string
}

type GetAgentDetailReportShrinkRequest struct {
	AgentIdsShrink *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// example:
	//
	// 1
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIdsShrink *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
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

func (s GetAgentDetailReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDetailReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportShrinkRequest) GetAgentIdsShrink() *string {
	return s.AgentIdsShrink
}

func (s *GetAgentDetailReportShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAgentDetailReportShrinkRequest) GetDepIdsShrink() *string {
	return s.DepIdsShrink
}

func (s *GetAgentDetailReportShrinkRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAgentDetailReportShrinkRequest) GetExistAgentGrouping() *bool {
	return s.ExistAgentGrouping
}

func (s *GetAgentDetailReportShrinkRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetAgentDetailReportShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentDetailReportShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentDetailReportShrinkRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAgentDetailReportShrinkRequest) GetTimeLatitudeType() *string {
	return s.TimeLatitudeType
}

func (s *GetAgentDetailReportShrinkRequest) SetAgentIdsShrink(v string) *GetAgentDetailReportShrinkRequest {
	s.AgentIdsShrink = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetCurrentPage(v int32) *GetAgentDetailReportShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetDepIdsShrink(v string) *GetAgentDetailReportShrinkRequest {
	s.DepIdsShrink = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetEndDate(v int64) *GetAgentDetailReportShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetExistAgentGrouping(v bool) *GetAgentDetailReportShrinkRequest {
	s.ExistAgentGrouping = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetExistDepartmentGrouping(v bool) *GetAgentDetailReportShrinkRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetInstanceId(v string) *GetAgentDetailReportShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetPageSize(v int32) *GetAgentDetailReportShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetStartDate(v int64) *GetAgentDetailReportShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) SetTimeLatitudeType(v string) *GetAgentDetailReportShrinkRequest {
	s.TimeLatitudeType = &v
	return s
}

func (s *GetAgentDetailReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
