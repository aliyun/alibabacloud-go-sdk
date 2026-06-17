// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *CreateNatFirewallControlPolicyResponseBody
	GetAclUuid() *string
	SetRequestId(v string) *CreateNatFirewallControlPolicyResponseBody
	GetRequestId() *string
}

type CreateNatFirewallControlPolicyResponseBody struct {
	// The unique ID of the access control policy.
	//
	// > To modify an access control policy, you must provide its unique ID. You can call the `DescribeNatFirewallControlPolicy` operation to obtain this ID.
	//
	// example:
	//
	// 6504d2fb-ab36-49c3-92a6-*****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0DC783F1-B3A7-578D-8A63-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallControlPolicyResponseBody) GetAclUuid() *string {
	return s.AclUuid
}

func (s *CreateNatFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatFirewallControlPolicyResponseBody) SetAclUuid(v string) *CreateNatFirewallControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *CreateNatFirewallControlPolicyResponseBody) SetRequestId(v string) *CreateNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatFirewallControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
