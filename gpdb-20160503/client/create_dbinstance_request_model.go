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
	// The AI node specifications information.
	AINodeSpecInfos []*CreateDBInstanceRequestAINodeSpecInfos `json:"AINodeSpecInfos,omitempty" xml:"AINodeSpecInfos,omitempty" type:"Repeated"`
	// The backup set ID.
	//
	// > You can call the [DescribeDataBackups](https://help.aliyun.com/document_detail/210093.html) operation to query the backup set IDs of all backup sets for the target instance.
	//
	// example:
	//
	// 1111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The Serverless cache storage size. Valid values: 800 to 102400. The step size is 32 GB. Unit: GB.
	//
	// example:
	//
	// 800
	CacheStorageSize *string `json:"CacheStorageSize,omitempty" xml:"CacheStorageSize,omitempty"`
	// The client token used for idempotence checks. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/327176.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to load a sample dataset after the instance is created. Valid values:
	//
	// - **true**: Load a sample dataset.
	//
	// - **false**: Do not load a sample dataset.
	//
	// > If you do not specify this parameter, the default value is false.
	//
	// example:
	//
	// false
	CreateSampleData *bool `json:"CreateSampleData,omitempty" xml:"CreateSampleData,omitempty"`
	// The instance edition. Valid values:
	//
	// - **HighAvailability**: high-availability edition.
	//
	// - **Basic**: basic edition.
	//
	// > This parameter is required when you create a storage elastic mode instance.
	//
	// example:
	//
	// HighAvailability
	DBInstanceCategory *string `json:"DBInstanceCategory,omitempty" xml:"DBInstanceCategory,omitempty"`
	// The instance specifications. For more information, see the supplementary description of the DBInstanceClass parameter.
	//
	// > This parameter is required when you create a storage reservation mode instance.
	//
	// example:
	//
	// gpdb.group.segsdx1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The number of compute groups. Valid values: 2, 4, 8, 12, 16, 24, 32, 64, 96, and 128.
	//
	// > This parameter is required when you create a storage reservation mode instance.
	//
	// example:
	//
	// 2
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// The resource type of the instance. Valid values:
	//
	// - **StorageElastic**: storage elastic mode.
	//
	// - **Serverless**: Serverless mode.
	//
	// - **ServerlessPro**: Serverless Pro mode.
	//
	// - **Classic**: storage reservation mode.
	//
	// > This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// StorageElastic
	DBInstanceMode *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	// The deployment mode. Valid values:
	//
	// - multiple: multi-zone deployment.
	//
	// - single: single-zone deployment.
	//
	// >
	//
	// > - If you do not specify this parameter, the default value is single-zone deployment.
	//
	// > - Currently, only single-zone deployment is supported.
	//
	// example:
	//
	// single
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// Indicates whether to enable Secure Sockets Layer (SSL) encryption. Valid values:
	//
	// - **true**: Enable SSL encryption.
	//
	// - **false*	- (default): Do not enable SSL encryption.
	//
	// example:
	//
	// false
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// The encryption key ID.
	//
	// > If the **EncryptionType*	- parameter is set to **CloudDisk**, specify the encryption key ID in the same region. Otherwise, leave this parameter empty.
	//
	// example:
	//
	// 0d2470df-da7b-4786-b981-88888888****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Valid values:
	//
	// - **NULL**: Encryption is not enabled. This is the default value.
	//
	// - **CloudDisk**: Cloud disk encryption is enabled. Use the **EncryptionKey*	- parameter to specify the encryption key.
	//
	// > Cloud disk encryption cannot be disabled after it is enabled.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The database engine. Set the value to **gpdb**.
	//
	// This parameter is required.
	//
	// example:
	//
	// gpdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version. Valid values:
	//
	// - **6.0**: version 6.0.
	//
	// - **7.0**: version 7.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The idle release wait time. The instance transitions to an idle state after no service traffic is detected for the specified duration. Unit: seconds. Minimum value: 60. Default value: 600.
	//
	// > This parameter is required only for Serverless instances in automatic scheduling mode.
	//
	// example:
	//
	// 600
	IdleTime *int32 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// The network type of the instance. Set the value to **VPC**.
	//
	// > - Only VPC is supported on the public cloud.
	//
	// > - If you do not specify this parameter, the default value is VPC.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The node specifications of compute nodes.
	//
	// Valid values for storage elastic mode (high-availability edition):
	//
	// - **2C16G**
	//
	// - **4C32G**
	//
	// - **16C128G**
	//
	// Valid values for storage elastic mode (basic edition):
	//
	// - **2C8G**
	//
	// - **4C16G**
	//
	// - **8C32G**
	//
	// - **16C64G**
	//
	// Valid values for Serverless mode:
	//
	// - **4C16G**
	//
	// - **8C32G**
	//
	// > This parameter is required when you create a storage elastic mode instance or a Serverless mode instance.
	//
	// example:
	//
	// 2C16G
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// Specify this parameter if you want to change the master node to a MasterAI node.
	//
	// > - This parameter and MasterCU cannot be specified at the same time.
	//
	// > - Only specific regions and zones support changing the master node to a MasterAI node.
	//
	// > - Only basic edition instances of AnalyticDB for PostgreSQL 7.0 support MasterAI nodes.
	//
	// > - You can view all valid values of this parameter on the specification change page for the master node.
	//
	// example:
	//
	// ADB.AIMedium.2
	MasterAISpec *string `json:"MasterAISpec,omitempty" xml:"MasterAISpec,omitempty"`
	// The master resources. Valid values:
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
	// > Master resources above 8 CU incur additional fees.
	//
	// example:
	//
	// 8 CU
	MasterCU *int32 `json:"MasterCU,omitempty" xml:"MasterCU,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	MasterNodeNum *string `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// > - If you do not specify this parameter, the default value is pay-as-you-go.
	//
	// > - Discounts are available for subscription instances when you purchase them for one year or longer. Select a billing method based on your business requirements.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - **Month**: month.
	//
	// - **Year**: year.
	//
	// > This parameter is required when you create a subscription instance.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The product type. Valid values:
	//
	// - **standard**: Standard Edition.
	//
	// - **cost-effective**: Economy Edition.
	//
	// > If you do not specify this parameter, the default value is Standard Edition.
	//
	// example:
	//
	// standard
	ProdType *string `json:"ProdType,omitempty" xml:"ProdType,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IP address whitelist.
	//
	// A value of 127.0.0.1 indicates that no external IP addresses are allowed to access the instance. After the instance is created, you can call the [ModifySecurityIps](https://help.aliyun.com/document_detail/86928.html) operation to modify the IP address whitelist.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. Valid values:
	//
	// - **pl0**: PL0.
	//
	// - **pl1**: PL1.
	//
	// - **pl2**: PL2.
	//
	// > - This parameter takes effect only when the cloud disk storage type is ESSD cloud disk.
	//
	// > - If you do not specify this parameter, the default value is PL1.
	//
	// example:
	//
	// pl1
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// The number of compute nodes. Valid values:
	//
	// - Storage elastic mode (high-availability edition): 4 to 512. The value must be a multiple of 4.
	//
	// - Storage elastic mode (basic edition): 2 to 512. The value must be a multiple of 2.
	//
	// - Serverless mode: 2 to 512. The value must be a multiple of 2.
	//
	// > This parameter is required when you create a storage elastic mode instance or a Serverless mode instance.
	//
	// example:
	//
	// 4
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The cloud disk storage type. Currently, only ESSD cloud disks are supported. Set the value to **cloud_essd**.
	//
	// > This parameter is required when you create a storage elastic mode instance.
	//
	// example:
	//
	// cloud_essd
	SegStorageType *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	// The mode of the Serverless instance. Valid values:
	//
	// - **Manual**: manual scheduling. This is the default value.
	//
	// - **Auto**: automatic scheduling.
	//
	// > - This parameter is required only for Serverless mode instances.
	//
	// > - The automatic scheduling mode for AnalyticDB for PostgreSQL Serverless is in invitational preview. To use this feature, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket?product=rds) to apply for access.
	//
	// example:
	//
	// Auto
	ServerlessMode *string `json:"ServerlessMode,omitempty" xml:"ServerlessMode,omitempty"`
	// The compute resource threshold. Valid values: 16 to 1024. The step size is 8. Unit: ACU.
	//
	// > This parameter is required only for Serverless instances in automatic scheduling mode.
	//
	// example:
	//
	// 16
	ServerlessResource *int32 `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	// The ID of the source instance to clone.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in the target region, including instance IDs.
	//
	// example:
	//
	// gp-bp***************
	SrcDbInstanceName *string `json:"SrcDbInstanceName,omitempty" xml:"SrcDbInstanceName,omitempty"`
	// The vSwitch ID in the secondary zone.
	//
	// >
	//
	// > - This parameter is required only for multi-zone deployment.
	//
	// > - The vSwitch must be in the same zone as the value specified by StandbyZoneId.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The secondary zone ID.
	//
	// >
	//
	// > - This parameter is required only for multi-zone deployment.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available zone IDs.
	//
	// > - The secondary zone ID must be different from the primary zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The storage capacity. Unit: GB. Valid values: <props="china">50 to 8000<props="intl">50 to 6000.
	//
	// > This parameter is required when you create a storage elastic mode instance.
	//
	// example:
	//
	// 200
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The Nth tag. Valid values of N: 1 to 20.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription duration. Valid values:
	//
	// - If **Period*	- is set to **Month**, valid values are 1 to 9.
	//
	// - If **Period*	- is set to **Year**, valid values are 1 to 3.
	//
	// > This parameter is required when you create a subscription instance.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID.
	//
	// > - **VPCId*	- is required.
	//
	// > - The VPC must be in the same region as the value specified by **RegionId**.
	//
	// example:
	//
	// vpc-bp19ame5m1r3oejns****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// > - **vSwitchId*	- is required.
	//
	// > - The vSwitch must be in the same zone as the value specified by **ZoneId**.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Specifies whether to enable vector engine optimization. Valid values:
	//
	// - **enabled**: Enable vector engine optimization.
	//
	// - **disabled*	- (default): Do not enable vector engine optimization.
	//
	// > - For mainstream analytics, data warehousing, and real-time data warehouse scenarios, **do not enable*	- vector engine optimization.
	//
	// > - For scenarios that use the vector analysis engine for AIGC and AISearch, **enable*	- vector engine optimization.
	//
	// example:
	//
	// enabled
	VectorConfigurationStatus *string `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available zone IDs.
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
	if s.AINodeSpecInfos != nil {
		for _, item := range s.AINodeSpecInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBInstanceRequestAINodeSpecInfos struct {
	// The number of AI nodes.
	//
	// example:
	//
	// 1
	AINodeNum *string `json:"AINodeNum,omitempty" xml:"AINodeNum,omitempty"`
	// The AI node specifications.
	//
	// example:
	//
	// ADB.AIMedium.2
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
	// The tag key. The following limits apply:
	//
	// - The tag key cannot be an empty string.
	//
	// - The tag key can be up to 128 characters in length.
	//
	// - The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The following limits apply:
	//
	// - The tag value can be an empty string.
	//
	// - The tag value can be up to 128 characters in length.
	//
	// - The tag value cannot start with `acs:`, and cannot contain `http://` or `https://`.
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
