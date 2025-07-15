// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainBpsDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainBpsDataByLayerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainBpsDataByLayerRequest
	GetStartTime() *string
}

type DescribeLiveDomainBpsDataByLayerRequest struct {
	// The streaming domain. You can specify multiple domain names by separating them with commas (,). If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// pull.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be displayed in UTC.
	//
	// example:
	//
	// 2022-03-16T16:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- **300**
	//
	// 	- **3600**
	//
	// 	- **86400**
	//
	// >
	//
	// 	- If the time range specified by the StartTime and EndTime parameters is smaller than or equal to 3 days, the supported time granularities include 300, 3,600, and 86,400 seconds.
	//
	// 	- If the time range is larger than 3 days but smaller than or equal to 31 days, the supported time granularities include 3,600 and 86,400 seconds.
	//
	// 	- If the time range is larger than 31 days, the supported time granularity is 86,400 seconds.
	//
	// 	- If you specify an invalid value or do not specify this parameter, the default time granularity of 300 seconds is used.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). If you do not specify this parameter, the data of all ISPs is returned.
	//
	// >  You can call the [DescribeLiveRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query available regions and ISPs.
	//
	// example:
	//
	// tele***
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The layer at which you want to query the data. Valid values:
	//
	// 	- IPv4 and IPv6 (network layer)
	//
	// 	- http, https, and quic (application layer)
	//
	// 	- all (default)
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The name of the region. If you do not specify this parameter, the data of all regions is returned.
	//
	// >  You can call the [DescribeLiveRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query available regions and ISPs.
	//
	// example:
	//
	// hangzhou
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be displayed in UTC.
	//
	// >  If you do not specify this parameter, the data of the last 24 hours is returned by default. The minimum time granularity is 5 minutes.
	//
	// example:
	//
	// 2022-03-15T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainBpsDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetDomainName(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetEndTime(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetInterval(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetIspNameEn(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetLayer(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetOwnerId(v int64) *DescribeLiveDomainBpsDataByLayerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetRegionId(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) SetStartTime(v string) *DescribeLiveDomainBpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
