// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallV2RoutePolicyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeTrFirewallV2RoutePolicyListRequest
	GetCurrentPage() *int32
	SetFirewallId(v string) *DescribeTrFirewallV2RoutePolicyListRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeTrFirewallV2RoutePolicyListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeTrFirewallV2RoutePolicyListRequest
	GetPageSize() *int32
	SetPolicyId(v string) *DescribeTrFirewallV2RoutePolicyListRequest
	GetPolicyId() *string
}

type DescribeTrFirewallV2RoutePolicyListRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-d5ba592ac6c84aff****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// policy-7b66257c14e141fb****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetCurrentPage(v int32) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetFirewallId(v string) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetLang(v string) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetPageSize(v int32) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetPolicyId(v string) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) Validate() error {
	return dara.Validate(s)
}
