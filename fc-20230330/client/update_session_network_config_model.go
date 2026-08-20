// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowOut(v []*string) *UpdateSessionNetworkConfig
	GetAllowOut() []*string
	SetDenyOut(v []*string) *UpdateSessionNetworkConfig
	GetDenyOut() []*string
	SetRules(v map[string][]*SessionNetworkRule) *UpdateSessionNetworkConfig
	GetRules() map[string][]*SessionNetworkRule
}

type UpdateSessionNetworkConfig struct {
	AllowOut []*string `json:"allowOut" xml:"allowOut" type:"Repeated"`
	DenyOut  []*string `json:"denyOut" xml:"denyOut" type:"Repeated"`
	// The request transform rules configured by exact target host. If omitted, existing rules are retained. An empty object clears all rules, and a non-empty object replaces all rules entirely. Null is not supported. The transform.headers and transform.headerValueReplacements fields are supported.
	Rules map[string][]*SessionNetworkRule `json:"rules" xml:"rules"`
}

func (s UpdateSessionNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionNetworkConfig) GoString() string {
	return s.String()
}

func (s *UpdateSessionNetworkConfig) GetAllowOut() []*string {
	return s.AllowOut
}

func (s *UpdateSessionNetworkConfig) GetDenyOut() []*string {
	return s.DenyOut
}

func (s *UpdateSessionNetworkConfig) GetRules() map[string][]*SessionNetworkRule {
	return s.Rules
}

func (s *UpdateSessionNetworkConfig) SetAllowOut(v []*string) *UpdateSessionNetworkConfig {
	s.AllowOut = v
	return s
}

func (s *UpdateSessionNetworkConfig) SetDenyOut(v []*string) *UpdateSessionNetworkConfig {
	s.DenyOut = v
	return s
}

func (s *UpdateSessionNetworkConfig) SetRules(v map[string][]*SessionNetworkRule) *UpdateSessionNetworkConfig {
	s.Rules = v
	return s
}

func (s *UpdateSessionNetworkConfig) Validate() error {
	return dara.Validate(s)
}
