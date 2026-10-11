// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEgressProxyGatewayRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteEgressProxyGatewayRequest
	GetDryRun() *bool
	SetEgressProxyGatewayId(v string) *DeleteEgressProxyGatewayRequest
	GetEgressProxyGatewayId() *string
}

type DeleteEgressProxyGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value of this parameter on the client side. Make sure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request is unique.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: Sends a check request and does not delete the egress proxy gateway instance. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the check fails, the system returns an error. If the check succeeds, the system returns the error code DryRunOperation.
	//
	// - **false*	- (default): Sends a regular request. If the check succeeds, the system returns an HTTP 2xx status code and performs the operation.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
}

func (s DeleteEgressProxyGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEgressProxyGatewayRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteEgressProxyGatewayRequest) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *DeleteEgressProxyGatewayRequest) SetClientToken(v string) *DeleteEgressProxyGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEgressProxyGatewayRequest) SetDryRun(v bool) *DeleteEgressProxyGatewayRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteEgressProxyGatewayRequest) SetEgressProxyGatewayId(v string) *DeleteEgressProxyGatewayRequest {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *DeleteEgressProxyGatewayRequest) Validate() error {
	return dara.Validate(s)
}
