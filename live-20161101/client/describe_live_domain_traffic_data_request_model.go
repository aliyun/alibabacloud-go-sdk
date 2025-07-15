// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeLiveDomainTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeLiveDomainTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainTrafficDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainTrafficDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainTrafficDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainTrafficDataRequest struct {
	// The streaming domain. You can query one or more domain names. If you specify multiple domain names, separate them with commas (,). If you do not specify this parameter, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- **300*	- (default)
	//
	// 	- **3600**
	//
	// 	- **86400**
	//
	// >  If you specify an invalid value or do not specify this parameter, the default value **300*	- is used.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query a list of available ISPs. If you do not specify this parameter, the data of all ISPs is returned.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query a list of available regions. If you do not specify this parameter, the data of all regions is returned.
	//
	// example:
	//
	// tianjin
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  You can query data in the last **90*	- days.
	//
	// example:
	//
	// 2017-12-10T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainTrafficDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainTrafficDataRequest) SetDomainName(v string) *DescribeLiveDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetEndTime(v string) *DescribeLiveDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetInterval(v string) *DescribeLiveDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeLiveDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeLiveDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetRegionId(v string) *DescribeLiveDomainTrafficDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) SetStartTime(v string) *DescribeLiveDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
