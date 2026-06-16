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
	// The descriptor ID for the rate limit rule. This parameter is required if the descriptor type is not `model`. For `model` types, the server assembles the ID.
	//
	// example:
	//
	// model:gpt-4
	DescriptorId *string `json:"descriptorId,omitempty" xml:"descriptorId,omitempty"`
	// The descriptor type for the rate limit rule, such as `model` or `user`.
	//
	// This parameter is required.
	//
	// example:
	//
	// model
	DescriptorType *string `json:"descriptorType,omitempty" xml:"descriptorType,omitempty"`
	// Whether to enable the rate limit rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// A list of time window configurations. At least one window is required.
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
