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
	// The ingest domain.
	//
	// Batch domain name queries are supported. Separate multiple domain names with commas (,).
	//
	// If this parameter is left empty, the merged data of all ingest domains is returned by default.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Unit: seconds. Valid values:
	//
	// - **300**
	//
	// - **3600**
	//
	// - **86400**
	//
	// If you do not specify this parameter or the specified value is not supported, the default value 300 is used.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP) in English.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to obtain ISP names. If you do not specify this parameter, data of all ISPs is returned.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region in English.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to obtain region names. If you do not specify this parameter, data of all regions is returned.
	//
	// example:
	//
	// tianjin
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// If you do not specify this parameter, data of the last 24 hours is returned by default.
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
