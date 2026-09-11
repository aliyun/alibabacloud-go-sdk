// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDtsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateDtsInstanceRequest
	GetAutoPay() *bool
	SetAutoStart(v bool) *CreateDtsInstanceRequest
	GetAutoStart() *bool
	SetComputeUnit(v int32) *CreateDtsInstanceRequest
	GetComputeUnit() *int32
	SetDatabaseCount(v int32) *CreateDtsInstanceRequest
	GetDatabaseCount() *int32
	SetDestinationEndpointEngineName(v string) *CreateDtsInstanceRequest
	GetDestinationEndpointEngineName() *string
	SetDestinationRegion(v string) *CreateDtsInstanceRequest
	GetDestinationRegion() *string
	SetDtsRegion(v string) *CreateDtsInstanceRequest
	GetDtsRegion() *string
	SetDu(v int32) *CreateDtsInstanceRequest
	GetDu() *int32
	SetFeeType(v string) *CreateDtsInstanceRequest
	GetFeeType() *string
	SetInsightModule(v bool) *CreateDtsInstanceRequest
	GetInsightModule() *bool
	SetInstanceClass(v string) *CreateDtsInstanceRequest
	GetInstanceClass() *string
	SetJobId(v string) *CreateDtsInstanceRequest
	GetJobId() *string
	SetMaxDu(v float64) *CreateDtsInstanceRequest
	GetMaxDu() *float64
	SetMinDu(v float64) *CreateDtsInstanceRequest
	GetMinDu() *float64
	SetPayType(v string) *CreateDtsInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDtsInstanceRequest
	GetPeriod() *string
	SetQuantity(v int32) *CreateDtsInstanceRequest
	GetQuantity() *int32
	SetRegionId(v string) *CreateDtsInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDtsInstanceRequest
	GetResourceGroupId() *string
	SetSourceEndpointEngineName(v string) *CreateDtsInstanceRequest
	GetSourceEndpointEngineName() *string
	SetSourceRegion(v string) *CreateDtsInstanceRequest
	GetSourceRegion() *string
	SetSyncArchitecture(v string) *CreateDtsInstanceRequest
	GetSyncArchitecture() *string
	SetType(v string) *CreateDtsInstanceRequest
	GetType() *string
	SetUsedTime(v int32) *CreateDtsInstanceRequest
	GetUsedTime() *int32
}

