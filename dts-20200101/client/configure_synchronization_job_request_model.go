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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter. This parameter will be deprecated soon.
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
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > Default value: **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// The reserved parameter for DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, such as whether to automatically start a precheck. For more information, see [MigrationReserved parameter description](https://help.aliyun.com/document_detail/176470.html).
	//
	// > For example, you can use this parameter for data synchronization between Redis Enterprise instances. For more information, see [Call OpenAPI operations to configure one-way or two-way data synchronization between Redis Enterprise instances](https://help.aliyun.com/document_detail/155967.html).
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to perform initial schema synchronization. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
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
	// - **Forward**: Forward synchronization.
	//
	// - **Reverse**: Reverse synchronization.
	//
	// > - Default value: **Forward**.
	//
	// - This parameter takes effect only when the synchronization topology of the data synchronization instance is two-way synchronization and the value is set to **Reverse**.
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
	// > Specify a descriptive name that is easy to identify. The name does not need to be unique.
	//
	// example:
	//
	// MySQL synchronization
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	// The objects to be synchronized. The value is a JSON string and supports specific regular expressions. For more information, see [Configure synchronization objects](https://help.aliyun.com/document_detail/141901.html).
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
	// The name of the database to which the synchronization objects in the destination instance belong.
	//
	// example:
	//
	// dtstestdata
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// The IP address of the destination database.
	//
	// > This parameter is available and required only when the **DestinationEndpoint.InstanceType*	- parameter is set to **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The instance ID of the destination instance for synchronization.
	//
	// > When the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute*	- or **DataHub**, set this parameter to the name of the MaxCompute project or DataHub project.
	//
	// If the destination instance is an Alibaba Cloud AnalyticDB for MySQL database, set this parameter to the cluster ID of the AnalyticDB for MySQL database.
	//
	// example:
	//
	// rm-bp1r46452ai50****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The destination instance type. Valid values:
	//
	// - **Redis**: Alibaba Cloud Redis instance.
	//
	// - **RDS**: ApsaraDB RDS instance.
	//
	// - **PolarDB**: PolarDB cluster (only MySQL or Oracle-compatible engines are supported).
	//
	// - **ECS**: Self-managed database hosted on ECS.
	//
	// - **Express**: Self-managed database connected over Express Connect.
	//
	// - **DataHub**: DataHub instance.
	//
	// - **MaxCompute**: MaxCompute instance.
	//
	// - **AnalyticDB**: AnalyticDB for MySQL 3.0 and 2.0.
	//
	// - **Greenplum**: AnalyticDB for PostgreSQL.
	//
	// > Default value: **RDS**.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The password of the database account for the destination database.
	//
	// > - This parameter is required when the **DestinationEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// Test654321
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port number of the destination database.
	//
	// > This parameter is available and required only when the **DestinationEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The database account of the destination database.
	//
	// > - The permissions required for synchronizing different databases vary. For more information, see the corresponding configuration examples in [Overview of DTS data synchronization solutions](https://help.aliyun.com/document_detail/140954.html).
	//
	// - This parameter is required when the **DestinationEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// - This parameter and the **DestinationEndpoint.Password*	- parameter are not required when the **DestinationEndpoint.InstanceType*	- parameter is set to **RDS*	- and the database engine version is MySQL 5.5 or MySQL 5.6.
	//
	// - This parameter is not required when the **DestinationEndpoint.InstanceType*	- parameter is set to **Redis**.
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
	// Specifies whether the incremental log table contains partitions defined by the day of the incremental update time. Valid values: **true*	- and **false**.
	//
	// > This parameter is available only when the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeDay *bool `json:"ModifyTime_Day,omitempty" xml:"ModifyTime_Day,omitempty"`
	// Specifies whether the incremental log table contains partitions defined by the hour of the incremental update time. Valid values: **true*	- and **false**.
	//
	// > This parameter is available only when the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeHour *bool `json:"ModifyTime_Hour,omitempty" xml:"ModifyTime_Hour,omitempty"`
	// Specifies whether the incremental log table contains partitions defined by the minute of the incremental update time. Valid values: **true*	- and **false**.
	//
	// > This parameter is available only when the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeMinute *bool `json:"ModifyTime_Minute,omitempty" xml:"ModifyTime_Minute,omitempty"`
	// Specifies whether the incremental log table contains partitions defined by the month of the incremental update time. Valid values: **true*	- and **false**.
	//
	// > This parameter is available only when the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeMonth *bool `json:"ModifyTime_Month,omitempty" xml:"ModifyTime_Month,omitempty"`
	// Specifies whether the incremental log table contains partitions defined by the year of the incremental update time. Valid values: **true*	- and **false**.
	//
	// > This parameter is available only when the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
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
	// The name of the database to which the synchronization objects in the source instance belong.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The IP address of the source database.
	//
	// > This parameter is available and required only when the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The instance ID of the source instance.
	//
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the source instance. Valid values:
	//
	// - **RDS**: ApsaraDB RDS instance.
	//
	// - **Redis**: Alibaba Cloud Redis instance.
	//
	// - **PolarDB**: PolarDB cluster (only MySQL or Oracle-compatible engines are supported).
	//
	// - **ECS**: Self-managed database hosted on Elastic Compute Service (ECS).
	//
	// - **Express**: Self-managed database connected over Express Connect.
	//
	// - **dg**: Self-managed database connected over Database Gateway.
	//
	// - **cen**: Self-managed database connected over Cloud Enterprise Network (CEN).
	//
	// > Default value: **RDS**.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The Alibaba Cloud account ID to which the source ApsaraDB RDS instance belongs.
	//
	// > Specifying this parameter indicates that cross-Alibaba Cloud account data synchronization is executed. You must also specify the **SourceEndpoint.Role*	- parameter.
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The password of the database account for the source database.
	//
	// > This parameter is required only when the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port number of the source database.
	//
	// > This parameter is available and required only when the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// > This parameter is required for cross-account data synchronization. For more information about the permissions required for the role and how to grant the permissions, see [Configure RAM authorization for cross-account data migration or synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The database account of the source database.
	//
	// > - This parameter is available and required only when the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// - This parameter is not required when the **SourceEndpoint.InstanceType*	- parameter is set to **Redis**.
	//
	// - The permissions required for synchronizing different databases vary. For more information, see the corresponding configuration examples in [Overview of DTS data synchronization solutions](https://help.aliyun.com/document_detail/140954.html).
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
