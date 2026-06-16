// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRateLimitRuleResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RateLimitRuleResult
	GetCode() *string
	SetData(v *RateLimitRule) *RateLimitRuleResult
	GetData() *RateLimitRule
	SetRequestId(v string) *RateLimitRuleResult
	GetRequestId() *string
}

type RateLimitRuleResult struct {
	// A value of `SUCCESS` indicates the request succeeded. On failure, this field returns the corresponding error type.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the rate limit rule.
	//
	// example:
	//
	// {}
	Data *RateLimitRule `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request identifier for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RateLimitRuleResult) String() string {
	return dara.Prettify(s)
}

func (s RateLimitRuleResult) GoString() string {
	return s.String()
}

func (s *RateLimitRuleResult) GetCode() *string {
	return s.Code
}

func (s *RateLimitRuleResult) GetData() *RateLimitRule {
	return s.Data
}

func (s *RateLimitRuleResult) GetRequestId() *string {
	return s.RequestId
}

func (s *RateLimitRuleResult) SetCode(v string) *RateLimitRuleResult {
	s.Code = &v
	return s
}

func (s *RateLimitRuleResult) SetData(v *RateLimitRule) *RateLimitRuleResult {
	s.Data = v
	return s
}

func (s *RateLimitRuleResult) SetRequestId(v string) *RateLimitRuleResult {
	s.RequestId = &v
	return s
}

func (s *RateLimitRuleResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
