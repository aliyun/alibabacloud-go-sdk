// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeClassRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetResourceOwnerAccount(v string) *ModifyDBNodeClassRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeClassRequest
	GetResourceOwnerId() *int64
	SetSubCategory(v string) *ModifyDBNodeClassRequest
	GetSubCategory() *string
}

type ModifyDBNodeClassRequest struct {
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
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The specifications of all nodes. For more information, see [Specifications of computing nodes](https://help.aliyun.com/document_detail/102542.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeTargetClass *string `json:"DBNodeTargetClass,omitempty" xml:"DBNodeTargetClass,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- RO
	//
	// 	- STANDBY
	//
	// 	- DLNode
	//
	// example:
	//
	// DLNode
	DBNodeType *string `json:"DBNodeType,omitempty" xml:"DBNodeType,omitempty"`
	// The type of the configuration change. Valid values:
	//
	// 	- **Upgrade**
	//
	// 	- **Downgrade**
	//
	// This parameter is required.
	//
	// example:
	//
	// Upgrade
	ModifyType   *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest start time to upgrade the specifications within the scheduled time period. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > 	- The value of this parameter must be at least 30 minutes later than the value of PlannedStartTime.
	//
	// >	- By default, if you specify `PlannedStartTime` but do not specify PlannedEndTime, the latest start time of the task is set to `Value of PlannedEndTime + 30 minutes`. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and you do not specify PlannedEndTime, the latest start time of the task is `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime         *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	PlannedFlashingOffTime *string `json:"PlannedFlashingOffTime,omitempty" xml:"PlannedFlashingOffTime,omitempty"`
	// The earliest start time to upgrade the specifications within the scheduled time period. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >	- This parameter takes effect only when `ModifyType` is set to `Upgrade`.
	//
	// >	- The earliest start time of the task can be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a point in the time that ranges from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >	- If this parameter is left empty, the upgrade task is immediately performed.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The category of the cluster. Valid values:
	//
	// 	- **normal_exclusive**: dedicated.
	//
	// 	- **normal_general**: genera-purpose.
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

func (s *ModifyDBNodeClassRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeClassRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeClassRequest) GetSubCategory() *string {
	return s.SubCategory
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
