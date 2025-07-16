// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAverageResponseTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainAverageResponseTimeRequest
	GetDomainName() *string
	SetDomainType(v string) *DescribeDomainAverageResponseTimeRequest
	GetDomainType() *string
	SetEndTime(v string) *DescribeDomainAverageResponseTimeRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainAverageResponseTimeRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainAverageResponseTimeRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainAverageResponseTimeRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainAverageResponseTimeRequest
	GetStartTime() *string
	SetTimeMerge(v string) *DescribeDomainAverageResponseTimeRequest
	GetTimeMerge() *string
}

type DescribeDomainAverageResponseTimeRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// By default, this operation queries the geographic distribution of users for all accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the query condition. When you set the value to dynamic, this operation queries the average response time of dynamic resources and static resources. If you do not set this parameter, this operation queries the average response time of only static resources.
	//
	// example:
	//
	// domaintype
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries. Unit: seconds. The value varies based on the values of the **StartTime*	- and **EndTime*	- parameters. Valid values:
	//
	// 	- If the time span between StartTime and EndTime is less than 3 days, valid values are **300**, **3600**, and **86400**. Default value: **300**.
	//
	// 	- If the time span between StartTime and EndTime is greater than or equal to 3 days and less than 31 days, valid values are **3600*	- and **86400**. Default value: **3600**.
	//
	// 	- If the time range between StartTime and EndTime is 31 days or longer, the valid value is **86400**. Default value: **86400**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP) for your Alibaba Cloud CDN service. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISPs. If you do not set this parameter, data of all ISPs is queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions. If you do not set this parameter, data in all regions is queried.
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
	// Specifies whether to automatically set the interval. If you set the value to 1, the value of the Interval parameter is automatically assigned based on the StartTime and EndTime parameters. You can set this parameter or the Interval parameter.
	//
	// example:
	//
	// 1
	TimeMerge *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDomainAverageResponseTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainAverageResponseTimeRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDomainAverageResponseTimeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainAverageResponseTimeRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainAverageResponseTimeRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainAverageResponseTimeRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainAverageResponseTimeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainAverageResponseTimeRequest) GetTimeMerge() *string {
	return s.TimeMerge
}

func (s *DescribeDomainAverageResponseTimeRequest) SetDomainName(v string) *DescribeDomainAverageResponseTimeRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetDomainType(v string) *DescribeDomainAverageResponseTimeRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetEndTime(v string) *DescribeDomainAverageResponseTimeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetInterval(v string) *DescribeDomainAverageResponseTimeRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetIspNameEn(v string) *DescribeDomainAverageResponseTimeRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetLocationNameEn(v string) *DescribeDomainAverageResponseTimeRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetStartTime(v string) *DescribeDomainAverageResponseTimeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetTimeMerge(v string) *DescribeDomainAverageResponseTimeRequest {
	s.TimeMerge = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) Validate() error {
	return dara.Validate(s)
}
