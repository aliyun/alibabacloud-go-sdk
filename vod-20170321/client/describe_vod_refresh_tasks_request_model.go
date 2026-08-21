// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRefreshTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodRefreshTasksRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodRefreshTasksRequest
	GetEndTime() *string
	SetObjectPath(v string) *DescribeVodRefreshTasksRequest
	GetObjectPath() *string
	SetObjectType(v string) *DescribeVodRefreshTasksRequest
	GetObjectType() *string
	SetOwnerId(v int64) *DescribeVodRefreshTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVodRefreshTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVodRefreshTasksRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeVodRefreshTasksRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeVodRefreshTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeVodRefreshTasksRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeVodRefreshTasksRequest
	GetTaskId() *string
}

type DescribeVodRefreshTasksRequest struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2017-01-01T12:30:20Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The path used to query. Exact match is used.
	//
	// example:
	//
	// http://example.com/***.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The task type. Valid values:
	//
	// - **file**: file purge.
	//
	// - **directory**: directory purge.
	//
	// - **preload**: file prefetch.
	//
	// > When DomainName or Status is specified, ObjectType is required.
	//
	// example:
	//
	// file
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**. Maximum value: **50**.
	//
	// example:
	//
	// 20
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// > Only data within the last 3 days can be queried.
	//
	// example:
	//
	// 2017-01-01T12:12:20Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - **Complete**: completed.
	//
	// - **Refreshing**: in progress.
	//
	// - **Failed**: failed.
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID used to query the purge status.
	//
	// example:
	//
	// 70422****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeVodRefreshTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodRefreshTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodRefreshTasksRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *DescribeVodRefreshTasksRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeVodRefreshTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodRefreshTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVodRefreshTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVodRefreshTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodRefreshTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodRefreshTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodRefreshTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeVodRefreshTasksRequest) SetDomainName(v string) *DescribeVodRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetEndTime(v string) *DescribeVodRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetObjectPath(v string) *DescribeVodRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetObjectType(v string) *DescribeVodRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetOwnerId(v int64) *DescribeVodRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetPageNumber(v int32) *DescribeVodRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetPageSize(v int32) *DescribeVodRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetSecurityToken(v string) *DescribeVodRefreshTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetStartTime(v string) *DescribeVodRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetStatus(v string) *DescribeVodRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetTaskId(v string) *DescribeVodRefreshTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) Validate() error {
	return dara.Validate(s)
}
