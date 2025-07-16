// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainBpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainBpsDataRequest
	GetStartTime() *string
}

type DescribeDomainBpsDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries bandwidth data for all accelerated domain names that belong to your Alibaba Cloud account.
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
	// 2020-05-14T10:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISP names.
	//
	// If you do not set this parameter, data of all ISPs is queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
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
	// 2020-05-14T09:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainBpsDataRequest) SetDomainName(v string) *DescribeDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetEndTime(v string) *DescribeDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetInterval(v string) *DescribeDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetIspNameEn(v string) *DescribeDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetStartTime(v string) *DescribeDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
