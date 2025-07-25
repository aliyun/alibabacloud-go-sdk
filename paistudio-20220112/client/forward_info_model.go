// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForwardInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEipAllocationId(v string) *ForwardInfo
	GetEipAllocationId() *string
	SetNatGatewayId(v string) *ForwardInfo
	GetNatGatewayId() *string
}

type ForwardInfo struct {
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	NatGatewayId    *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s ForwardInfo) String() string {
	return dara.Prettify(s)
}

func (s ForwardInfo) GoString() string {
	return s.String()
}

func (s *ForwardInfo) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *ForwardInfo) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ForwardInfo) SetEipAllocationId(v string) *ForwardInfo {
	s.EipAllocationId = &v
	return s
}

func (s *ForwardInfo) SetNatGatewayId(v string) *ForwardInfo {
	s.NatGatewayId = &v
	return s
}

func (s *ForwardInfo) Validate() error {
	return dara.Validate(s)
}
