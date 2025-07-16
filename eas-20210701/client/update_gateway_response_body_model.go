// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *UpdateGatewayResponseBody
	GetGatewayId() *string
	SetMessage(v string) *UpdateGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayResponseBody
	GetRequestId() *string
}

type UpdateGatewayResponseBody struct {
	// The ID of the gateway.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Gateway is updated
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayResponseBody) SetGatewayId(v string) *UpdateGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetMessage(v string) *UpdateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetRequestId(v string) *UpdateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
