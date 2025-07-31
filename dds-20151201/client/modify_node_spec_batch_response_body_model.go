// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeSpecBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyNodeSpecBatchResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyNodeSpecBatchResponseBody
	GetRequestId() *string
}

type ModifyNodeSpecBatchResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 21012719476****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0637BC25-6895-5500-871F-1127CA34****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodeSpecBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeSpecBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecBatchResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyNodeSpecBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodeSpecBatchResponseBody) SetOrderId(v string) *ModifyNodeSpecBatchResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyNodeSpecBatchResponseBody) SetRequestId(v string) *ModifyNodeSpecBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeSpecBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
