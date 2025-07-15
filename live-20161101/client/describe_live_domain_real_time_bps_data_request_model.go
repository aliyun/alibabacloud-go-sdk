// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainRealTimeBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainRealTimeBpsDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeLiveDomainRealTimeBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeLiveDomainRealTimeBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainRealTimeBpsDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainRealTimeBpsDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainRealTimeBpsDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainRealTimeBpsDataRequest struct {
	// The streaming domain name or names.
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
	// >  If you do not set this parameter, the data of the hour following the specified start time is returned.
	//
	// example:
	//
	// 2015-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// To query ISPs, call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region in which you want to query data.
	//
	// To query regions, call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation.
	//
	// example:
	//
	// tianjin
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2015-11-30T05:39:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetOwnerId(v int64) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetRegionId(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeLiveDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
