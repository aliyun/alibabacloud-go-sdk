// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAINodeSpecInfos(v []*CreateDBInstanceRequestAINodeSpecInfos) *CreateDBInstanceRequest
	GetAINodeSpecInfos() []*CreateDBInstanceRequestAINodeSpecInfos
	SetBackupId(v string) *CreateDBInstanceRequest
	GetBackupId() *string
	SetCacheStorageSize(v string) *CreateDBInstanceRequest
	GetCacheStorageSize() *string
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetCreateSampleData(v bool) *CreateDBInstanceRequest
	GetCreateSampleData() *bool
	SetDBInstanceCategory(v string) *CreateDBInstanceRequest
	GetDBInstanceCategory() *string
	SetDBInstanceClass(v string) *CreateDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDBInstanceGroupCount(v string) *CreateDBInstanceRequest
	GetDBInstanceGroupCount() *string
	SetDBInstanceMode(v string) *CreateDBInstanceRequest
	GetDBInstanceMode() *string
	SetDeployMode(v string) *CreateDBInstanceRequest
	GetDeployMode() *string
	SetEnableSSL(v bool) *CreateDBInstanceRequest
	GetEnableSSL() *bool
	SetEncryptionKey(v string) *CreateDBInstanceRequest
	GetEncryptionKey() *string
	SetEncryptionType(v string) *CreateDBInstanceRequest
	GetEncryptionType() *string
	SetEngine(v string) *CreateDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceRequest
	GetEngineVersion() *string
	SetIdleTime(v int32) *CreateDBInstanceRequest
	GetIdleTime() *int32
	SetInstanceNetworkType(v string) *CreateDBInstanceRequest
	GetInstanceNetworkType() *string
	SetInstanceSpec(v string) *CreateDBInstanceRequest
	GetInstanceSpec() *string
	SetMasterAISpec(v string) *CreateDBInstanceRequest
	GetMasterAISpec() *string
	SetMasterCU(v int32) *CreateDBInstanceRequest
	GetMasterCU() *int32
	SetMasterNodeNum(v string) *CreateDBInstanceRequest
	GetMasterNodeNum() *string
	SetOwnerId(v int64) *CreateDBInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *CreateDBInstanceRequest
	GetPrivateIpAddress() *string
	SetProdType(v string) *CreateDBInstanceRequest
	GetProdType() *string
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetSecurityIPList(v string) *CreateDBInstanceRequest
	GetSecurityIPList() *string
	SetSegDiskPerformanceLevel(v string) *CreateDBInstanceRequest
	GetSegDiskPerformanceLevel() *string
	SetSegNodeNum(v string) *CreateDBInstanceRequest
	GetSegNodeNum() *string
	SetSegStorageType(v string) *CreateDBInstanceRequest
	GetSegStorageType() *string
	SetServerlessMode(v string) *CreateDBInstanceRequest
	GetServerlessMode() *string
	SetServerlessResource(v int32) *CreateDBInstanceRequest
	GetServerlessResource() *int32
	SetSrcDbInstanceName(v string) *CreateDBInstanceRequest
	GetSrcDbInstanceName() *string
	SetStandbyVSwitchId(v string) *CreateDBInstanceRequest
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *CreateDBInstanceRequest
	GetStandbyZoneId() *string
	SetStorageSize(v int64) *CreateDBInstanceRequest
	GetStorageSize() *int64
	SetStorageType(v string) *CreateDBInstanceRequest
	GetStorageType() *string
	SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest
	GetTag() []*CreateDBInstanceRequestTag
	SetUsedTime(v string) *CreateDBInstanceRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceRequest
	GetVSwitchId() *string
	SetVectorConfigurationStatus(v string) *CreateDBInstanceRequest
	GetVectorConfigurationStatus() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
}

