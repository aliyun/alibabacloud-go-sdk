// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBMiniEngineVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBMiniEngineVersionsRequest
	GetDBInstanceId() *string
	SetDedicatedHostGroupId(v string) *DescribeDBMiniEngineVersionsRequest
	GetDedicatedHostGroupId() *string
	SetEngine(v string) *DescribeDBMiniEngineVersionsRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDBMiniEngineVersionsRequest
	GetEngineVersion() *string
	SetMinorVersionTag(v string) *DescribeDBMiniEngineVersionsRequest
	GetMinorVersionTag() *string
	SetNodeType(v string) *DescribeDBMiniEngineVersionsRequest
	GetNodeType() *string
	SetRegionId(v string) *DescribeDBMiniEngineVersionsRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeDBMiniEngineVersionsRequest
	GetResourceOwnerId() *int64
	SetStorageType(v string) *DescribeDBMiniEngineVersionsRequest
	GetStorageType() *string
}

type DescribeDBMiniEngineVersionsRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// example:
	//
	// rm-uf6wjk5*******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The dedicated cluster ID. You can call the DescribeDedicatedHostGroups operation to query the dedicated cluster ID.
	//
	// example:
	//
	// dhg-4n*****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The database engine of the instance. Valid values: **MySQL*	- and **PostgreSQL**.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// 	- Valid values when you set the Engine parameter to MySQL: **8.0**, **5.7**, **5.6**, and **5.5**
	//
	// 	- Valid values when you set the Engine parameter to PostgreSQL: **15.0**, **14.0**, **13.0**, **12.0**, **11.0**, and **10.0**
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The minor engine version of the instance. You can specify this parameter to query the minor engine version of the instance.
	//
	// example:
	//
	// rds_20220731
	MinorVersionTag *string `json:"MinorVersionTag,omitempty" xml:"MinorVersionTag,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// HighAvailability
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: local SSD
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: enhanced SSD (ESSD) of performance level 1 (PL1)
	//
	// 	- **cloud_essd2**: ESSD of PL2
	//
	// 	- **cloud_essd3**: ESSD of PL3
	//
	// example:
	//
	// local_ssd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeDBMiniEngineVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBMiniEngineVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBMiniEngineVersionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBMiniEngineVersionsRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBMiniEngineVersionsRequest) GetMinorVersionTag() *string {
	return s.MinorVersionTag
}

func (s *DescribeDBMiniEngineVersionsRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBMiniEngineVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBMiniEngineVersionsRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBMiniEngineVersionsRequest) SetDBInstanceId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetDedicatedHostGroupId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetEngine(v string) *DescribeDBMiniEngineVersionsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetEngineVersion(v string) *DescribeDBMiniEngineVersionsRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetMinorVersionTag(v string) *DescribeDBMiniEngineVersionsRequest {
	s.MinorVersionTag = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetNodeType(v string) *DescribeDBMiniEngineVersionsRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetRegionId(v string) *DescribeDBMiniEngineVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetResourceOwnerId(v int64) *DescribeDBMiniEngineVersionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) SetStorageType(v string) *DescribeDBMiniEngineVersionsRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeDBMiniEngineVersionsRequest) Validate() error {
	return dara.Validate(s)
}
