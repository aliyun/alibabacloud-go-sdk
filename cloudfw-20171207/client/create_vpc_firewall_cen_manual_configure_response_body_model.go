// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallCenManualConfigureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVpcFirewallCenManualConfigureResponseBody
	GetRequestId() *string
	SetVpcFirewallId(v string) *CreateVpcFirewallCenManualConfigureResponseBody
	GetVpcFirewallId() *string
}

type CreateVpcFirewallCenManualConfigureResponseBody struct {
	// example:
	//
	// B14757D0-4640-4B44-AC67-7F558F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallCenManualConfigureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallCenManualConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenManualConfigureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcFirewallCenManualConfigureResponseBody) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *CreateVpcFirewallCenManualConfigureResponseBody) SetRequestId(v string) *CreateVpcFirewallCenManualConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureResponseBody) SetVpcFirewallId(v string) *CreateVpcFirewallCenManualConfigureResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureResponseBody) Validate() error {
	return dara.Validate(s)
}
