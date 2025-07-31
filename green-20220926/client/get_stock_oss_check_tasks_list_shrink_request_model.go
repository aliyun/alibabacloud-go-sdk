// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStockOssCheckTasksListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetStockOssCheckTasksListShrinkRequest
	GetCurrentPage() *int32
	SetEndTime(v string) *GetStockOssCheckTasksListShrinkRequest
	GetEndTime() *string
	SetIsInc(v bool) *GetStockOssCheckTasksListShrinkRequest
	GetIsInc() *bool
	SetMediaType(v int32) *GetStockOssCheckTasksListShrinkRequest
	GetMediaType() *int32
	SetPageSize(v int32) *GetStockOssCheckTasksListShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetStockOssCheckTasksListShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *GetStockOssCheckTasksListShrinkRequest
	GetSortShrink() *string
	SetStartTime(v string) *GetStockOssCheckTasksListShrinkRequest
	GetStartTime() *string
	SetTaskType(v string) *GetStockOssCheckTasksListShrinkRequest
	GetTaskType() *string
}

type GetStockOssCheckTasksListShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-06-18 02:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// image
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-06-17 02:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetStockOssCheckTasksListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStockOssCheckTasksListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetIsInc() *bool {
	return s.IsInc
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetMediaType() *int32 {
	return s.MediaType
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStockOssCheckTasksListShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetCurrentPage(v int32) *GetStockOssCheckTasksListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetEndTime(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetIsInc(v bool) *GetStockOssCheckTasksListShrinkRequest {
	s.IsInc = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetMediaType(v int32) *GetStockOssCheckTasksListShrinkRequest {
	s.MediaType = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetPageSize(v int32) *GetStockOssCheckTasksListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetRegionId(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetSortShrink(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetStartTime(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetTaskType(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
