// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityProxyResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSecurityProxyResourcesRequest
	GetLang() *string
	SetMemberUid(v int64) *DescribeSecurityProxyResourcesRequest
	GetMemberUid() *int64
	SetNatGatewayId(v string) *DescribeSecurityProxyResourcesRequest
	GetNatGatewayId() *string
}

type DescribeSecurityProxyResourcesRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 136481150091****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// ngw-uf6y16l23fm8hq0****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s DescribeSecurityProxyResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResourcesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSecurityProxyResourcesRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeSecurityProxyResourcesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSecurityProxyResourcesRequest) SetLang(v string) *DescribeSecurityProxyResourcesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecurityProxyResourcesRequest) SetMemberUid(v int64) *DescribeSecurityProxyResourcesRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeSecurityProxyResourcesRequest) SetNatGatewayId(v string) *DescribeSecurityProxyResourcesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSecurityProxyResourcesRequest) Validate() error {
	return dara.Validate(s)
}
