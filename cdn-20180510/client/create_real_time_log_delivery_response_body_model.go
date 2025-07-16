// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealTimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRealTimeLogDeliveryResponseBody
	GetRequestId() *string
}

type CreateRealTimeLogDeliveryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRealTimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRealTimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRealTimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRealTimeLogDeliveryResponseBody) SetRequestId(v string) *CreateRealTimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRealTimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
