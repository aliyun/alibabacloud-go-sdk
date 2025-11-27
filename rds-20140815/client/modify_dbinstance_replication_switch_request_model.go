// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceReplicationSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceReplicationSwitchRequest
	GetDBInstanceId() *string
	SetExternalReplication(v string) *ModifyDBInstanceReplicationSwitchRequest
	GetExternalReplication() *string
	SetRegionId(v string) *ModifyDBInstanceReplicationSwitchRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDBInstanceReplicationSwitchRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceReplicationSwitchRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceReplicationSwitchRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable the native replication feature. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	ExternalReplication *string `json:"ExternalReplication,omitempty" xml:"ExternalReplication,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can leave this parameter empty.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceReplicationSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceReplicationSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceReplicationSwitchRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceReplicationSwitchRequest) GetExternalReplication() *string {
	return s.ExternalReplication
}

func (s *ModifyDBInstanceReplicationSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceReplicationSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceReplicationSwitchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceReplicationSwitchRequest) SetDBInstanceId(v string) *ModifyDBInstanceReplicationSwitchRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchRequest) SetExternalReplication(v string) *ModifyDBInstanceReplicationSwitchRequest {
	s.ExternalReplication = &v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchRequest) SetRegionId(v string) *ModifyDBInstanceReplicationSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchRequest) SetResourceGroupId(v string) *ModifyDBInstanceReplicationSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceReplicationSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceReplicationSwitchRequest) Validate() error {
	return dara.Validate(s)
}
