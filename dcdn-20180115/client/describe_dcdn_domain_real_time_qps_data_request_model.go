// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeQpsDataRequest
	GetEndTime() *string
	SetIspNameEn(v string) *DescribeDcdnDomainRealTimeQpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeQpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeQpsDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeQpsDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
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
	// 2018-01-02T11:27:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query ISPs.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region.
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query regions.
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
	// 2018-01-02T11:26:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
