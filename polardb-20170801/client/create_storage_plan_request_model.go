// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoragePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *CreateStoragePlanRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *CreateStoragePlanRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CreateStoragePlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateStoragePlanRequest
	GetOwnerId() *int64
	SetPeriod(v string) *CreateStoragePlanRequest
	GetPeriod() *string
	SetPromotionCode(v string) *CreateStoragePlanRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *CreateStoragePlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateStoragePlanRequest
	GetResourceOwnerId() *int64
	SetStorageClass(v string) *CreateStoragePlanRequest
	GetStorageClass() *string
	SetStorageType(v string) *CreateStoragePlanRequest
	GetStorageType() *string
	SetUsedTime(v string) *CreateStoragePlanRequest
	GetUsedTime() *string
}

type CreateStoragePlanRequest struct {
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the subscription duration for the storage plan. Valid values:
	//
	// 	- **Month**
	//
	// 	- **Year**
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 727xxxxxx934
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The capacity of the storage plan. Unit: GB. Valid values: 50, 100, 200, 300, 500, 1000, 2000, 3000, 5000, 10000, 15000, 20000, 25000, 30000, 50000, 100000, and 200000
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The type of the storage plan. Valid values:
	//
	// 	- **Mainland**: The storage plan is used inside the Chinese mainland.
	//
	// 	- **Overseas**: The storage plan is used outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mainland
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The subscription duration of the storage plan.
	//
	// 	- If **Period*	- is set to **Month**, the value ranges from 1 to 9.
	//
	// 	- If **Period*	- is set to **Year**, the value can be 1, 2, 3, or 5.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s CreateStoragePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStoragePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateStoragePlanRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateStoragePlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStoragePlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateStoragePlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateStoragePlanRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateStoragePlanRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateStoragePlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateStoragePlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateStoragePlanRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *CreateStoragePlanRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateStoragePlanRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateStoragePlanRequest) SetAutoUseCoupon(v bool) *CreateStoragePlanRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateStoragePlanRequest) SetClientToken(v string) *CreateStoragePlanRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStoragePlanRequest) SetOwnerAccount(v string) *CreateStoragePlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateStoragePlanRequest) SetOwnerId(v int64) *CreateStoragePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateStoragePlanRequest) SetPeriod(v string) *CreateStoragePlanRequest {
	s.Period = &v
	return s
}

func (s *CreateStoragePlanRequest) SetPromotionCode(v string) *CreateStoragePlanRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateStoragePlanRequest) SetResourceOwnerAccount(v string) *CreateStoragePlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateStoragePlanRequest) SetResourceOwnerId(v int64) *CreateStoragePlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateStoragePlanRequest) SetStorageClass(v string) *CreateStoragePlanRequest {
	s.StorageClass = &v
	return s
}

func (s *CreateStoragePlanRequest) SetStorageType(v string) *CreateStoragePlanRequest {
	s.StorageType = &v
	return s
}

func (s *CreateStoragePlanRequest) SetUsedTime(v string) *CreateStoragePlanRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateStoragePlanRequest) Validate() error {
	return dara.Validate(s)
}
