// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateInstanceRequest
	GetChargeType() *string
	SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest
	GetComponents() []*CreateInstanceRequestComponents
	SetOrder(v *CreateInstanceRequestOrder) *CreateInstanceRequest
	GetOrder() *CreateInstanceRequestOrder
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
}

type CreateInstanceRequest struct {
	// The billing method of the instance. Valid values: PREPAY: subscription. If you set this parameter to PREPAY, make sure that your Alibaba Cloud account supports balance payment or credit card payment. Otherwise, the system returns the InvalidPayMethod error message. If you set this parameter to PREPAY, you must also specify paymentInfo. POSTPAY: pay-as-you-go. This billing method is not supported.
	//
	// example:
	//
	// ""
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The information about the instance specification.
	Components []*CreateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The billing information.
	Order           *CreateInstanceRequestOrder  `json:"order,omitempty" xml:"order,omitempty" type:"Struct"`
	ResourceGroupId *string                      `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*CreateInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceRequest) GetComponents() []*CreateInstanceRequestComponents {
	return s.Components
}

func (s *CreateInstanceRequest) GetOrder() *CreateInstanceRequestOrder {
	return s.Order
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest {
	s.Components = v
	return s
}

func (s *CreateInstanceRequest) SetOrder(v *CreateInstanceRequestOrder) *CreateInstanceRequest {
	s.Order = v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestComponents struct {
	// The code of the specification, which must be consistent with the value that you specify on the buy page.
	//
	// example:
	//
	// ""
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The value of the specification.
	//
	// example:
	//
	// ""
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateInstanceRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestComponents) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceRequestComponents) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestComponents) SetCode(v string) *CreateInstanceRequestComponents {
	s.Code = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetValue(v string) *CreateInstanceRequestComponents {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestComponents) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestOrder struct {
	// Specifies whether to enable auto-renewal. Valid values: true and false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The billing duration. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, and 12.
	//
	// example:
	//
	// 29
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The unit of the billing duration. Valid values: Month and Year.
	//
	// example:
	//
	// ""
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s CreateInstanceRequestOrder) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestOrder) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestOrder) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequestOrder) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateInstanceRequestOrder) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInstanceRequestOrder) SetAutoRenew(v bool) *CreateInstanceRequestOrder {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequestOrder) SetDuration(v int64) *CreateInstanceRequestOrder {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequestOrder) SetPricingCycle(v string) *CreateInstanceRequestOrder {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequestOrder) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
