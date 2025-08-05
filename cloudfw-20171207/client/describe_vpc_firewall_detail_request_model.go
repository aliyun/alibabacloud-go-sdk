// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcFirewallDetailRequest
	GetLang() *string
	SetLocalVpcId(v string) *DescribeVpcFirewallDetailRequest
	GetLocalVpcId() *string
	SetPeerVpcId(v string) *DescribeVpcFirewallDetailRequest
	GetPeerVpcId() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallDetailRequest
	GetVpcFirewallId() *string
}

type DescribeVpcFirewallDetailRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the local VPC.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	LocalVpcId *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	// The ID of the peer VPC.
	//
	// example:
	//
	// vpc-90rq0anm6t8vbwbo****
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallList](https://help.aliyun.com/document_detail/342932.html) operation to query the instance IDs of VPC firewalls.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallDetailRequest) GetLocalVpcId() *string {
	return s.LocalVpcId
}

func (s *DescribeVpcFirewallDetailRequest) GetPeerVpcId() *string {
	return s.PeerVpcId
}

func (s *DescribeVpcFirewallDetailRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallDetailRequest) SetLang(v string) *DescribeVpcFirewallDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetLocalVpcId(v string) *DescribeVpcFirewallDetailRequest {
	s.LocalVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetPeerVpcId(v string) *DescribeVpcFirewallDetailRequest {
	s.PeerVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallDetailRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) Validate() error {
	return dara.Validate(s)
}
