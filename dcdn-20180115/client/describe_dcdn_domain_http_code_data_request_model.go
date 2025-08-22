// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainHttpCodeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainHttpCodeDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainHttpCodeDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainHttpCodeDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainHttpCodeDataRequest struct {
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
	// 2019-03-02T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see **Usage notes**.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query ISPs.
	//
	// example:
	//
	// unicom
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
	// 2019-03-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainHttpCodeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetDomainName(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetEndTime(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetInterval(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) SetStartTime(v string) *DescribeDcdnDomainHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainHttpCodeDataRequest) Validate() error {
	return dara.Validate(s)
}
