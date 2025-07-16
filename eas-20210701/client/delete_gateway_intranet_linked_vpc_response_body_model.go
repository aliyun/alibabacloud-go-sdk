// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIntranetLinkedVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DeleteGatewayIntranetLinkedVpcResponseBody
	GetGatewayId() *string
	SetMessage(v string) *DeleteGatewayIntranetLinkedVpcResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayIntranetLinkedVpcResponseBody
	GetRequestId() *string
}

type DeleteGatewayIntranetLinkedVpcResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully delete intranet linked vpc for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) SetGatewayId(v string) *DeleteGatewayIntranetLinkedVpcResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) SetMessage(v string) *DeleteGatewayIntranetLinkedVpcResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) SetRequestId(v string) *DeleteGatewayIntranetLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
