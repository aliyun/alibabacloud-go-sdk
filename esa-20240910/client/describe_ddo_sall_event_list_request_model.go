// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSAllEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDDoSAllEventListRequest
	GetEndTime() *string
	SetEventType(v string) *DescribeDDoSAllEventListRequest
	GetEventType() *string
	SetPageNumber(v int32) *DescribeDDoSAllEventListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDDoSAllEventListRequest
	GetPageSize() *int32
	SetSiteId(v int64) *DescribeDDoSAllEventListRequest
	GetSiteId() *int64
	SetStartTime(v string) *DescribeDDoSAllEventListRequest
	GetStartTime() *string
}

type DescribeDDoSAllEventListRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// This parameter is required.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSAllEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSAllEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSAllEventListRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDoSAllEventListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDDoSAllEventListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDDoSAllEventListRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSAllEventListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSAllEventListRequest) SetEndTime(v string) *DescribeDDoSAllEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetEventType(v string) *DescribeDDoSAllEventListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetPageNumber(v int32) *DescribeDDoSAllEventListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetPageSize(v int32) *DescribeDDoSAllEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetSiteId(v int64) *DescribeDDoSAllEventListRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) SetStartTime(v string) *DescribeDDoSAllEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSAllEventListRequest) Validate() error {
	return dara.Validate(s)
}
