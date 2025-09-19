// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAssetTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSyncAssetTaskListRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeSyncAssetTaskListRequest
	GetEndTime() *int64
	SetPageSize(v int32) *DescribeSyncAssetTaskListRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeSyncAssetTaskListRequest
	GetStartTime() *int64
	SetTaskName(v string) *DescribeSyncAssetTaskListRequest
	GetTaskName() *string
}

type DescribeSyncAssetTaskListRequest struct {
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The timestamp when the IDC scan task ends. Unit: milliseconds.
	//
	// example:
	//
	// 1662430077000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The timestamp when the IDC scan task starts. Unit: milliseconds.
	//
	// example:
	//
	// 1652063828796
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the IDC scan task.
	//
	// example:
	//
	// IDC_PROBE_SCAN***
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeSyncAssetTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSyncAssetTaskListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSyncAssetTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSyncAssetTaskListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSyncAssetTaskListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSyncAssetTaskListRequest) SetCurrentPage(v int32) *DescribeSyncAssetTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSyncAssetTaskListRequest) SetEndTime(v int64) *DescribeSyncAssetTaskListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSyncAssetTaskListRequest) SetPageSize(v int32) *DescribeSyncAssetTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSyncAssetTaskListRequest) SetStartTime(v int64) *DescribeSyncAssetTaskListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSyncAssetTaskListRequest) SetTaskName(v string) *DescribeSyncAssetTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeSyncAssetTaskListRequest) Validate() error {
	return dara.Validate(s)
}
