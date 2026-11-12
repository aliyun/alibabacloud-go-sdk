// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFirewallRuleResponseBody
	GetRequestId() *string
}

type ModifyFirewallRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E1FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFirewallRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFirewallRuleResponseBody) SetRequestId(v string) *ModifyFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFirewallRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
