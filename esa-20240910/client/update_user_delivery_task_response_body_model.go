// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserDeliveryTaskResponseBody
	GetRequestId() *string
}

type UpdateUserDeliveryTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserDeliveryTaskResponseBody) SetRequestId(v string) *UpdateUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
