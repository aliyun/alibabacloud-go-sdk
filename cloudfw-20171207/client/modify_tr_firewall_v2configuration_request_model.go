// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrFirewallV2ConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *ModifyTrFirewallV2ConfigurationRequest
	GetFirewallId() *string
	SetFirewallName(v string) *ModifyTrFirewallV2ConfigurationRequest
	GetFirewallName() *string
	SetLang(v string) *ModifyTrFirewallV2ConfigurationRequest
	GetLang() *string
}

type ModifyTrFirewallV2ConfigurationRequest struct {
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-bcdf89d405ce4bd2****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The instance name of the VPC firewall.
	//
	// example:
	//
	// Test instance
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ModifyTrFirewallV2ConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2ConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2ConfigurationRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *ModifyTrFirewallV2ConfigurationRequest) GetFirewallName() *string {
	return s.FirewallName
}

func (s *ModifyTrFirewallV2ConfigurationRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyTrFirewallV2ConfigurationRequest) SetFirewallId(v string) *ModifyTrFirewallV2ConfigurationRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationRequest) SetFirewallName(v string) *ModifyTrFirewallV2ConfigurationRequest {
	s.FirewallName = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationRequest) SetLang(v string) *ModifyTrFirewallV2ConfigurationRequest {
	s.Lang = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
