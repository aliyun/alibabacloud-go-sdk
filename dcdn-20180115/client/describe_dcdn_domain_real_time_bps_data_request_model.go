// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeBpsDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDcdnDomainRealTimeBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeBpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeBpsDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeBpsDataRequest struct {
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
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2018-01-02T11:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// If you do not set this parameter, all ISPs are queried. You can call [DescribeDcdnRegionAndIsp](~~DescribeDcdnRegionAndIsp~~) to query ISP names.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// If you do not set this parameter, all regions are queried. You can call [DescribeDcdnRegionAndIsp](~~DescribeDcdnRegionAndIsp~~) to query regions.
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
	// 2018-01-02T11:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
