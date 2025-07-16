// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainBpsDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainBpsDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainBpsDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainBpsDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeDomainBpsDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeDomainBpsDataByLayerRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainBpsDataByLayerRequest
	GetStartTime() *string
}

type DescribeDomainBpsDataByLayerRequest struct {
	// The accelerated domain name. You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
	//
	// > If you do not specify this parameter, the bandwidth data about all accelerated domain names that belong to your Alibaba Cloud account is queried.
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
	// 2020-05-06T07:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). You can call the [DescribeCdnRegionAndIsp](~~DescribeCdnRegionAndIsp~~) operation to query ISPs. If you do not specify an ISP, data of all ISPs is queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The layer at which you want to query the bandwidth data. Valid values:
	//
	// 	- Network layer: **IPv4*	- and **IPv6**.
	//
	// 	- Application layer: **http**, **https**, and **quic**.
	//
	// 	- **all**: specifies that both the network and application layers are included.
	//
	// Default value: **all**.
	//
	// example:
	//
	// IPv4
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](~~DescribeCdnRegionAndIsp~~) operation to query regions. If you do not specify a region, data in all regions is queried.
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
	// 2020-05-06T07:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainBpsDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainBpsDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainBpsDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainBpsDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainBpsDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDomainBpsDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainBpsDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainBpsDataByLayerRequest) SetDomainName(v string) *DescribeDomainBpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetEndTime(v string) *DescribeDomainBpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetInterval(v string) *DescribeDomainBpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainBpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetLayer(v string) *DescribeDomainBpsDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainBpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetStartTime(v string) *DescribeDomainBpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
