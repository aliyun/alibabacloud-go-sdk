// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEgressProxyGatewayResponseBody
	GetRequestId() *string
}

type DeleteEgressProxyGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 01A05AB5-5BED-56ED-8EE9-304FBC2F8537
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEgressProxyGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEgressProxyGatewayResponseBody) SetRequestId(v string) *DeleteEgressProxyGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEgressProxyGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
