// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDnsFirewallPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDnsFirewallPolicyResponseBody
	GetRequestId() *string
}

type ModifyDnsFirewallPolicyResponseBody struct {
	// example:
	//
	// 9C50C2A9-4BBB-5504-8ADA-C41A79****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDnsFirewallPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDnsFirewallPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDnsFirewallPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDnsFirewallPolicyResponseBody) SetRequestId(v string) *ModifyDnsFirewallPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDnsFirewallPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
