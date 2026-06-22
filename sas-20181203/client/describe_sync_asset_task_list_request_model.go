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
	// The page number of the page to return. Default value: 1, which indicates that the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end timestamp of the IDC scan task to query. Unit: milliseconds.
	//
	// example:
	//
	// 1662430077000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries per page in a paged query. Default value: 20. If the PageSize parameter is left empty, 20 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start timestamp of the IDC scan task to query. Unit: milliseconds.
	//
	// example:
	//
	// 1652063828796
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task name.
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
