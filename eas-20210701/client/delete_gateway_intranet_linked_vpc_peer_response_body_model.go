// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIntranetLinkedVpcPeerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DeleteGatewayIntranetLinkedVpcPeerResponseBody
	GetGatewayId() *string
	SetMessage(v string) *DeleteGatewayIntranetLinkedVpcPeerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayIntranetLinkedVpcPeerResponseBody
	GetRequestId() *string
}

type DeleteGatewayIntranetLinkedVpcPeerResponseBody struct {
	// The ID of the private gateway.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// Successfully delete intranet linked vpc Peer for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcPeerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcPeerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponseBody) SetGatewayId(v string) *DeleteGatewayIntranetLinkedVpcPeerResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponseBody) SetMessage(v string) *DeleteGatewayIntranetLinkedVpcPeerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponseBody) SetRequestId(v string) *DeleteGatewayIntranetLinkedVpcPeerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcPeerResponseBody) Validate() error {
	return dara.Validate(s)
}
