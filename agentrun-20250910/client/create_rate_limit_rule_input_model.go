// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRateLimitRuleInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescriptorId(v string) *CreateRateLimitRuleInput
	GetDescriptorId() *string
	SetDescriptorType(v string) *CreateRateLimitRuleInput
	GetDescriptorType() *string
	SetEnabled(v bool) *CreateRateLimitRuleInput
	GetEnabled() *bool
	SetWindows(v []*WindowLimit) *CreateRateLimitRuleInput
	GetWindows() []*WindowLimit
}

type CreateRateLimitRuleInput struct {
	// 限流规则的描述符标识；非model维时必填，model维由服务端拼接
	//
	// example:
	//
	// model:gpt-4
	DescriptorId *string `json:"descriptorId,omitempty" xml:"descriptorId,omitempty"`
	// 限流规则的描述符类型，如model、user等
	//
	// This parameter is required.
	//
	// example:
	//
	// model
	DescriptorType *string `json:"descriptorType,omitempty" xml:"descriptorType,omitempty"`
	// 是否启用该限流规则
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 限流时间窗口配置列表，至少需要一个窗口
	//
	// This parameter is required.
	Windows []*WindowLimit `json:"windows" xml:"windows" type:"Repeated"`
}

func (s CreateRateLimitRuleInput) String() string {
	return dara.Prettify(s)
}

func (s CreateRateLimitRuleInput) GoString() string {
	return s.String()
}

func (s *CreateRateLimitRuleInput) GetDescriptorId() *string {
	return s.DescriptorId
}

func (s *CreateRateLimitRuleInput) GetDescriptorType() *string {
	return s.DescriptorType
}

func (s *CreateRateLimitRuleInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateRateLimitRuleInput) GetWindows() []*WindowLimit {
	return s.Windows
}

func (s *CreateRateLimitRuleInput) SetDescriptorId(v string) *CreateRateLimitRuleInput {
	s.DescriptorId = &v
	return s
}

func (s *CreateRateLimitRuleInput) SetDescriptorType(v string) *CreateRateLimitRuleInput {
	s.DescriptorType = &v
	return s
}

func (s *CreateRateLimitRuleInput) SetEnabled(v bool) *CreateRateLimitRuleInput {
	s.Enabled = &v
	return s
}

func (s *CreateRateLimitRuleInput) SetWindows(v []*WindowLimit) *CreateRateLimitRuleInput {
	s.Windows = v
	return s
}

func (s *CreateRateLimitRuleInput) Validate() error {
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
