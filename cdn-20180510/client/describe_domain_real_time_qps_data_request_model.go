// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeQpsDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDomainRealTimeQpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainRealTimeQpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDomainRealTimeQpsDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeQpsDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// > You can specify up to 500 domain names in each request.
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
	// 2019-12-02T11:26:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// If you do not set this parameter, data of all ISPs is queried. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISP names.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// If you do not set this parameter, data in all regions is queried. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
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
	// 2019-12-02T11:25:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeQpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainRealTimeQpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainRealTimeQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetDomainName(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetEndTime(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetStartTime(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
