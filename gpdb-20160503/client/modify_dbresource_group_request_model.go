// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBResourceGroupRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyDBResourceGroupRequest
	GetOwnerId() *int64
	SetResourceGroupItems(v []*ModifyDBResourceGroupRequestResourceGroupItems) *ModifyDBResourceGroupRequest
	GetResourceGroupItems() []*ModifyDBResourceGroupRequestResourceGroupItems
}

type ModifyDBResourceGroupRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The information about the resource group.
	//
	// This parameter is required.
	ResourceGroupItems []*ModifyDBResourceGroupRequestResourceGroupItems `json:"ResourceGroupItems,omitempty" xml:"ResourceGroupItems,omitempty" type:"Repeated"`
}

func (s ModifyDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBResourceGroupRequest) GetResourceGroupItems() []*ModifyDBResourceGroupRequestResourceGroupItems {
	return s.ResourceGroupItems
}

func (s *ModifyDBResourceGroupRequest) SetDBInstanceId(v string) *ModifyDBResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetOwnerId(v int64) *ModifyDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetResourceGroupItems(v []*ModifyDBResourceGroupRequestResourceGroupItems) *ModifyDBResourceGroupRequest {
	s.ResourceGroupItems = v
	return s
}

func (s *ModifyDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyDBResourceGroupRequestResourceGroupItems struct {
	// The configurations of the resource group to which you want to modify.
	//
	// >
	//
	// 	- CpuRateLimit: the percentage of CPU resources that are available for the resource group. Unit: %.
	//
	// 	- MemoryLimit: the percentage of memory resources that are available for the resource group. Unit: %.
	//
	// 	- MemorySharedQuota: the percentage of memory resources shared among transactions that are submitted to the resource group. Unit: %. Default value: 80.
	//
	// 	- MemorySpillRatio: the memory spill ratio for memory-intensive transactions. When the memory that is used by memory-intensive transactions reaches this value, data is spilled to disks. Unit: %. Default value: 0.
	//
	// 	- Concurrency: the maximum number of concurrent transactions or parallel queries that are allowed for a resource group. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"CpuRateLimit":"10","MemoryLimit":"12","MemorySharedQuota":"20","MemorySpillRatio":"75","Concurrency":"3"}
	ResourceGroupConfig *string `json:"ResourceGroupConfig,omitempty" xml:"ResourceGroupConfig,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// testgroup
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ModifyDBResourceGroupRequestResourceGroupItems) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestResourceGroupItems) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestResourceGroupItems) GetResourceGroupConfig() *string {
	return s.ResourceGroupConfig
}

func (s *ModifyDBResourceGroupRequestResourceGroupItems) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ModifyDBResourceGroupRequestResourceGroupItems) SetResourceGroupConfig(v string) *ModifyDBResourceGroupRequestResourceGroupItems {
	s.ResourceGroupConfig = &v
	return s
}

func (s *ModifyDBResourceGroupRequestResourceGroupItems) SetResourceGroupName(v string) *ModifyDBResourceGroupRequestResourceGroupItems {
	s.ResourceGroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestResourceGroupItems) Validate() error {
	return dara.Validate(s)
}
