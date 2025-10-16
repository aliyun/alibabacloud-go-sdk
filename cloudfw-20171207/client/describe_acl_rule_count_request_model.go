// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclRuleCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAclRuleCountRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAclRuleCountRequest
	GetSourceIp() *string
}

type DescribeAclRuleCountRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 27.151.85.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAclRuleCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclRuleCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclRuleCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAclRuleCountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAclRuleCountRequest) SetLang(v string) *DescribeAclRuleCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclRuleCountRequest) SetSourceIp(v string) *DescribeAclRuleCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAclRuleCountRequest) Validate() error {
	return dara.Validate(s)
}
