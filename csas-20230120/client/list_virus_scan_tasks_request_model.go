// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListVirusScanTasksRequest
	GetCurrentPage() *int64
	SetEndTime(v int64) *ListVirusScanTasksRequest
	GetEndTime() *int64
	SetPageSize(v int64) *ListVirusScanTasksRequest
	GetPageSize() *int64
	SetPerformanceModes(v []*string) *ListVirusScanTasksRequest
	GetPerformanceModes() []*string
	SetScanModes(v []*string) *ListVirusScanTasksRequest
	GetScanModes() []*string
	SetStartTime(v int64) *ListVirusScanTasksRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListVirusScanTasksRequest
	GetStatus() *int32
	SetTaskIds(v []*string) *ListVirusScanTasksRequest
	GetTaskIds() []*string
	SetUserGroupId(v string) *ListVirusScanTasksRequest
	GetUserGroupId() *string
}

type ListVirusScanTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1762135466
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize         *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PerformanceModes []*string `json:"PerformanceModes,omitempty" xml:"PerformanceModes,omitempty" type:"Repeated"`
	ScanModes        []*string `json:"ScanModes,omitempty" xml:"ScanModes,omitempty" type:"Repeated"`
	// example:
	//
	// 1754150421
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	Status  *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// usergroup-9d4f2a7b3c1e****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListVirusScanTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTasksRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanTasksRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListVirusScanTasksRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListVirusScanTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVirusScanTasksRequest) GetPerformanceModes() []*string {
	return s.PerformanceModes
}

func (s *ListVirusScanTasksRequest) GetScanModes() []*string {
	return s.ScanModes
}

func (s *ListVirusScanTasksRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListVirusScanTasksRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListVirusScanTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListVirusScanTasksRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListVirusScanTasksRequest) SetCurrentPage(v int64) *ListVirusScanTasksRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanTasksRequest) SetEndTime(v int64) *ListVirusScanTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListVirusScanTasksRequest) SetPageSize(v int64) *ListVirusScanTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanTasksRequest) SetPerformanceModes(v []*string) *ListVirusScanTasksRequest {
	s.PerformanceModes = v
	return s
}

func (s *ListVirusScanTasksRequest) SetScanModes(v []*string) *ListVirusScanTasksRequest {
	s.ScanModes = v
	return s
}

func (s *ListVirusScanTasksRequest) SetStartTime(v int64) *ListVirusScanTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListVirusScanTasksRequest) SetStatus(v int32) *ListVirusScanTasksRequest {
	s.Status = &v
	return s
}

func (s *ListVirusScanTasksRequest) SetTaskIds(v []*string) *ListVirusScanTasksRequest {
	s.TaskIds = v
	return s
}

func (s *ListVirusScanTasksRequest) SetUserGroupId(v string) *ListVirusScanTasksRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListVirusScanTasksRequest) Validate() error {
	return dara.Validate(s)
}
