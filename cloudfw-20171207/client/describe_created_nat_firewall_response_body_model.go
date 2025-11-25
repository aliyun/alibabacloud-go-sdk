// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreatedNatFirewallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedNatFirewalls(v []*DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) *DescribeCreatedNatFirewallResponseBody
	GetCreatedNatFirewalls() []*DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls
	SetRequestId(v string) *DescribeCreatedNatFirewallResponseBody
	GetRequestId() *string
}

type DescribeCreatedNatFirewallResponseBody struct {
	CreatedNatFirewalls []*DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls `json:"CreatedNatFirewalls,omitempty" xml:"CreatedNatFirewalls,omitempty" type:"Repeated"`
	// example:
	//
	// 072B5287-8A85-529E-BD47-F8AC2DB1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCreatedNatFirewallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatedNatFirewallResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreatedNatFirewallResponseBody) GetCreatedNatFirewalls() []*DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls {
	return s.CreatedNatFirewalls
}

func (s *DescribeCreatedNatFirewallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreatedNatFirewallResponseBody) SetCreatedNatFirewalls(v []*DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) *DescribeCreatedNatFirewallResponseBody {
	s.CreatedNatFirewalls = v
	return s
}

func (s *DescribeCreatedNatFirewallResponseBody) SetRequestId(v string) *DescribeCreatedNatFirewallResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreatedNatFirewallResponseBody) Validate() error {
	if s.CreatedNatFirewalls != nil {
		for _, item := range s.CreatedNatFirewalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls struct {
	// example:
	//
	// cfw-adk2ad45sf4t8****
	NatFirewallId *string `json:"NatFirewallId,omitempty" xml:"NatFirewallId,omitempty"`
	// example:
	//
	// ngw-uf6i0zkjtz4t2sttf****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// ngw-text
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
}

func (s DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) GetNatFirewallId() *string {
	return s.NatFirewallId
}

func (s *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) SetNatFirewallId(v string) *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls {
	s.NatFirewallId = &v
	return s
}

func (s *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) SetNatGatewayId(v string) *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) SetNatGatewayName(v string) *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeCreatedNatFirewallResponseBodyCreatedNatFirewalls) Validate() error {
	return dara.Validate(s)
}