type CreateDtsInstanceRequest struct {
	// Specifies whether to enable auto-renewal upon expiration. Valid values:
	//
	// - **false**: no. This is the default value.
	//
	// - **true**: yes.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to automatically start the task after the purchase is complete. Valid values:
	//
	// - **false**: no. This is the default value.
	//
	// - **true**: yes.
	//
	// > This parameter takes effect only when **JobId*	- is set to a valid task ID and this parameter is set to **true**.
	//
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// The specifications of the ETL instance. Unit: compute unit (CU). 1 CU = 1 vCPU + 4 GB memory. Valid values: integers that are greater than or equal to 2.
	//
	// <props="china">
	//
	// > If you specify this parameter, the [ETL feature](https://help.aliyun.com/document_detail/212324.html) is enabled for data cleaning and transformation..
	//
	// example:
	//
	// 5
	ComputeUnit *int32 `json:"ComputeUnit,omitempty" xml:"ComputeUnit,omitempty"`
	// The number of private custom ApsaraDB RDS instances under PolarDB-X. Default value: **1**.
	//
	// > This parameter is required only when **SourceEndpointEngineName*	- is set to **drds**.
	//
	// example:
	//
	// 3
	DatabaseCount *int32 `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	// The database engine type of the destination instance. Valid values:
	//
	// - **MySQL**: MySQL database, including ApsaraDB RDS for MySQL and self-managed MySQL.
	//
	// - **PolarDB**: PolarDB for MySQL.
	//
	// - **polardb_o**: PolarDB for Oracle.
	//
	// - **polardb_pg**: PolarDB for PostgreSQL.
	//
	// - **Redis**: Redis database, including Tair (Redis® OSS-Compatible) and self-managed Redis.
	//
	// - **DRDS**: cloud-native distributed database PolarDB-X 1.0 and 2.0.
	//
	// - **PostgreSQL**: self-managed PostgreSQL.
	//
	// - **odps**: MaxCompute.
	//
	// - **oracle**: self-managed Oracle.
	//
	// - **mongodb**: MongoDB database, including ApsaraDB for MongoDB and self-managed MongoDB.
	//
	// - **tidb**: TiDB database.
	//
	// - **ADS**: AnalyticDB for MySQL 2.0.
	//
	// - **ADB30**: AnalyticDB for MySQL 3.0.
	//
	// - **Greenplum**: AnalyticDB for PostgreSQL.
	//
	// - **MSSQL**: SQL Server database, including ApsaraDB RDS for SQL Server and self-managed SQL Server.
	//
	// - **kafka**: Kafka database, including ApsaraMQ for Kafka and self-managed Kafka.
	//
	// - **DataHub**: Alibaba Cloud DataHub.
	//
	// - **DB2**: self-managed Db2 for LUW.
	//
	// - **as400**: AS/400.
	//
	// - **Tablestore**: Tablestore.
	//
	// > - Default value: **MySQL**.
	//
	// - For more information about the supported source and destination database combinations, see [Databases, initial synchronization types, and synchronization topologies](https://help.aliyun.com/document_detail/130744.html) and [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
	//
	// - You must specify this parameter or **JobId**.
	//
	// example:
	//
	// MySQL
	DestinationEndpointEngineName *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	// The region of the destination instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > You must specify this parameter or **JobId**.
	//
	// example:
	//
	// cn-hangzhou
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// The region to which the instance belongs. The value must be the same as the value of **RegionId**.
	//
	// example:
	//
	// cn-hangzhou
	DtsRegion *string `json:"DtsRegion,omitempty" xml:"DtsRegion,omitempty"`
	// The number of DU resources to allocate to the DTS task on a DTS dedicated cluster. Valid values: **1*	- to **100**.
	//
	// > - The value must be within the range of available DUs in the DTS dedicated cluster.
	//
	// - For more information about DTS dedicated clusters, see [What is a DTS dedicated cluster](https://help.aliyun.com/document_detail/417481.html).
	//
	// example:
	//
	// 30
	Du *int32 `json:"Du,omitempty" xml:"Du,omitempty"`
	// The billing type for change tracking. Valid values: ONLY_CONFIGURATION_FEE, which indicates that only configuration fees are charged and data traffic fees are waived. CONFIGURATION_FEE_AND_DATA_FEE, which indicates that data traffic fees are additionally charged.
	//
	// example:
	//
	// ONLY_CONFIGURATION_FEE
	FeeType       *string `json:"FeeType,omitempty" xml:"FeeType,omitempty"`
	InsightModule *bool   `json:"InsightModule,omitempty" xml:"InsightModule,omitempty"`
	// The specification of the data migration or data synchronization instance.
	//
	// - Specifications supported by data migration instances: **xxlarge**, **xlarge**, **large**, **medium**, and **small**.
	//
	// - Specifications supported by data synchronization instances: **large**, **medium**, **small**, and **micro**.
	//
	// > For more information about the performance of each specification, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html) and [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// xxlarge
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The task ID (**DtsJobId**) obtained by calling the **ConfigureDtsJob*	- operation.
	//
	// > If you specify this parameter, you do not need to specify **SourceRegion**, **DestinationRegion**, **Type**, **SourceEndpointEngineName**, or **DestinationEndpointEngineName**. Even if you specify these parameters, the configurations in **JobId*	- take precedence.
	//
	// example:
	//
	// bi6e22ay243****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The minimum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The billing method. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// > Correction: This parameter is required.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing method of the subscription instance. Valid values: **Year*	- and **Month**.
	//
	// > This parameter is valid and required only when **PayType*	- is set to **PrePaid*	- (subscription).
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The number of instances to purchase.
	//
	// > A maximum of one instance can be purchased per call.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region ID of the instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The database engine type of the source instance. Valid values:
	//
	// - **MySQL**: MySQL database, including ApsaraDB RDS for MySQL and self-managed MySQL.
	//
	// - **PolarDB**: PolarDB for MySQL.
	//
	// - **polardb_o**: PolarDB for Oracle.
	//
	// - **polardb_pg**: PolarDB for PostgreSQL.
	//
	// - **Redis**: Redis database, including Tair (Redis® OSS-Compatible) and self-managed Redis.
	//
	// - **DRDS**: cloud-native distributed database PolarDB-X 1.0 and 2.0.
	//
	// - **PostgreSQL**: self-managed PostgreSQL.
	//
	// - **odps**: MaxCompute.
	//
	// - **oracle**: self-managed Oracle.
	//
	// - **mongodb**: MongoDB database, including ApsaraDB for MongoDB and self-managed MongoDB.
	//
	// - **tidb**: TiDB database.
	//
	// - **ADS**: AnalyticDB for MySQL 2.0.
	//
	// - **ADB30**: AnalyticDB for MySQL 3.0.
	//
	// - **Greenplum**: AnalyticDB for PostgreSQL.
	//
	// - **MSSQL**: SQL Server database, including ApsaraDB RDS for SQL Server and self-managed SQL Server.
	//
	// - **kafka**: Kafka database, including ApsaraMQ for Kafka and self-managed Kafka.
	//
	// - **DataHub**: Alibaba Cloud DataHub.
	//
	// - **DB2**: self-managed Db2 for LUW.
	//
	// - **as400**: AS/400.
	//
	// - **Tablestore**: Tablestore.
	//
	// - **OceanBase**: OceanBase (MySQL). Only data migration instances are supported.
	//
	// > - Default value: **MySQL**.
	//
	// - For more information about the supported source and destination database combinations, see [Databases, initial synchronization types, and synchronization topologies](https://help.aliyun.com/document_detail/130744.html) and [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
	//
	// - You must specify this parameter or **JobId**.
	//
	// example:
	//
	// MySQL
	SourceEndpointEngineName *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	// The region of the source instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > You must specify this parameter or **JobId**.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The synchronization topology. Valid values:
	//
	// - **oneway**: one-way synchronization. This is the default value.
	//
	// - **bidirectional**: two-way synchronization.
	//
	// example:
	//
	// oneway
	SyncArchitecture *string `json:"SyncArchitecture,omitempty" xml:"SyncArchitecture,omitempty"`
	// The instance type. Valid values:
	//
	// - **MIGRATION**: data migration.
	//
	// - **SYNC**: data synchronization.
	//
	// - **SUBSCRIBE**: change tracking.
	//
	// > You must specify this parameter or **JobId**.
	//
	// example:
	//
	// SYNC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The subscription duration of the subscription instance.
	//
	// - If **Period*	- is set to **Month**, valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// - If **Period*	- is set to **Year**, valid values are 1, 2, 3, and 5.
	//
	// > - This parameter is valid and required only when **PayType*	- is set to **PrePaid*	- (subscription).
	//
	// - You can set the billing method of the subscription instance by using the **Period*	- parameter.
	//
	// example:
	//
	// 5
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s CreateDtsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDtsInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDtsInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateDtsInstanceRequest) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *CreateDtsInstanceRequest) GetComputeUnit() *int32 {
	return s.ComputeUnit
}

func (s *CreateDtsInstanceRequest) GetDatabaseCount() *int32 {
	return s.DatabaseCount
}

func (s *CreateDtsInstanceRequest) GetDestinationEndpointEngineName() *string {
	return s.DestinationEndpointEngineName
}

func (s *CreateDtsInstanceRequest) GetDestinationRegion() *string {
	return s.DestinationRegion
}

func (s *CreateDtsInstanceRequest) GetDtsRegion() *string {
	return s.DtsRegion
}

func (s *CreateDtsInstanceRequest) GetDu() *int32 {
	return s.Du
}

func (s *CreateDtsInstanceRequest) GetFeeType() *string {
	return s.FeeType
}

func (s *CreateDtsInstanceRequest) GetInsightModule() *bool {
	return s.InsightModule
}

func (s *CreateDtsInstanceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateDtsInstanceRequest) GetJobId() *string {
	return s.JobId
}

func (s *CreateDtsInstanceRequest) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *CreateDtsInstanceRequest) GetMinDu() *float64 {
	return s.MinDu
}

func (s *CreateDtsInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDtsInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDtsInstanceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *CreateDtsInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDtsInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDtsInstanceRequest) GetSourceEndpointEngineName() *string {
	return s.SourceEndpointEngineName
}

func (s *CreateDtsInstanceRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *CreateDtsInstanceRequest) GetSyncArchitecture() *string {
	return s.SyncArchitecture
}

func (s *CreateDtsInstanceRequest) GetType() *string {
	return s.Type
}

func (s *CreateDtsInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateDtsInstanceRequest) SetAutoPay(v bool) *CreateDtsInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetAutoStart(v bool) *CreateDtsInstanceRequest {
	s.AutoStart = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetComputeUnit(v int32) *CreateDtsInstanceRequest {
	s.ComputeUnit = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDatabaseCount(v int32) *CreateDtsInstanceRequest {
	s.DatabaseCount = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDestinationEndpointEngineName(v string) *CreateDtsInstanceRequest {
	s.DestinationEndpointEngineName = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDestinationRegion(v string) *CreateDtsInstanceRequest {
	s.DestinationRegion = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDtsRegion(v string) *CreateDtsInstanceRequest {
	s.DtsRegion = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDu(v int32) *CreateDtsInstanceRequest {
	s.Du = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetFeeType(v string) *CreateDtsInstanceRequest {
	s.FeeType = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetInsightModule(v bool) *CreateDtsInstanceRequest {
	s.InsightModule = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetInstanceClass(v string) *CreateDtsInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetJobId(v string) *CreateDtsInstanceRequest {
	s.JobId = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetMaxDu(v float64) *CreateDtsInstanceRequest {
	s.MaxDu = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetMinDu(v float64) *CreateDtsInstanceRequest {
	s.MinDu = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetPayType(v string) *CreateDtsInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetPeriod(v string) *CreateDtsInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetQuantity(v int32) *CreateDtsInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetRegionId(v string) *CreateDtsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetResourceGroupId(v string) *CreateDtsInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetSourceEndpointEngineName(v string) *CreateDtsInstanceRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetSourceRegion(v string) *CreateDtsInstanceRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetSyncArchitecture(v string) *CreateDtsInstanceRequest {
	s.SyncArchitecture = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetType(v string) *CreateDtsInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetUsedTime(v int32) *CreateDtsInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDtsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
