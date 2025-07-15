// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainBpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainBpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeLiveDomainBpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeLiveDomainBpsDataRequest
	GetLocationNameEn() *string
	SetOwnerId(v int64) *DescribeLiveDomainBpsDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainBpsDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainBpsDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainBpsDataRequest struct {
	// The streaming domain. You can query one or more domain names. If you specify multiple domain names, separate them with commas (,). If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-10T09:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- **300*	- (default)
	//
	// 	- **3600**
	//
	// 	- **86400**
	//
	// >
	//
	// 	- If you specify an invalid value or do not specify this parameter, the default value **300*	- is used.
	//
	// 	- When the time granularity is **300*	- seconds, the returned bandwidth is the average bandwidth within the 300 seconds.
	//
	// 	- When the time granularity is **3600*	- or **86400*	- seconds, the returned bandwidth is the peak value of all average bandwidths within each 300-second period.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP). You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query a list of available ISPs.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query a list of available regions.
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
	// 2017-12-10T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainBpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeLiveDomainBpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeLiveDomainBpsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainBpsDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainBpsDataRequest) SetDomainName(v string) *DescribeLiveDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetEndTime(v string) *DescribeLiveDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetInterval(v string) *DescribeLiveDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetIspNameEn(v string) *DescribeLiveDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeLiveDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetOwnerId(v int64) *DescribeLiveDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetRegionId(v string) *DescribeLiveDomainBpsDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) SetStartTime(v string) *DescribeLiveDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
