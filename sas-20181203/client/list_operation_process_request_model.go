// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOperationProcessRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListOperationProcessRequest
	GetEndTime() *int64
	SetPageSize(v int32) *ListOperationProcessRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListOperationProcessRequest
	GetStartTime() *int64
	SetStatusCodes(v []*int32) *ListOperationProcessRequest
	GetStatusCodes() []*int32
	SetTaskIds(v []*string) *ListOperationProcessRequest
	GetTaskIds() []*string
	SetTaskSources(v []*string) *ListOperationProcessRequest
	GetTaskSources() []*string
	SetTaskTypes(v []*string) *ListOperationProcessRequest
	GetTaskTypes() []*string
}

type ListOperationProcessRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1635575219000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1680919232000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status codes.
	StatusCodes []*int32 `json:"StatusCodes,omitempty" xml:"StatusCodes,omitempty" type:"Repeated"`
	// The task IDs.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// List of task sources.
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
	// The task types. Valid values:
	//
	// 	- CHECK_ALL: full check.
	//
	// 	- CHECK_POLICY: policy-based check for which check items are configured.
	//
	// 	- CHECK_SCHEDULE: scheduled check.
	//
	// 	- CHECK_ITEM: specific check item-based check.
	//
	// 	- CHECK_INSTANCE: specific check item-based check on specific instances.
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s ListOperationProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessRequest) GoString() string {
	return s.String()
}

func (s *ListOperationProcessRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOperationProcessRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationProcessRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationProcessRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationProcessRequest) GetStatusCodes() []*int32 {
	return s.StatusCodes
}

func (s *ListOperationProcessRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListOperationProcessRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *ListOperationProcessRequest) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *ListOperationProcessRequest) SetCurrentPage(v int32) *ListOperationProcessRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationProcessRequest) SetEndTime(v int64) *ListOperationProcessRequest {
	s.EndTime = &v
	return s
}

func (s *ListOperationProcessRequest) SetPageSize(v int32) *ListOperationProcessRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationProcessRequest) SetStartTime(v int64) *ListOperationProcessRequest {
	s.StartTime = &v
	return s
}

func (s *ListOperationProcessRequest) SetStatusCodes(v []*int32) *ListOperationProcessRequest {
	s.StatusCodes = v
	return s
}

func (s *ListOperationProcessRequest) SetTaskIds(v []*string) *ListOperationProcessRequest {
	s.TaskIds = v
	return s
}

func (s *ListOperationProcessRequest) SetTaskSources(v []*string) *ListOperationProcessRequest {
	s.TaskSources = v
	return s
}

func (s *ListOperationProcessRequest) SetTaskTypes(v []*string) *ListOperationProcessRequest {
	s.TaskTypes = v
	return s
}

func (s *ListOperationProcessRequest) Validate() error {
	return dara.Validate(s)
}
