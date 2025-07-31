// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyNodeSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyNodeSpecResponseBody
	GetRequestId() *string
}

type ModifyNodeSpecResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 21084641369****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 61B05CCF-EBAB-4BF3-BA67-D77256A721E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodeSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyNodeSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodeSpecResponseBody) SetOrderId(v string) *ModifyNodeSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyNodeSpecResponseBody) SetRequestId(v string) *ModifyNodeSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodeSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
