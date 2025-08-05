// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLProtectTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeACLProtectTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeACLProtectTrendRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeACLProtectTrendRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeACLProtectTrendRequest
	GetStartTime() *string
}

type DescribeACLProtectTrendRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp that is accurate to seconds.
	//
	// example:
	//
	// 1670397599
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 223.95.87.130
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp that is accurate to seconds.
	//
	// example:
	//
	// 1677050306
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeACLProtectTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLProtectTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeACLProtectTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeACLProtectTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeACLProtectTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeACLProtectTrendRequest) SetEndTime(v string) *DescribeACLProtectTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeACLProtectTrendRequest) SetLang(v string) *DescribeACLProtectTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeACLProtectTrendRequest) SetSourceIp(v string) *DescribeACLProtectTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeACLProtectTrendRequest) SetStartTime(v string) *DescribeACLProtectTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeACLProtectTrendRequest) Validate() error {
	return dara.Validate(s)
}
