// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
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
	// 2019-11-30T05:39:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
