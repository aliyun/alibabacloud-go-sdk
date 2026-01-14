// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackageAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *DescribeBandwidthPackageAutoRenewAttributeResponseBody
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *DescribeBandwidthPackageAutoRenewAttributeResponseBody
	GetAutoRenewDuration() *int32
	SetInstanceId(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody
	GetInstanceId() *string
	SetRenewalStatus(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody
	GetRenewalStatus() *string
	SetRequestId(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type DescribeBandwidthPackageAutoRenewAttributeResponseBody struct {
	// Indicates whether auto-renewal is enabled.
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: month. Valid values: **1*	- to **12**.
	//
	// >  This parameter is returned only if the value of **AutoRenew*	- is **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1iquvlp8khla5emb3ia
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The auto-renewal status of the bandwidth plan.
	//
	// 	- **AutoRenewal**: The bandwidth plan is automatically renewed.
	//
	// 	- **Normal**: You must manually renew the bandwidth plan.
	//
	// 	- **NotRenewal**: The bandwidth plan is not renewed after it expires. The system sends a non-renewal reminder three days before the expiration date but no longer sends reminders to renew the bandwidth plan. You can change the auto-renewal status of a bandwidth plan from NotRenewal to **Normal*	- or **AutoRenewal**.
	//
	// >  **RenewalStatus*	- takes precedence over **AutoRenew**. If you do not specify **RenewalStatus**, **AutoRenew*	- is automatically used.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B820E9F9-D459-5AE7-8F08-A368B53C1AC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetAutoRenew(v bool) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetAutoRenewDuration(v int32) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.AutoRenewDuration = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetInstanceId(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetRenewalStatus(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
