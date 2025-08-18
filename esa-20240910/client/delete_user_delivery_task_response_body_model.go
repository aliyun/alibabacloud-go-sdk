// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserDeliveryTaskResponseBody
	GetRequestId() *string
}

type DeleteUserDeliveryTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 952ea16b-1f05-4a76-bb32-420282d8aeb9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserDeliveryTaskResponseBody) SetRequestId(v string) *DeleteUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
