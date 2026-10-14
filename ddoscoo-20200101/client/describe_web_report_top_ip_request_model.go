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
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Top       *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
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
