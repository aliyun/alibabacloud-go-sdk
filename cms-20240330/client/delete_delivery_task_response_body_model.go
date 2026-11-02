// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeliveryTaskResponseBody
	GetRequestId() *string
}

type DeleteDeliveryTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeliveryTaskResponseBody) SetRequestId(v string) *DeleteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
