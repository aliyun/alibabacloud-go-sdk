// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecLogDeliveryStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApisecLogDeliveryStatusResponseBody
	GetRequestId() *string
}

type ModifyApisecLogDeliveryStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D****E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApisecLogDeliveryStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecLogDeliveryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApisecLogDeliveryStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApisecLogDeliveryStatusResponseBody) SetRequestId(v string) *ModifyApisecLogDeliveryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApisecLogDeliveryStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
