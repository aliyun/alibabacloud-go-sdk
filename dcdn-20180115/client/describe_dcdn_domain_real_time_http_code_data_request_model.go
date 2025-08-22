// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeHttpCodeDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query ISPs.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query regions.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
