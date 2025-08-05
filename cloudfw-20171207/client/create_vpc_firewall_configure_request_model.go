// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallConfigureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallSwitch(v string) *CreateVpcFirewallConfigureRequest
	GetFirewallSwitch() *string
	SetLang(v string) *CreateVpcFirewallConfigureRequest
	GetLang() *string
	SetLocalVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest
	GetLocalVpcCidrTableList() *string
	SetLocalVpcId(v string) *CreateVpcFirewallConfigureRequest
	GetLocalVpcId() *string
	SetLocalVpcRegion(v string) *CreateVpcFirewallConfigureRequest
	GetLocalVpcRegion() *string
	SetMemberUid(v string) *CreateVpcFirewallConfigureRequest
	GetMemberUid() *string
	SetPeerVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest
	GetPeerVpcCidrTableList() *string
	SetPeerVpcId(v string) *CreateVpcFirewallConfigureRequest
	GetPeerVpcId() *string
	SetPeerVpcRegion(v string) *CreateVpcFirewallConfigureRequest
	GetPeerVpcRegion() *string
	SetVpcFirewallName(v string) *CreateVpcFirewallConfigureRequest
	GetVpcFirewallName() *string
}

type CreateVpcFirewallConfigureRequest struct {
	// The status of the VPC firewall after you create the firewall. Valid values:
	//
	// 	- **open**: After you create the VPC firewall, the VPC firewall is automatically enabled. This is the default value.
	//
	// 	- **close**: After you create the VPC firewall, the VPC firewall is disabled. To enable the firewall, you can call the [ModifyVpcFirewallSwitchStatus](https://help.aliyun.com/document_detail/342935.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The CIDR blocks of the local VPC. The value is a JSON string that contains the following parameters:
	//
	// 	- **RouteTableId**: the ID of the route table for the local VPC.
	//
	// 	- **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the local VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the local VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"RouteTableId":"vtb-1234","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]},{"RouteTableId":"vtb-1235","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]}]
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	// The ID of the local VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	LocalVpcId *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	// The region ID of the local VPC.
	//
	// >  For more information about the regions in which Cloud Firewall is available, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	LocalVpcRegion *string `json:"LocalVpcRegion,omitempty" xml:"LocalVpcRegion,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The CIDR blocks of the peer VPC. The value is a JSON string that contains the following parameters:
	//
	// 	- **RouteTableId**: the ID of the route table for the peer VPC.
	//
	// 	- **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the peer VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the peer VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"RouteTableId":"vtb-1234","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]},{"RouteTableId":"vtb-1235","RouteEntryList":[{"DestinationCidr":"192.168.XX.XX/24","NextHopInstanceId":"vrt-m5eb5me6c3l5sezae****"}]}]
	PeerVpcCidrTableList *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	// The ID of the peer VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-wb8vbo90rq0anm6t****
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The region ID of the peer VPC.
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	PeerVpcRegion *string `json:"PeerVpcRegion,omitempty" xml:"PeerVpcRegion,omitempty"`
	// The instance name of the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-vpc-firewall
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s CreateVpcFirewallConfigureRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureRequest) GetFirewallSwitch() *string {
	return s.FirewallSwitch
}

func (s *CreateVpcFirewallConfigureRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateVpcFirewallConfigureRequest) GetLocalVpcCidrTableList() *string {
	return s.LocalVpcCidrTableList
}

func (s *CreateVpcFirewallConfigureRequest) GetLocalVpcId() *string {
	return s.LocalVpcId
}

func (s *CreateVpcFirewallConfigureRequest) GetLocalVpcRegion() *string {
	return s.LocalVpcRegion
}

func (s *CreateVpcFirewallConfigureRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *CreateVpcFirewallConfigureRequest) GetPeerVpcCidrTableList() *string {
	return s.PeerVpcCidrTableList
}

func (s *CreateVpcFirewallConfigureRequest) GetPeerVpcId() *string {
	return s.PeerVpcId
}

func (s *CreateVpcFirewallConfigureRequest) GetPeerVpcRegion() *string {
	return s.PeerVpcRegion
}

func (s *CreateVpcFirewallConfigureRequest) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *CreateVpcFirewallConfigureRequest) SetFirewallSwitch(v string) *CreateVpcFirewallConfigureRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLang(v string) *CreateVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcCidrTableList = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcId(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcId = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcRegion(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcRegion = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetMemberUid(v string) *CreateVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcCidrTableList = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcId(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcId = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcRegion(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcRegion = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) Validate() error {
	return dara.Validate(s)
}
