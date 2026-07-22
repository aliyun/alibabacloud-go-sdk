// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2RouteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeTrFirewallsV2RouteListRequest
	GetCurrentPage() *string
	SetFirewallId(v string) *DescribeTrFirewallsV2RouteListRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeTrFirewallsV2RouteListRequest
	GetLang() *string
	SetPageSize(v string) *DescribeTrFirewallsV2RouteListRequest
	GetPageSize() *string
	SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallsV2RouteListRequest
	GetTrFirewallRoutePolicyId() *string
}

type DescribeTrFirewallsV2RouteListRequest struct {
	// The page number in a paging query. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the virtual private cloud (VPC) firewalls.
	//
	// > FirewallId is required. If it is not specified, ErrorParameters(400) is returned. TrFirewallRoutePolicyId must be used together with FirewallId. If TrFirewallRoutePolicyId is specified without FirewallId, ErrorParameters(400) is returned. If the value is invalid or does not exist, ErrorTrFirewallPolicyNotFound is returned.
	//
	// example:
	//
	// vfw-tr-8bcfa0f2f12d411e****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page in a paging query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the firewall routing policy.
	//
	// example:
	//
	// policy-7d5c672e37ee4175****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeTrFirewallsV2RouteListRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallsV2RouteListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTrFirewallsV2RouteListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeTrFirewallsV2RouteListRequest) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetCurrentPage(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetFirewallId(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetLang(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetPageSize(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) Validate() error {
	return dara.Validate(s)
}
