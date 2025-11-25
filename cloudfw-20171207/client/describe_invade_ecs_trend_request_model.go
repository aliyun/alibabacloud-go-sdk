// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEcsTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInvadeEcsTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInvadeEcsTrendRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeInvadeEcsTrendRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInvadeEcsTrendRequest
	GetStartTime() *string
}

type DescribeInvadeEcsTrendRequest struct {
	// example:
	//
	// 1733796528
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 120.230.45.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1736561456
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEcsTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEcsTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEcsTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInvadeEcsTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInvadeEcsTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInvadeEcsTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvadeEcsTrendRequest) SetEndTime(v string) *DescribeInvadeEcsTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEcsTrendRequest) SetLang(v string) *DescribeInvadeEcsTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEcsTrendRequest) SetSourceIp(v string) *DescribeInvadeEcsTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEcsTrendRequest) SetStartTime(v string) *DescribeInvadeEcsTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInvadeEcsTrendRequest) Validate() error {
	return dara.Validate(s)
}
