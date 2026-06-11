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
	// The action to take when a request exceeds the token rate limit.
	LimitMode *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	// The scope of the rate limit, such as per user or per project.
	LimitType *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	// The maximum number of tokens allowed within the defined time period. For example, if the time unit is one minute, this value represents the tokens-per-minute (TPM) limit.
	LimitValue *int32 `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	// The key that identifies the request source. Its value is extracted from the request context to apply the rule.
	MatchKey *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	// The matching logic applied to the value of `matchKey`.
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The value to match. The rate limit applies only when the value of `matchKey` in the request matches this value, according to the `matchType`.
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
