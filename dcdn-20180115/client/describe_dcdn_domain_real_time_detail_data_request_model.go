// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeDetailDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetEndTime() *string
	SetField(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetField() *string
	SetIspNameEn(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetLocationNameEn() *string
	SetMerge(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetMerge() *string
	SetMergeLocIsp(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetMergeLocIsp() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeDetailDataRequest
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeDetailDataRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time, and the maximum time range to query is 10 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of information that you want to query. Separate multiple types with commas (,). Valid values:
	//
	// 	- **qps**: queries per second (QPS)
	//
	// 	- **bps**: bandwidth
	//
	// 	- **http_code**: HTTP status code
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The name of the ISP. You can call the [DescribeDcdnRegionAndIsp](~~DescribeDcdnRegionAndIsp~~) operation to obtain the ISP name.
	//
	// If you do not set this parameter, data of all ISPs is queried.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. You can call the [DescribeDcdnRegionAndIsp](~~DescribeDcdnRegionAndIsp~~) operation to obtain the region name.
	//
	// If you do not set this parameter, all regions are queried.
	//
	// example:
	//
	// beijing
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// Specifies whether to return a summary value. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// true
	Merge *string `json:"Merge,omitempty" xml:"Merge,omitempty"`
	// Specifies whether to return a summary value of **LocationNameEn*	- and **IspNameEn**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// true
	MergeLocIsp *string `json:"MergeLocIsp,omitempty" xml:"MergeLocIsp,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeDetailDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeDetailDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetField() *string {
	return s.Field
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetMerge() *string {
	return s.Merge
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetMergeLocIsp() *string {
	return s.MergeLocIsp
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetDomainName(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetEndTime(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetField(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetIspNameEn(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetLocationNameEn(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetMerge(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.Merge = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetMergeLocIsp(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.MergeLocIsp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) SetStartTime(v string) *DescribeDcdnDomainRealTimeDetailDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataRequest) Validate() error {
	return dara.Validate(s)
}
