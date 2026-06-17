// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallControlPolicyPositionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *ModifyVpcFirewallControlPolicyPositionRequest
	GetAclUuid() *string
	SetLang(v string) *ModifyVpcFirewallControlPolicyPositionRequest
	GetLang() *string
	SetNewOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest
	GetNewOrder() *string
	SetOldOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest
	GetOldOrder() *string
	SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyPositionRequest
	GetVpcFirewallId() *string
}

type ModifyVpcFirewallControlPolicyPositionRequest struct {
	// The unique ID of the access control policy.
	//
	// To modify an access control policy, provide the unique ID of the policy. Call the [DescribeVpcFirewallControlPolicy](https://help.aliyun.com/document_detail/159758.html) operation to obtain the ID.
	//
	// example:
	//
	// 2746d9ff-5d7c-449d-a2a9-ccaa15fe****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The language of the request and response.
	//
	// Valid values:
	//
	// - **zh*	- (Default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new priority of the access control policy.
	//
	// > For more information about the valid range of priorities, see [DescribePolicyPriorities](https://help.aliyun.com/document_detail/474145.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The original priority of the access control policy.
	//
	// > This parameter is deprecated. Use the AclUuid parameter to specify the policy to modify.
	//
	// example:
	//
	// 5
	OldOrder *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	// The ID of the policy group for the VPC firewall. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to obtain the ID.
	//
	// Valid values:
	//
	// - If the VPC firewall protects a Cloud Enterprise Network (CEN) instance, the ID of the policy group is the ID of the CEN instance.
	//
	//   Example: cen-ervw0g12b5jbw\\*\\*\\*\\*
	//
	// - If the VPC firewall protects an Express Connect circuit, the ID of the policy group is the ID of the VPC firewall instance.
	//
	//   Example: vfw-a42bbb7b887148c9\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) GetNewOrder() *string {
	return s.NewOrder
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) GetOldOrder() *string {
	return s.OldOrder
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetAclUuid(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetNewOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetOldOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.OldOrder = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) Validate() error {
	return dara.Validate(s)
}
