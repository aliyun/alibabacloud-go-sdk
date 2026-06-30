// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBNodeClassRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *ModifyDBNodeClassRequest
	GetClientToken() *string
	SetCloudProvider(v string) *ModifyDBNodeClassRequest
	GetCloudProvider() *string
	SetDBClusterId(v string) *ModifyDBNodeClassRequest
	GetDBClusterId() *string
	SetDBNodeTargetClass(v string) *ModifyDBNodeClassRequest
	GetDBNodeTargetClass() *string
	SetDBNodeType(v string) *ModifyDBNodeClassRequest
	GetDBNodeType() *string
	SetModifyType(v string) *ModifyDBNodeClassRequest
	GetModifyType() *string
	SetOwnerAccount(v string) *ModifyDBNodeClassRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeClassRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *ModifyDBNodeClassRequest
	GetPlannedEndTime() *string
	SetPlannedFlashingOffTime(v string) *ModifyDBNodeClassRequest
	GetPlannedFlashingOffTime() *string
	SetPlannedStartTime(v string) *ModifyDBNodeClassRequest
	GetPlannedStartTime() *string
	SetPromotionCode(v string) *ModifyDBNodeClassRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *ModifyDBNodeClassRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeClassRequest
	GetResourceOwnerId() *int64
	SetSubCategory(v string) *ModifyDBNodeClassRequest
	GetSubCategory() *string
}

type ModifyDBNodeClassRequest struct {
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- true (default): Uses coupons.
	//
	// 	- false: Does not use coupons.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The token is case-sensitive and can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cloud service provider of the instance.
	//
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The target specifications for all nodes. For more information, see [Compute node specifications](https://help.aliyun.com/document_detail/102542.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeTargetClass *string `json:"DBNodeTargetClass,omitempty" xml:"DBNodeTargetClass,omitempty"`
	// The node type. Set this parameter to DLNode only when you change the node specifications of an AI node.
	//
	// example:
	//
	// DLNode
	DBNodeType *string `json:"DBNodeType,omitempty" xml:"DBNodeType,omitempty"`
	// The type of the specification change. Valid values:
	//
	// 	- **Upgrade**: upgrades the specifications.
	//
	// 	- **Downgrade**: downgrades the specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// Upgrade
	ModifyType   *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest start time of the scheduled specification change task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format (UTC).
	//
	// > 	- The latest time must be at least 30 minutes later than the start time.
	//
	// > 	- If `PlannedStartTime` is set but this parameter is not specified, the latest time defaults to `start time + 30 minutes`. For example, if `PlannedStartTime` is set to `2021-01-14T09:00:00Z` and this parameter is left empty, the task starts no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The planned transient disconnection time.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedFlashingOffTime *string `json:"PlannedFlashingOffTime,omitempty" xml:"PlannedFlashingOffTime,omitempty"`
	// The earliest start time of the scheduled specification change task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format (UTC).
	//
	// > 	- This parameter takes effect when `ModifyType` is set to `Upgrade` or `Downgrade`.
	//
	// > 	- The start time must be within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, the valid range for the start time is from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// > 	- If this parameter is left empty, the specification change task is immediately executed.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The coupon code. If this parameter is not specified, the default coupon is used.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The cluster sub-series. Valid values:
	//
	// - **normal_exclusive**: Dedicated
	//
	// - **normal_general**: General-purpose
	//
	// This parameter is required when you change specifications between Dedicated and General-purpose.
	//
	// example:
	//
	// normal_general
	SubCategory *string `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
}

func (s ModifyDBNodeClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBNodeClassRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBNodeClassRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *ModifyDBNodeClassRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeClassRequest) GetDBNodeTargetClass() *string {
	return s.DBNodeTargetClass
}

func (s *ModifyDBNodeClassRequest) GetDBNodeType() *string {
	return s.DBNodeType
}

func (s *ModifyDBNodeClassRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyDBNodeClassRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeClassRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeClassRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBNodeClassRequest) GetPlannedFlashingOffTime() *string {
	return s.PlannedFlashingOffTime
}

func (s *ModifyDBNodeClassRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBNodeClassRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBNodeClassRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeClassRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeClassRequest) GetSubCategory() *string {
	return s.SubCategory
}

func (s *ModifyDBNodeClassRequest) SetAutoUseCoupon(v bool) *ModifyDBNodeClassRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetClientToken(v string) *ModifyDBNodeClassRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetCloudProvider(v string) *ModifyDBNodeClassRequest {
	s.CloudProvider = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetDBClusterId(v string) *ModifyDBNodeClassRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetDBNodeTargetClass(v string) *ModifyDBNodeClassRequest {
	s.DBNodeTargetClass = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetDBNodeType(v string) *ModifyDBNodeClassRequest {
	s.DBNodeType = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetModifyType(v string) *ModifyDBNodeClassRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetOwnerAccount(v string) *ModifyDBNodeClassRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetOwnerId(v int64) *ModifyDBNodeClassRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetPlannedEndTime(v string) *ModifyDBNodeClassRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetPlannedFlashingOffTime(v string) *ModifyDBNodeClassRequest {
	s.PlannedFlashingOffTime = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetPlannedStartTime(v string) *ModifyDBNodeClassRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetPromotionCode(v string) *ModifyDBNodeClassRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeClassRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetResourceOwnerId(v int64) *ModifyDBNodeClassRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetSubCategory(v string) *ModifyDBNodeClassRequest {
	s.SubCategory = &v
	return s
}

func (s *ModifyDBNodeClassRequest) Validate() error {
	return dara.Validate(s)
}
