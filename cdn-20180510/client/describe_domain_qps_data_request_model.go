// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainQpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainQpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainQpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainQpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainQpsDataRequest
	GetStartTime() *string
}

type DescribeDomainQpsDataRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries QPS data for all accelerated domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP) for your Alibaba Cloud CDN service. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISPs. If you do not specify an ISP, data of all ISPs is queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions. If you do not specify a region, data in all regions is queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainQpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainQpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainQpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainQpsDataRequest) SetDomainName(v string) *DescribeDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetEndTime(v string) *DescribeDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetInterval(v string) *DescribeDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetIspNameEn(v string) *DescribeDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetStartTime(v string) *DescribeDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
