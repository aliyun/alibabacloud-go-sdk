// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeIspInfoRequest
	GetLang() *string
	SetSourceCode(v string) *DescribeIspInfoRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeIspInfoRequest
	GetSourceIp() *string
}

type DescribeIspInfoRequest struct {
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
	// 39.91.37.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeIspInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeIspInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIspInfoRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeIspInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeIspInfoRequest) SetLang(v string) *DescribeIspInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIspInfoRequest) SetSourceCode(v string) *DescribeIspInfoRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeIspInfoRequest) SetSourceIp(v string) *DescribeIspInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIspInfoRequest) Validate() error {
	return dara.Validate(s)
}
