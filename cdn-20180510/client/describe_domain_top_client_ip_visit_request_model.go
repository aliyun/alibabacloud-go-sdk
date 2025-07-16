// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopClientIpVisitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainTopClientIpVisitRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainTopClientIpVisitRequest
	GetEndTime() *string
	SetLimit(v string) *DescribeDomainTopClientIpVisitRequest
	GetLimit() *string
	SetLocationNameEn(v string) *DescribeDomainTopClientIpVisitRequest
	GetLocationNameEn() *string
	SetSortBy(v string) *DescribeDomainTopClientIpVisitRequest
	GetSortBy() *string
	SetStartTime(v string) *DescribeDomainTopClientIpVisitRequest
	GetStartTime() *string
}

type DescribeDomainTopClientIpVisitRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// By default, this operation queries client IP addresses for all accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2019-10-01T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return. Maximum value: 100.
	//
	// Default value: 20. The default value specifies that the top 20 IP addresses are returned.
	//
	// example:
	//
	// 20
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the region. Separate multiple region names with commas (,).
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The criterion by which you want to sort client IP addresses. Valid values:
	//
	// 	- **traf**: by network traffic. This is the default value.
	//
	// 	- **acc**: by the number of requests.
	//
	// example:
	//
	// traf
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopClientIpVisitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainTopClientIpVisitRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainTopClientIpVisitRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeDomainTopClientIpVisitRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainTopClientIpVisitRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDomainTopClientIpVisitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainTopClientIpVisitRequest) SetDomainName(v string) *DescribeDomainTopClientIpVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetEndTime(v string) *DescribeDomainTopClientIpVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetLimit(v string) *DescribeDomainTopClientIpVisitRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetLocationNameEn(v string) *DescribeDomainTopClientIpVisitRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetSortBy(v string) *DescribeDomainTopClientIpVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetStartTime(v string) *DescribeDomainTopClientIpVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) Validate() error {
	return dara.Validate(s)
}
