// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPushTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainPushTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPushTrafficDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainPushTrafficDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeLiveDomainPushTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeLiveDomainPushTrafficDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainPushTrafficDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainPushTrafficDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainPushTrafficDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainPushTrafficDataRequest struct {
	// The ingest domain. You can specify multiple ingest domains and separate them with commas (,). If you do not specify this parameter, the merged data of all your ingest domains is returned.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
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

func (s DescribeLiveDomainPushTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainPushTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetDomainName(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetEndTime(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetInterval(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetIspNameEn(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetOwnerId(v int64) *DescribeLiveDomainPushTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetRegionId(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) SetStartTime(v string) *DescribeLiveDomainPushTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
