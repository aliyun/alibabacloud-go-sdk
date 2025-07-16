// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeDetailDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainRealTimeDetailDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeDetailDataRequest
	GetEndTime() *string
	SetField(v string) *DescribeDomainRealTimeDetailDataRequest
	GetField() *string
	SetIspNameEn(v string) *DescribeDomainRealTimeDetailDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeDomainRealTimeDetailDataRequest
	GetLocationNameEn() *string
	SetMerge(v string) *DescribeDomainRealTimeDetailDataRequest
	GetMerge() *string
	SetMergeLocIsp(v string) *DescribeDomainRealTimeDetailDataRequest
	GetMergeLocIsp() *string
	SetStartTime(v string) *DescribeDomainRealTimeDetailDataRequest
	GetStartTime() *string
}

type DescribeDomainRealTimeDetailDataRequest struct {
	// The accelerated domain name that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2019-11-30T05:40:00Z.
	//
	// > The end time must be later than the start time. The difference between the end time and the start time cannot exceed 10 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of data that you want to query. You can specify multiple data types and separate them with commas (,). Valid values:
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
	// qps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The name of the Internet service provider (ISP). You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query ISPs. If you do not specify an ISP, data of all ISPs is queried.
	//
	// example:
	//
	// telecom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. If you do not specify a region, data in all regions is queried. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
	//
	// example:
	//
	// Guangdong
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// Specifies whether to return a summary value based on ISPs and regions. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	Merge *string `json:"Merge,omitempty" xml:"Merge,omitempty"`
	// Specifies whether to return a summary value based on ISPs and regions. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	MergeLocIsp *string `json:"MergeLocIsp,omitempty" xml:"MergeLocIsp,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2019-11-30T05:33:00Z.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeDetailDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeDetailDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetField() *string {
	return s.Field
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetMerge() *string {
	return s.Merge
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetMergeLocIsp() *string {
	return s.MergeLocIsp
}

func (s *DescribeDomainRealTimeDetailDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetDomainName(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetEndTime(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetField(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetMerge(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.Merge = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetMergeLocIsp(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.MergeLocIsp = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetStartTime(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) Validate() error {
	return dara.Validate(s)
}
