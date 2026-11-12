// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFirewallRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableFirewallRuleResponseBody
	GetRequestId() *string
}

type DisableFirewallRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableFirewallRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFirewallRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableFirewallRuleResponseBody) SetRequestId(v string) *DisableFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableFirewallRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
