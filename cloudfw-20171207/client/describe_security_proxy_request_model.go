// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSecurityProxyRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeSecurityProxyRequest
	GetMemberUid() *string
	SetNatGatewayId(v string) *DescribeSecurityProxyRequest
	GetNatGatewayId() *string
	SetPageNo(v string) *DescribeSecurityProxyRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeSecurityProxyRequest
	GetPageSize() *string
	SetProxyId(v string) *DescribeSecurityProxyRequest
	GetProxyId() *string
	SetProxyName(v string) *DescribeSecurityProxyRequest
	GetProxyName() *string
	SetRegionNo(v string) *DescribeSecurityProxyRequest
	GetRegionNo() *string
	SetStatus(v string) *DescribeSecurityProxyRequest
	GetStatus() *string
	SetVpcId(v string) *DescribeSecurityProxyRequest
	GetVpcId() *string
}

type DescribeSecurityProxyRequest struct {
	// The language type for the request and response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member accounts of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-zm0h3c1exm5bifuorg8c5
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of entries per page in a paged query. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the NAT firewall.
	//
	// example:
	//
	// proxy-nat80d763eb0dee4eacaec9
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The name of the NAT firewall. The name can contain uppercase and lowercase letters, Chinese characters, digits, and underscores (_). The name must be 4 to 50 characters in length and cannot start with an underscore.
	//
	// example:
	//
	// nat-idmp-fir
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The status of Cloud Firewall. Valid values:
	//
	// - **configuring**: Being created.
	//
	// - **deleting**: Being deleted.
	//
	// - **normal**: Normal.
	//
	// - **abnormal**: Abnormal.
	//
	// - **opening**: Being enabled.
	//
	// - **closing**: Being disabled.
	//
	// - **closed**: Disabled.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VPC-connected instance ID.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeSecurityProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSecurityProxyRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeSecurityProxyRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSecurityProxyRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeSecurityProxyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSecurityProxyRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeSecurityProxyRequest) GetProxyName() *string {
	return s.ProxyName
}

func (s *DescribeSecurityProxyRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeSecurityProxyRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSecurityProxyRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSecurityProxyRequest) SetLang(v string) *DescribeSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetMemberUid(v string) *DescribeSecurityProxyRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetNatGatewayId(v string) *DescribeSecurityProxyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetPageNo(v string) *DescribeSecurityProxyRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetPageSize(v string) *DescribeSecurityProxyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetProxyId(v string) *DescribeSecurityProxyRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetProxyName(v string) *DescribeSecurityProxyRequest {
	s.ProxyName = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetRegionNo(v string) *DescribeSecurityProxyRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetStatus(v string) *DescribeSecurityProxyRequest {
	s.Status = &v
	return s
}

func (s *DescribeSecurityProxyRequest) SetVpcId(v string) *DescribeSecurityProxyRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeSecurityProxyRequest) Validate() error {
	return dara.Validate(s)
}
