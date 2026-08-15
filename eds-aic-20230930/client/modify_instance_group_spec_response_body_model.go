// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceGroupSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderInfo(v []*ModifyInstanceGroupSpecResponseBodyOrderInfo) *ModifyInstanceGroupSpecResponseBody
	GetOrderInfo() []*ModifyInstanceGroupSpecResponseBodyOrderInfo
	SetOrderTaskId(v string) *ModifyInstanceGroupSpecResponseBody
	GetOrderTaskId() *string
	SetRequestId(v string) *ModifyInstanceGroupSpecResponseBody
	GetRequestId() *string
}

type ModifyInstanceGroupSpecResponseBody struct {
	// The order information.
	OrderInfo []*ModifyInstanceGroupSpecResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Repeated"`
	// The order task ID that is returned when specifications of more than 10 instance groups are changed in a batch. You can call the **DescribeOrderTasks*	- operation to query the information about each order.
	//
	// example:
	//
	// t-aycabdsjsbgd****
	OrderTaskId *string `json:"OrderTaskId,omitempty" xml:"OrderTaskId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceGroupSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceGroupSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceGroupSpecResponseBody) GetOrderInfo() []*ModifyInstanceGroupSpecResponseBodyOrderInfo {
	return s.OrderInfo
}

func (s *ModifyInstanceGroupSpecResponseBody) GetOrderTaskId() *string {
	return s.OrderTaskId
}

func (s *ModifyInstanceGroupSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceGroupSpecResponseBody) SetOrderInfo(v []*ModifyInstanceGroupSpecResponseBodyOrderInfo) *ModifyInstanceGroupSpecResponseBody {
	s.OrderInfo = v
	return s
}

func (s *ModifyInstanceGroupSpecResponseBody) SetOrderTaskId(v string) *ModifyInstanceGroupSpecResponseBody {
	s.OrderTaskId = &v
	return s
}

func (s *ModifyInstanceGroupSpecResponseBody) SetRequestId(v string) *ModifyInstanceGroupSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceGroupSpecResponseBody) Validate() error {
	if s.OrderInfo != nil {
		for _, item := range s.OrderInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceGroupSpecResponseBodyOrderInfo struct {
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The order ID.
	//
	// example:
	//
	// 296325540190****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyInstanceGroupSpecResponseBodyOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceGroupSpecResponseBodyOrderInfo) GoString() string {
	return s.String()
}

func (s *ModifyInstanceGroupSpecResponseBodyOrderInfo) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyInstanceGroupSpecResponseBodyOrderInfo) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyInstanceGroupSpecResponseBodyOrderInfo) SetInstanceIds(v []*string) *ModifyInstanceGroupSpecResponseBodyOrderInfo {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstanceGroupSpecResponseBodyOrderInfo) SetOrderId(v string) *ModifyInstanceGroupSpecResponseBodyOrderInfo {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceGroupSpecResponseBodyOrderInfo) Validate() error {
	return dara.Validate(s)
}
