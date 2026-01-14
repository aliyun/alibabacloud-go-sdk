// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody
	GetAcceleratorId() *string
	SetAutoRenew(v bool) *DescribeAcceleratorAutoRenewAttributeResponseBody
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *DescribeAcceleratorAutoRenewAttributeResponseBody
	GetAutoRenewDuration() *int32
	SetRenewalStatus(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody
	GetRenewalStatus() *string
	SetRequestId(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type DescribeAcceleratorAutoRenewAttributeResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Indicates whether auto-renewal is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: month.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Indicates how the GA instance is renewed. Valid values:
	//
	// 	- **AutoRenewal**: The GA instance is automatically renewed.
	//
	// 	- **Normal**: You must manually renew the GA instance.
	//
	// 	- **NotRenewal**: The GA instance is not renewed after it expires. The system sends only a non-renewal reminder three days before the expiration date. The system no longer sends notifications to remind you to renew the GA instance.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAcceleratorAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetAcceleratorId(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetAutoRenew(v bool) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetAutoRenewDuration(v int32) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.AutoRenewDuration = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetRenewalStatus(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