type CreateDBInstanceRequest struct {
	AINodeSpecInfos []*CreateDBInstanceRequestAINodeSpecInfos `json:"AINodeSpecInfos,omitempty" xml:"AINodeSpecInfos,omitempty" type:"Repeated"`
	// Backup set ID.
	//
	// > You can call the [DescribeDataBackups](https://help.aliyun.com/document_detail/210093.html) interface to view the backup set IDs of all backup sets under the target instance.
	//
	// example:
	//
	// 1111111111
	BackupId         *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	CacheStorageSize *string `json:"CacheStorageSize,omitempty" xml:"CacheStorageSize,omitempty"`
	// Idempotence check. For more information, see [How to Ensure Idempotence](https://help.aliyun.com/document_detail/327176.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Whether to load sample datasets after the instance is created. The values are as follows:
	//
	// - **true**: Load sample datasets.
	//
	// - **false**: Do not load sample datasets.
	//
	// > If this parameter is not specified, it defaults to not loading sample datasets.
	//
	// example:
	//
	// false
	CreateSampleData *bool `json:"CreateSampleData,omitempty" xml:"CreateSampleData,omitempty"`
	// Instance series. The value description is as follows:
	//
	// - **HighAvailability**: High availability version.
	//
	// - **Basic**: Basic version.
	//
	// > This parameter is required when creating an instance in the storage elastic mode.
	//
	// example:
	//
	// HighAvailability
	DBInstanceCategory *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	// Instance type. For more details, see the supplementary description of the DBInstanceClass parameter.
	//
	// > This parameter is required when creating a reserved storage mode instance.
	//
	// example:
	//
	// gpdb.group.segsdx1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// Instance description.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// Number of compute groups. The values are: 2, 4, 8, 12, 16, 24, 32, 64, 96, 128.
	//
	// > This parameter is required when creating a reserved storage mode instance.
	//
	// example:
	//
	// 2
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// Instance resource type. The value description is as follows:
	//
	// - **StorageElastic**: Storage elastic mode.
	//
	// - **Serverless**: Serverless mode.
	//
	// - **Classic**: Storage reserved mode.
	//
	// > This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// StorageElastic
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// Deployment mode. The values are as follows:
	//
	// - multiple: Multi-zone deployment.
	//
	// - single: Single-zone deployment.
	//
	// >
	//
	// > - If this parameter is not specified, the default value is single-zone deployment.
	//
	// > - Currently, only single-zone deployment is supported.
	//
	// example:
	//
	// single
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// Specifies whether to enable SSL encryption. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// Key ID.
	//
	// > If the value of the **EncryptionType*	- parameter is **CloudDisk**, you need to specify the encryption key ID within the same region through this parameter; otherwise, it should be empty.
	//
	// example:
	//
	// 0d2470df-da7b-4786-b981-88888888****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// Encryption type. The value description is as follows:
	//
	// - **NULL**: No encryption (default).
	//
	// - **CloudDisk**: Enable cloud disk encryption and specify the key through the **EncryptionKey*	- parameter.
	//
	// > Once cloud disk encryption is enabled, it cannot be disabled.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// Database engine, with the value **gpdb**.
	//
	// This parameter is required.
	//
	// example:
	//
	// gpdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Engine version. The values are as follows:
	//
	// - **6.0**: Version 6.0.
	//
	// - **7.0**: Version 7.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The idle release wait time. When the duration without business traffic reaches the specified time, the instance will enter the idle state. The unit is seconds, with a minimum value of 60, and the default value is 600.
	//
	// > This parameter is required only for Serverless auto-scheduling mode instances.
	//
	// example:
	//
	// 600
	IdleTime *int32 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// Instance network type, with the value **VPC**.
	//
	// > - Only VPC networks are supported in public cloud.
	//
	// > - If not specified, it defaults to VPC type.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// Compute node specifications.
	//
	// For high-availability versions of the elastic storage mode, the values are as follows:
	//
	// - **2C16G**
	//
	// - **4C32G**
	//
	// - **16C128G**
	//
	// For basic versions of the elastic storage mode, the values are as follows:
	//
	// - **2C8G**
	//
	// - **4C16G**
	//
	// - **8C32G**
	//
	// - **16C64G**
	//
	// For Serverless mode, the values are as follows:
	//
	// - **4C16G**
	//
	// - **8C32G**
	//
	// > This parameter is required when creating an elastic storage mode instance or a Serverless mode instance.
	//
	// example:
	//
	// 2C16G
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// This parameter must be specified if you want to change coordinator nodes to AI coordinator nodes.
	//
	// >-  You cannot specify the MasterAISpec and MasterCU parameters at the same time.
	//
	// >- You can change coordinator nodes to AI coordinator nodes only in specific regions and zones.
	//
	// >- Only AnalyticDB for PostgreSQL V7.0 instances of Basic Edition support AI coordinator nodes.
	//
	// >- You can view the valid values of this parameter on the configuration change page of coordinator nodes.
	//
	// example:
	//
	// ADB.AIMedium.2
	MasterAISpec *string `json:"MasterAISpec,omitempty" xml:"MasterAISpec,omitempty"`
	// Master resources, with the following values:
	//
	// - 2 CU
	//
	// - 4 CU
	//
	// - 8 CU
	//
	// - 16 CU
	//
	// - 32 CU
	//
	// > Master resources above 8 CU will incur charges.
	//
	// example:
	//
	// 8 CU
	MasterCU *int32 `json:"MasterCU,omitempty" xml:"MasterCU,omitempty"`
	// This parameter is deprecated and should not be passed.
	//
	// example:
	//
	// null
	MasterNodeNum *string `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// >
	//
	// 	- If you do not specify this parameter, Postpaid is used.
	//
	// 	- You can obtain more cost savings if you create a subscription instance for one year or longer. We recommend that you select the billing method that best suits your needs.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Unit of the duration for which resources are purchased. The values are as follows:
	//
	// - **Month**: Month
	//
	// - **Year**: Year
	//
	// > This parameter is required when creating a subscription-billed instance.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is deprecated and should not be passed.
	//
	// example:
	//
	// null
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// Product type. The values are as follows:
	//
	// - **standard**: Standard Edition.
	//
	// - **cost-effective**: Cost-Effective Edition.
	//
	// > If this parameter is not specified, the default value is Standard Edition.
	//
	// example:
	//
	// standard
	ProdType *string `json:"ProdType,omitempty" xml:"ProdType,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface to view available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group where the instance is located.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IP address whitelist of the instance.
	//
	// A value of 127.0.0.1 denies access from any external IP address. You can call the [ModifySecurityIps](https://help.aliyun.com/document_detail/86928.html) operation to modify the IP address whitelist after you create an instance.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The performance level of ESSDs. Valid values:
	//
	// 	- **pl0**
	//
	// 	- **pl1**
	//
	// 	- **pl2**
	//
	// >
	//
	// 	- This parameter takes effect only when SegStorageType is set to cloud_essd.
	//
	// 	- If you do not specify this parameter, pl1 is used.
	//
	// example:
	//
	// pl1
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// The number of compute nodes. The value description is as follows:
	//
	// - For the high-availability version of the storage elastic mode, the value range is 4 to 512, and the value must be a multiple of 4.
	//
	// - For the basic version of the storage elastic mode, the value range is 2 to 512, and the value must be a multiple of 2.
	//
	// - For the Serverless mode, the value range is 2 to 512, and the value must be a multiple of 2.
	//
	// > This parameter is required when creating instances in the storage elastic mode or Serverless mode.
	//
	// example:
	//
	// 4
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// Disk storage type, currently only ESSD cloud disks are supported, with the value **cloud_essd**.
	//
	// > This parameter is required when creating an elastic storage mode instance.
	//
	// example:
	//
	// cloud_essd
	SegStorageType *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	// The mode of the Serverless instance. The values are as follows:
	//
	// - **Manual**: Manual scheduling (default).
	//
	// - **Auto**: Auto scheduling.
	//
	// > This parameter is required only for Serverless mode instances.
	//
	// example:
	//
	// Auto
	ServerlessMode *string `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	// The threshold for computing resources. The value range is 8 to 32, with a step of 8, and the unit is ACU. The default value is 32.
	//
	// > This parameter is required only for Serverless auto-scheduling mode instances.
	//
	// example:
	//
	// 32
	ServerlessResource *int32 `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	// ID of the source instance to be cloned.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) interface to view details of all AnalyticDB for PostgreSQL instances in the target region, including the instance ID.
	//
	// example:
	//
	// gp-bp***************
	SrcDbInstanceName *string `json:"SrcDbInstanceName,omitempty" xml:"SrcDbInstanceName,omitempty"`
	// VSwitch ID of the standby zone.
	//
	// >
	//
	// > - This parameter is required for multi-zone deployment.
	//
	// > - The VSwitch ID of the standby zone must be in the same zone as the StandbyZoneId.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// ID of the standby zone.
	//
	// >
	//
	// > - This parameter is required for multi-zone deployment.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface to view available zone IDs.
	//
	// > - The ID of the standby zone must be different from the ID of the primary zone.
	//
	// example:
	//
	// cn-hangzhou-j
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The size of the storage space, in GB, with a value range of <props="china">50~8000<props="intl">50~6000.
	//
	// > This parameter is required when creating an instance in the storage elastic mode.
	//
	// example:
	//
	// 200
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// This parameter is deprecated and should not be passed.
	//
	// example:
	//
	// null
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The Nth tag. The value of N ranges from 1 to 20.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Duration for which resources are purchased. The values are as follows:
	//
	// - When **Period*	- is **Month**, the value ranges from 1 to 9.
	//
	// - When **Period*	- is **Year**, the value ranges from 1 to 3.
	//
	// > This parameter is required when creating a subscription-billed instance.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID.
	//
	// > - **VPCId*	- is required.
	//
	// > - The region of the **VPC*	- must be consistent with **RegionId**.
	//
	// example:
	//
	// vpc-bp19ame5m1r3oejns****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// vSwitch ID.
	//
	// > - **vSwitchId*	- is required.
	//
	// > - The availability zone of the **vSwitch*	- must be consistent with **ZoneId**.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Whether to enable vector engine optimization. The value description is as follows:
	//
	// - **enabled**: Enable vector engine optimization.
	//
	// - **disabled*	- (default): Do not enable vector engine optimization.
	//
	// > - For mainstream analysis scenarios, data warehouse scenarios, and real-time data warehouse scenarios, it is recommended to **not enable*	- vector engine optimization.
	//
	// > - For users using the vector analysis engine for AIGC, vector retrieval, and other scenarios, it is recommended to **enable*	- vector engine optimization.
	//
	// example:
	//
	// enabled
	VectorConfigurationStatus *string `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
	// Zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface to view available zone IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetAINodeSpecInfos() []*CreateDBInstanceRequestAINodeSpecInfos {
	return s.AINodeSpecInfos
}

func (s *CreateDBInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateDBInstanceRequest) GetCacheStorageSize() *string {
	return s.CacheStorageSize
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetCreateSampleData() *bool {
	return s.CreateSampleData
}

func (s *CreateDBInstanceRequest) GetDBInstanceCategory() *string {
	return s.DBInstanceCategory
}

func (s *CreateDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceRequest) GetDBInstanceGroupCount() *string {
	return s.DBInstanceGroupCount
}

func (s *CreateDBInstanceRequest) GetDBInstanceMode() *string {
	return s.DBInstanceMode
}

func (s *CreateDBInstanceRequest) GetDeployMode() *string {
	return s.DeployMode
}

func (s *CreateDBInstanceRequest) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *CreateDBInstanceRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateDBInstanceRequest) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *CreateDBInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceRequest) GetIdleTime() *int32 {
	return s.IdleTime
}

func (s *CreateDBInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateDBInstanceRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *CreateDBInstanceRequest) GetMasterAISpec() *string {
	return s.MasterAISpec
}

func (s *CreateDBInstanceRequest) GetMasterCU() *int32 {
	return s.MasterCU
}

func (s *CreateDBInstanceRequest) GetMasterNodeNum() *string {
	return s.MasterNodeNum
}

func (s *CreateDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateDBInstanceRequest) GetProdType() *string {
	return s.ProdType
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceRequest) GetSegDiskPerformanceLevel() *string {
	return s.SegDiskPerformanceLevel
}

func (s *CreateDBInstanceRequest) GetSegNodeNum() *string {
	return s.SegNodeNum
}

func (s *CreateDBInstanceRequest) GetSegStorageType() *string {
	return s.SegStorageType
}

func (s *CreateDBInstanceRequest) GetServerlessMode() *string {
	return s.ServerlessMode
}

func (s *CreateDBInstanceRequest) GetServerlessResource() *int32 {
	return s.ServerlessResource
}

func (s *CreateDBInstanceRequest) GetSrcDbInstanceName() *string {
	return s.SrcDbInstanceName
}

func (s *CreateDBInstanceRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *CreateDBInstanceRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *CreateDBInstanceRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *CreateDBInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBInstanceRequest) GetTag() []*CreateDBInstanceRequestTag {
	return s.Tag
}

func (s *CreateDBInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceRequest) GetVectorConfigurationStatus() *string {
	return s.VectorConfigurationStatus
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) SetAINodeSpecInfos(v []*CreateDBInstanceRequestAINodeSpecInfos) *CreateDBInstanceRequest {
	s.AINodeSpecInfos = v
	return s
}

func (s *CreateDBInstanceRequest) SetBackupId(v string) *CreateDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCacheStorageSize(v string) *CreateDBInstanceRequest {
	s.CacheStorageSize = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCreateSampleData(v bool) *CreateDBInstanceRequest {
	s.CreateSampleData = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceCategory(v string) *CreateDBInstanceRequest {
	s.DBInstanceCategory = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceGroupCount(v string) *CreateDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceMode(v string) *CreateDBInstanceRequest {
	s.DBInstanceMode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDeployMode(v string) *CreateDBInstanceRequest {
	s.DeployMode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEnableSSL(v bool) *CreateDBInstanceRequest {
	s.EnableSSL = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionKey(v string) *CreateDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionType(v string) *CreateDBInstanceRequest {
	s.EncryptionType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetIdleTime(v int32) *CreateDBInstanceRequest {
	s.IdleTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceNetworkType(v string) *CreateDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceSpec(v string) *CreateDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateDBInstanceRequest) SetMasterAISpec(v string) *CreateDBInstanceRequest {
	s.MasterAISpec = &v
	return s
}

func (s *CreateDBInstanceRequest) SetMasterCU(v int32) *CreateDBInstanceRequest {
	s.MasterCU = &v
	return s
}

func (s *CreateDBInstanceRequest) SetMasterNodeNum(v string) *CreateDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrivateIpAddress(v string) *CreateDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDBInstanceRequest) SetProdType(v string) *CreateDBInstanceRequest {
	s.ProdType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegDiskPerformanceLevel(v string) *CreateDBInstanceRequest {
	s.SegDiskPerformanceLevel = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegNodeNum(v string) *CreateDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSegStorageType(v string) *CreateDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetServerlessMode(v string) *CreateDBInstanceRequest {
	s.ServerlessMode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetServerlessResource(v int32) *CreateDBInstanceRequest {
	s.ServerlessResource = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSrcDbInstanceName(v string) *CreateDBInstanceRequest {
	s.SrcDbInstanceName = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStandbyVSwitchId(v string) *CreateDBInstanceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStandbyZoneId(v string) *CreateDBInstanceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageSize(v int64) *CreateDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageType(v string) *CreateDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVectorConfigurationStatus(v string) *CreateDBInstanceRequest {
	s.VectorConfigurationStatus = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDBInstanceRequestAINodeSpecInfos struct {
	AINodeNum  *string `json:"AINodeNum,omitempty" xml:"AINodeNum,omitempty"`
	AINodeSpec *string `json:"AINodeSpec,omitempty" xml:"AINodeSpec,omitempty"`
}

func (s CreateDBInstanceRequestAINodeSpecInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequestAINodeSpecInfos) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestAINodeSpecInfos) GetAINodeNum() *string {
	return s.AINodeNum
}

func (s *CreateDBInstanceRequestAINodeSpecInfos) GetAINodeSpec() *string {
	return s.AINodeSpec
}

func (s *CreateDBInstanceRequestAINodeSpecInfos) SetAINodeNum(v string) *CreateDBInstanceRequestAINodeSpecInfos {
	s.AINodeNum = &v
	return s
}

func (s *CreateDBInstanceRequestAINodeSpecInfos) SetAINodeSpec(v string) *CreateDBInstanceRequestAINodeSpecInfos {
	s.AINodeSpec = &v
	return s
}

func (s *CreateDBInstanceRequestAINodeSpecInfos) Validate() error {
	return dara.Validate(s)
}

type CreateDBInstanceRequestTag struct {
	// Tag key. The restrictions are as follows:
	//
	// - It cannot be an empty string.
	//
	// - It supports up to 128 characters.
	//
	// - It cannot start with `aliyun` or `acs:`, and it cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value. The restrictions are as follows:
	//
	// - It can be an empty string.
	//
	// - It supports up to 128 characters.
	//
	// - It cannot start with `acs:`, and it cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDBInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDBInstanceRequestTag) SetKey(v string) *CreateDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceRequestTag) SetValue(v string) *CreateDBInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDBInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
