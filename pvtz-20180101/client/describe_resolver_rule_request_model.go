// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeResolverRuleRequest
	GetLang() *string
	SetRuleId(v string) *DescribeResolverRuleRequest
	GetRuleId() *string
}

type DescribeResolverRuleRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// hr****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeResolverRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeResolverRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeResolverRuleRequest) SetLang(v string) *DescribeResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverRuleRequest) SetRuleId(v string) *DescribeResolverRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeResolverRuleRequest) Validate() error {
	return dara.Validate(s)
}
