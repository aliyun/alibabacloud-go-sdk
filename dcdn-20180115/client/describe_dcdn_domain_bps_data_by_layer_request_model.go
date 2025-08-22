// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainBpsDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainBpsDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainBpsDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainBpsDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainBpsDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeDcdnDomainBpsDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainBpsDataByLayerRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainBpsDataByLayerRequest
	GetStartTime() *string
}

type DescribeDcdnDomainBpsDataByLayerRequest struct {
	// The accelerated domain name. Separate mutiple domain names with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated.
	//
	// If you do not specify a domain name, data of all domain names is queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
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
	// The Internet service provider (ISP) name. You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query the ISP name. If you do not specify this parameter, all ISPs are queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The layer at which you want to query the bandwidth data. The network layer supports IPv4 and IPv6. The application layer supports http, https, and quic. You can also set the value to all.
	//
	// Default value: all.
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The region name. You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query regions. If you do not specify this parameter, all regions are queried.
	//
	// example:
	//
	// hangzhou
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The minimum data granularity is 5 minutes.
	//
	// If you do not set this parameter, data in the last 24 hours is queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainBpsDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) SetDomainName(v string) *DescribeDcdnDomainBpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) SetEndTime(v string) *DescribeDcdnDomainBpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) SetInterval(v string) *DescribeDcdnDomainBpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) SetIspNameEn(v string) *DescribeDcdnDomainBpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) SetLayer(v string) *DescribeDcdnDomainBpsDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeDcdnDomainBpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) SetStartTime(v string) *DescribeDcdnDomainBpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
