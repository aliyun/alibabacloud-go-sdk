// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallDropTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeFirewallDropTrendRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeFirewallDropTrendRequest
	GetLang() *string
	SetStartTime(v int64) *DescribeFirewallDropTrendRequest
	GetStartTime() *int64
}

type DescribeFirewallDropTrendRequest struct {
	// The end time of the query. Specify the value as a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1758474000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start time of the query. Specify the value as a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1758470400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFirewallDropTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallDropTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallDropTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeFirewallDropTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFirewallDropTrendRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeFirewallDropTrendRequest) SetEndTime(v int64) *DescribeFirewallDropTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFirewallDropTrendRequest) SetLang(v string) *DescribeFirewallDropTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFirewallDropTrendRequest) SetStartTime(v int64) *DescribeFirewallDropTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFirewallDropTrendRequest) Validate() error {
	return dara.Validate(s)
}
