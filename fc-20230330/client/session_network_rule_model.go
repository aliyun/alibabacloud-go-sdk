// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSessionNetworkRule interface {
	dara.Model
	String() string
	GoString() string
	SetTransform(v *SessionNetworkRuleTransform) *SessionNetworkRule
	GetTransform() *SessionNetworkRuleTransform
}

type SessionNetworkRule struct {
	Transform *SessionNetworkRuleTransform `json:"transform,omitempty" xml:"transform,omitempty"`
}

func (s SessionNetworkRule) String() string {
	return dara.Prettify(s)
}

func (s SessionNetworkRule) GoString() string {
	return s.String()
}

func (s *SessionNetworkRule) GetTransform() *SessionNetworkRuleTransform {
	return s.Transform
}

func (s *SessionNetworkRule) SetTransform(v *SessionNetworkRuleTransform) *SessionNetworkRule {
	s.Transform = v
	return s
}

func (s *SessionNetworkRule) Validate() error {
	if s.Transform != nil {
		if err := s.Transform.Validate(); err != nil {
			return err
		}
	}
	return nil
}
