// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterStorageSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBClusterStorageSpaceRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *ModifyDBClusterStorageSpaceRequest
	GetClientToken() *string
	SetCloudProvider(v string) *ModifyDBClusterStorageSpaceRequest
	GetCloudProvider() *string
	SetDBClusterId(v string) *ModifyDBClusterStorageSpaceRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterStorageSpaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterStorageSpaceRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *ModifyDBClusterStorageSpaceRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *ModifyDBClusterStorageSpaceRequest
	GetPlannedStartTime() *string
	SetPromotionCode(v string) *ModifyDBClusterStorageSpaceRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterStorageSpaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterStorageSpaceRequest
	GetResourceOwnerId() *int64
	SetStorageSpace(v int64) *ModifyDBClusterStorageSpaceRequest
	GetStorageSpace() *int64
	SetSubCategory(v string) *ModifyDBClusterStorageSpaceRequest
	GetSubCategory() *string
}

type ModifyDBClusterStorageSpaceRequest struct {
	// Specifies whether to automatically use a coupon. Valid values:
	//
	// - `true` (default): A coupon is automatically used.
	//
	// - `false`: A coupon is not used.
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// A client-generated token that ensures the idempotence of the request. The token must be unique for each request, case-sensitive, and a maximum of 64 ASCII characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies the cloud provider of the instance.
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
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies the latest time to start the scheduled task. Specify the time in UTC in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// > - The latest start time must be at least 30 minutes later than the earliest start time.
	//
	// >
	//
	// > - If you specify `PlannedStartTime` but not this parameter, the latest start time is `PlannedStartTime + 30 minutes` by default. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and leave this parameter empty, the task starts no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// Specifies the earliest time to start the scheduled task. Specify the time in UTC in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// > - This parameter takes effect only when `ModifyType` is set to `Upgrade`.
	//
	// >
	//
	// > - The start time can be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can set the start time to a value that is between `2021-01-14T09:00:00Z` and `2021-01-15T09:00:00Z`.
	//
	// >
	//
	// > - If you leave this parameter empty, the task runs immediately.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// Specifies the promotion code. If you do not specify this parameter, the system uses the default coupon.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies the new storage space. Unit: GB.
	//
	// > For PolarDB for MySQL Standard Edition clusters, the storage space must be between 20 GB and 32,000 GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// Specifies the subcategory of the cluster. Valid values:
	//
	// - **normal_exclusive**: dedicated
	//
	// - **normal_general**: general-purpose
	//
	// example:
	//
	// normal_general
	SubCategory *string `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
}

func (s ModifyDBClusterStorageSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterStorageSpaceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterStorageSpaceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBClusterStorageSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBClusterStorageSpaceRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *ModifyDBClusterStorageSpaceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterStorageSpaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterStorageSpaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterStorageSpaceRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBClusterStorageSpaceRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBClusterStorageSpaceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBClusterStorageSpaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterStorageSpaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterStorageSpaceRequest) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *ModifyDBClusterStorageSpaceRequest) GetSubCategory() *string {
	return s.SubCategory
}

func (s *ModifyDBClusterStorageSpaceRequest) SetAutoUseCoupon(v bool) *ModifyDBClusterStorageSpaceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetClientToken(v string) *ModifyDBClusterStorageSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetCloudProvider(v string) *ModifyDBClusterStorageSpaceRequest {
	s.CloudProvider = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetDBClusterId(v string) *ModifyDBClusterStorageSpaceRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetOwnerAccount(v string) *ModifyDBClusterStorageSpaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetOwnerId(v int64) *ModifyDBClusterStorageSpaceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetPlannedEndTime(v string) *ModifyDBClusterStorageSpaceRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetPlannedStartTime(v string) *ModifyDBClusterStorageSpaceRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetPromotionCode(v string) *ModifyDBClusterStorageSpaceRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterStorageSpaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetResourceOwnerId(v int64) *ModifyDBClusterStorageSpaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetStorageSpace(v int64) *ModifyDBClusterStorageSpaceRequest {
	s.StorageSpace = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) SetSubCategory(v string) *ModifyDBClusterStorageSpaceRequest {
	s.SubCategory = &v
	return s
}

func (s *ModifyDBClusterStorageSpaceRequest) Validate() error {
	return dara.Validate(s)
}
