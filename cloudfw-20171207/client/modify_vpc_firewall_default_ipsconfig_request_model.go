// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallDefaultIPSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasicRules(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetBasicRules() *string
	SetEnableAllPatch(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetEnableAllPatch() *string
	SetLang(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetLang() *string
	SetMemberUid(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetMemberUid() *string
	SetRuleClass(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetRuleClass() *string
	SetRunMode(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetRunMode() *string
	SetSourceIp(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetSourceIp() *string
	SetVpcFirewallId(v string) *ModifyVpcFirewallDefaultIPSConfigRequest
	GetVpcFirewallId() *string
}

type ModifyVpcFirewallDefaultIPSConfigRequest struct {
	// Specifies whether to enable basic protection. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- **0**: no.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BasicRules *string `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Specifies whether to enable virtual patching. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- **0**: no.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnableAllPatch *string `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The level of the rule group for the IPS. Valid values:
	//
	// 	- **1**: loose
	//
	// 	- **2**: medium
	//
	// 	- **3**: strict
	//
	// example:
	//
	// 1
	RuleClass *string `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The mode of the intrusion prevention system (IPS). Valid values:
	//
	// 	- **1**: block mode.
	//
	// 	- **0**: monitor mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// 	- If the VPC firewall protects traffic between a VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. You can call the [DescribeVpcFirewallCenList](https://help.aliyun.com/document_detail/345777.html) operation to query the IDs of CEN instances.
	//
	// 	- If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall. You can call the [DescribeVpcFirewallList](https://help.aliyun.com/document_detail/342932.html) operation to query the instance IDs of VPC firewalls.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetBasicRules() *string {
	return s.BasicRules
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetEnableAllPatch() *string {
	return s.EnableAllPatch
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetRuleClass() *string {
	return s.RuleClass
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetBasicRules(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.BasicRules = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetEnableAllPatch(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.EnableAllPatch = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetLang(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetMemberUid(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetRuleClass(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.RuleClass = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetRunMode(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.RunMode = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetSourceIp(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) Validate() error {
	return dara.Validate(s)
}
