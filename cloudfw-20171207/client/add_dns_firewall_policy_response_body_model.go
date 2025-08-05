// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsFirewallPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *AddDnsFirewallPolicyResponseBody
	GetAclUuid() *string
	SetRequestId(v string) *AddDnsFirewallPolicyResponseBody
	GetRequestId() *string
}

type AddDnsFirewallPolicyResponseBody struct {
	// example:
	//
	// f88dae6f-XXX-XXX-613de9ab2be8
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// example:
	//
	// 71209DFE-XXX-XXX-52B4A4E9DA3B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDnsFirewallPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDnsFirewallPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsFirewallPolicyResponseBody) GetAclUuid() *string {
	return s.AclUuid
}

func (s *AddDnsFirewallPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDnsFirewallPolicyResponseBody) SetAclUuid(v string) *AddDnsFirewallPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *AddDnsFirewallPolicyResponseBody) SetRequestId(v string) *AddDnsFirewallPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsFirewallPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
