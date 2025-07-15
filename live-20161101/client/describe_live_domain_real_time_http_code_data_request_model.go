// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainRealTimeHttpCodeDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainRealTimeHttpCodeDataRequest struct {
	// The streaming domain.
	//
	// Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If you specify neither the StartTime parameter nor the EndTime parameter, the data of the last **1*	- hour is returned.
	//
	// example:
	//
	// 2015-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/448109.html) operation to query a list of available ISPs.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/448109.html) operation to query a list of available regions.
	//
	// example:
	//
	// tianjin
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If you specify neither the StartTime parameter nor the EndTime parameter, the data of the last **1*	- hour is returned.
	//
	// example:
	//
	// 2015-11-30T05:39:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetDomainName(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetEndTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetIspNameEn(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetOwnerId(v int64) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetRegionId(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) SetStartTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
