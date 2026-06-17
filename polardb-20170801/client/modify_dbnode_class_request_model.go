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
	// Specifies whether to automatically use a coupon. Valid values:
	//
	// - `true` (default): A coupon is automatically applied.
	//
	// - `false`: A coupon is not applied.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// A client-generated token that ensures the idempotence of the request. The token must be unique across requests. It is case-sensitive and can be up to 64 ASCII characters long.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cloud provider of the instance.
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
	// The target node specifications for all nodes in the cluster. For more information, see [compute node specifications](https://help.aliyun.com/document_detail/102542.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeTargetClass *string `json:"DBNodeTargetClass,omitempty" xml:"DBNodeTargetClass,omitempty"`
	// To modify the specifications of an AI node, you must set this parameter to `DLNode`.
	//
	// example:
	//
	// DLNode
	DBNodeType *string `json:"DBNodeType,omitempty" xml:"DBNodeType,omitempty"`
	// The modification type. Valid values:
	//
	// - **Upgrade**: Upgrades the node specifications.
	//
	// - **Downgrade**: Downgrades the node specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// Upgrade
	ModifyType   *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest time to start the scheduled task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - The latest start time must be at least 30 minutes later than the earliest start time.
	//
	// >
	//
	// > - If you specify `PlannedStartTime` but omit this parameter, the latest start time defaults to `PlannedStartTime + 30 minutes`. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and leave this parameter empty, the task starts no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The planned time for a transient disconnection. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedFlashingOffTime *string `json:"PlannedFlashingOffTime,omitempty" xml:"PlannedFlashingOffTime,omitempty"`
	// The earliest time to start the scheduled upgrade or downgrade task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - This parameter is valid only when `ModifyType` is set to `Upgrade` or `Downgrade`.
	//
	// >
	//
	// > - The start time must be within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can set the start time to a value in the range from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >
	//
	// > - If you leave this parameter empty, the task is immediately executed.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The coupon code. If you omit this parameter, the system applies the default coupon.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The subcategory of the cluster. Valid values:
	//
	// - **normal_exclusive**: dedicated specifications
	//
	// - **normal_general**: general-purpose specifications
	//
	// This parameter is required when switching between dedicated and general-purpose specifications.
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
