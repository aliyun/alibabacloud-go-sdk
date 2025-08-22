// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainQpsDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainQpsDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainQpsDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainQpsDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainQpsDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeDcdnDomainQpsDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainQpsDataByLayerRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainQpsDataByLayerRequest
	GetStartTime() *string
}

type DescribeDcdnDomainQpsDataByLayerRequest struct {
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify a domain name, data of all domain names is queried.
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
	// The name of the ISP. You can call the DescribeDcdnRegionAndIsp operation to query the ISP name. If you do not specify a value for this parameter, all ISPs are queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The layers at which you want to query the QPS. The network layer supports IPv4 and IPv6. The application layer supports http, https, and quic. You can also set the value to all. Default value: all.
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
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The minimum data granularity is 5 minutes. If you do not set this parameter, data in the last 24 hours is queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainQpsDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) SetDomainName(v string) *DescribeDcdnDomainQpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) SetEndTime(v string) *DescribeDcdnDomainQpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) SetInterval(v string) *DescribeDcdnDomainQpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) SetIspNameEn(v string) *DescribeDcdnDomainQpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) SetLayer(v string) *DescribeDcdnDomainQpsDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeDcdnDomainQpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) SetStartTime(v string) *DescribeDcdnDomainQpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
