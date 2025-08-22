// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHttpCodeDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest
	GetStartTime() *string
}

type DescribeDcdnDomainHttpCodeDataByLayerRequest struct {
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time needs to be in UTC.
	//
	// > The end time needs to be later than the start time.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). You can call the DescribeDcdnRegionAndIsp operation to query the ISP name. If you do not specify a value for this parameter, all ISPs are queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The layer at which you want to query the bandwidth data. The network layer supports IPv4 and IPv6. The application layer supports http, https, and quic. You can also set the value to all. Default value: all.
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The name of the region. You can call the DescribeDcdnRegionAndIsp operation to query the region name. If you do not specify a value for this parameter, all regions are queried.
	//
	// example:
	//
	// hangzhou
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time needs to be in UTC. The minimum data granularity is 5 minutes. If you do not set this parameter, data in the last 24 hours is queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) SetInterval(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) SetIspNameEn(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) SetLayer(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) SetLocationNameEn(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
