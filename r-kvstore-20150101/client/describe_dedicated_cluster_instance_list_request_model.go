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
	// The dedicated cluster ID. You can obtain the ID on the dedicated cluster information page in the ApsaraDB for MyBase console.
	//
	// > To specify multiple IDs, separate them with commas (,).
	//
	// example:
	//
	// dhg-5f2v98840ioq****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The host ID in the dedicated cluster. You can call [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/200944.html) to query the host ID.
	//
	// > To specify multiple IDs, separate them with commas (,).
	//
	// example:
	//
	// ch-t4n664a9mal4c****
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The database engine. Set the value to **Redis**.
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version. Set the value to **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance ID.
	//
	// > The instance must be created in a dedicated cluster. For more information, see [ApsaraDB for MyBase](https://help.aliyun.com/document_detail/141455.html).
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type of the instance endpoint. Valid values:
	//
	// 	- **0**: Internet.
	//
	// 	- **1**: Classic network.
	//
	// 	- **2**: Virtual Private Cloud (VPC).
	//
	// example:
	//
	// 2
	InstanceNetType *string `json:"InstanceNetType,omitempty" xml:"InstanceNetType,omitempty"`
	// The instance status. Valid values:
	//
	// 	- **0**: Creating.
	//
	// 	- **1**: Running.
	//
	// 	- **3**: Deleting.
	//
	// 	- **5**: Changing specifications.
	//
	// 	- **6**: Migrating.
	//
	// 	- **7**: Restoring from backup.
	//
	// 	- **8**: Switching.
	//
	// 	- **9**: Deleting expired data.
	//
	// example:
	//
	// 1
	InstanceStatus *int32  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than **0*	- and less than or equal to the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) to query region IDs.
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
	// The zone ID. You can call [DescribeZones](https://help.aliyun.com/document_detail/473764.html) to query zone IDs.
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
