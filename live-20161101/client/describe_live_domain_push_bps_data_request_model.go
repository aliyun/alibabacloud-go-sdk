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
	// The end of the time range to query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// The end time must be later than the start time.
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
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to obtain ISP names. If you do not specify this parameter, data for all ISPs is returned.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region in English.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to obtain region names. If you do not specify this parameter, data for all regions is returned.
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
	// The beginning of the time range to query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// If you do not specify this parameter, data from the last 24 hours is returned by default.
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
