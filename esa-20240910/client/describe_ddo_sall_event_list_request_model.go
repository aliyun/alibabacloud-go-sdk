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
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The maximum time range is 31 days.
	//
	// If you do not configure this parameter, the current time is used as the end of the time range to query.
	//
	// example:
	//
	// 2023-02-22T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of DDoS attacks to query. Valid values:
	//
	// 	- **web-cc**: web resource exhaustion attacks.
	//
	// 	- **cc**: connection flood attacks.
	//
	// 	- **traffic**: volumetric attacks.
	//
	// Default value: web-cc.
	//
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The page number. Valid values: **1*	- to **100000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: 5, 10, and 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
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
