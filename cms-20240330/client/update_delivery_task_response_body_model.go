// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDeliveryTaskResponseBody
	GetRequestId() *string
}

type UpdateDeliveryTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9335D7CB-2725-53DD-A2B3-187821C8F1B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeliveryTaskResponseBody) SetRequestId(v string) *UpdateDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
