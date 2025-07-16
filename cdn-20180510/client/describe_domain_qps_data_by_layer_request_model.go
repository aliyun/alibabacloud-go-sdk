// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainQpsDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainQpsDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDomainQpsDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDomainQpsDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeDomainQpsDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeDomainQpsDataByLayerRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainQpsDataByLayerRequest
	GetStartTime() *string
}

type DescribeDomainQpsDataByLayerRequest struct {
	// The accelerated domain name. You can specify a maximum of 500 domain names in a request. Separate multiple domain names with commas (,).
	//
	// By default, this operation queries the QPS of all accelerated domain names that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format in the ISO 8601 standard. The time is displayed in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries to return. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Description**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP) for your Alibaba Cloud CDN service. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISPs. If you do not set this parameter, all ISPs are queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The layers at which you want to query the number of queries per second. Valid values:
	//
	// 	- **Network layer**: **IPv4**and **IPv6**.
	//
	// 	- **Application layer**: **http**, **https**, and **quic**.
	//
	// 	- **all**: The default value. Both the network and application layers are included.
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query the most recent region list. If you do not set this parameter, all regions are queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format in the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainQpsDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainQpsDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDomainQpsDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainQpsDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDomainQpsDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainQpsDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainQpsDataByLayerRequest) SetDomainName(v string) *DescribeDomainQpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetEndTime(v string) *DescribeDomainQpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetInterval(v string) *DescribeDomainQpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainQpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetLayer(v string) *DescribeDomainQpsDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainQpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetStartTime(v string) *DescribeDomainQpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
