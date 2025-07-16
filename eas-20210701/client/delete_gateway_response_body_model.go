// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DeleteGatewayResponseBody
	GetGatewayId() *string
	SetMessage(v string) *DeleteGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayResponseBody
	GetRequestId() *string
}

type DeleteGatewayResponseBody struct {
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
	// Gateway is deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeleteGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayResponseBody) SetGatewayId(v string) *DeleteGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetMessage(v string) *DeleteGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
