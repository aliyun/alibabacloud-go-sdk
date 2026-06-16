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
	// Specifies whether to enable the rate limit rule.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// A list of time window configurations. Providing this parameter replaces the entire existing list.
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
