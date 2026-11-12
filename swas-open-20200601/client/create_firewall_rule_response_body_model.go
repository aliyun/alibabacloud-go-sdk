// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *CreateFirewallRuleResponseBody
	GetFirewallId() *string
	SetRequestId(v string) *CreateFirewallRuleResponseBody
	GetRequestId() *string
}

type CreateFirewallRuleResponseBody struct {
	// The ID of the firewall rule.
	//
	// example:
	//
	// 8007e18c61024aafbd776d52d0****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleResponseBody) GetFirewallId() *string {
	return s.FirewallId
}

func (s *CreateFirewallRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFirewallRuleResponseBody) SetFirewallId(v string) *CreateFirewallRuleResponseBody {
	s.FirewallId = &v
	return s
}

func (s *CreateFirewallRuleResponseBody) SetRequestId(v string) *CreateFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFirewallRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
