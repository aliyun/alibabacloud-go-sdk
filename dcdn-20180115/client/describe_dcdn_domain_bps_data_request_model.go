// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainBpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainBpsDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainBpsDataRequest struct {
	// The accelerated domain name.
	//
	// Separate multiple domain names with commas (,). If you do not specify a value for this parameter, bandwidth data of all accelerated domain names is queried.
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
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query ISPs. If you do not specify an ISP, bandwidth data of all ISPs is queried.
	//
	// example:
	//
	// beijing
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query regions. If you do not specify a region, bandwidth data in all regions is queried.
	//
	// example:
	//
	// unicom
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

func (s DescribeDcdnDomainBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetInterval(v string) *DescribeDcdnDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
