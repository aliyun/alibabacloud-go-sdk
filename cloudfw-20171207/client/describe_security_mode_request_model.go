// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSecurityModeRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeSecurityModeRequest
	GetSourceIp() *string
}

type DescribeSecurityModeRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 218.108.54.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSecurityModeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityModeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSecurityModeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSecurityModeRequest) SetLang(v string) *DescribeSecurityModeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecurityModeRequest) SetSourceIp(v string) *DescribeSecurityModeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityModeRequest) Validate() error {
	return dara.Validate(s)
}
