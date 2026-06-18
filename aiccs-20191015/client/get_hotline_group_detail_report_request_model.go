// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineGroupDetailReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetHotlineGroupDetailReportRequest
	GetCurrentPage() *int32
	SetDepIds(v []*int64) *GetHotlineGroupDetailReportRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetHotlineGroupDetailReportRequest
	GetEndDate() *int64
	SetGroupIds(v []*int64) *GetHotlineGroupDetailReportRequest
	GetGroupIds() []*int64
	SetInstanceId(v string) *GetHotlineGroupDetailReportRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetHotlineGroupDetailReportRequest
	GetPageSize() *int32
	SetStartDate(v int64) *GetHotlineGroupDetailReportRequest
	GetStartDate() *int64
}

type GetHotlineGroupDetailReportRequest struct {
	// Current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// End date as a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614824972
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// List of skill group IDs.
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID.
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
	// Start date as a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614824872
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHotlineGroupDetailReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineGroupDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetHotlineGroupDetailReportRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetHotlineGroupDetailReportRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetHotlineGroupDetailReportRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *GetHotlineGroupDetailReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineGroupDetailReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHotlineGroupDetailReportRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetHotlineGroupDetailReportRequest) SetCurrentPage(v int32) *GetHotlineGroupDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetDepIds(v []*int64) *GetHotlineGroupDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetEndDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetGroupIds(v []*int64) *GetHotlineGroupDetailReportRequest {
	s.GroupIds = v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetInstanceId(v string) *GetHotlineGroupDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetPageSize(v int32) *GetHotlineGroupDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetStartDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) Validate() error {
	return dara.Validate(s)
}
