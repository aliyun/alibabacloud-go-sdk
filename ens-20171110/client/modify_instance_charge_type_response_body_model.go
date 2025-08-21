// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *ModifyInstanceChargeTypeResponseBody
	GetInstanceIds() []*string
	SetOrderId(v int64) *ModifyInstanceChargeTypeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyInstanceChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyInstanceChargeTypeResponseBody struct {
	// The IDs of the instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 21522202681****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E2CE5821-7A18-5F7B-A18A-1C751B933D2A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyInstanceChargeTypeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceChargeTypeResponseBody) SetInstanceIds(v []*string) *ModifyInstanceChargeTypeResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) SetOrderId(v int64) *ModifyInstanceChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) SetRequestId(v string) *ModifyInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
