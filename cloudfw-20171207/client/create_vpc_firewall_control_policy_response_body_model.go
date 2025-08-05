// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *CreateVpcFirewallControlPolicyResponseBody
	GetAclUuid() *string
	SetRequestId(v string) *CreateVpcFirewallControlPolicyResponseBody
	GetRequestId() *string
}

type CreateVpcFirewallControlPolicyResponseBody struct {
	// The ID of the access control policy.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221ad84c
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpcFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyResponseBody) GetAclUuid() *string {
	return s.AclUuid
}

func (s *CreateVpcFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetAclUuid(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
