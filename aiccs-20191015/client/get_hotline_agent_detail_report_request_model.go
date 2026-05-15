// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentDetailReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetHotlineAgentDetailReportRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetHotlineAgentDetailReportRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetHotlineAgentDetailReportRequest
	GetEndDate() *int64
	SetGroupIds(v []*int64) *GetHotlineAgentDetailReportRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetHotlineAgentDetailReportRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetHotlineAgentDetailReportRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetHotlineAgentDetailReportRequest
	GetStartDate() *int64
}

type GetHotlineAgentDetailReportRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIds      []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1614824972
	EndDate  *int64   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614824872
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHotlineAgentDetailReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetHotlineAgentDetailReportRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetHotlineAgentDetailReportRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetHotlineAgentDetailReportRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetHotlineAgentDetailReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineAgentDetailReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHotlineAgentDetailReportRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetHotlineAgentDetailReportRequest) SetCurrentPage(v int32) *GetHotlineAgentDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetDepIds(v []*int64) *GetHotlineAgentDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetEndDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetGroupIds(v []*int64) *GetHotlineAgentDetailReportRequest {
	s.GroupIds = v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetInstanceId(v string) *GetHotlineAgentDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetPageSize(v int32) *GetHotlineAgentDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetStartDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) Validate() error {
	return dara.Validate(s)
}
