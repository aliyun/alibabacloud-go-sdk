// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationEndpoint(v *ConfigureSynchronizationJobRequestDestinationEndpoint) *ConfigureSynchronizationJobRequest
	GetDestinationEndpoint() *ConfigureSynchronizationJobRequestDestinationEndpoint
	SetPartitionKey(v *ConfigureSynchronizationJobRequestPartitionKey) *ConfigureSynchronizationJobRequest
	GetPartitionKey() *ConfigureSynchronizationJobRequestPartitionKey
	SetSourceEndpoint(v *ConfigureSynchronizationJobRequestSourceEndpoint) *ConfigureSynchronizationJobRequest
	GetSourceEndpoint() *ConfigureSynchronizationJobRequestSourceEndpoint
	SetAccountId(v string) *ConfigureSynchronizationJobRequest
	GetAccountId() *string
	SetCheckpoint(v string) *ConfigureSynchronizationJobRequest
	GetCheckpoint() *string
	SetDataInitialization(v bool) *ConfigureSynchronizationJobRequest
	GetDataInitialization() *bool
	SetMigrationReserved(v string) *ConfigureSynchronizationJobRequest
	GetMigrationReserved() *string
	SetOwnerId(v string) *ConfigureSynchronizationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureSynchronizationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConfigureSynchronizationJobRequest
	GetResourceGroupId() *string
	SetStructureInitialization(v bool) *ConfigureSynchronizationJobRequest
	GetStructureInitialization() *bool
	SetSynchronizationDirection(v string) *ConfigureSynchronizationJobRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *ConfigureSynchronizationJobRequest
	GetSynchronizationJobId() *string
	SetSynchronizationJobName(v string) *ConfigureSynchronizationJobRequest
	GetSynchronizationJobName() *string
	SetSynchronizationObjects(v string) *ConfigureSynchronizationJobRequest
	GetSynchronizationObjects() *string
}

type ConfigureSynchronizationJobRequest struct {
	DestinationEndpoint *ConfigureSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	PartitionKey        *ConfigureSynchronizationJobRequestPartitionKey        `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty" type:"Struct"`
	SourceEndpoint      *ConfigureSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because it will be discontinued.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The synchronization checkpoint.
	//
	// example:
	//
	// 1610540493
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// Specifies whether to perform initial full data synchronization. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// > Default value: **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet special requirements, such as specifying whether to automatically start the precheck. For more information, see [MigrationReserved parameter description](https://help.aliyun.com/document_detail/176470.html).
	//
	// > For example, you can use this parameter for data synchronization between ApsaraDB for Redis Enhanced Edition (Tair) instances. For more information, see [Use OpenAPI to configure one-way or bidirectional data synchronization between ApsaraDB for Redis Enhanced Edition instances](https://help.aliyun.com/document_detail/155967.html).
	//
	// example:
	//
	// {     "autoStartModulesAfterConfig": "none",     "targetTableMode": 2 }
	MigrationReserved *string `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data synchronization instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to perform initial schema synchronization. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// > Default value: **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - Default value: **Forward**.
	//
	// - This parameter takes effect only if you set it to **Reverse*	- and the synchronization topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The ID of the data synchronization instance. You can call the [DescribeSynchronizationJobs](https://help.aliyun.com/document_detail/49454.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsz4ao1dor13d****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// The name of the synchronization task.
	//
	// > Specify a descriptive name that makes it easy to identify the task. It does not need to be unique.
	//
	// example:
	//
	// MySQL同步
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	// The objects to be synchronized. The value is a JSON string and supports certain regular expressions. For more information, see [Synchronization object configuration](https://help.aliyun.com/document_detail/141901.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"DBName":"dtstestdata","TableIncludes":[{"TableName":"customer"}]}]
	SynchronizationObjects *string `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
}

func (s ConfigureSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequest) GetDestinationEndpoint() *ConfigureSynchronizationJobRequestDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *ConfigureSynchronizationJobRequest) GetPartitionKey() *ConfigureSynchronizationJobRequestPartitionKey {
	return s.PartitionKey
}

func (s *ConfigureSynchronizationJobRequest) GetSourceEndpoint() *ConfigureSynchronizationJobRequestSourceEndpoint {
	return s.SourceEndpoint
}

func (s *ConfigureSynchronizationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ConfigureSynchronizationJobRequest) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *ConfigureSynchronizationJobRequest) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *ConfigureSynchronizationJobRequest) GetMigrationReserved() *string {
	return s.MigrationReserved
}

func (s *ConfigureSynchronizationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureSynchronizationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureSynchronizationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureSynchronizationJobRequest) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *ConfigureSynchronizationJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ConfigureSynchronizationJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *ConfigureSynchronizationJobRequest) GetSynchronizationJobName() *string {
	return s.SynchronizationJobName
}

