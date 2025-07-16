// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIntranetLinkedVpcPeerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *CreateGatewayIntranetLinkedVpcPeerResponseBody
	GetGatewayId() *string
	SetMessage(v string) *CreateGatewayIntranetLinkedVpcPeerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGatewayIntranetLinkedVpcPeerResponseBody
	GetRequestId() *string
}

type CreateGatewayIntranetLinkedVpcPeerResponseBody struct {
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
	// Successfully add intranet linked vpc Peer for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcPeerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcPeerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponseBody) SetGatewayId(v string) *CreateGatewayIntranetLinkedVpcPeerResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponseBody) SetMessage(v string) *CreateGatewayIntranetLinkedVpcPeerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponseBody) SetRequestId(v string) *CreateGatewayIntranetLinkedVpcPeerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcPeerResponseBody) Validate() error {
	return dara.Validate(s)
}
