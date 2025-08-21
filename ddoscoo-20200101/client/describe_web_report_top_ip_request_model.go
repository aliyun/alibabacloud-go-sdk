// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebReportTopIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeWebReportTopIpRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeWebReportTopIpRequest
	GetEndTime() *int64
	SetInterval(v int32) *DescribeWebReportTopIpRequest
	GetInterval() *int32
	SetQueryType(v string) *DescribeWebReportTopIpRequest
	GetQueryType() *string
	SetStartTime(v int64) *DescribeWebReportTopIpRequest
	GetStartTime() *int64
	SetTop(v int32) *DescribeWebReportTopIpRequest
	GetTop() *int32
}

type DescribeWebReportTopIpRequest struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query the domain names for which forwarding rules are configured.
	//
	// example:
	//
	// app.bmjqxvb.cn
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1687228200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval at which data is collected. Unit: seconds. Valid values are 300, 3600, and 86400.
	//
	// 	- If the time span between StartTime and EndTime is less than 3 days (3 days excluded), valid values are 300, 3600, and 86400.
	//
	// 	- If the time span between StartTime and EndTime is from 3 to 31 days (31 days excluded), valid values are 3600 and 86400.
	//
	// 	- If the time span between StartTime and EndTime is 31 days or longer, the valid value is 86400.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The source of the statistics. Valid values:
	//
	// 	- **visit**: indicates all IP addresses.
	//
	// 	- **block**: indicates blocked IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1680424200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 5
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeWebReportTopIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebReportTopIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebReportTopIpRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebReportTopIpRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeWebReportTopIpRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeWebReportTopIpRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeWebReportTopIpRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeWebReportTopIpRequest) GetTop() *int32 {
	return s.Top
}

func (s *DescribeWebReportTopIpRequest) SetDomain(v string) *DescribeWebReportTopIpRequest {
	s.Domain = &v
	return s
}

func (s *DescribeWebReportTopIpRequest) SetEndTime(v int64) *DescribeWebReportTopIpRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeWebReportTopIpRequest) SetInterval(v int32) *DescribeWebReportTopIpRequest {
	s.Interval = &v
	return s
}

func (s *DescribeWebReportTopIpRequest) SetQueryType(v string) *DescribeWebReportTopIpRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeWebReportTopIpRequest) SetStartTime(v int64) *DescribeWebReportTopIpRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeWebReportTopIpRequest) SetTop(v int32) *DescribeWebReportTopIpRequest {
	s.Top = &v
	return s
}

func (s *DescribeWebReportTopIpRequest) Validate() error {
	return dara.Validate(s)
}
