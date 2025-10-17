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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
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
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  Default value: **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet special requirements, for example, whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// >  This parameter can be used for data synchronization between ApsaraDB for Redis Enterprise Edition instances. For more information, see [Use OpenAPI Explorer to configure one-way or two-way data synchronization between ApsaraDB for Redis Enterprise Edition instances](https://help.aliyun.com/document_detail/155967.html).
	//
	// example:
	//
	// {     "autoStartModulesAfterConfig": "none",     "targetTableMode": 2 }
	MigrationReserved *string `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组ID。
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to perform initial schema synchronization. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  Default value: **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// >
	//
	// 	- Default value: **Forward**.
	//
	// 	- The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
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
	// The name of the data synchronization task.
	//
	// >  We recommend that you specify an informative name for easy identification. You do not need to use a unique task name.
	//
	// example:
	//
	// MySQL同步
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	// The objects that you want to synchronize. The value is a JSON string and can contain regular expressions. For more information, see [SynchronizationObjects](https://help.aliyun.com/document_detail/141901.html).
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
	// The name of the database to which the synchronization object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// The IP address of the destination database.
	//
	// >  You must specify this parameter only if the **DestinationEndpoint.InstanceType*	- parameter is set to **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the destination instance.
	//
	// >  If the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute*	- or **DataHub**, you must specify the name of the MaxCompute project or the DataHub project.
	//
	// If the destination instance is an AnalyticDB for MySQL cluster, specify the ID of the AnalyticDB for MySQL cluster.
	//
	// example:
	//
	// rm-bp1r46452ai50****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the destination instance. Valid values:
	//
	// 	- **Redis**: ApsaraDB for Redis instance
	//
	// 	- **RDS**: ApsaraDB RDS instance
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster or PolarDB O Edition cluster
	//
	// 	- **ECS**: self-managed database that is hosted on ECS
	//
	// 	- **Express**: self-managed database that is connected over Express Connect
	//
	// 	- **DataHub**: DataHub project
	//
	// 	- **MaxCompute**: MaxCompute project
	//
	// 	- **AnalyticDB**: AnalyticDB for MySQL cluster V3.0 or V2.0
	//
	// 	- **Greenplum**: AnalyticDB for PostgreSQL instance
	//
	// >  The default value is **RDS**.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The password of the destination database account.
	//
	// >
	//
	// 	- If the **DestinationEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**, you must specify the DestinationEndpoint.Password parameter.
	//
	// example:
	//
	// Test654321
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port number of the destination database.
	//
	// >  You must specify this parameter only if the **DestinationEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The database account of the destination database.
	//
	// >
	//
	// 	- The permissions that are required for database accounts vary with the synchronization scenario. For more information, see [Overview of data synchronization scenarios](https://help.aliyun.com/document_detail/140954.html).
	//
	// 	- If the **DestinationEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**, you must specify the DestinationEndpoint.UserName parameter.
	//
	// 	- If the **DestinationEndpoint.InstanceType*	- parameter is set to RDS and the database version is MySQL 5.5 or MySQL 5.6, you do not need to specify the DestinationEndpoint.UserName and **DestinationEndpoint.Password*	- parameters.
	//
	// 	- If the **DestinationEndpoint.InstanceType*	- parameter is set to **Redis**, you do not need to specify the DestinationEndpoint.UserName parameter.
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
	// Specifies whether the incremental data table contains partitions defined by the modifytime_day field. Valid values: **true*	- and **false**.
	//
	// >  This parameter is available only if the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeDay *bool `json:"ModifyTime_Day,omitempty" xml:"ModifyTime_Day,omitempty"`
	// Specifies whether the incremental data table contains partitions defined by the modifytime_hour field. Valid values: **true*	- and **false**.
	//
	// >  This parameter is available only if the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeHour *bool `json:"ModifyTime_Hour,omitempty" xml:"ModifyTime_Hour,omitempty"`
	// Specifies whether the incremental data table contains partitions defined by the modifytime_minute field. Valid values: **true*	- and **false**.
	//
	// >  This parameter is available only if the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeMinute *bool `json:"ModifyTime_Minute,omitempty" xml:"ModifyTime_Minute,omitempty"`
	// Specifies whether the incremental data table contains partitions defined by the modifytime_month field. Valid values: **true*	- and **false**.
	//
	// >  This parameter is available only if the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
	//
	// example:
	//
	// true
	ModifyTimeMonth *bool `json:"ModifyTime_Month,omitempty" xml:"ModifyTime_Month,omitempty"`
	// Specifies whether the incremental data table contains partitions defined by the modifytime_year field. Valid values: **true*	- and **false**.
	//
	// >  This parameter is available only if the **DestinationEndpoint.InstanceType*	- parameter is set to **MaxCompute**.
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
	// The name of the database to which the synchronization object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The IP address of the source database.
	//
	// >  You must specify this parameter only if the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the source instance. Valid values:
	//
	// 	- **RDS**: ApsaraDB RDS instance
	//
	// 	- **Redis**: ApsaraDB for Redis instance
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster or PolarDB O Edition cluster
	//
	// 	- **ECS**: self-managed database that is hosted on Elastic Compute Service (ECS)
	//
	// 	- **Express**: self-managed database that is connected over Express Connect
	//
	// 	- **dg**: self-managed database that is connected over Database Gateway
	//
	// 	- **cen**: self-managed database that is connected over Cloud Enterprise Network (CEN)
	//
	// >  The default value is **RDS**.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the Alibaba Cloud account that owns the source RDS instance.
	//
	// >  You can specify this parameter to synchronize data across different Alibaba Cloud accounts. In this case, you also need to specify the **SourceEndpoint.Role*	- parameter.
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The password of the source database account.
	//
	// >  You must specify this parameter only if the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port number of the source database.
	//
	// >  You must specify this parameter only if the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account that owns the source instance.
	//
	// >  You must specify this parameter when you synchronize data across different Alibaba Cloud accounts. For information about the permissions and authorization methods of the RAM role, see [Configure RAM authorization for cross-account data migration and synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The database account of the source database.
	//
	// >
	//
	// 	- You must specify this parameter only if the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **Express**, **dg**, or **cen**.
	//
	// 	- If the **SourceEndpoint.InstanceType*	- parameter is set to **Redis**, you do not need to specify the database account.
	//
	// 	- The permissions that are required for database accounts vary with the synchronization scenario. For more information, see [Overview of data synchronization scenarios](https://help.aliyun.com/document_detail/140954.html).
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
