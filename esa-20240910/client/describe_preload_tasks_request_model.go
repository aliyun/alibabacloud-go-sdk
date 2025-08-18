// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreloadTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribePreloadTasksRequest
	GetContent() *string
	SetEndTime(v string) *DescribePreloadTasksRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribePreloadTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePreloadTasksRequest
	GetPageSize() *int32
	SetSiteId(v int64) *DescribePreloadTasksRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribePreloadTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribePreloadTasksRequest
	GetStatus() *string
}

type DescribePreloadTasksRequest struct {
	// The content to prefetch. Exact match is supported.
	//
	// example:
	//
	// http://a.com/1.jpg?b=2
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The end time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2023-03-23T06:23:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-03-22T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// 	- **Complete**: The task is complete.
	//
	// 	- **Refreshing**: The task is running.
	//
	// 	- **Failed**: The task failed.
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePreloadTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribePreloadTasksRequest) GetContent() *string {
	return s.Content
}

func (s *DescribePreloadTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePreloadTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePreloadTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePreloadTasksRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribePreloadTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePreloadTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribePreloadTasksRequest) SetContent(v string) *DescribePreloadTasksRequest {
	s.Content = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetEndTime(v string) *DescribePreloadTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetPageNumber(v int32) *DescribePreloadTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetPageSize(v int32) *DescribePreloadTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetSiteId(v int64) *DescribePreloadTasksRequest {
	s.SiteId = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetStartTime(v string) *DescribePreloadTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePreloadTasksRequest) SetStatus(v string) *DescribePreloadTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribePreloadTasksRequest) Validate() error {
	return dara.Validate(s)
}
