// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *CreateGatewayResponseBody
	GetGwClusterId() *string
	SetOrderId(v string) *CreateGatewayResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateGatewayResponseBody
	GetRequestId() *string
}

type CreateGatewayResponseBody struct {
	// example:
	//
	// pg-2ze2079ueg20****
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 2035629******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F9F1CB1A-B1D5-4EF5-A53A-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateGatewayResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayResponseBody) SetGwClusterId(v string) *CreateGatewayResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetOrderId(v string) *CreateGatewayResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
