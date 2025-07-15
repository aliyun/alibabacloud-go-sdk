// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *AllocateEipAddressProResponseBody
	GetAllocationId() *string
	SetEipAddress(v string) *AllocateEipAddressProResponseBody
	GetEipAddress() *string
	SetOrderId(v int64) *AllocateEipAddressProResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *AllocateEipAddressProResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *AllocateEipAddressProResponseBody
	GetResourceGroupId() *string
}

type AllocateEipAddressProResponseBody struct {
	// The EIP ID.
	//
	// example:
	//
	// eip-25877c70gddh****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The IP address that is allocated to the EIP. This parameter is returned only when **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// 192.0.XX.XX
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The order ID.
	//
	// This parameter is returned when InstanceChargeType is set to PrePaid. If AutoPay is set to false, you must manually complete the payment in the [Order Center](https://usercenter2-intl.aliyun.com/order/list).
	//
	// example:
	//
	// 20190000
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group. This parameter is returned only when **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// rg-resourcegroup****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AllocateEipAddressProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressProResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressProResponseBody) GetAllocationId() *string {
	return s.AllocationId
}

func (s *AllocateEipAddressProResponseBody) GetEipAddress() *string {
	return s.EipAddress
}

func (s *AllocateEipAddressProResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *AllocateEipAddressProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateEipAddressProResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateEipAddressProResponseBody) SetAllocationId(v string) *AllocateEipAddressProResponseBody {
	s.AllocationId = &v
	return s
}

func (s *AllocateEipAddressProResponseBody) SetEipAddress(v string) *AllocateEipAddressProResponseBody {
	s.EipAddress = &v
	return s
}

func (s *AllocateEipAddressProResponseBody) SetOrderId(v int64) *AllocateEipAddressProResponseBody {
	s.OrderId = &v
	return s
}

func (s *AllocateEipAddressProResponseBody) SetRequestId(v string) *AllocateEipAddressProResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateEipAddressProResponseBody) SetResourceGroupId(v string) *AllocateEipAddressProResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateEipAddressProResponseBody) Validate() error {
	return dara.Validate(s)
}
