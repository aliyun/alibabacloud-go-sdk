// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHttpCodeDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainHttpCodeDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainHttpCodeDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainHttpCodeDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainHttpCodeDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeDomainHttpCodeDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeDomainHttpCodeDataByLayerRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainHttpCodeDataByLayerRequest
	GetStartTime() *string
}

type DescribeDomainHttpCodeDataByLayerRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// If you do not specify this parameter, data of all accelerated domain names under your account is queried.
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
	// 2020-07-06T22:00:00Z
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
	// If you do not specify an ISP, data of all ISPs is queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The protocol by which you want to query HTTP status codes. The network layer supports **IPv4*	- and **IPv6**. The application layer supports **http**, **https**, and **quic**. You can also set the value to **all**.
	//
	// Default value: **all**
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
	//
	// If you do not specify a region, data in all regions is queried.
	//
	// example:
	//
	// hangzhou
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-07-05T22:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHttpCodeDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetDomainName(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetEndTime(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetInterval(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetLayer(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetStartTime(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
