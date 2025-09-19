// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAssetTaskLogDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSyncAssetTaskLogDetailRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeSyncAssetTaskLogDetailRequest
	GetEndTime() *int64
	SetPageSize(v int32) *DescribeSyncAssetTaskLogDetailRequest
	GetPageSize() *int32
	SetRootTaskId(v string) *DescribeSyncAssetTaskLogDetailRequest
	GetRootTaskId() *string
	SetStartTime(v int64) *DescribeSyncAssetTaskLogDetailRequest
	GetStartTime() *int64
	SetTaskName(v string) *DescribeSyncAssetTaskLogDetailRequest
	GetTaskName() *string
}

type DescribeSyncAssetTaskLogDetailRequest struct {
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end timestamp of the task.
	//
	// example:
	//
	// 1668064495000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the IDC scan task. You can call the [DescribeSyncAssetTaskList](https://help.aliyun.com/document_detail/141932.html) operation to obtain the ID.
	//
	// example:
	//
	// 7e9565f537146fdf6bfb4e01f6f08818
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	// The start timestamp of the task.
	//
	// example:
	//
	// 1644027670
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the IDC scan task.
	//
	// example:
	//
	// IDC_PROBE_SCAN-TEST_001
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeSyncAssetTaskLogDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskLogDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskLogDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSyncAssetTaskLogDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSyncAssetTaskLogDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSyncAssetTaskLogDetailRequest) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *DescribeSyncAssetTaskLogDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSyncAssetTaskLogDetailRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSyncAssetTaskLogDetailRequest) SetCurrentPage(v int32) *DescribeSyncAssetTaskLogDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailRequest) SetEndTime(v int64) *DescribeSyncAssetTaskLogDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailRequest) SetPageSize(v int32) *DescribeSyncAssetTaskLogDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailRequest) SetRootTaskId(v string) *DescribeSyncAssetTaskLogDetailRequest {
	s.RootTaskId = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailRequest) SetStartTime(v int64) *DescribeSyncAssetTaskLogDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailRequest) SetTaskName(v string) *DescribeSyncAssetTaskLogDetailRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailRequest) Validate() error {
	return dara.Validate(s)
}
