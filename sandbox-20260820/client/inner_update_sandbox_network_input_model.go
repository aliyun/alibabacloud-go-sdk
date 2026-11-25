// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInnerUpdateSandboxNetworkInput interface {
	dara.Model
	String() string
	GoString() string
	SetAllowInternetAccess(v bool) *InnerUpdateSandboxNetworkInput
	GetAllowInternetAccess() *bool
	SetAllowOut(v []*string) *InnerUpdateSandboxNetworkInput
	GetAllowOut() []*string
	SetDenyOut(v []*string) *InnerUpdateSandboxNetworkInput
	GetDenyOut() []*string
	SetRules(v map[string][]map[string]interface{}) *InnerUpdateSandboxNetworkInput
	GetRules() map[string][]map[string]interface{}
}

type InnerUpdateSandboxNetworkInput struct {
	// The switch for outbound public network access from the sandbox. When set to false, it collapses to the default deny policy of the platform and shares the same outbound policy with allowOut and denyOut. This field is included in the full replacement. If omitted, no deny entries are appended based on this switch. However, if allowOut is not empty and denyOut is empty, denyOut is automatically set to 0.0.0.0/0, and IPv4 outbound traffic to destinations other than those in allowOut is denied.
	//
	// example:
	//
	// true
	AllowInternetAccess *bool `json:"allowInternetAccess,omitempty" xml:"allowInternetAccess,omitempty"`
	// The destinations allowed for outbound access, which can be IP addresses, CIDR blocks, exact domain names, or domain names with a wildcard only in the leftmost label (such as *.example.com). allowOut takes precedence over denyOut. The maximum list length depends on the account quota. If omitted, the existing allow list is cleared.
	//
	// example:
	//
	// ["api.example.com","10.0.0.0/8"]
	AllowOut []*string `json:"allowOut" xml:"allowOut" type:"Repeated"`
	// The destinations denied for outbound access, in the same format as allowOut. If a destination matches both allow and deny rules, the allow rule takes precedence. The maximum list length depends on the account quota. If omitted, the existing deny list is cleared.
	//
	// example:
	//
	// ["blocked.example.com","192.168.0.0/16"]
	DenyOut []*string `json:"denyOut" xml:"denyOut" type:"Repeated"`
	// The outbound network rules where the key is an exact destination hostname or the E2B reserved key @fc.sandbox.identity, and the value is a list of rule objects. The number of hosts and replacement entries is subject to the account quota limit. Omitting rules or passing a null object clears the standard host rules. OAuth options tables are updated individually. Omitting the corresponding reserved header retains the original table, while passing the character string "{}" clears the table.
	//
	// example:
	//
	// {"api.example.com":[{"transform":{"headers":{"X-Example":"demo"}}}]}
	Rules map[string][]map[string]interface{} `json:"rules" xml:"rules"`
}

func (s InnerUpdateSandboxNetworkInput) String() string {
	return dara.Prettify(s)
}

func (s InnerUpdateSandboxNetworkInput) GoString() string {
	return s.String()
}

func (s *InnerUpdateSandboxNetworkInput) GetAllowInternetAccess() *bool {
	return s.AllowInternetAccess
}

func (s *InnerUpdateSandboxNetworkInput) GetAllowOut() []*string {
	return s.AllowOut
}

func (s *InnerUpdateSandboxNetworkInput) GetDenyOut() []*string {
	return s.DenyOut
}

func (s *InnerUpdateSandboxNetworkInput) GetRules() map[string][]map[string]interface{} {
	return s.Rules
}

func (s *InnerUpdateSandboxNetworkInput) SetAllowInternetAccess(v bool) *InnerUpdateSandboxNetworkInput {
	s.AllowInternetAccess = &v
	return s
}

func (s *InnerUpdateSandboxNetworkInput) SetAllowOut(v []*string) *InnerUpdateSandboxNetworkInput {
	s.AllowOut = v
	return s
}

func (s *InnerUpdateSandboxNetworkInput) SetDenyOut(v []*string) *InnerUpdateSandboxNetworkInput {
	s.DenyOut = v
	return s
}

func (s *InnerUpdateSandboxNetworkInput) SetRules(v map[string][]map[string]interface{}) *InnerUpdateSandboxNetworkInput {
	s.Rules = v
	return s
}

func (s *InnerUpdateSandboxNetworkInput) Validate() error {
	return dara.Validate(s)
}
