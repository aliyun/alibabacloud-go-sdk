// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAclWhitelistRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAclWhitelistRequest
	GetSourceIp() *string
}

type DescribeAclWhitelistRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 101.36.65.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAclWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAclWhitelistRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAclWhitelistRequest) SetLang(v string) *DescribeAclWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclWhitelistRequest) SetSourceIp(v string) *DescribeAclWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAclWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
