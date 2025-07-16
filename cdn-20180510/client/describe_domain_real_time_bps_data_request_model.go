// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeBpsDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDomainRealTimeBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainRealTimeBpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainRealTimeBpsDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeBpsDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// > You can specify up to 500 domain names in each request.
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
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISPs. If you do not set this parameter, all ISPs are queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions. If you do not set this parameter, all regions are queried.
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
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainRealTimeBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainRealTimeBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
