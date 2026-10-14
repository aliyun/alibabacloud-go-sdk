// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7GlobalRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeL7GlobalRuleRequest
	GetDomain() *string
	SetLang(v string) *DescribeL7GlobalRuleRequest
	GetLang() *string
}

type DescribeL7GlobalRuleRequest struct {
	// The domain name of the website service to operate on.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The language type for the request and response messages. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeL7GlobalRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7GlobalRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeL7GlobalRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeL7GlobalRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeL7GlobalRuleRequest) SetDomain(v string) *DescribeL7GlobalRuleRequest {
	s.Domain = &v
	return s
}

func (s *DescribeL7GlobalRuleRequest) SetLang(v string) *DescribeL7GlobalRuleRequest {
	s.Lang = &v
	return s
}

func (s *DescribeL7GlobalRuleRequest) Validate() error {
	return dara.Validate(s)
}
