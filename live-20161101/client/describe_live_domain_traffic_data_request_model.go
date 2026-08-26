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
	// The streaming domain. You can specify a single domain name or multiple domain names. Separate multiple domain names with commas (,). If this parameter is left empty, the merged data of all live streaming domains is returned by default.
	//
	// > - When you specify DomainName, make sure that the specified domain names are live streaming domains and that you have the required permissions to operate on the specified domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC.
	//
	// example:
	//
	// 2017-12-10T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for querying data. Unit: seconds. Valid values:
	//
	// - **300*	- (default).
	//
	// - **3600**.
	//
	// - **86400**.
	//
	// > If you do not set this parameter or set it to an unsupported value, the default value **300*	- seconds is used.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP) in English. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to obtain the ISP name. If you do not set this parameter, data of all ISPs is returned.
	//
	// example:
	//
	// alibaba
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region in English. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to obtain the region name. If you do not set this parameter, data of all regions is returned.
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
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC.
	//
	// >You can query data from the past **90*	- days.
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
