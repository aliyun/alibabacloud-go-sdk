// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCanAllocateVpcPrivateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanAllocate(v bool) *CheckCanAllocateVpcPrivateIpAddressResponseBody
	GetCanAllocate() *bool
	SetRequestId(v string) *CheckCanAllocateVpcPrivateIpAddressResponseBody
	GetRequestId() *string
}

type CheckCanAllocateVpcPrivateIpAddressResponseBody struct {
	// Indicates whether the private IP address is available. Valid values:
	//
	// - **true**: The private IP address is not occupied in the specified vSwitch and can be used.
	//
	// - **false**: The private IP address is already occupied in the specified vSwitch and cannot be used.
	//
	// example:
	//
	// true
	CanAllocate *bool `json:"CanAllocate,omitempty" xml:"CanAllocate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 93360B0-2969-40BF-8542-EBB34FD358AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCanAllocateVpcPrivateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCanAllocateVpcPrivateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponseBody) GetCanAllocate() *bool {
	return s.CanAllocate
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponseBody) SetCanAllocate(v bool) *CheckCanAllocateVpcPrivateIpAddressResponseBody {
	s.CanAllocate = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponseBody) SetRequestId(v string) *CheckCanAllocateVpcPrivateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
