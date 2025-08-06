// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeDetailDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetEndTime() *string
	SetField(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetField() *string
	SetIspNameEn(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetIspNameEn() *string
	SetLocationNameEn(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetLocationNameEn() *string
	SetMerge(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetMerge() *string
	SetMergeLocIsp(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetMergeLocIsp() *string
	SetOwnerId(v int64) *DescribeVodDomainRealTimeDetailDataRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodDomainRealTimeDetailDataRequest
	GetStartTime() *string
}

type DescribeVodDomainRealTimeDetailDataRequest struct {
	// The accelerated domain name. You can specify a maximum of 20 accelerated domain names in each call. Separate domain names with commas (,).
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
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-23T12:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of data that you want to query. You can specify multiple data types and separate them with commas (,). Valid values:
	//
	// qps: the number of queries per second bps: bandwidth data http_code: HTTP status codes
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The name of the region. If you do not specify a region, data in all regions is queried.
	//
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// Specifies whether to return a summary value. Valid values:
	//
	// true false (default)
	//
	// example:
	//
	// false
	Merge *string `json:"Merge,omitempty" xml:"Merge,omitempty"`
	// Specifies whether to return a summary value. Valid values:
	//
	// 	- **true**: groups the results by domain name and merges the results by region and ISP.
	//
	// 	- **false**: groups the results by domain name.
	//
	// Default value: **false**.
	//
	// example:
	//
	// true
	MergeLocIsp *string `json:"MergeLocIsp,omitempty" xml:"MergeLocIsp,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2019-11-30T05:33:00Z.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-23T12:35:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeDetailDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeDetailDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetField() *string {
	return s.Field
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetIspNameEn() *string {
	return s.IspNameEn
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetLocationNameEn() *string {
	return s.LocationNameEn
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetMerge() *string {
	return s.Merge
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetMergeLocIsp() *string {
	return s.MergeLocIsp
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetDomainName(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetEndTime(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetField(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetIspNameEn(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetLocationNameEn(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetMerge(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.Merge = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetMergeLocIsp(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.MergeLocIsp = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetOwnerId(v int64) *DescribeVodDomainRealTimeDetailDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) SetStartTime(v string) *DescribeVodDomainRealTimeDetailDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataRequest) Validate() error {
	return dara.Validate(s)
}
