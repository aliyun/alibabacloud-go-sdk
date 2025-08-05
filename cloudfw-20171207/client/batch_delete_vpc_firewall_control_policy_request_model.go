// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteVpcFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuidList(v []*string) *BatchDeleteVpcFirewallControlPolicyRequest
	GetAclUuidList() []*string
	SetVpcFirewallId(v string) *BatchDeleteVpcFirewallControlPolicyRequest
	GetVpcFirewallId() *string
}

type BatchDeleteVpcFirewallControlPolicyRequest struct {
	// The UUIDs of access control policies.
	//
	// This parameter is required.
	AclUuidList []*string `json:"AclUuidList,omitempty" xml:"AclUuidList,omitempty" type:"Repeated"`
	// The instance ID of the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s BatchDeleteVpcFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteVpcFirewallControlPolicyRequest) GetAclUuidList() []*string {
	return s.AclUuidList
}

func (s *BatchDeleteVpcFirewallControlPolicyRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *BatchDeleteVpcFirewallControlPolicyRequest) SetAclUuidList(v []*string) *BatchDeleteVpcFirewallControlPolicyRequest {
	s.AclUuidList = v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *BatchDeleteVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
