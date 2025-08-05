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
	// The UUID of the access control policy.
	//
	// If you want to modify the configurations of an access control policy, you must provide the UUID of the policy. You can call the [DescribeVpcFirewallControlPolicy](https://help.aliyun.com/document_detail/159758.html) operation to query the UUID.
	//
	// example:
	//
	// 2746d9ff-5d7c-449d-a2a9-ccaa15fe****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The language of the content within the request and the response.
	//
	// Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new priority of the access control policy.
	//
	// >  For more information about the valid values of the new priority, see [DescribeVpcFirewallPolicyPriorUsed](https://help.aliyun.com/document_detail/474145.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The original priority of the access control policy.
	//
	// > This parameter is not recommended. We recommend that you use the AclUuid parameter to specify the policy that you want to modify.
	//
	// example:
	//
	// 5
	OldOrder *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	// The ID of the group to which the access control policy belongs. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// Valid values:
	//
	// 	- If the VPC firewall is used to protect a CEN instance, the value of this parameter must be the ID of the CEN instance.
	//
	//     Example: cen-ervw0g12b5jbw\\*\\*\\*\\*
	//
	// 	- If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	//     Example: vfw-a42bbb7b887148c9\\*\\*\\*\\*
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
