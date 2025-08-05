// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrFirewallV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *CreateTrFirewallV2ResponseBody
	GetFirewallId() *string
	SetRequestId(v string) *CreateTrFirewallV2ResponseBody
	GetRequestId() *string
}

type CreateTrFirewallV2ResponseBody struct {
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-37e22bf0d9b34870****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 822B9125-6E1A-551C-8EAF-6E7AE7444B00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTrFirewallV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2ResponseBody) GetFirewallId() *string {
	return s.FirewallId
}

func (s *CreateTrFirewallV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrFirewallV2ResponseBody) SetFirewallId(v string) *CreateTrFirewallV2ResponseBody {
	s.FirewallId = &v
	return s
}

func (s *CreateTrFirewallV2ResponseBody) SetRequestId(v string) *CreateTrFirewallV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrFirewallV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
