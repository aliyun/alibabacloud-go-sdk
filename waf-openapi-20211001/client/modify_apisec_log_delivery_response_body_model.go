// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApisecLogDeliveryResponseBody
	GetRequestId() *string
}

type ModifyApisecLogDeliveryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D****E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApisecLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApisecLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApisecLogDeliveryResponseBody) SetRequestId(v string) *ModifyApisecLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApisecLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
