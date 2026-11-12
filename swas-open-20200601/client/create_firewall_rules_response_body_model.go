// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallRuleIds(v []*string) *CreateFirewallRulesResponseBody
	GetFirewallRuleIds() []*string
	SetRequestId(v string) *CreateFirewallRulesResponseBody
	GetRequestId() *string
}

type CreateFirewallRulesResponseBody struct {
	// The list of firewall rule IDs.
	FirewallRuleIds []*string `json:"FirewallRuleIds,omitempty" xml:"FirewallRuleIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesResponseBody) GetFirewallRuleIds() []*string {
	return s.FirewallRuleIds
}

func (s *CreateFirewallRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFirewallRulesResponseBody) SetFirewallRuleIds(v []*string) *CreateFirewallRulesResponseBody {
	s.FirewallRuleIds = v
	return s
}

func (s *CreateFirewallRulesResponseBody) SetRequestId(v string) *CreateFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFirewallRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
