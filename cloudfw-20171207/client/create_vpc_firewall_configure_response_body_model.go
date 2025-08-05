// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallConfigureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVpcFirewallConfigureResponseBody
	GetRequestId() *string
	SetVpcFirewallId(v string) *CreateVpcFirewallConfigureResponseBody
	GetVpcFirewallId() *string
}

type CreateVpcFirewallConfigureResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125h4j6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallConfigureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcFirewallConfigureResponseBody) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *CreateVpcFirewallConfigureResponseBody) SetRequestId(v string) *CreateVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallConfigureResponseBody) SetVpcFirewallId(v string) *CreateVpcFirewallConfigureResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *CreateVpcFirewallConfigureResponseBody) Validate() error {
	return dara.Validate(s)
}
