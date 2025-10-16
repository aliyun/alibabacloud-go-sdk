// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRegionInfoRequest
	GetLang() *string
	SetSourceCode(v string) *DescribeRegionInfoRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeRegionInfoRequest
	GetSourceIp() *string
}

type DescribeRegionInfoRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// example:
	//
	// 59.82.59.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRegionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRegionInfoRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeRegionInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRegionInfoRequest) SetLang(v string) *DescribeRegionInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionInfoRequest) SetSourceCode(v string) *DescribeRegionInfoRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeRegionInfoRequest) SetSourceIp(v string) *DescribeRegionInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRegionInfoRequest) Validate() error {
	return dara.Validate(s)
}
