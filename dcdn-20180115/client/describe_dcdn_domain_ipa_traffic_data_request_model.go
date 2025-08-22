// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetEndTime() *string
	SetFixTimeGap(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetFixTimeGap() *string
	SetInterval(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetStartTime() *string
	SetTimeMerge(v string) *DescribeDcdnDomainIpaTrafficDataRequest
	GetTimeMerge() *string
}

type DescribeDcdnDomainIpaTrafficDataRequest struct {
	// The accelerated domain name.
	//
	// Separate multiple domain names with commas (,). If you do not specify a value for this parameter, data for all accelerated domain names is queried.
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
	// Specify whether to implement padding with zeros. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	FixTimeGap *string `json:"FixTimeGap,omitempty" xml:"FixTimeGap,omitempty"`
	// The time granularity of data entries. Unit: seconds.
	//
	// The time granularity varies with the time range specified by **StartTime*	- and **EndTime**.
	//
	// 	- If the time range between StartTime and EndTime is less than 3 days, the valid values are **300**, **3600**, and **86400**. If you do not specify a value for this parameter, **300*	- is used.
	//
	// 	- If the time range between StartTime and EndTime is greater than or equal to 3 days and less than 31 days, the valid values are **3600*	- and **86400**. Default value: **3600**.
	//
	// 	- If the time range between StartTime and EndTime is 31 days or longer, the valid value is **86400**. Default value: **86400**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
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
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether to automatically calculate the value of the **interval**. If the **timeMerge*	- parameter is set to **1**, the value of **inteval*	- is calculated based on **StartTime*	- and **EndTime**. You can set either this parameter or the **interval*	- parameter.
	//
	// example:
	//
	// 1
	TimeMerge *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDcdnDomainIpaTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetFixTimeGap() *string {
	return s.FixTimeGap
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) GetTimeMerge() *string {
	return s.TimeMerge
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetDomainName(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetEndTime(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetFixTimeGap(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.FixTimeGap = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetInterval(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetStartTime(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) SetTimeMerge(v string) *DescribeDcdnDomainIpaTrafficDataRequest {
	s.TimeMerge = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
