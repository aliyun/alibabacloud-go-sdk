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
	// 限流规则的创建时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 限流规则的描述符标识，用于关联具体的限流对象
	//
	// example:
	//
	// model:gpt-4
	DescriptorId *string `json:"descriptorId,omitempty" xml:"descriptorId,omitempty"`
	// 限流规则的描述符类型，如model、user等
	//
	// example:
	//
	// model
	DescriptorType *string `json:"descriptorType,omitempty" xml:"descriptorType,omitempty"`
	// 限流规则是否启用，true表示启用，false表示禁用
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 限流规则最后一次更新的时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// 限流规则的唯一标识符
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789abc
	RateLimitRuleId *string `json:"rateLimitRuleId,omitempty" xml:"rateLimitRuleId,omitempty"`
	// 限流时间窗口配置列表，支持多个窗口叠加限流
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
