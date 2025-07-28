// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResolverRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteResolverRuleRequest
	GetLang() *string
	SetRuleId(v string) *DeleteResolverRuleRequest
	GetRuleId() *string
}

type DeleteResolverRuleRequest struct {
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hr****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteResolverRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteResolverRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteResolverRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteResolverRuleRequest) SetLang(v string) *DeleteResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteResolverRuleRequest) SetRuleId(v string) *DeleteResolverRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteResolverRuleRequest) Validate() error {
	return dara.Validate(s)
}
