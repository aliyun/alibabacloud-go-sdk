// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainRealTimeTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainRealTimeTrafficDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeLiveDomainRealTimeTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeLiveDomainRealTimeTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainRealTimeTrafficDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainRealTimeTrafficDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainRealTimeTrafficDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainRealTimeTrafficDataRequest struct {
	// The streaming domain.
	//
	// Separate multiple streaming domains with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-10T15:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the ISP.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query a list of available ISPs.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query a list of available regions.
	//
	// example:
	//
	// tianjin
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If you do not specify this parameter, the data of the last hour is returned.
	//
	// example:
	//
	// 2015-12-10T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetIspNameEn(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetOwnerId(v int64) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetRegionId(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeLiveDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
