// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallConfigureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyVpcFirewallConfigureRequest
	GetLang() *string
	SetLocalVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest
	GetLocalVpcCidrTableList() *string
	SetMemberUid(v string) *ModifyVpcFirewallConfigureRequest
	GetMemberUid() *string
	SetPeerVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest
	GetPeerVpcCidrTableList() *string
	SetVpcFirewallId(v string) *ModifyVpcFirewallConfigureRequest
	GetVpcFirewallId() *string
	SetVpcFirewallName(v string) *ModifyVpcFirewallConfigureRequest
	GetVpcFirewallName() *string
}

type ModifyVpcFirewallConfigureRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The CIDR block information of the local VPC, in JSON format. The following parameters are included:
	//
	// - **RouteTableId**: the routing table ID of the local VPC.
	//
	// - **RouteEntryList**: specified in JSON format and contains DestinationCidr (the destination CIDR block of the local VPC) and NextHopInstanceId (the next hop instance ID of the local VPC).
	//
	// > You can invoke the [DescribeVpcFirewallDetail](https://help.aliyun.com/document_detail/342892.html) operation to query the CIDR block information of the local VPC for the virtual private cloud (VPC) firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"RouteTableId":"vtb-1234","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]},{"RouteTableId":"vtb-1235","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]}]
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	// The UID of the member accounts of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The CIDR block information of the peer VPC, in JSON format. The following parameters are included:
	//
	// - **RouteTableId**: the routing table ID of the peer VPC.
	//
	// - **RouteEntryList**: specified in JSON format and contains DestinationCidr (the destination CIDR block of the peer VPC) and NextHopInstanceId (the next hop instance ID of the peer VPC).
	//
	// > You can invoke the [DescribeVpcFirewallDetail](https://help.aliyun.com/document_detail/342892.html) operation to query the CIDR block information of the peer VPC for the virtual private cloud (VPC) firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"RouteTableId":"vtb-1234","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]},{"RouteTableId":"vtb-1235","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]}]
	PeerVpcCidrTableList *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	// The instance ID of the virtual private cloud (VPC) firewall.
	//
	// > You can invoke the [DescribeVpcFirewallList](https://help.aliyun.com/document_detail/342932.html) operation to query the instance ID of the virtual private cloud (VPC) firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the virtual private cloud (VPC) firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-vpc-firewall
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s ModifyVpcFirewallConfigureRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyVpcFirewallConfigureRequest) GetLocalVpcCidrTableList() *string {
	return s.LocalVpcCidrTableList
}

func (s *ModifyVpcFirewallConfigureRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *ModifyVpcFirewallConfigureRequest) GetPeerVpcCidrTableList() *string {
	return s.PeerVpcCidrTableList
}

func (s *ModifyVpcFirewallConfigureRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *ModifyVpcFirewallConfigureRequest) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *ModifyVpcFirewallConfigureRequest) SetLang(v string) *ModifyVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetLocalVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest {
	s.LocalVpcCidrTableList = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetMemberUid(v string) *ModifyVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetPeerVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest {
	s.PeerVpcCidrTableList = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallConfigureRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetVpcFirewallName(v string) *ModifyVpcFirewallConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) Validate() error {
	return dara.Validate(s)
}
