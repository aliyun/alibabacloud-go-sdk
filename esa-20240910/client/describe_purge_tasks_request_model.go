// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurgeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribePurgeTasksRequest
	GetContent() *string
	SetEndTime(v string) *DescribePurgeTasksRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribePurgeTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePurgeTasksRequest
	GetPageSize() *int32
	SetSiteId(v int64) *DescribePurgeTasksRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribePurgeTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribePurgeTasksRequest
	GetStatus() *string
	SetType(v string) *DescribePurgeTasksRequest
	GetType() *string
}

type DescribePurgeTasksRequest struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePurgeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePurgeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribePurgeTasksRequest) GetContent() *string {
	return s.Content
}

func (s *DescribePurgeTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePurgeTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePurgeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePurgeTasksRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribePurgeTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePurgeTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribePurgeTasksRequest) GetType() *string {
	return s.Type
}

func (s *DescribePurgeTasksRequest) SetContent(v string) *DescribePurgeTasksRequest {
	s.Content = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetEndTime(v string) *DescribePurgeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetPageNumber(v int32) *DescribePurgeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetPageSize(v int32) *DescribePurgeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetSiteId(v int64) *DescribePurgeTasksRequest {
	s.SiteId = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetStartTime(v string) *DescribePurgeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetStatus(v string) *DescribePurgeTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribePurgeTasksRequest) SetType(v string) *DescribePurgeTasksRequest {
	s.Type = &v
	return s
}

func (s *DescribePurgeTasksRequest) Validate() error {
	return dara.Validate(s)
}
