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
	// The ID of the VPC firewall instance. This parameter is required. If this parameter is not specified, the ErrorMissingFirewallId (-360444, firewall id is required) error is returned.
	//
	// > You can call DescribeTrFirewallsV2List to query existing FirewallIds (in the format of vfw-tr-*). If the firewall has not been created, you must first create an Enterprise Edition transit router in CEN and then call CreateTrFirewallV2.
	//
	// example:
	//
	// vfw-tr-a5a6b89f46764928****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the response message. Valid values:
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
