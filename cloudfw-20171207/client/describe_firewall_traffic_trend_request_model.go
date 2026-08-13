// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTrafficTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeFirewallTrafficTrendRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeFirewallTrafficTrendRequest
	GetLang() *string
	SetStartTime(v int64) *DescribeFirewallTrafficTrendRequest
	GetStartTime() *int64
}

type DescribeFirewallTrafficTrendRequest struct {
	// The end time of the query. Specify a UNIX timestamp in seconds. This parameter is required. If this parameter is not specified, ErrorTimeError (400) is returned.
	//
	// > The query interval (EndTime − StartTime) cannot exceed 90 days. If the interval exceeds 90 days, ErrorTimeError is returned. If the value is later than the current time, it is silently adjusted to the current time.
	//
	// example:
	//
	// 1758474000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response message.
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start time of the query. Specify a UNIX timestamp in seconds. This parameter is required. If this parameter is not specified, ErrorTimeError (400) is returned.
	//
	// > The query interval (EndTime − StartTime) cannot exceed 90 days. If the interval exceeds 90 days, ErrorTimeError is returned. If the value is later than the current time, it is silently adjusted to the current time. If StartTime is later than EndTime, no error is returned, but the response contains empty data.
	//
	// example:
	//
	// 1758470400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFirewallTrafficTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTrafficTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeFirewallTrafficTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFirewallTrafficTrendRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeFirewallTrafficTrendRequest) SetEndTime(v int64) *DescribeFirewallTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFirewallTrafficTrendRequest) SetLang(v string) *DescribeFirewallTrafficTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFirewallTrafficTrendRequest) SetStartTime(v int64) *DescribeFirewallTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFirewallTrafficTrendRequest) Validate() error {
	return dara.Validate(s)
}
