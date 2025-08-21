// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7GlobalRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ConfigL7GlobalRuleRequest
	GetDomain() *string
	SetRuleAttr(v string) *ConfigL7GlobalRuleRequest
	GetRuleAttr() *string
}

type ConfigL7GlobalRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{\\"RuleId\\":\\"global_01\\",\\"Action\\":\\"block\\",\\"Enabled\\":0}]
	RuleAttr *string `json:"RuleAttr,omitempty" xml:"RuleAttr,omitempty"`
}

func (s ConfigL7GlobalRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7GlobalRuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigL7GlobalRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigL7GlobalRuleRequest) GetRuleAttr() *string {
	return s.RuleAttr
}

func (s *ConfigL7GlobalRuleRequest) SetDomain(v string) *ConfigL7GlobalRuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigL7GlobalRuleRequest) SetRuleAttr(v string) *ConfigL7GlobalRuleRequest {
	s.RuleAttr = &v
	return s
}

func (s *ConfigL7GlobalRuleRequest) Validate() error {
	return dara.Validate(s)
}
