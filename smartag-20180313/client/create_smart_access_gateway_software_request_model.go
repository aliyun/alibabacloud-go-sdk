// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewaySoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateSmartAccessGatewaySoftwareRequest
	GetAutoPay() *bool
	SetChargeType(v string) *CreateSmartAccessGatewaySoftwareRequest
	GetChargeType() *string
	SetDataPlan(v int64) *CreateSmartAccessGatewaySoftwareRequest
	GetDataPlan() *int64
	SetOwnerAccount(v string) *CreateSmartAccessGatewaySoftwareRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSmartAccessGatewaySoftwareRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateSmartAccessGatewaySoftwareRequest
	GetPeriod() *int32
	SetRegionId(v string) *CreateSmartAccessGatewaySoftwareRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSmartAccessGatewaySoftwareRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmartAccessGatewaySoftwareRequest
	GetResourceOwnerId() *int64
	SetUserCount(v int32) *CreateSmartAccessGatewaySoftwareRequest
	GetUserCount() *int32
}

type CreateSmartAccessGatewaySoftwareRequest struct {
	// Specifies whether to automatically complete the payment of the order. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// If you set the parameter to false, go to Billing Management to complete the payment after you call this operation. The instance is created only after you complete the payment.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The billing method of the SAG app instance. Set the value to **PREPAY**. This value indicates that the SAG app instance is a subscription resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The size of the free data plan that is allocated to each client account per month. Unit: GB. Valid value: **5**.
	//
	// >  This value specifies that a free data plan of 5 GB is allocated to each client account per month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DataPlan     *int64  `json:"DataPlan,omitempty" xml:"DataPlan,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the SAG app instance. Unit: months.
	//
	// Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region where you want to create the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum number of client accounts that can be created on the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s CreateSmartAccessGatewaySoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewaySoftwareRequest) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetDataPlan() *int64 {
	return s.DataPlan
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmartAccessGatewaySoftwareRequest) GetUserCount() *int32 {
	return s.UserCount
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetAutoPay(v bool) *CreateSmartAccessGatewaySoftwareRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetChargeType(v string) *CreateSmartAccessGatewaySoftwareRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetDataPlan(v int64) *CreateSmartAccessGatewaySoftwareRequest {
	s.DataPlan = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetOwnerAccount(v string) *CreateSmartAccessGatewaySoftwareRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetOwnerId(v int64) *CreateSmartAccessGatewaySoftwareRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetPeriod(v int32) *CreateSmartAccessGatewaySoftwareRequest {
	s.Period = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetRegionId(v string) *CreateSmartAccessGatewaySoftwareRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetResourceOwnerAccount(v string) *CreateSmartAccessGatewaySoftwareRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetResourceOwnerId(v int64) *CreateSmartAccessGatewaySoftwareRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) SetUserCount(v int32) *CreateSmartAccessGatewaySoftwareRequest {
	s.UserCount = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareRequest) Validate() error {
	return dara.Validate(s)
}
