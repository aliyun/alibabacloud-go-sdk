// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRateLimitRule interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *RateLimitRule
	GetCreatedAt() *string
	SetDescriptorId(v string) *RateLimitRule
	GetDescriptorId() *string
	SetDescriptorType(v string) *RateLimitRule
	GetDescriptorType() *string
	SetEnabled(v bool) *RateLimitRule
	GetEnabled() *bool
	SetLastUpdatedAt(v string) *RateLimitRule
	GetLastUpdatedAt() *string
	SetRateLimitRuleId(v string) *RateLimitRule
	GetRateLimitRuleId() *string
	SetWindows(v []*WindowLimit) *RateLimitRule
	GetWindows() []*WindowLimit
}

type RateLimitRule struct {
	// The creation time of the rate limit rule, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The descriptor ID for the rate limit rule, which associates the rule with a specific throttling target.
	//
	// example:
	//
	// model:gpt-4
	DescriptorId *string `json:"descriptorId,omitempty" xml:"descriptorId,omitempty"`
	// The descriptor type for the rate limit rule, such as \\"model\\" or \\"user\\".
	//
	// example:
	//
	// model
	DescriptorType *string `json:"descriptorType,omitempty" xml:"descriptorType,omitempty"`
	// Indicates whether the rate limit rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The last update time of the rate limit rule, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The unique identifier for the rate limit rule.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789abc
	RateLimitRuleId *string `json:"rateLimitRuleId,omitempty" xml:"rateLimitRuleId,omitempty"`
	// A list of time window configurations. Multiple windows can be used to enforce layered rate limiting.
	Windows []*WindowLimit `json:"windows" xml:"windows" type:"Repeated"`
}

func (s RateLimitRule) String() string {
	return dara.Prettify(s)
}

func (s RateLimitRule) GoString() string {
	return s.String()
}

func (s *RateLimitRule) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *RateLimitRule) GetDescriptorId() *string {
	return s.DescriptorId
}

func (s *RateLimitRule) GetDescriptorType() *string {
	return s.DescriptorType
}

func (s *RateLimitRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *RateLimitRule) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *RateLimitRule) GetRateLimitRuleId() *string {
	return s.RateLimitRuleId
}

func (s *RateLimitRule) GetWindows() []*WindowLimit {
	return s.Windows
}

func (s *RateLimitRule) SetCreatedAt(v string) *RateLimitRule {
	s.CreatedAt = &v
	return s
}

func (s *RateLimitRule) SetDescriptorId(v string) *RateLimitRule {
	s.DescriptorId = &v
	return s
}

func (s *RateLimitRule) SetDescriptorType(v string) *RateLimitRule {
	s.DescriptorType = &v
	return s
}

func (s *RateLimitRule) SetEnabled(v bool) *RateLimitRule {
	s.Enabled = &v
	return s
}

func (s *RateLimitRule) SetLastUpdatedAt(v string) *RateLimitRule {
	s.LastUpdatedAt = &v
	return s
}

func (s *RateLimitRule) SetRateLimitRuleId(v string) *RateLimitRule {
	s.RateLimitRuleId = &v
	return s
}

func (s *RateLimitRule) SetWindows(v []*WindowLimit) *RateLimitRule {
	s.Windows = v
	return s
}

func (s *RateLimitRule) Validate() error {
	if s.Windows != nil {
		for _, item := range s.Windows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
