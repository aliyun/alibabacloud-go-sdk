// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeRefreshTasksRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeRefreshTasksRequest
	GetEndTime() *string
	SetObjectPath(v string) *DescribeRefreshTasksRequest
	GetObjectPath() *string
	SetObjectType(v string) *DescribeRefreshTasksRequest
	GetObjectType() *string
	SetOwnerId(v int64) *DescribeRefreshTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRefreshTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRefreshTasksRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeRefreshTasksRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *DescribeRefreshTasksRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeRefreshTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeRefreshTasksRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeRefreshTasksRequest
	GetTaskId() *string
}

type DescribeRefreshTasksRequest struct {
	// The accelerated domain name. You can specify only one accelerated domain name in each call. By default, this operation queries the status of tasks for all accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The path of the object. The path is used as a condition for exact matching.
	//
	// example:
	//
	// http://example.com/1.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **file**: refreshes one or more files.
	//
	// 	- **directory**: refreshes files in specific directories.
	//
	// 	- **regex**: refreshes content based on a regular expression.
	//
	// 	- **preload**: prefetches one or more files.
	//
	// > If you set the **DomainName*	- or **Status*	- parameter, you must also set the **ObjectType*	- parameter.
	//
	// example:
	//
	// file
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**. Maximum value: **100**. Valid values: **1*	- to **100**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The start of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **Complete**: The task is complete.
	//
	// 	- **Refreshing**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task that you want to query.
	//
	// example:
	//
	// 1234321
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRefreshTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRefreshTasksRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeRefreshTasksRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeRefreshTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRefreshTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRefreshTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRefreshTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRefreshTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeRefreshTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRefreshTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeRefreshTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRefreshTasksRequest) SetDomainName(v string) *DescribeRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetEndTime(v string) *DescribeRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetObjectPath(v string) *DescribeRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetObjectType(v string) *DescribeRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetOwnerId(v int64) *DescribeRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetPageNumber(v int32) *DescribeRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetPageSize(v int32) *DescribeRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetResourceGroupId(v string) *DescribeRefreshTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetSecurityToken(v string) *DescribeRefreshTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetStartTime(v string) *DescribeRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetStatus(v string) *DescribeRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetTaskId(v string) *DescribeRefreshTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) Validate() error {
	return dara.Validate(s)
}
