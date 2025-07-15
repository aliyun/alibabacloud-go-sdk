// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPushBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainPushBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPushBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainPushBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeLiveDomainPushBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeLiveDomainPushBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainPushBpsDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainPushBpsDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainPushBpsDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainPushBpsDataRequest struct {
	// The ingest domain. You can specify multiple ingest domains and separate them with commas (,). If you do not specify this parameter, the merged data of all your ingest domains is returned.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- **300**
	//
	// 	- **3600**
	//
	// 	- **86400**
	//
	// The default value is 300. If you specify an invalid value or do not specify this parameter, the default value is used.
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
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC. If you do not specify this parameter, the data of the last 24 hours is returned.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainPushBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainPushBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetDomainName(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetEndTime(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetInterval(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetIspNameEn(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetOwnerId(v int64) *DescribeLiveDomainPushBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetRegionId(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) SetStartTime(v string) *DescribeLiveDomainPushBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
