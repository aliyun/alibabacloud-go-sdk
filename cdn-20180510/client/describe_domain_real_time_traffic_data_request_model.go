// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeTrafficDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeTrafficDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeTrafficDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDomainRealTimeTrafficDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainRealTimeTrafficDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainRealTimeTrafficDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeTrafficDataRequest struct {
	// The accelerated domain name. You can specify up to 100 domain names in each call. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query the most recent region list. If you do not set this parameter, all regions are queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query the most recent region list. If you do not set this parameter, all regions are queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeTrafficDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeTrafficDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeTrafficDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainRealTimeTrafficDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainRealTimeTrafficDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) Validate() error {
	return dara.Validate(s)
}
