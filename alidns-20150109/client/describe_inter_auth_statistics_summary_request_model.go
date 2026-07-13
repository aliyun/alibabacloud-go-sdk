// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetDirection() *string
	SetDomainName(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetDomainName() *string
	SetEndTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest
	GetEndTimestamp() *int64
	SetGrowType(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetGrowType() *string
	SetOrderBy(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeInterAuthStatisticsSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInterAuthStatisticsSummaryRequest
	GetPageSize() *int32
	SetPeriod(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetPeriod() *string
	SetRcode(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetRcode() *string
	SetServerRegion(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetServerRegion() *string
	SetSourceType(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetSourceType() *string
	SetStartTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest
	GetStartTimestamp() *int64
	SetStatisticalType(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetStatisticalType() *string
	SetZoneName(v string) *DescribeInterAuthStatisticsSummaryRequest
	GetZoneName() *string
}

type DescribeInterAuthStatisticsSummaryRequest struct {
	// The sort direction. Valid values:
	//
	// - DESC (default): descending order
	//
	// - ASC: ascending order.
	//
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in milliseconds.
	//
	// 	Warning: If the query time range is large and the domain name has an excessive volume of resolution logs, the query may time out or return inaccurate results..
	//
	// example:
	//
	// 1741526400000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Used for the report of subdomain names with sudden increases or decreases in request volume.
	//
	// example:
	//
	// up
	//
	// down
	GrowType *string `json:"GrowType,omitempty" xml:"GrowType,omitempty"`
	// The sort parameter. Valid values:
	//
	// - createDate: sorts by creation time (default if left empty)
	//
	// - expireDate: sorts by expiration time.
	//
	// example:
	//
	// default
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. The value starts from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sampling period. Valid values:
	//
	// - DAY: day
	//
	// - WEEK: week
	//
	// - MONTH: month.
	//
	// example:
	//
	// DAY
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The DNS response code.
	//
	// example:
	//
	// 0
	Rcode *string `json:"Rcode,omitempty" xml:"Rcode,omitempty"`
	// The region of the resolution cluster.
	//
	// example:
	//
	// ap-southeast-3
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
	// The route type. Valid values: cloud: cloud route. local: on-premises route.
	//
	// example:
	//
	// File
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1516779348000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The type of statistical metric.
	//
	// example:
	//
	// QTYPE_RATIO
	//
	// 枚举值：
	//
	// QTYPE_RATIO：DNS请求记录类型占比
	//
	// RCODE_DOMAIN：否定应答域名TOP排行
	//
	// RCODE_SOURCE_IP：某否定应答域名的请求源地址TOP排行
	//
	// REQUEST_ZONE：域名请求量排行（zone级别）
	//
	// REQUEST_DOMAIN：子域名请求量排行（domain name级别）
	//
	// SOURCE_REGION：请求来源地域分布
	//
	// SOURCE_ISP：请求来源运营商（ISP）分布
	//
	// SOURCE_IP：请求源IP详情（含地域、运营商、占比）
	//
	// LINE_HIT：解析线路命中详情
	//
	// LINE_RATIO：解析线路流量占比
	StatisticalType *string `json:"StatisticalType,omitempty" xml:"StatisticalType,omitempty"`
	// The zone name.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInterAuthStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetGrowType() *string {
	return s.GrowType
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetRcode() *string {
	return s.Rcode
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetStatisticalType() *string {
	return s.StatisticalType
}

func (s *DescribeInterAuthStatisticsSummaryRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetDirection(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetDomainName(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetEndTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetGrowType(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.GrowType = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetOrderBy(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeInterAuthStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetPageSize(v int32) *DescribeInterAuthStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetPeriod(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.Period = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetRcode(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.Rcode = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetServerRegion(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetSourceType(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetStartTimestamp(v int64) *DescribeInterAuthStatisticsSummaryRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetStatisticalType(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.StatisticalType = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) SetZoneName(v string) *DescribeInterAuthStatisticsSummaryRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
