// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainTrafficDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainTrafficDataRequest
	GetStartTime() *string
}

type DescribeDomainTrafficDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries the network traffic for all accelerated domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). You can call the [DescribeCdnRegionAndIsp](~~DescribeCdnRegionAndIsp~~) operation to query ISPs.
	//
	// If you do not specify an ISP, data of all ISPs is queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](~~DescribeCdnRegionAndIsp~~) operation to query regions.
	//
	// If you do not specify a region, data in all regions is queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainTrafficDataRequest) SetDomainName(v string) *DescribeDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetEndTime(v string) *DescribeDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetInterval(v string) *DescribeDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetStartTime(v string) *DescribeDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
