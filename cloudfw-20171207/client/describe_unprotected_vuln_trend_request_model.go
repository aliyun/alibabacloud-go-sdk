// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnprotectedVulnTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeUnprotectedVulnTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeUnprotectedVulnTrendRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeUnprotectedVulnTrendRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeUnprotectedVulnTrendRequest
	GetStartTime() *string
}

type DescribeUnprotectedVulnTrendRequest struct {
	// example:
	//
	// 1754878752
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 112.15.190.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1740623016
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUnprotectedVulnTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedVulnTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedVulnTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUnprotectedVulnTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUnprotectedVulnTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeUnprotectedVulnTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUnprotectedVulnTrendRequest) SetEndTime(v string) *DescribeUnprotectedVulnTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendRequest) SetLang(v string) *DescribeUnprotectedVulnTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendRequest) SetSourceIp(v string) *DescribeUnprotectedVulnTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendRequest) SetStartTime(v string) *DescribeUnprotectedVulnTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUnprotectedVulnTrendRequest) Validate() error {
	return dara.Validate(s)
}
