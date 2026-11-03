// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCopyVpcFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *BatchCopyVpcFirewallControlPolicyRequest
	GetLang() *string
	SetSourceIp(v string) *BatchCopyVpcFirewallControlPolicyRequest
	GetSourceIp() *string
	SetSourceVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest
	GetSourceVpcFirewallId() *string
	SetTargetVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest
	GetTargetVpcFirewallId() *string
}

type BatchCopyVpcFirewallControlPolicyRequest struct {
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the access control policy group for the source virtual private cloud (VPC) border firewall. Valid values:
	//
	// - If the VPC border firewall protects traffic between network instances (including VPCs, VBRs, and CCNs) in a Cloud Enterprise Network (CEN) instance and a specified VPC, set the value to the ID of the CEN instance.
	//
	// - If the VPC border firewall protects traffic between two VPCs that are connected by using Express Connect, set the value to the ID of the VPC border firewall instance.
	//
	// > You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	SourceVpcFirewallId *string `json:"SourceVpcFirewallId,omitempty" xml:"SourceVpcFirewallId,omitempty"`
	// The ID of the access control policy group for the destination virtual private cloud (VPC) border firewall. Valid values:
	//
	// - If the VPC border firewall protects traffic between network instances (including VPCs, VBRs, and CCNs) in a CEN instance and a specified VPC, set the value to the ID of the CEN instance.
	//
	// - If the VPC border firewall protects traffic between two VPCs that are connected by using Express Connect, set the value to the ID of the VPC border firewall instance.
	//
	// > You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-e37d3a04cf79446a****
	TargetVpcFirewallId *string `json:"TargetVpcFirewallId,omitempty" xml:"TargetVpcFirewallId,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) GetSourceVpcFirewallId() *string {
	return s.SourceVpcFirewallId
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) GetTargetVpcFirewallId() *string {
	return s.TargetVpcFirewallId
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetLang(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetSourceIp(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetSourceVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.SourceVpcFirewallId = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetTargetVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.TargetVpcFirewallId = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
