// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTokenRateLimitConfigRule interface {
	dara.Model
	String() string
	GoString() string
	SetLimitMode(v string) *AiTokenRateLimitConfigRule
	GetLimitMode() *string
	SetLimitType(v string) *AiTokenRateLimitConfigRule
	GetLimitType() *string
	SetLimitValue(v int32) *AiTokenRateLimitConfigRule
	GetLimitValue() *int32
	SetMatchKey(v string) *AiTokenRateLimitConfigRule
	GetMatchKey() *string
	SetMatchType(v string) *AiTokenRateLimitConfigRule
	GetMatchType() *string
	SetMatchValue(v string) *AiTokenRateLimitConfigRule
	GetMatchValue() *string
}

type AiTokenRateLimitConfigRule struct {
	LimitMode  *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	LimitType  *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	LimitValue *int32  `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	MatchKey   *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	MatchType  *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
}

func (s AiTokenRateLimitConfigRule) String() string {
	return dara.Prettify(s)
}

func (s AiTokenRateLimitConfigRule) GoString() string {
	return s.String()
}

func (s *AiTokenRateLimitConfigRule) GetLimitMode() *string {
	return s.LimitMode
}

func (s *AiTokenRateLimitConfigRule) GetLimitType() *string {
	return s.LimitType
}

func (s *AiTokenRateLimitConfigRule) GetLimitValue() *int32 {
	return s.LimitValue
}

func (s *AiTokenRateLimitConfigRule) GetMatchKey() *string {
	return s.MatchKey
}

func (s *AiTokenRateLimitConfigRule) GetMatchType() *string {
	return s.MatchType
}

func (s *AiTokenRateLimitConfigRule) GetMatchValue() *string {
	return s.MatchValue
}

func (s *AiTokenRateLimitConfigRule) SetLimitMode(v string) *AiTokenRateLimitConfigRule {
	s.LimitMode = &v
	return s
}

func (s *AiTokenRateLimitConfigRule) SetLimitType(v string) *AiTokenRateLimitConfigRule {
	s.LimitType = &v
	return s
}

func (s *AiTokenRateLimitConfigRule) SetLimitValue(v int32) *AiTokenRateLimitConfigRule {
	s.LimitValue = &v
	return s
}

func (s *AiTokenRateLimitConfigRule) SetMatchKey(v string) *AiTokenRateLimitConfigRule {
	s.MatchKey = &v
	return s
}

func (s *AiTokenRateLimitConfigRule) SetMatchType(v string) *AiTokenRateLimitConfigRule {
	s.MatchType = &v
	return s
}

func (s *AiTokenRateLimitConfigRule) SetMatchValue(v string) *AiTokenRateLimitConfigRule {
	s.MatchValue = &v
	return s
}

func (s *AiTokenRateLimitConfigRule) Validate() error {
	return dara.Validate(s)
}
