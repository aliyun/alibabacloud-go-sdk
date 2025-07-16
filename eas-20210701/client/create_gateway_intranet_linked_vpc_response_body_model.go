// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIntranetLinkedVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *CreateGatewayIntranetLinkedVpcResponseBody
	GetGatewayId() *string
	SetMessage(v string) *CreateGatewayIntranetLinkedVpcResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGatewayIntranetLinkedVpcResponseBody
	GetRequestId() *string
}

type CreateGatewayIntranetLinkedVpcResponseBody struct {
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
	// Successfully add intranet linked vpc for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) SetGatewayId(v string) *CreateGatewayIntranetLinkedVpcResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) SetMessage(v string) *CreateGatewayIntranetLinkedVpcResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) SetRequestId(v string) *CreateGatewayIntranetLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
