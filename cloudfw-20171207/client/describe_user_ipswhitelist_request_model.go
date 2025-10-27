// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserIPSWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeUserIPSWhitelistRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeUserIPSWhitelistRequest
	GetSourceIp() *string
}

type DescribeUserIPSWhitelistRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 47.100.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUserIPSWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUserIPSWhitelistRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeUserIPSWhitelistRequest) SetLang(v string) *DescribeUserIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserIPSWhitelistRequest) SetSourceIp(v string) *DescribeUserIPSWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserIPSWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
