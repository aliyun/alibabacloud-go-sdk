// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConvertChargeTypeRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *ConvertChargeTypeRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *ConvertChargeTypeRequest
	GetPeriod() *int32
	SetProductCode(v string) *ConvertChargeTypeRequest
	GetProductCode() *string
	SetProductType(v string) *ConvertChargeTypeRequest
	GetProductType() *string
	SetSubscriptionType(v string) *ConvertChargeTypeRequest
	GetSubscriptionType() *string
}

type ConvertChargeTypeRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-kasjgfjshgf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Unit: months. This parameter is required if you switch the billing method to subscription. Valid values:
	//
	// 	- 1 to 9
	//
	// 	- 12
	//
	// 	- 24
	//
	// 	- 36
	//
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The code of the service to which the instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service to which the instance belongs.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// >  After the call is successful, the billing method of the instance is switched.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s ConvertChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConvertChargeTypeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ConvertChargeTypeRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ConvertChargeTypeRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ConvertChargeTypeRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *ConvertChargeTypeRequest) SetInstanceId(v string) *ConvertChargeTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetOwnerId(v int64) *ConvertChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetPeriod(v int32) *ConvertChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetProductCode(v string) *ConvertChargeTypeRequest {
	s.ProductCode = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetProductType(v string) *ConvertChargeTypeRequest {
	s.ProductType = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetSubscriptionType(v string) *ConvertChargeTypeRequest {
	s.SubscriptionType = &v
	return s
}

func (s *ConvertChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
