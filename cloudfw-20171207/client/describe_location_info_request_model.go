// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeLocationInfoRequest
	GetLang() *string
	SetSourceCode(v string) *DescribeLocationInfoRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeLocationInfoRequest
	GetSourceIp() *string
}

type DescribeLocationInfoRequest struct {
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
	// 180.169.141.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLocationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocationInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLocationInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeLocationInfoRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeLocationInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeLocationInfoRequest) SetLang(v string) *DescribeLocationInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLocationInfoRequest) SetSourceCode(v string) *DescribeLocationInfoRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeLocationInfoRequest) SetSourceIp(v string) *DescribeLocationInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLocationInfoRequest) Validate() error {
	return dara.Validate(s)
}
