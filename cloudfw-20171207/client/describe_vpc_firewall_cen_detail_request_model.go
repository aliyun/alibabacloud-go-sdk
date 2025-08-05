// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcFirewallCenDetailRequest
	GetLang() *string
	SetNetworkInstanceId(v string) *DescribeVpcFirewallCenDetailRequest
	GetNetworkInstanceId() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailRequest
	GetVpcFirewallId() *string
}

type DescribeVpcFirewallCenDetailRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the VPC for which the VPC firewall is created.
	//
	// example:
	//
	// vpc-2zefk9fbn8j7v585g****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallCenList](https://help.aliyun.com/document_detail/345777.html) operation to query the instance IDs of VPC firewalls.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallCenDetailRequest) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeVpcFirewallCenDetailRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallCenDetailRequest) SetLang(v string) *DescribeVpcFirewallCenDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailRequest) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenDetailRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailRequest) Validate() error {
	return dara.Validate(s)
}
