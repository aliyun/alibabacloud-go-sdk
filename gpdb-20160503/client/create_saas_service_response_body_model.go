// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaasServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateSaasServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateSaasServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CreateSaasServiceResponseBody
	GetServiceId() *string
}

type CreateSaasServiceResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// *********
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2CAD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// agdb-eqxwj5tj5ojx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateSaasServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSaasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSaasServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSaasServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSaasServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateSaasServiceResponseBody) SetOrderId(v string) *CreateSaasServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSaasServiceResponseBody) SetRequestId(v string) *CreateSaasServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSaasServiceResponseBody) SetServiceId(v string) *CreateSaasServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateSaasServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
