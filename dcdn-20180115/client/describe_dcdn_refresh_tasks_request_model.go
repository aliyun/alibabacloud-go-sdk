// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnRefreshTasksRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnRefreshTasksRequest
	GetEndTime() *string
	SetObjectPath(v string) *DescribeDcdnRefreshTasksRequest
	GetObjectPath() *string
	SetObjectType(v string) *DescribeDcdnRefreshTasksRequest
	GetObjectType() *string
	SetOwnerId(v int64) *DescribeDcdnRefreshTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDcdnRefreshTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnRefreshTasksRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeDcdnRefreshTasksRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeDcdnRefreshTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeDcdnRefreshTasksRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeDcdnRefreshTasksRequest
	GetTaskId() *string
}

type DescribeDcdnRefreshTasksRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2017-01-01T12:13:20Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The path of the object. The path is used as a condition for exact matching.
	//
	// example:
	//
	// http://example.com/examplefile.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the task.
	//
	// 	- **file**: URL-based refresh
	//
	// 	- **directory**: directory-based refresh
	//
	// 	- **preload**: URL-based prefetch
	//
	// If you set **DomainName*	- or **Status**, you must also set this parameter.
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
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**. Maximum value: **50**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 20
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-01T12:12:20Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task.
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
	// The ID of the task. A task ID is assigned when you create a refresh or prefetch task.
	//
	// example:
	//
	// 704225667
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnRefreshTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnRefreshTasksRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeDcdnRefreshTasksRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeDcdnRefreshTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnRefreshTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnRefreshTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnRefreshTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnRefreshTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnRefreshTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnRefreshTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDcdnRefreshTasksRequest) SetDomainName(v string) *DescribeDcdnRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetEndTime(v string) *DescribeDcdnRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetObjectPath(v string) *DescribeDcdnRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetObjectType(v string) *DescribeDcdnRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetOwnerId(v int64) *DescribeDcdnRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetPageNumber(v int32) *DescribeDcdnRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetPageSize(v int32) *DescribeDcdnRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetSecurityToken(v string) *DescribeDcdnRefreshTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetStartTime(v string) *DescribeDcdnRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetStatus(v string) *DescribeDcdnRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) SetTaskId(v string) *DescribeDcdnRefreshTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDcdnRefreshTasksRequest) Validate() error {
	return dara.Validate(s)
}
