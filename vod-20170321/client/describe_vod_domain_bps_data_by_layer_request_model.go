// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainBpsDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainBpsDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainBpsDataByLayerRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodDomainBpsDataByLayerRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeVodDomainBpsDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeVodDomainBpsDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeVodDomainBpsDataByLayerRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeVodDomainBpsDataByLayerRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainBpsDataByLayerRequest
	GetStartTime() *string
}

type DescribeVodDomainBpsDataByLayerRequest struct {
	// The accelerated domain name to query.
	//
	// - If you do not specify this parameter, the pooled data of all accelerated domain names is returned by default.
	//
	// - Batch queries are supported. Separate multiple domain names with commas (,). You can specify up to 500 domain names at a time.
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com), and choose **Configuration Management > CDN Configuration > Domain Names*	- in the left-side navigation pane to view the accelerated domain names that you have added to ApsaraVideo VOD. You can also call the [DescribeVodUserDomains](~~DescribeVodUserDomains~~) operation to query the list of accelerated domain names.
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
	// 2019-01-23T12:40:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data. Unit: seconds. Valid values: **300**, **3600**, and **86400**. If you do not specify this parameter or specify an unsupported value, the default value is used. The supported time granularity varies based on the time range specified by `StartTime` and `EndTime`:
	//
	// - Less than 3 days (exclusive): **300*	- (default), **3600**, and **86400**.
	//
	// - 3 to 31 days (exclusive): **3600*	- (default) and **86400**.
	//
	// - 31 days or more: **86400*	- (default).
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP) in English. If you do not specify this parameter, data of all ISPs is queried by default.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The protocol type. You can specify the protocol type at the network layer or application layer.
	//
	// Default value:
	//
	// - all: includes both network layer and application layer
	//
	// Network layer values:
	//
	// - IPv4
	//
	// - IPv6
	//
	// Application layer values:
	//
	// - http
	//
	// - https
	//
	// - quic
	//
	// example:
	//
	// IPv4
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The name of the region in English. If you do not specify this parameter, data of all regions is queried by default.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-01-23T12:35:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainBpsDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainBpsDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetDomainName(v string) *DescribeVodDomainBpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetEndTime(v string) *DescribeVodDomainBpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetInterval(v string) *DescribeVodDomainBpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetIspNameEn(v string) *DescribeVodDomainBpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetLayer(v string) *DescribeVodDomainBpsDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeVodDomainBpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetOwnerId(v int64) *DescribeVodDomainBpsDataByLayerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) SetStartTime(v string) *DescribeVodDomainBpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
