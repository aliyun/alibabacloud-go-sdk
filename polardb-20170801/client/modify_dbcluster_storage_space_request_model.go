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
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// The latest time to upgrade the specifications within the scheduled time period. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >- The value of this parameter must be at least 30 minutes later than PlannedStartTime.
	//
	// >- By default, if you specify `PlannedStartTime` but do not specify PlannedEndTime, the latest start time of the task is set to `PlannedEndTime + 30 minutes`. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and you do not specify PlannedEndTime, the latest start time of the task is `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to upgrade the specifications within the scheduled time period. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >- This parameter takes effect only when `ModifyType` is set to `Upgrade`.
	//
	// >- The earliest start time of the task can be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a point in the time that ranges from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >- If this parameter is left empty, the upgrade task is immediately performed.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// example:
	//
	// 727xxxxxx934
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage capacity that you can select when you change the cluster. Unit: GB.
	//
	// >  You can set this parameter for PolarDB for MySQL clusters of Standard Edition to a value that ranges from 20 to 32000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The category of the cluster. Default value: ON. Valid values:
	//
	// 	- **normal_exclusive**: dedicated
	//
	// 	- **normal_general**: general-purpose
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
