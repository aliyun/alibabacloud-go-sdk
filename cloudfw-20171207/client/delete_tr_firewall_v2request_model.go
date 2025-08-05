// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrFirewallV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *DeleteTrFirewallV2Request
	GetFirewallId() *string
	SetLang(v string) *DeleteTrFirewallV2Request
	GetLang() *string
}

type DeleteTrFirewallV2Request struct {
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-99bc4f0fc88b4d00****
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
}

func (s DeleteTrFirewallV2Request) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrFirewallV2Request) GoString() string {
	return s.String()
}

func (s *DeleteTrFirewallV2Request) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DeleteTrFirewallV2Request) GetLang() *string {
	return s.Lang
}

func (s *DeleteTrFirewallV2Request) SetFirewallId(v string) *DeleteTrFirewallV2Request {
	s.FirewallId = &v
	return s
}

func (s *DeleteTrFirewallV2Request) SetLang(v string) *DeleteTrFirewallV2Request {
	s.Lang = &v
	return s
}

func (s *DeleteTrFirewallV2Request) Validate() error {
	return dara.Validate(s)
}
