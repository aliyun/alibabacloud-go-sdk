// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetEndTime() *string
	SetFixTimeGap(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetFixTimeGap() *string
	SetInterval(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetStartTime() *string
	SetTimeMerge(v string) *DescribeDcdnDomainIpaBpsDataRequest
	GetTimeMerge() *string
}

type DescribeDcdnDomainIpaBpsDataRequest struct {
	// The accelerated domain name.
	//
	// Separate multiple domain names with commas (,). If you leave this parameter empty, all accelerated domain names are queried.
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
	// Specifies whether to implement padding with zeros. Valid values:
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
	// 	- If the time range between StartTime and EndTime is less than 3 days, the valid values are **300**, **3600**, and **86400**. If you leave this parameter empty, **300*	- is used.
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
	// Unicom
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
	// Specifies whether to automatically set the interval. If you set **TimeMerge*	- to **1**, the value of the **Interval*	- parameter is automatically assigned based on the **startTime*	- and **endTime*	- parameters. You can specify either this parameter or the **Interval*	- parameter.
	//
	// example:
	//
	// 1
	TimeMerge *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
}

func (s DescribeDcdnDomainIpaBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetFixTimeGap() *string {
	return s.FixTimeGap
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) GetTimeMerge() *string {
	return s.TimeMerge
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetFixTimeGap(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.FixTimeGap = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetInterval(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) SetTimeMerge(v string) *DescribeDcdnDomainIpaBpsDataRequest {
	s.TimeMerge = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
