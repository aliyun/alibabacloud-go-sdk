// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRateLimitRuleInput interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateRateLimitRuleInput
	GetEnabled() *bool
	SetWindows(v []*WindowLimit) *UpdateRateLimitRuleInput
	GetWindows() []*WindowLimit
}

type UpdateRateLimitRuleInput struct {
	// 是否启用该限流规则
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 限流时间窗口配置列表，提供时将整组覆盖
	Windows []*WindowLimit `json:"windows" xml:"windows" type:"Repeated"`
}

func (s UpdateRateLimitRuleInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateRateLimitRuleInput) GoString() string {
	return s.String()
}

func (s *UpdateRateLimitRuleInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateRateLimitRuleInput) GetWindows() []*WindowLimit {
	return s.Windows
}

func (s *UpdateRateLimitRuleInput) SetEnabled(v bool) *UpdateRateLimitRuleInput {
	s.Enabled = &v
	return s
}

func (s *UpdateRateLimitRuleInput) SetWindows(v []*WindowLimit) *UpdateRateLimitRuleInput {
	s.Windows = v
	return s
}

func (s *UpdateRateLimitRuleInput) Validate() error {
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
