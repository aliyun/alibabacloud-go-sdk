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
	// The accelerated domain name to query.
	//
	// - Batch queries are supported. Separate multiple domain names with commas (,). You can specify up to 20 domain names at a time.
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com), and in the left-side navigation pane, choose **Configuration Management > CDN Configuration > Domain Names*	- to view the accelerated domain names that you have added to ApsaraVideo VOD. Alternatively, call the [DescribeVodUserDomains](~~DescribeVodUserDomains~~) operation to query the list of accelerated domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time, and the difference between the end time and the start time cannot exceed 10 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-23T12:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of access data to query. You can specify multiple types. Separate multiple types with commas (,). Valid values:
	//
	// - **qps**: queries per second (QPS).
	//
	// - **bps**: bandwidth data.
	//
	// - **http_code**: HTTP status codes.
	//
	// This parameter is required.
	//
	// example:
	//
	// qps
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The Internet service provider (ISP) name in English. If you do not specify this parameter, data for all ISPs is queried by default.
	//
	// example:
	//
	// unicom
	IspNameEn *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	// The region name in English. If you do not specify this parameter, data for all regions is queried by default.
	//
	// example:
	//
	// shanghai
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	// Specifies whether to return aggregated data by domain name. Valid values:
	//
	// - **true**: Returns aggregated data across all domain names.
	//
	// - **false*	- (default): Returns data grouped by domain name.
	//
	// example:
	//
	// false
	Merge *string `json:"Merge,omitempty" xml:"Merge,omitempty"`
	// Specifies whether to return aggregated data by region and ISP. Valid values:
	//
	// - **true**: Returns data grouped only by domain name, with region and ISP values aggregated.
	//
	// - **false*	- (default): Returns data grouped by domain name, region, and ISP.
	//
	// example:
	//
	// true
	MergeLocIsp *string `json:"MergeLocIsp,omitempty" xml:"MergeLocIsp,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
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
