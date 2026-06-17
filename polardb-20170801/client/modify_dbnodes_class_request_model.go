// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodesClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBNodesClassRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *ModifyDBNodesClassRequest
	GetClientToken() *string
	SetCloudProvider(v string) *ModifyDBNodesClassRequest
	GetCloudProvider() *string
	SetDBClusterId(v string) *ModifyDBNodesClassRequest
	GetDBClusterId() *string
	SetDBNode(v []*ModifyDBNodesClassRequestDBNode) *ModifyDBNodesClassRequest
	GetDBNode() []*ModifyDBNodesClassRequestDBNode
	SetModifyType(v string) *ModifyDBNodesClassRequest
	GetModifyType() *string
	SetOwnerAccount(v string) *ModifyDBNodesClassRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodesClassRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *ModifyDBNodesClassRequest
	GetPlannedEndTime() *string
	SetPlannedFlashingOffTime(v string) *ModifyDBNodesClassRequest
	GetPlannedFlashingOffTime() *string
	SetPlannedStartTime(v string) *ModifyDBNodesClassRequest
	GetPlannedStartTime() *string
	SetPromotionCode(v string) *ModifyDBNodesClassRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *ModifyDBNodesClassRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodesClassRequest
	GetResourceOwnerId() *int64
	SetSubCategory(v string) *ModifyDBNodesClassRequest
	GetSubCategory() *string
}

type ModifyDBNodesClassRequest struct {
	// Specifies whether to automatically apply a coupon. Valid values:
	//
	// - true (Default): A coupon is automatically applied.
	//
	// - false: A coupon is not applied.
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// A client-generated token to ensure request idempotence. This token must be unique for each request and must be a case-sensitive string of up to 64 ASCII characters.
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
	// The list of cluster nodes.
	//
	// This parameter is required.
	DBNode []*ModifyDBNodesClassRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
	// The modification type. Valid values:
	//
	// - **Upgrade**: Upgrades the specifications.
	//
	// - **Downgrade**: Downgrades the specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// Upgrade
	ModifyType   *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest time to begin the scheduled task. Specify the time in UTC using the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// > - The latest start time must be at least 30 minutes later than the earliest start time.
	//
	// >
	//
	// > - If you specify `PlannedStartTime` but not this parameter, the task starts within 30 minutes of the `PlannedStartTime`. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and leave this parameter empty, the task will start by `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The planned time for the transient disconnection.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedFlashingOffTime *string `json:"PlannedFlashingOffTime,omitempty" xml:"PlannedFlashingOffTime,omitempty"`
	// The earliest time to begin the scheduled upgrade of the node specifications. Specify the time in UTC using the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// > - This parameter takes effect only when `ModifyType` is set to `Upgrade`.
	//
	// >
	//
	// > - The specified time must be within the next 24 hours.
	//
	// >
	//
	// > - If this parameter is not specified, the upgrade task runs immediately.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The coupon code. If you do not specify this parameter, a default coupon is applied.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The sub-category of the cluster. Valid values:
	//
	// - **normal_exclusive**: dedicated specifications
	//
	// - **normal_general**: general-purpose specifications
	//
	// example:
	//
	// normal_general
	SubCategory *string `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
}

func (s ModifyDBNodesClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodesClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodesClassRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBNodesClassRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBNodesClassRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *ModifyDBNodesClassRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodesClassRequest) GetDBNode() []*ModifyDBNodesClassRequestDBNode {
	return s.DBNode
}

func (s *ModifyDBNodesClassRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyDBNodesClassRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodesClassRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodesClassRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBNodesClassRequest) GetPlannedFlashingOffTime() *string {
	return s.PlannedFlashingOffTime
}

func (s *ModifyDBNodesClassRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBNodesClassRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBNodesClassRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodesClassRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodesClassRequest) GetSubCategory() *string {
	return s.SubCategory
}

func (s *ModifyDBNodesClassRequest) SetAutoUseCoupon(v bool) *ModifyDBNodesClassRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetClientToken(v string) *ModifyDBNodesClassRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetCloudProvider(v string) *ModifyDBNodesClassRequest {
	s.CloudProvider = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetDBClusterId(v string) *ModifyDBNodesClassRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetDBNode(v []*ModifyDBNodesClassRequestDBNode) *ModifyDBNodesClassRequest {
	s.DBNode = v
	return s
}

func (s *ModifyDBNodesClassRequest) SetModifyType(v string) *ModifyDBNodesClassRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetOwnerAccount(v string) *ModifyDBNodesClassRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetOwnerId(v int64) *ModifyDBNodesClassRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetPlannedEndTime(v string) *ModifyDBNodesClassRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetPlannedFlashingOffTime(v string) *ModifyDBNodesClassRequest {
	s.PlannedFlashingOffTime = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetPlannedStartTime(v string) *ModifyDBNodesClassRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetPromotionCode(v string) *ModifyDBNodesClassRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetResourceOwnerAccount(v string) *ModifyDBNodesClassRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetResourceOwnerId(v int64) *ModifyDBNodesClassRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodesClassRequest) SetSubCategory(v string) *ModifyDBNodesClassRequest {
	s.SubCategory = &v
	return s
}

func (s *ModifyDBNodesClassRequest) Validate() error {
	if s.DBNode != nil {
		for _, item := range s.DBNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBNodesClassRequestDBNode struct {
	// The ID of the cluster node.
	//
	// > If you specify this parameter, you must also specify DBNode.N.TargetClass. N represents the index of the node in the request, starting from 1.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The target specifications of the node. For more information about node specifications, see [compute node specifications](https://help.aliyun.com/document_detail/102542.html).
	//
	// > If you specify this parameter, you must also specify DBNode.N.DBNodeId. N represents the index of the node in the request, starting from 1.
	//
	// example:
	//
	// polar.mysql.x4.medium
	TargetClass *string `json:"TargetClass,omitempty" xml:"TargetClass,omitempty"`
}

func (s ModifyDBNodesClassRequestDBNode) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodesClassRequestDBNode) GoString() string {
	return s.String()
}

func (s *ModifyDBNodesClassRequestDBNode) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *ModifyDBNodesClassRequestDBNode) GetTargetClass() *string {
	return s.TargetClass
}

func (s *ModifyDBNodesClassRequestDBNode) SetDBNodeId(v string) *ModifyDBNodesClassRequestDBNode {
	s.DBNodeId = &v
	return s
}

func (s *ModifyDBNodesClassRequestDBNode) SetTargetClass(v string) *ModifyDBNodesClassRequestDBNode {
	s.TargetClass = &v
	return s
}

func (s *ModifyDBNodesClassRequestDBNode) Validate() error {
	return dara.Validate(s)
}
