// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2DetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *DescribeTrFirewallsV2DetailRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeTrFirewallsV2DetailRequest
	GetLang() *string
}

type DescribeTrFirewallsV2DetailRequest struct {
	// The instance ID of the virtual private cloud (VPC) firewall.
	//
	// > FirewallId is required. If this parameter is not specified, ErrorParameters(400) is returned. You can call DescribeTrFirewallsV2List to query existing FirewallId values. If no VPC firewall exists, create a CEN instance and a transit router first, and then call CreateTrFirewallV2 to obtain a FirewallId.
	//
	// example:
	//
	// vfw-tr-a5a6b89f46764928****
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
}

func (s DescribeTrFirewallsV2DetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2DetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2DetailRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallsV2DetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTrFirewallsV2DetailRequest) SetFirewallId(v string) *DescribeTrFirewallsV2DetailRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailRequest) SetLang(v string) *DescribeTrFirewallsV2DetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailRequest) Validate() error {
	return dara.Validate(s)
}
