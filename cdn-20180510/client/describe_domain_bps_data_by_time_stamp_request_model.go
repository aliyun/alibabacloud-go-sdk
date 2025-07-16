// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataByTimeStampRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainBpsDataByTimeStampRequest
	GetDomainName() *string
	SetIspNames(v string) *DescribeDomainBpsDataByTimeStampRequest
	GetIspNames() *string
	SetLocationNames(v string) *DescribeDomainBpsDataByTimeStampRequest
	GetLocationNames() *string
	SetTimePoint(v string) *DescribeDomainBpsDataByTimeStampRequest
	GetTimePoint() *string
}

type DescribeDomainBpsDataByTimeStampRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The names of the Internet service providers (ISPs). Separate multiple ISPs with commas (,).
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// uni***,tele***
	IspNames *string `json:"IspNames,omitempty" xml:"IspNames,omitempty"`
	// The regions. Separate multiple regions with commas (,).
	//
	// You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// liaoning,guangxi
	LocationNames *string `json:"LocationNames,omitempty" xml:"LocationNames,omitempty"`
	// The point in time to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The data is collected every 5 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainBpsDataByTimeStampRequest) GetIspNames() *string {
	return s.IspNames
}

func (s *DescribeDomainBpsDataByTimeStampRequest) GetLocationNames() *string {
	return s.LocationNames
}

func (s *DescribeDomainBpsDataByTimeStampRequest) GetTimePoint() *string {
	return s.TimePoint
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetDomainName(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetIspNames(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.IspNames = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetLocationNames(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.LocationNames = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetTimePoint(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) Validate() error {
	return dara.Validate(s)
}
