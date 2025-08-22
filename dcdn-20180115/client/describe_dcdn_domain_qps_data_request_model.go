// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainQpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainQpsDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainQpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnDomainQpsDataRequest
	GetInterval() *string
	SetIspNameEn(v string) *DescribeDcdnDomainQpsDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainQpsDataRequest
	GetLocationNameEn() *string
	SetStartTime(v string) *DescribeDcdnDomainQpsDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainQpsDataRequest struct {
	// The accelerated domain name.
	//
	// Separate multiple domain names with commas (,). If you do not specify a value for this parameter, all accelerated domain names are queried.
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
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for a query. Unit: seconds.
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
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainQpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainQpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainQpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnDomainQpsDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainQpsDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainQpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainQpsDataRequest) SetDomainName(v string) *DescribeDcdnDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetEndTime(v string) *DescribeDcdnDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetInterval(v string) *DescribeDcdnDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) SetStartTime(v string) *DescribeDcdnDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataRequest) Validate() error {
	return dara.Validate(s)
}
