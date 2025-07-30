// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeDedicatedClusterInstanceListRequest
	GetClusterId() *string
	SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListRequest
	GetDedicatedHostName() *string
	SetEngine(v string) *DescribeDedicatedClusterInstanceListRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListRequest
	GetEngineVersion() *string
	SetInstanceId(v string) *DescribeDedicatedClusterInstanceListRequest
	GetInstanceId() *string
	SetInstanceNetType(v string) *DescribeDedicatedClusterInstanceListRequest
	GetInstanceNetType() *string
	SetInstanceStatus(v int32) *DescribeDedicatedClusterInstanceListRequest
	GetInstanceStatus() *int32
	SetOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedClusterInstanceListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDedicatedClusterInstanceListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDedicatedClusterInstanceListRequest
	GetSecurityToken() *string
	SetZoneId(v string) *DescribeDedicatedClusterInstanceListRequest
	GetZoneId() *string
}

type DescribeDedicatedClusterInstanceListRequest struct {
	// The ID of the dedicated cluster. You can view the dedicated cluster ID on the Dedicated Clusters page in the ApsaraDB for MyBase console.
	//
	// > Separate multiple IDs with commas (,).
	//
	// example:
	//
	// dhg-5f2v98840ioq****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the host in the dedicated cluster. You can call the [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/200944.html) operation to query the host ID.
	//
	// > Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ch-t4n664a9mal4c****
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The database engine of the instance. Set the value to **Redis**.
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. Set the value to **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The ID of the instance.
	//
	// > The instance must be created by using a dedicated cluster. For more information, see [What is ApsaraDB for MyBase?](https://help.aliyun.com/document_detail/141455.html)
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **0**: Internet
	//
	// 	- **1**: classic network
	//
	// 	- **2**: Virtual Private Cloud (VPC)
	//
	// example:
	//
	// 2
	InstanceNetType *string `json:"InstanceNetType,omitempty" xml:"InstanceNetType,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- **0**: The instance is being created.
	//
	// 	- **1**: The instance is running.
	//
	// 	- **3**: The instance is being deleted.
	//
	// 	- **5**: The configurations of the instance are being changed.
	//
	// 	- **6**: The instance is being migrated.
	//
	// 	- **7**: The instance is being restored from a backup.
	//
	// 	- **8**: A master-replica switchover is in progress.
	//
	// 	- **9**: Expired data of the instance is being deleted.
	//
	// example:
	//
	// 1
	InstanceStatus *int32  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The zone ID of the instance. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetInstanceNetType() *string {
	return s.InstanceNetType
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetInstanceStatus() *int32 {
	return s.InstanceStatus
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDedicatedClusterInstanceListRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetClusterId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetEngine(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceNetType(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceNetType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetInstanceStatus(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetPageSize(v int32) *DescribeDedicatedClusterInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetRegionId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetResourceOwnerId(v int64) *DescribeDedicatedClusterInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetSecurityToken(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) SetZoneId(v string) *DescribeDedicatedClusterInstanceListRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
