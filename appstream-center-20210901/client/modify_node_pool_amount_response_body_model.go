// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyNodePoolAmountResponseBodyData) *ModifyNodePoolAmountResponseBody
	GetData() *ModifyNodePoolAmountResponseBodyData
	SetRequestId(v string) *ModifyNodePoolAmountResponseBody
	GetRequestId() *string
}

type ModifyNodePoolAmountResponseBody struct {
	// The returned data.
	Data *ModifyNodePoolAmountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodePoolAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountResponseBody) GetData() *ModifyNodePoolAmountResponseBodyData {
	return s.Data
}

func (s *ModifyNodePoolAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNodePoolAmountResponseBody) SetData(v *ModifyNodePoolAmountResponseBodyData) *ModifyNodePoolAmountResponseBody {
	s.Data = v
	return s
}

func (s *ModifyNodePoolAmountResponseBody) SetRequestId(v string) *ModifyNodePoolAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNodePoolAmountResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyNodePoolAmountResponseBodyData struct {
	// The order ID.
	//
	// example:
	//
	// 23429322113****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyNodePoolAmountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAmountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyNodePoolAmountResponseBodyData) SetOrderId(v string) *ModifyNodePoolAmountResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ModifyNodePoolAmountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
