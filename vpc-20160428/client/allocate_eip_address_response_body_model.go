// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *AllocateEipAddressResponseBody
	GetAllocationId() *string
	SetEipAddress(v string) *AllocateEipAddressResponseBody
	GetEipAddress() *string
	SetOrderId(v int64) *AllocateEipAddressResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *AllocateEipAddressResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *AllocateEipAddressResponseBody
	GetResourceGroupId() *string
}

type AllocateEipAddressResponseBody struct {
	// The ID of the EIP instance.
	//
	// example:
	//
	// eip-25877c70gddh****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The allocated EIP. This parameter is returned only if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// 192.0.XX.XX
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The order ID. <props="china">This parameter is returned if InstanceChargeType (the billing method of the EIP) is set to PrePaid (subscription). If AutoPay (automatic payment) is not enabled, go to the [Order Center](https://usercenter2.aliyun.com/order/list) to complete the payment.
	//
	// <props="intl">This parameter is returned if InstanceChargeType (the billing method of the EIP) is set to PrePaid (subscription). If AutoPay (automatic payment) is not enabled, go to the [Order Center](https://usercenter2-intl.aliyun.com/order/list) to complete the payment.
	//
	// example:
	//
	// 10
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID. This parameter is returned only if **InstanceChargeType*	- is set to **PostPaid**.
	//
	// example:
	//
	// rg-acfmxazfdgdg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AllocateEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponseBody) GetAllocationId() *string {
	return s.AllocationId
}

func (s *AllocateEipAddressResponseBody) GetEipAddress() *string {
	return s.EipAddress
}

func (s *AllocateEipAddressResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *AllocateEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateEipAddressResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateEipAddressResponseBody) SetAllocationId(v string) *AllocateEipAddressResponseBody {
	s.AllocationId = &v
	return s
}

func (s *AllocateEipAddressResponseBody) SetEipAddress(v string) *AllocateEipAddressResponseBody {
	s.EipAddress = &v
	return s
}

func (s *AllocateEipAddressResponseBody) SetOrderId(v int64) *AllocateEipAddressResponseBody {
	s.OrderId = &v
	return s
}

func (s *AllocateEipAddressResponseBody) SetRequestId(v string) *AllocateEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateEipAddressResponseBody) SetResourceGroupId(v string) *AllocateEipAddressResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
