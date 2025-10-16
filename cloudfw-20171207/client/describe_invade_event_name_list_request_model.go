// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInvadeEventNameListRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInvadeEventNameListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeInvadeEventNameListRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInvadeEventNameListRequest
	GetStartTime() *string
}

type DescribeInvadeEventNameListRequest struct {
	// example:
	//
	// 1738780437
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 36.112.73.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1757620800
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEventNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventNameListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInvadeEventNameListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInvadeEventNameListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInvadeEventNameListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvadeEventNameListRequest) SetEndTime(v string) *DescribeInvadeEventNameListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEventNameListRequest) SetLang(v string) *DescribeInvadeEventNameListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEventNameListRequest) SetSourceIp(v string) *DescribeInvadeEventNameListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEventNameListRequest) SetStartTime(v string) *DescribeInvadeEventNameListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInvadeEventNameListRequest) Validate() error {
	return dara.Validate(s)
}
