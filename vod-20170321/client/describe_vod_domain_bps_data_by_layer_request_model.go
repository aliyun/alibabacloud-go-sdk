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
	// The accelerated domain name. You can specify a maximum of 500 accelerated domain names. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-01-23T12:40:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries. Unit: seconds.
	//
	// The time granularity varies based on the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see the supported time granularity described in Usage notes.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The layer at which you want to query the data.
	//
	// Network layer: IPv4 and IPv6. Application layer: http, https, and quic. all: specifies that both the network and application layers are included. Default value: all.
	//
	// example:
	//
	// IPv4
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The name of the region.
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
