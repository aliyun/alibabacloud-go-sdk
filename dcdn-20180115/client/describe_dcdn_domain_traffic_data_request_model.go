// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainTrafficDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainTrafficDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainTrafficDataRequest struct {
	// The accelerated domain name.
	//
	// Separate multiple domain names with commas (,). If you do not specify a value for this parameter, network traffic of all accelerated domain names is queried.
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
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query ISPs. If you do not specify an ISP, network traffic of all ISPs is queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query regions. If you do not specify a region, network traffic in all regions is queried.
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

func (s DescribeDcdnDomainTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
