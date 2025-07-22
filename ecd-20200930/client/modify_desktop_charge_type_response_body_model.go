// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *ModifyDesktopChargeTypeResponseBody
	GetDesktopId() []*string
	SetOrderId(v string) *ModifyDesktopChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDesktopChargeTypeResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDesktopChargeTypeResponseBody
	GetTaskId() *string
}

type ModifyDesktopChargeTypeResponseBody struct {
	// The IDs of the cloud computers.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 123456789
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDesktopChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopChargeTypeResponseBody) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *ModifyDesktopChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDesktopChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopChargeTypeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDesktopChargeTypeResponseBody) SetDesktopId(v []*string) *ModifyDesktopChargeTypeResponseBody {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopChargeTypeResponseBody) SetOrderId(v string) *ModifyDesktopChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDesktopChargeTypeResponseBody) SetRequestId(v string) *ModifyDesktopChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopChargeTypeResponseBody) SetTaskId(v string) *ModifyDesktopChargeTypeResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDesktopChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
