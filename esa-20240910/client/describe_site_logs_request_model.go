// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteLogsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeSiteLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSiteLogsRequest
	GetPageSize() *int64
	SetSiteId(v int64) *DescribeSiteLogsRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeSiteLogsRequest
	GetStartTime() *string
}

type DescribeSiteLogsRequest struct {
	// example:
	//
	// 2022-11-06T17:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2022-11-06T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSiteLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSiteLogsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeSiteLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteLogsRequest) SetEndTime(v string) *DescribeSiteLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteLogsRequest) SetPageNumber(v int64) *DescribeSiteLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSiteLogsRequest) SetPageSize(v int64) *DescribeSiteLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSiteLogsRequest) SetSiteId(v int64) *DescribeSiteLogsRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteLogsRequest) SetStartTime(v string) *DescribeSiteLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteLogsRequest) Validate() error {
	return dara.Validate(s)
}
