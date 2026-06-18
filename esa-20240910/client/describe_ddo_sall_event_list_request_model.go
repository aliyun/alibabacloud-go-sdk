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
	// The end time of the query.
	//
	// The time must be in ISO 8601 format and in UTC. Format: `yyyy-MM-ddTHH:mm:ssZ`. The time range between `StartTime` and `EndTime` cannot exceed 31 days.
	//
	// If this parameter is not specified, it defaults to the current time.
	//
	// example:
	//
	// 2023-02-22T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of DDoS attack events to query. Valid values:
	//
	// - **web-cc**: A web resource exhaustion attack.
	//
	// - **cc**: A connection-based attack.
	//
	// - **traffic**: A traffic-based attack.
	//
	// If you do not specify this parameter, the operation queries `web-cc` events by default.
	//
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The page number to return. Valid range: **1*	- to **100000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: **5**, **10**, and **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the site. You can obtain this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time of the query.
	//
	// The time must be in ISO 8601 format and in UTC. Format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-12T15:00:00Z
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
