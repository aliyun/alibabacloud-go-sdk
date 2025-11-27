// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *MigrateDBNodesRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *MigrateDBNodesRequest
	GetDBInstanceId() *string
	SetDBNode(v []*MigrateDBNodesRequestDBNode) *MigrateDBNodesRequest
	GetDBNode() []*MigrateDBNodesRequestDBNode
	SetEffectiveTime(v string) *MigrateDBNodesRequest
	GetEffectiveTime() *string
	SetOwnerAccount(v string) *MigrateDBNodesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateDBNodesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MigrateDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateDBNodesRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *MigrateDBNodesRequest
	GetSwitchTime() *string
	SetVSwitchId(v string) *MigrateDBNodesRequest
	GetVSwitchId() *string
}

type MigrateDBNodesRequest struct {
	// Specifies the client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3a****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The details of the nodes.
	DBNode []*MigrateDBNodesRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
	// The time when you want the system to start the migration. Valid value:
	//
	// 	- **Immediately**: The system immediately starts the migration. This is the default value.
	//
	// 	- **MaintainTime**: The system starts the migration during the specified maintenance window.
	//
	// 	- **Specified**: The system starts the migration at the specified point in time.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime        *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies the time when the modification is performed. We recommend that you apply the specification during off-peak hours. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s MigrateDBNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBNodesRequest) GoString() string {
	return s.String()
}

func (s *MigrateDBNodesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *MigrateDBNodesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateDBNodesRequest) GetDBNode() []*MigrateDBNodesRequestDBNode {
	return s.DBNode
}

func (s *MigrateDBNodesRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *MigrateDBNodesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateDBNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateDBNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateDBNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateDBNodesRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *MigrateDBNodesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *MigrateDBNodesRequest) SetClientToken(v string) *MigrateDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *MigrateDBNodesRequest) SetDBInstanceId(v string) *MigrateDBNodesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateDBNodesRequest) SetDBNode(v []*MigrateDBNodesRequestDBNode) *MigrateDBNodesRequest {
	s.DBNode = v
	return s
}

func (s *MigrateDBNodesRequest) SetEffectiveTime(v string) *MigrateDBNodesRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateDBNodesRequest) SetOwnerAccount(v string) *MigrateDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateDBNodesRequest) SetOwnerId(v int64) *MigrateDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateDBNodesRequest) SetResourceOwnerAccount(v string) *MigrateDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateDBNodesRequest) SetResourceOwnerId(v int64) *MigrateDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateDBNodesRequest) SetSwitchTime(v string) *MigrateDBNodesRequest {
	s.SwitchTime = &v
	return s
}

func (s *MigrateDBNodesRequest) SetVSwitchId(v string) *MigrateDBNodesRequest {
	s.VSwitchId = &v
	return s
}

func (s *MigrateDBNodesRequest) Validate() error {
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

type MigrateDBNodesRequestDBNode struct {
	// The node ID.
	//
	// example:
	//
	// rn-6256r4a87xvv7****
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The ID of the zone in which the node resides.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s MigrateDBNodesRequestDBNode) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBNodesRequestDBNode) GoString() string {
	return s.String()
}

func (s *MigrateDBNodesRequestDBNode) GetNodeId() *string {
	return s.NodeId
}

func (s *MigrateDBNodesRequestDBNode) GetZoneId() *string {
	return s.ZoneId
}

func (s *MigrateDBNodesRequestDBNode) SetNodeId(v string) *MigrateDBNodesRequestDBNode {
	s.NodeId = &v
	return s
}

func (s *MigrateDBNodesRequestDBNode) SetZoneId(v string) *MigrateDBNodesRequestDBNode {
	s.ZoneId = &v
	return s
}

func (s *MigrateDBNodesRequestDBNode) Validate() error {
	return dara.Validate(s)
}
