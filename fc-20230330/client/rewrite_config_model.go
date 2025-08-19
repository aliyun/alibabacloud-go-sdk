// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRewriteConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEqualRules(v []*EqualRule) *RewriteConfig
	GetEqualRules() []*EqualRule
	SetRegexRules(v []*RegexRule) *RewriteConfig
	GetRegexRules() []*RegexRule
	SetWildcardRules(v []*WildcardRule) *RewriteConfig
	GetWildcardRules() []*WildcardRule
}

type RewriteConfig struct {
	EqualRules    []*EqualRule    `json:"equalRules" xml:"equalRules" type:"Repeated"`
	RegexRules    []*RegexRule    `json:"regexRules" xml:"regexRules" type:"Repeated"`
	WildcardRules []*WildcardRule `json:"wildcardRules" xml:"wildcardRules" type:"Repeated"`
}

func (s RewriteConfig) String() string {
	return dara.Prettify(s)
}

func (s RewriteConfig) GoString() string {
	return s.String()
}

func (s *RewriteConfig) GetEqualRules() []*EqualRule {
	return s.EqualRules
}

func (s *RewriteConfig) GetRegexRules() []*RegexRule {
	return s.RegexRules
}

func (s *RewriteConfig) GetWildcardRules() []*WildcardRule {
	return s.WildcardRules
}

func (s *RewriteConfig) SetEqualRules(v []*EqualRule) *RewriteConfig {
	s.EqualRules = v
	return s
}

func (s *RewriteConfig) SetRegexRules(v []*RegexRule) *RewriteConfig {
	s.RegexRules = v
	return s
}

func (s *RewriteConfig) SetWildcardRules(v []*WildcardRule) *RewriteConfig {
	s.WildcardRules = v
	return s
}

func (s *RewriteConfig) Validate() error {
	return dara.Validate(s)
}
