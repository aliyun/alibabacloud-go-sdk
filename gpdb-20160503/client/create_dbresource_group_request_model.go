// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBResourceGroupRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *CreateDBResourceGroupRequest
	GetOwnerId() *int64
	SetResourceGroupConfig(v string) *CreateDBResourceGroupRequest
	GetResourceGroupConfig() *string
	SetResourceGroupName(v string) *CreateDBResourceGroupRequest
	GetResourceGroupName() *string
}

type CreateDBResourceGroupRequest struct {
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
	// The configurations of the resource group.
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

func (s CreateDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBResourceGroupRequest) GetResourceGroupConfig() *string {
	return s.ResourceGroupConfig
}

func (s *CreateDBResourceGroupRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CreateDBResourceGroupRequest) SetDBInstanceId(v string) *CreateDBResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetOwnerId(v int64) *CreateDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetResourceGroupConfig(v string) *CreateDBResourceGroupRequest {
	s.ResourceGroupConfig = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetResourceGroupName(v string) *CreateDBResourceGroupRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
