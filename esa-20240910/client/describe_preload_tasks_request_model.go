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
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
