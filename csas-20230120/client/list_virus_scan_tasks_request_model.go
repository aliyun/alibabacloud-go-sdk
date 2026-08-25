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
	// The page number of the current page in paging. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end point for filtering by task expiration time. The value is a UNIX timestamp in seconds. The value must be greater than StartTime.
	//
	// example:
	//
	// 1762135466
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page in paging. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The collection of scan performance modes. Duplicate values are not allowed.
	PerformanceModes []*string `json:"PerformanceModes,omitempty" xml:"PerformanceModes,omitempty" type:"Repeated"`
	// The collection of scan path scopes. Duplicate values are not allowed.
	ScanModes []*string `json:"ScanModes,omitempty" xml:"ScanModes,omitempty" type:"Repeated"`
	// The start point for filtering by task expiration time. The value is a UNIX timestamp in seconds. This parameter must be specified together with EndTime. Specifying this parameter alone does not take effect.
	//
	// example:
	//
	// 1754150421
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - **0**: Not canceled. This is the default value.
	//
	// - **1**: Canceled.
	//
	// - **-1**: No status filter. All tasks are returned.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of virus scan task IDs. Duplicate values are not allowed.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The user group ID, used to filter tasks whose effective scope includes the specified user group. You can obtain the value from:
	//
	// - [ListUserGroups](~~ListUserGroups~~): Lists user groups.
	//
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
