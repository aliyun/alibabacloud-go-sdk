// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailDataByLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainDetailDataByLayerRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainDetailDataByLayerRequest
	GetEndTime() *string
	SetField(v string) *DescribeDomainDetailDataByLayerRequest
	GetField() *string
	SetIspNameEn(v string) *DescribeDomainDetailDataByLayerRequest
	GetIspNameEn() *string
	SetLayer(v string) *DescribeDomainDetailDataByLayerRequest
	GetLayer() *string
	SetLocationNameEn(v string) *DescribeDomainDetailDataByLayerRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainDetailDataByLayerRequest
	GetStartTime() *string
}

type DescribeDomainDetailDataByLayerRequest struct {
	// The name of the Internet service provider (ISP) for your Alibaba Cloud CDN service. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISP names.
	//
	// If you do not specify an ISP, data of all ISPs is queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The protocol by which you want to query data. Valid values: **http**, **https**, **quic**, and **all**.
	//
	// The default value is **all**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-07-05T22:05:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// bps,ipv6_traf,traf,http_code,qps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// all
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The detailed data of the accelerated domain names.
	//
	// example:
	//
	// hangzhou
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
	//
	// If you do not specify a region, data in all regions is queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-07-05T22:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainDetailDataByLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainDetailDataByLayerRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainDetailDataByLayerRequest) GetField() *string {
	return s.Field
}

func (s *DescribeDomainDetailDataByLayerRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainDetailDataByLayerRequest) GetLayer() *string {
	return s.Layer
}

func (s *DescribeDomainDetailDataByLayerRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainDetailDataByLayerRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainDetailDataByLayerRequest) SetDomainName(v string) *DescribeDomainDetailDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetEndTime(v string) *DescribeDomainDetailDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetField(v string) *DescribeDomainDetailDataByLayerRequest {
	s.Field = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainDetailDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetLayer(v string) *DescribeDomainDetailDataByLayerRequest {
	s.Layer = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainDetailDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetStartTime(v string) *DescribeDomainDetailDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) Validate() error {
	return dara.Validate(s)
}
