// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyGatewayId(v string) *CreateEgressProxyGatewayResponseBody
	GetEgressProxyGatewayId() *string
	SetRequestId(v string) *CreateEgressProxyGatewayResponseBody
	GetRequestId() *string
}

type CreateEgressProxyGatewayResponseBody struct {
	// The ID of the egress proxy gateway instance.
	//
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A05AB5-98DC-5175-9BBB-33BA653BC0C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEgressProxyGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyGatewayResponseBody) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *CreateEgressProxyGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEgressProxyGatewayResponseBody) SetEgressProxyGatewayId(v string) *CreateEgressProxyGatewayResponseBody {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *CreateEgressProxyGatewayResponseBody) SetRequestId(v string) *CreateEgressProxyGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEgressProxyGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
