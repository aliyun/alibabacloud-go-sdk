// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainHttpCodeDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainHttpCodeDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainHttpCodeDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeDomainHttpCodeDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries the number and proportions of HTTP status codes for all accelerated domain names that belong to your Alibaba Cloud account.
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
	// 2021-06-29T05:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the region. You can call the DescribeCdnRegionAndIsp operation to query regions. If you do not specify this parameter, data in all regions is queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the Internet service provider (ISP). You can call the DescribeCdnRegionAndIsp operation to query ISPs. If you do not specify this parameter, data of all ISPs is queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-06-29T05:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainHttpCodeDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetInterval(v string) *DescribeDomainHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDomainHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDomainHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
