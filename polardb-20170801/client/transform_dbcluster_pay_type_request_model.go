// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformDBClusterPayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *TransformDBClusterPayTypeRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *TransformDBClusterPayTypeRequest
	GetClientToken() *string
	SetDBClusterId(v string) *TransformDBClusterPayTypeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *TransformDBClusterPayTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformDBClusterPayTypeRequest
	GetOwnerId() *int64
	SetPayType(v string) *TransformDBClusterPayTypeRequest
	GetPayType() *string
	SetPeriod(v string) *TransformDBClusterPayTypeRequest
	GetPeriod() *string
	SetPromotionCode(v string) *TransformDBClusterPayTypeRequest
	GetPromotionCode() *string
	SetRegionId(v string) *TransformDBClusterPayTypeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *TransformDBClusterPayTypeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *TransformDBClusterPayTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformDBClusterPayTypeRequest
	GetResourceOwnerId() *int64
	SetUsedTime(v string) *TransformDBClusterPayTypeRequest
	GetUsedTime() *string
}

type TransformDBClusterPayTypeRequest struct {
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The renewal cycle of the cluster. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// >  This parameter is required if you set the **PayType*	- parameter to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3f4un32****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The subscription duration of the cluster. Valid values:
	//
	// 	- If the **Period*	- parameter is set to **Year**, the **UsedTime*	- parameter can be set to 1, 2, or 3.
	//
	// 	- If the **Period*	- parameter is set to **Month**, the **UsedTime*	- parameter can be set to 1, 2, 3, 4, 5, 6, 7, 8, or 9.
	//
	// >  This parameter is required if you set the **PayType*	- parameter to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s TransformDBClusterPayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformDBClusterPayTypeRequest) GoString() string {
	return s.String()
}

func (s *TransformDBClusterPayTypeRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *TransformDBClusterPayTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TransformDBClusterPayTypeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *TransformDBClusterPayTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformDBClusterPayTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformDBClusterPayTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *TransformDBClusterPayTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *TransformDBClusterPayTypeRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *TransformDBClusterPayTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TransformDBClusterPayTypeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransformDBClusterPayTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformDBClusterPayTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformDBClusterPayTypeRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *TransformDBClusterPayTypeRequest) SetAutoUseCoupon(v bool) *TransformDBClusterPayTypeRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetClientToken(v string) *TransformDBClusterPayTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetDBClusterId(v string) *TransformDBClusterPayTypeRequest {
	s.DBClusterId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetOwnerAccount(v string) *TransformDBClusterPayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetOwnerId(v int64) *TransformDBClusterPayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetPayType(v string) *TransformDBClusterPayTypeRequest {
	s.PayType = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetPeriod(v string) *TransformDBClusterPayTypeRequest {
	s.Period = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetPromotionCode(v string) *TransformDBClusterPayTypeRequest {
	s.PromotionCode = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetRegionId(v string) *TransformDBClusterPayTypeRequest {
	s.RegionId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetResourceGroupId(v string) *TransformDBClusterPayTypeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetResourceOwnerAccount(v string) *TransformDBClusterPayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetResourceOwnerId(v int64) *TransformDBClusterPayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetUsedTime(v string) *TransformDBClusterPayTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) Validate() error {
	return dara.Validate(s)
}
