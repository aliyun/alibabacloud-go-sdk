// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateGatewayResponseBody
	GetClusterId() *string
	SetGatewayId(v string) *CreateGatewayResponseBody
	GetGatewayId() *string
	SetMessage(v string) *CreateGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGatewayResponseBody
	GetRequestId() *string
}

type CreateGatewayResponseBody struct {
	// The region ID of the private gateway.
	//
	// example:
	//
	// cn-hangzhou
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// Successfully create gateway.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateGatewayResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayResponseBody) SetClusterId(v string) *CreateGatewayResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetGatewayId(v string) *CreateGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetMessage(v string) *CreateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