func (s *ConfigureSynchronizationJobRequest) GetSynchronizationObjects() *string {
	return s.SynchronizationObjects
}

func (s *ConfigureSynchronizationJobRequest) SetDestinationEndpoint(v *ConfigureSynchronizationJobRequestDestinationEndpoint) *ConfigureSynchronizationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetPartitionKey(v *ConfigureSynchronizationJobRequestPartitionKey) *ConfigureSynchronizationJobRequest {
	s.PartitionKey = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSourceEndpoint(v *ConfigureSynchronizationJobRequestSourceEndpoint) *ConfigureSynchronizationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetAccountId(v string) *ConfigureSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetCheckpoint(v string) *ConfigureSynchronizationJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetDataInitialization(v bool) *ConfigureSynchronizationJobRequest {
	s.DataInitialization = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetMigrationReserved(v string) *ConfigureSynchronizationJobRequest {
	s.MigrationReserved = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetOwnerId(v string) *ConfigureSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetRegionId(v string) *ConfigureSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetResourceGroupId(v string) *ConfigureSynchronizationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetStructureInitialization(v bool) *ConfigureSynchronizationJobRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationJobName(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationJobName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationObjects(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationObjects = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) Validate() error {
	if s.DestinationEndpoint != nil {
		if err := s.DestinationEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.PartitionKey != nil {
		if err := s.PartitionKey.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigureSynchronizationJobRequestDestinationEndpoint struct {
	// 目标实例中的同步对象所属数据库名称。
	//
	// example:
	//
	// dtstestdata
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// 目标库的IP地址。
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**Express**、**dg**或**cen**时，本参数必须传入本参数才可用且必须传入。
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 同步目标实例的实例ID
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**MaxCompute**或**DataHub**时，本参数传入MaxCompute实例或DataHub的Project名称。
	//
	// 当目标实例为阿里云分析型数据库MySQL版时，传入分析型数据库MySQL版的集群ID。
	//
	// example:
	//
	// rm-bp1r46452ai50****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 目标实例类型，取值为：
	//
	// - **Redis**：阿里云Redis实例。
	//
	// - **RDS**：阿里云RDS实例。
	//
	// - **PolarDB**：阿里云PolarDB集群（仅支持MySQL或兼容Oracle语法的引擎）。
	//
	// - **ECS**：ECS上的自建数据库。
	//
	// - **Express**：通过专线接入的本地数据库。
	//
	// - **DataHub**：阿里云DataHub实例。
	//
	// - **MaxCompute**：阿里云MaxCompute实例。
	//
	// - **AnalyticDB**：云原生数据仓库AnalyticDB MySQL  3.0和2.0版本。
	//
	// - **Greenplum**：云原生数据仓库ADB PostgreSQL版（原分析型数据库PostgreSQL版）。
	//
	// > 默认取值为**RDS**。
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 目标库数据库账号密码。
	//
	// > - 当**DestinationEndpoint.InstanceType**取值为**ECS**、**Express**、**dg**或**cen**时，本参数必须传入。
	//
	// example:
	//
	// Test654321
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 目标库的数据库服务端口。
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**ECS**、**Express**、**dg**或**cen**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 目标库的数据库账号。
	//
	// > - 同步不同的数据库所需的权限有所差异，详情请参见[DTS数据同步方案概览](https://help.aliyun.com/document_detail/140954.html)中对应的配置案例。
	//
	// - 当**DestinationEndpoint.InstanceType**取值为**ECS**、**Express**、**dg**或**cen**时，本参数必须传入。
	//
	// - 当**DestinationEndpoint.InstanceType**取值为RDS且数据库版本为MySQL 5.5或MySQL 5.6，无需传入本参数和**DestinationEndpoint.Password**参数。
	//
	// - 当**DestinationEndpoint.InstanceType**取值为**Redis**时，无需传入本参数。
	//
	// example:
	//
	// dtstestaccount
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureSynchronizationJobRequestDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) GetDataBaseName() *string {
	return s.DataBaseName
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) GetIP() *string {
	return s.IP
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) GetPassword() *string {
	return s.Password
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetDataBaseName(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.DataBaseName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetIP(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetInstanceId(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetInstanceType(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetPassword(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetPort(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetUserName(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type ConfigureSynchronizationJobRequestPartitionKey struct {
	// 设置增量日志表是否包含以增量更新时间对应日期信息定义的分区，取值：**true**或**false**。
	//
	// > 当**DestinationEndpoint.InstanceType**参数取值为**Maxcompute**时，本参数才可用。
	//
	// example:
	//
	// true
	ModifyTimeDay *bool `json:"ModifyTime_Day,omitempty" xml:"ModifyTime_Day,omitempty"`
	// 设置增量日志表是否包含以增量更新时间对应小时信息定义的分区，取值：**true**或**false**。
	//
	// > 当**DestinationEndpoint.InstanceType**参数取值为**Maxcompute**时，本参数才可用。
	//
	// example:
	//
	// true
	ModifyTimeHour *bool `json:"ModifyTime_Hour,omitempty" xml:"ModifyTime_Hour,omitempty"`
	// 设置增量日志表是否包含以增量更新时间对应分钟信息定义的分区，取值：**true**或**false**。
	//
	// > 当**DestinationEndpoint.InstanceType**参数取值为**Maxcompute**时，本参数才可用。
	//
	// example:
	//
	// true
	ModifyTimeMinute *bool `json:"ModifyTime_Minute,omitempty" xml:"ModifyTime_Minute,omitempty"`
	// 设置增量日志表是否包含以增量更新时间对应月份信息定义的分区，取值：**true**或**false**。
	//
	// > 当**DestinationEndpoint.InstanceType**参数取值为**Maxcompute**时，本参数才可用。
	//
	// example:
	//
	// true
	ModifyTimeMonth *bool `json:"ModifyTime_Month,omitempty" xml:"ModifyTime_Month,omitempty"`
	// 设置增量日志表是否包含以增量更新时间对应年份信息定义的分区，取值：**true**或**false**。
	//
	// > 当**DestinationEndpoint.InstanceType**参数取值为**Maxcompute**时，本参数才可用。
	//
	// example:
	//
	// true
	ModifyTimeYear *bool `json:"ModifyTime_Year,omitempty" xml:"ModifyTime_Year,omitempty"`
}

func (s ConfigureSynchronizationJobRequestPartitionKey) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestPartitionKey) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) GetModifyTimeDay() *bool {
	return s.ModifyTimeDay
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) GetModifyTimeHour() *bool {
	return s.ModifyTimeHour
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) GetModifyTimeMinute() *bool {
	return s.ModifyTimeMinute
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) GetModifyTimeMonth() *bool {
	return s.ModifyTimeMonth
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) GetModifyTimeYear() *bool {
	return s.ModifyTimeYear
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeDay(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeDay = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeHour(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeHour = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeMinute(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeMinute = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeMonth(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeMonth = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeYear(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeYear = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) Validate() error {
	return dara.Validate(s)
}

type ConfigureSynchronizationJobRequestSourceEndpoint struct {
	// 源实例中的同步对象所属数据库名称。
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 源库的IP地址。
	//
	// > 当**SourceEndpoint.InstanceType**取值为**ECS**、**Express**、**dg**或**cen**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 源实例ID。
	//
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 源实例类型，取值为：
	//
	// - **RDS**：阿里云RDS实例。
	//
	// - **Redis**：阿里云Redis实例。
	//
	// - **PolarDB**：阿里云PolarDB集群（仅支持MySQL或兼容Oracle语法的引擎）。
	//
	// - **ECS**：ECS上的自建数据库。
	//
	// - **Express**：通过专线接入的自建数据库。
	//
	// - **dg**：通过数据库网关DG接入的自建数据库。
	//
	// - **cen**：通过云企业网CEN接入的自建数据库。
	//
	// > 默认取值为**RDS**。
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 源RDS实例所属的阿里云账号ID。
	//
	// > 传入本参数即代表执行跨阿里云账号的数据同步，同时您还需要传入**SourceEndpoint.Role**参数。
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// 源库数据库账号密码。
	//
	// > 当**SourceEndpoint.InstanceType**取值为**ECS**、**Express**、**dg**或**cen**时，本参数必须传入。
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 源库的数据库服务端口。
	//
	// > 当**SourceEndpoint.InstanceType**取值为**ECS**、**Express**、**dg**或**cen**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 源实例所属云账号配置的角色名称。
	//
	// > 执行跨阿里云账号的数据同步时须传入本参数，该角色所需的权限及授权方式请参见[跨阿里云账号数据迁移或同步时如何配置RAM授权](https://help.aliyun.com/document_detail/48468.html)。
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 源库的数据库账号。
	//
	// > - 当**SourceEndpoint.InstanceType**取值为**ECS**、**Express**、**dg**或**cen**时，本参数才可用且必须传入。
	//
	// - 当**SourceEndpoint.InstanceType**取值为**Redis**时，本参数无需传入。
	//
	// - 同步不同的数据库所需的权限有所差异，详情请参见[DTS数据同步方案概览](https://help.aliyun.com/document_detail/140954.html)中对应的配置案例。
	//
	// example:
	//
	// dtstestaccount
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureSynchronizationJobRequestSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetOwnerID() *string {
	return s.OwnerID
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetPassword() *string {
	return s.Password
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetRole() *string {
	return s.Role
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetDatabaseName(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetIP(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetInstanceId(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetInstanceType(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetOwnerID(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetPassword(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetPort(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetRole(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Role = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetUserName(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) Validate() error {
	return dara.Validate(s)
}
