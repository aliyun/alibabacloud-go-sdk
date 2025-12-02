// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStockOssCheckTasksListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetStockOssCheckTasksListRequest
	GetCurrentPage() *int32
	SetEndTime(v string) *GetStockOssCheckTasksListRequest
	GetEndTime() *string
	SetIsInc(v bool) *GetStockOssCheckTasksListRequest
	GetIsInc() *bool
	SetMediaType(v int32) *GetStockOssCheckTasksListRequest
	GetMediaType() *int32
	SetPageSize(v int32) *GetStockOssCheckTasksListRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetStockOssCheckTasksListRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *GetStockOssCheckTasksListRequest
	GetSort() map[string]*string
	SetStartTime(v string) *GetStockOssCheckTasksListRequest
	GetStartTime() *string
	SetTaskType(v string) *GetStockOssCheckTasksListRequest
	GetTaskType() *string
}

type GetStockOssCheckTasksListRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End time.
	//
	// example:
	//
	// 2023-06-18 02:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Whether it is a scheduled scan task.
	//
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// Media type.
	//
	// example:
	//
	// image
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Sort field.
	Sort map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Start time.
	//
	// example:
	//
	// 2023-06-17 02:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Task type.
	//
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetStockOssCheckTasksListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListRequest) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetStockOssCheckTasksListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStockOssCheckTasksListRequest) GetIsInc() *bool {
	return s.IsInc
}

func (s *GetStockOssCheckTasksListRequest) GetMediaType() *int32 {
	return s.MediaType
}

func (s *GetStockOssCheckTasksListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStockOssCheckTasksListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStockOssCheckTasksListRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *GetStockOssCheckTasksListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStockOssCheckTasksListRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetStockOssCheckTasksListRequest) SetCurrentPage(v int32) *GetStockOssCheckTasksListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetEndTime(v string) *GetStockOssCheckTasksListRequest {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetIsInc(v bool) *GetStockOssCheckTasksListRequest {
	s.IsInc = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetMediaType(v int32) *GetStockOssCheckTasksListRequest {
	s.MediaType = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetPageSize(v int32) *GetStockOssCheckTasksListRequest {
	s.PageSize = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetRegionId(v string) *GetStockOssCheckTasksListRequest {
	s.RegionId = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetSort(v map[string]*string) *GetStockOssCheckTasksListRequest {
	s.Sort = v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetStartTime(v string) *GetStockOssCheckTasksListRequest {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetTaskType(v string) *GetStockOssCheckTasksListRequest {
	s.TaskType = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) Validate() error {
	return dara.Validate(s)
}
