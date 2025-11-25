// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterResourcesListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeTransitRouterResourcesListRequest
	GetCenId() *string
	SetFirewallId(v string) *DescribeTransitRouterResourcesListRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeTransitRouterResourcesListRequest
	GetLang() *string
	SetRegionNo(v string) *DescribeTransitRouterResourcesListRequest
	GetRegionNo() *string
	SetResourceType(v string) *DescribeTransitRouterResourcesListRequest
	GetResourceType() *string
	SetTransitRouterId(v string) *DescribeTransitRouterResourcesListRequest
	GetTransitRouterId() *string
	SetVpcId(v string) *DescribeTransitRouterResourcesListRequest
	GetVpcId() *string
}

type DescribeTransitRouterResourcesListRequest struct {
	// example:
	//
	// cen-djz1i6p8shzioz****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// vfw-tr-741de4c8956341****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// TR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// tr-2zefgvkcl2qcexbb7****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// example:
	//
	// vpc-wz9lllsbftdm0svpj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeTransitRouterResourcesListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterResourcesListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterResourcesListRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeTransitRouterResourcesListRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTransitRouterResourcesListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTransitRouterResourcesListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeTransitRouterResourcesListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTransitRouterResourcesListRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeTransitRouterResourcesListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTransitRouterResourcesListRequest) SetCenId(v string) *DescribeTransitRouterResourcesListRequest {
	s.CenId = &v
	return s
}

func (s *DescribeTransitRouterResourcesListRequest) SetFirewallId(v string) *DescribeTransitRouterResourcesListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTransitRouterResourcesListRequest) SetLang(v string) *DescribeTransitRouterResourcesListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTransitRouterResourcesListRequest) SetRegionNo(v string) *DescribeTransitRouterResourcesListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeTransitRouterResourcesListRequest) SetResourceType(v string) *DescribeTransitRouterResourcesListRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTransitRouterResourcesListRequest) SetTransitRouterId(v string) *DescribeTransitRouterResourcesListRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTransitRouterResourcesListRequest) SetVpcId(v string) *DescribeTransitRouterResourcesListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeTransitRouterResourcesListRequest) Validate() error {
	return dara.Validate(s)
}
