// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWildcardRule interface {
	dara.Model
	String() string
	GoString() string
	SetMatch(v string) *WildcardRule
	GetMatch() *string
	SetReplacement(v string) *WildcardRule
	GetReplacement() *string
}

type WildcardRule struct {
	// This parameter is required.
	//
	// example:
	//
	// /api/*
	Match *string `json:"match,omitempty" xml:"match,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /$1
	Replacement *string `json:"replacement,omitempty" xml:"replacement,omitempty"`
}

func (s WildcardRule) String() string {
	return dara.Prettify(s)
}

func (s WildcardRule) GoString() string {
	return s.String()
}

func (s *WildcardRule) GetMatch() *string {
	return s.Match
}

func (s *WildcardRule) GetReplacement() *string {
	return s.Replacement
}

func (s *WildcardRule) SetMatch(v string) *WildcardRule {
	s.Match = &v
	return s
}

func (s *WildcardRule) SetReplacement(v string) *WildcardRule {
	s.Replacement = &v
	return s
}

func (s *WildcardRule) Validate() error {
	return dara.Validate(s)
}
