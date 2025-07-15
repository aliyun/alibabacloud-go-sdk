// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceGroupIds(v []*string) *ModifyInstanceChargeTypeResponseBody
	GetInstanceGroupIds() []*string
	SetOrderId(v string) *ModifyInstanceChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyInstanceChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyInstanceChargeTypeResponseBody struct {
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The ID of the order.
	//
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponseBody) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *ModifyInstanceChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceChargeTypeResponseBody) SetInstanceGroupIds(v []*string) *ModifyInstanceChargeTypeResponseBody {
	s.InstanceGroupIds = v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) SetOrderId(v string) *ModifyInstanceChargeTypeResponseBody {
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
