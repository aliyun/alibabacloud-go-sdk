// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNatFirewallControlPolicyResponseBody
	GetRequestId() *string
}

type ModifyNatFirewallControlPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3868197C-E6E8-52CD-8358-05E3308430E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatFirewallControlPolicyResponseBody) SetRequestId(v string) *ModifyNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
