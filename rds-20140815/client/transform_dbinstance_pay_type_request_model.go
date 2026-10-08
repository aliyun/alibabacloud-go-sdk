// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformDBInstancePayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *TransformDBInstancePayTypeRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *TransformDBInstancePayTypeRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *TransformDBInstancePayTypeRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *TransformDBInstancePayTypeRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *TransformDBInstancePayTypeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *TransformDBInstancePayTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformDBInstancePayTypeRequest
	GetOwnerId() *int64
	SetPayType(v string) *TransformDBInstancePayTypeRequest
	GetPayType() *string
	SetPeriod(v string) *TransformDBInstancePayTypeRequest
	GetPeriod() *string
	SetPromotionCode(v string) *TransformDBInstancePayTypeRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *TransformDBInstancePayTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformDBInstancePayTypeRequest
	GetResourceOwnerId() *int64
	SetUsedTime(v int32) *TransformDBInstancePayTypeRequest
	GetUsedTime() *int32
}

type TransformDBInstancePayTypeRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: Enabled.
	//
	// 	- **false**: Disabled.
	//
	// > 	- This parameter takes effect only when you change the billing method from pay-as-you-go to subscription.
	//
	// > 	- All non-**true*	- strings are treated as **false**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use coupons to offset fees. Valid values:
	//
	// - **true**: Uses coupons to offset fees.
	//
	// - **false**: Does not use coupons to offset fees. This is the default value.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The additional business information about the instance.
	//
	// example:
	//
	// 123456789
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance after the change. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The renewal cycle of the subscription instance. Valid values:
	//
	// 	- **Year**: yearly subscription
	//
	// 	- **Month**: monthly subscription
	//
	// > This parameter is required if **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 726122650073
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, the value of UsedTime ranges from **1 to 5**.
	//
	// 	- If **Period*	- is set to **Month**, the value of UsedTime ranges from **1 to 11**.
	//
	// > This parameter is required if **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s TransformDBInstancePayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformDBInstancePayTypeRequest) GoString() string {
	return s.String()
}

func (s *TransformDBInstancePayTypeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *TransformDBInstancePayTypeRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *TransformDBInstancePayTypeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *TransformDBInstancePayTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TransformDBInstancePayTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *TransformDBInstancePayTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformDBInstancePayTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformDBInstancePayTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *TransformDBInstancePayTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *TransformDBInstancePayTypeRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *TransformDBInstancePayTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformDBInstancePayTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformDBInstancePayTypeRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *TransformDBInstancePayTypeRequest) SetAutoRenew(v string) *TransformDBInstancePayTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetAutoUseCoupon(v bool) *TransformDBInstancePayTypeRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetBusinessInfo(v string) *TransformDBInstancePayTypeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetClientToken(v string) *TransformDBInstancePayTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetDBInstanceId(v string) *TransformDBInstancePayTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetOwnerAccount(v string) *TransformDBInstancePayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetOwnerId(v int64) *TransformDBInstancePayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetPayType(v string) *TransformDBInstancePayTypeRequest {
	s.PayType = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetPeriod(v string) *TransformDBInstancePayTypeRequest {
	s.Period = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetPromotionCode(v string) *TransformDBInstancePayTypeRequest {
	s.PromotionCode = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetResourceOwnerAccount(v string) *TransformDBInstancePayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetResourceOwnerId(v int64) *TransformDBInstancePayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetUsedTime(v int32) *TransformDBInstancePayTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) Validate() error {
	return dara.Validate(s)
}
