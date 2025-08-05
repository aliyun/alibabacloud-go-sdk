// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallCenConfigureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVpcFirewallCenConfigureResponseBody
	GetRequestId() *string
	SetVpcFirewallId(v string) *CreateVpcFirewallCenConfigureResponseBody
	GetVpcFirewallId() *string
}

type CreateVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
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

func (s CreateVpcFirewallCenConfigureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcFirewallCenConfigureResponseBody) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *CreateVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *CreateVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponseBody) SetVpcFirewallId(v string) *CreateVpcFirewallCenConfigureResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponseBody) Validate() error {
	return dara.Validate(s)
}
