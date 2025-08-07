// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempModifyDBNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *TempModifyDBNodeRequest
	GetClientToken() *string
	SetDBClusterId(v string) *TempModifyDBNodeRequest
	GetDBClusterId() *string
	SetDBNode(v []*TempModifyDBNodeRequestDBNode) *TempModifyDBNodeRequest
	GetDBNode() []*TempModifyDBNodeRequestDBNode
	SetModifyType(v string) *TempModifyDBNodeRequest
	GetModifyType() *string
	SetOperationType(v string) *TempModifyDBNodeRequest
	GetOperationType() *string
	SetOwnerAccount(v string) *TempModifyDBNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TempModifyDBNodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TempModifyDBNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TempModifyDBNodeRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *TempModifyDBNodeRequest
	GetRestoreTime() *string
}

type TempModifyDBNodeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
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
	// pc-xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about the scaled/added node.
	//
	// This parameter is required.
	DBNode []*TempModifyDBNodeRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
	// The type of configuration change. Set the value to **TempUpgrade**.
	//
	// This parameter is required.
	//
	// example:
	//
	// TempUpgrade
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The operation type. Valid values:
	//
	// 	- **Modify**: temporarily upgrades the configuration of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// Modify
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The rollback time of the configuration for the temporary upgrade. Specify the time in the ISO 8601 standard in the YYYY-MM-DD hh:mm:ss format.
	//
	// >  The rollback time cannot be 1 hour earlier than the current time and cannot be later than one day before the time when the cluster expires.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-23 18:16:00
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s TempModifyDBNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s TempModifyDBNodeRequest) GoString() string {
	return s.String()
}

func (s *TempModifyDBNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TempModifyDBNodeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *TempModifyDBNodeRequest) GetDBNode() []*TempModifyDBNodeRequestDBNode {
	return s.DBNode
}

func (s *TempModifyDBNodeRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *TempModifyDBNodeRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *TempModifyDBNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TempModifyDBNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TempModifyDBNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TempModifyDBNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TempModifyDBNodeRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *TempModifyDBNodeRequest) SetClientToken(v string) *TempModifyDBNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetDBClusterId(v string) *TempModifyDBNodeRequest {
	s.DBClusterId = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetDBNode(v []*TempModifyDBNodeRequestDBNode) *TempModifyDBNodeRequest {
	s.DBNode = v
	return s
}

func (s *TempModifyDBNodeRequest) SetModifyType(v string) *TempModifyDBNodeRequest {
	s.ModifyType = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetOperationType(v string) *TempModifyDBNodeRequest {
	s.OperationType = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetOwnerAccount(v string) *TempModifyDBNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetOwnerId(v int64) *TempModifyDBNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetResourceOwnerAccount(v string) *TempModifyDBNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetResourceOwnerId(v int64) *TempModifyDBNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TempModifyDBNodeRequest) SetRestoreTime(v string) *TempModifyDBNodeRequest {
	s.RestoreTime = &v
	return s
}

func (s *TempModifyDBNodeRequest) Validate() error {
	return dara.Validate(s)
}

type TempModifyDBNodeRequestDBNode struct {
	// The specifications of the scaled/added node.
	//
	// >
	//
	// 	- The specification of the new node must be consistent with the specifications of the original nodes.
	//
	// 	- You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view the specifications of the original nodes.
	//
	// example:
	//
	// polar.mysql.x4.medium
	TargetClass *string `json:"TargetClass,omitempty" xml:"TargetClass,omitempty"`
	// The ID of the zone in which the added node is deployed. It must be the same zone as the original nodes.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s TempModifyDBNodeRequestDBNode) String() string {
	return dara.Prettify(s)
}

func (s TempModifyDBNodeRequestDBNode) GoString() string {
	return s.String()
}

func (s *TempModifyDBNodeRequestDBNode) GetTargetClass() *string {
	return s.TargetClass
}

func (s *TempModifyDBNodeRequestDBNode) GetZoneId() *string {
	return s.ZoneId
}

func (s *TempModifyDBNodeRequestDBNode) SetTargetClass(v string) *TempModifyDBNodeRequestDBNode {
	s.TargetClass = &v
	return s
}

func (s *TempModifyDBNodeRequestDBNode) SetZoneId(v string) *TempModifyDBNodeRequestDBNode {
	s.ZoneId = &v
	return s
}

func (s *TempModifyDBNodeRequestDBNode) Validate() error {
	return dara.Validate(s)
}
