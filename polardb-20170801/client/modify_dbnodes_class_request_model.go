// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodesClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBNodesClassRequest
	GetClientToken() *string
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
	SetResourceOwnerAccount(v string) *ModifyDBNodesClassRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodesClassRequest
	GetResourceOwnerId() *int64
	SetSubCategory(v string) *ModifyDBNodesClassRequest
	GetSubCategory() *string
}

type ModifyDBNodesClassRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The details of the nodes.
	//
	// This parameter is required.
	DBNode []*ModifyDBNodesClassRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
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
	// >	- The value of this parameter must be at least 30 minutes later than the value of PlannedStartTime.
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
	// > 	- This parameter takes effect only when `ModifyType` is set to `Upgrade`.
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
	// 	- **normal_exclusive**: dedicated
	//
	// 	- **normal_general**: genera-purpose
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

func (s *ModifyDBNodesClassRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *ModifyDBNodesClassRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodesClassRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodesClassRequest) GetSubCategory() *string {
	return s.SubCategory
}

func (s *ModifyDBNodesClassRequest) SetClientToken(v string) *ModifyDBNodesClassRequest {
	s.ClientToken = &v
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
	return dara.Validate(s)
}

type ModifyDBNodesClassRequestDBNode struct {
	// The ID of the node.
	//
	// >  If you specify this parameter, DBNode.N.TargetClass is required. N is an integer that starts from 1. The maximum value of N is calculated by using the following formula:16 - The number of current nodes.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The specifications of the node that you want to change. For more information, see [Specifications of compute nodes](https://help.aliyun.com/document_detail/102542.html).
	//
	// >  If you specify this parameter, DBNode.N.DBNodeId is required. N is an integer that starts from 1. The maximum value of N is calculated by using the following formula:16 - The number of current nodes.
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
