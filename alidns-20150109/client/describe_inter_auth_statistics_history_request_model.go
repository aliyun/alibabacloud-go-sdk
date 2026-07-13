// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetDomainName() *string
	SetEndTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest
	GetEndTimestamp() *int64
	SetRcode(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetRcode() *string
	SetServerRegion(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetServerRegion() *string
	SetStartTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest
	GetStartTimestamp() *int64
	SetStatisticalType(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetStatisticalType() *string
	SetZoneName(v string) *DescribeInterAuthStatisticsHistoryRequest
	GetZoneName() *string
}

type DescribeInterAuthStatisticsHistoryRequest struct {
	// The domain name.<props="china">You can get this value by calling the [DescribeDomains](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0) operation.
	//
	// <props="intl">You can get this value by calling the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation.
	//
	// example:
	//
	// dns-example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. This value is a Unix timestamp in milliseconds.
	//
	// 	Warning: A large time range for a domain with a high volume of resolution logs may cause query timeouts or inaccurate results.
	//
	// example:
	//
	// 1741526400000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The DNS response code.
	//
	// example:
	//
	// 0
	Rcode *string `json:"Rcode,omitempty" xml:"Rcode,omitempty"`
	// The server region.
	//
	// example:
	//
	// ap-southeast-3
	ServerRegion *string `json:"ServerRegion,omitempty" xml:"ServerRegion,omitempty"`
	// The start of the time range to query. This value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1474335170000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The statistical metric type.
	//
	// example:
	//
	// PROTOCOL
	//
	// 其它：
	//
	// PROTOCOL：DNS请求UDP协议、TCP协议维度统计（仅全局域名统计）
	//
	// QTYPE：DNS请求记录类型分布，如A、AAAA、CNAME、MX 等
	//
	// RCODE：DNS否定应答统计
	//
	// SUCCESS_RATIO：SUCCESS_RATIO：平均解析成功率（解析成功率 = 返回与请求记录类型（QTYPE）匹配的有效应答次数 / 总解析次数）
	//
	// REQUEST：DNS请求量统计，展示QPS查询趋势
	StatisticalType *string `json:"StatisticalType,omitempty" xml:"StatisticalType,omitempty"`
	// The zone name.
	//
	// example:
	//
	// lisheng999.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInterAuthStatisticsHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetRcode() *string {
	return s.Rcode
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetServerRegion() *string {
	return s.ServerRegion
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetStatisticalType() *string {
	return s.StatisticalType
}

func (s *DescribeInterAuthStatisticsHistoryRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetDomainName(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetEndTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetRcode(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.Rcode = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetServerRegion(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.ServerRegion = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetStartTimestamp(v int64) *DescribeInterAuthStatisticsHistoryRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetStatisticalType(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.StatisticalType = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) SetZoneName(v string) *DescribeInterAuthStatisticsHistoryRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryRequest) Validate() error {
	return dara.Validate(s)
}
