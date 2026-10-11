// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyGatewayId(v string) *GetEgressProxyGatewayAttributeRequest
	GetEgressProxyGatewayId() *string
}

type GetEgressProxyGatewayAttributeRequest struct {
	// The ID of the egress proxy gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
}

func (s GetEgressProxyGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetEgressProxyGatewayAttributeRequest) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *GetEgressProxyGatewayAttributeRequest) SetEgressProxyGatewayId(v string) *GetEgressProxyGatewayAttributeRequest {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
