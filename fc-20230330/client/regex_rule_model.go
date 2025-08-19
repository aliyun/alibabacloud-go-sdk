// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegexRule interface {
	dara.Model
	String() string
	GoString() string
	SetMatch(v string) *RegexRule
	GetMatch() *string
	SetReplacement(v string) *RegexRule
	GetReplacement() *string
}

type RegexRule struct {
	// This parameter is required.
	//
	// example:
	//
	// ^/api/.+?/(.*)
	Match *string `json:"match,omitempty" xml:"match,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /api/v2/$1
	Replacement *string `json:"replacement,omitempty" xml:"replacement,omitempty"`
}

func (s RegexRule) String() string {
	return dara.Prettify(s)
}

func (s RegexRule) GoString() string {
	return s.String()
}

func (s *RegexRule) GetMatch() *string {
	return s.Match
}

func (s *RegexRule) GetReplacement() *string {
	return s.Replacement
}

func (s *RegexRule) SetMatch(v string) *RegexRule {
	s.Match = &v
	return s
}

func (s *RegexRule) SetReplacement(v string) *RegexRule {
	s.Replacement = &v
	return s
}

func (s *RegexRule) Validate() error {
	return dara.Validate(s)
}
