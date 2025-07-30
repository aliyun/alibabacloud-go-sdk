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
	// Specifies whether to automatically renew the DTS instance when it expires. Valid values:
	//
	// 	- **false**: does not automatically renew the DTS instance when it expires. This is the default value.
	//
	// 	- **true**: automatically renews the DTS instance when it expires.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to automatically start the task after the DTS instance is purchased. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// >  This parameter can be set to **true*	- and take effect only if you specify a valid value for **JobId**.
	//
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// The specification of the extract, transform, and load (ETL) instance. The unit is compute unit (CU). One CU is equal to 1 vCPU and 4 GB of memory. The value of this parameter must be an integer greater than or equal to 2.
	//
	// example:
	//
	// 5
	ComputeUnit *int32 `json:"ComputeUnit,omitempty" xml:"ComputeUnit,omitempty"`
	// The number of custom ApsaraDB RDS instances in the PolarDB-X instance. Default value: **1**.
	//
	// >  This parameter is required only if **SourceEndpointEngineName*	- is set to **drds**.
	//
	// example:
	//
	// 3
	DatabaseCount *int32 `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	// The database engine of the destination instance.
	//
	// 	- **MySQL**: ApsaraDB RDS for MySQL instance or self-managed MySQL database
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster
	//
	// 	- **polardb_o**: PolarDB for Oracle cluster
	//
	// 	- **polardb_pg**: PolarDB for PostgreSQL cluster
	//
	// 	- **Redis**: ApsaraDB for Redis instance or self-managed Redis database
	//
	// 	- **DRDS**: PolarDB-X 1.0 or PolarDB-X 2.0 instance
	//
	// 	- **PostgreSQL**: self-managed PostgreSQL database
	//
	// 	- **odps**: MaxCompute project
	//
	// 	- **oracle**: self-managed Oracle database
	//
	// 	- **mongodb**: ApsaraDB for MongoDB instance or self-managed MongoDB database
	//
	// 	- **tidb**: TiDB database
	//
	// 	- **ADS**: AnalyticDB for MySQL V2.0 cluster
	//
	// 	- **ADB30**: AnalyticDB for MySQL V3.0 cluster
	//
	// 	- **Greenplum**: AnalyticDB for PostgreSQL instance
	//
	// 	- **MSSQL**: ApsaraDB RDS for SQL Server instance or self-managed SQL Server database
	//
	// 	- **kafka**: Message Queue for Apache Kafka instance or self-managed Kafka cluster
	//
	// 	- **DataHub**: DataHub project
	//
	// 	- **DB2**: self-managed Db2 for LUW database
	//
	// 	- **as400**: AS/400
	//
	// 	- **Tablestore**: Tablestore instance
	//
	// >
	//
	// 	- The default value is **MySQL**.
	//
	// 	- For more information about the supported source and destination databases, see [Overview of data synchronization scenarios](https://help.aliyun.com/document_detail/130744.html) and [Overview of data migration scenarios](https://help.aliyun.com/document_detail/26618.html).
	//
	// 	- You must specify one of this parameter and the **JobId*	- parameter.
	//
	// example:
	//
	// MySQL
	DestinationEndpointEngineName *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// >  You must specify one of this parameter and the **JobId*	- parameter.
	//
	// example:
	//
	// cn-hangzhou
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// The region ID of the DTS instance. Set this parameter to the value of **RegionId**.
	//
	// example:
	//
	// cn-hangzhou
	DtsRegion *string `json:"DtsRegion,omitempty" xml:"DtsRegion,omitempty"`
	// The number of DTS units (DUs) that are assigned to a DTS task that is run on a DTS dedicated cluster. Valid values: **1*	- to **100**.
	//
	// >
	//
	// 	- The value of this parameter must be within the range of the number of DUs available for the DTS dedicated cluster.
	//
	// example:
	//
	// 30
	Du *int32 `json:"Du,omitempty" xml:"Du,omitempty"`
	// The billing type for a change tracking instance. Valid values: ONLY_CONFIGURATION_FEE and CONFIGURATION_FEE_AND_DATA_FEE. ONLY_CONFIGURATION_FEE: charges only configuration fees. CONFIGURATION_FEE_AND_DATA_FEE: charges configuration fees and data traffic fees.
	//
	// example:
	//
	// ONLY_CONFIGURATION_FEE
	FeeType *string `json:"FeeType,omitempty" xml:"FeeType,omitempty"`
	// The instance class.
	//
	// 	- DTS supports the following instance classes for a data migration instance: **xxlarge**, **xlarge**, **large**, **medium**, and **small**.
	//
	// 	- DTS supports the following instance classes for a data synchronization instance: **large**, **medium**, **small**, and **micro**.
	//
	// >  For more information about the test performance of each instance class, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html) and [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// xxlarge
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The ID of the task. You can call the **ConfigureDtsJob*	- operation to obtain the task ID from the **DtsJobId*	- parameter.
	//
	// >  If this parameter is specified, you do not need to specify the **SourceRegion**, **DestinationRegion**, **Type**, **SourceEndpointEngineName**, or **DestinationEndpointEngineName*	- parameter. Even if these parameters are specified, the value of the **JobId*	- parameter takes precedence.
	//
	// example:
	//
	// bi6e22ay243****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Upper limit of DU.
	//
	// > Only supported by Serverless instances.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// Lower limit of DU.
	//
	// > Only supported by Serverless instances.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// >  This parameter must be specified.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values: **Year*	- and **Month**.
	//
	// >  You must specify this parameter only if the **PayType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The number of DTS instances that you want to purchase.
	//
	// >  You can purchase only one DTS instance each time you call this operation.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The database engine of the source instance.
	//
	// 	- **MySQL**: ApsaraDB RDS for MySQL instance or self-managed MySQL database
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster
	//
	// 	- **polardb_o**: PolarDB for Oracle cluster
	//
	// 	- **polardb_pg**: PolarDB for PostgreSQL cluster
	//
	// 	- **Redis**: ApsaraDB for Redis instance or self-managed Redis database
	//
	// 	- **DRDS**: PolarDB-X 1.0 or PolarDB-X 2.0 instance
	//
	// 	- **PostgreSQL**: self-managed PostgreSQL database
	//
	// 	- **odps**: MaxCompute project
	//
	// 	- **oracle**: self-managed Oracle database
	//
	// 	- **mongodb**: ApsaraDB for MongoDB instance or self-managed MongoDB database
	//
	// 	- **tidb**: TiDB database
	//
	// 	- **ADS**: AnalyticDB for MySQL V2.0 cluster
	//
	// 	- **ADB30**: AnalyticDB for MySQL V3.0 cluster
	//
	// 	- **Greenplum**: AnalyticDB for PostgreSQL instance
	//
	// 	- **MSSQL**: ApsaraDB RDS for SQL Server instance or self-managed SQL Server database
	//
	// 	- **kafka**: Message Queue for Apache Kafka instance or self-managed Kafka cluster
	//
	// 	- **DataHub**: DataHub project
	//
	// 	- **DB2**: self-managed Db2 for LUW database
	//
	// 	- **as400**: AS/400
	//
	// 	- **Tablestore**: Tablestore instance
	//
	// >
	//
	// 	- The default value is **MySQL**.
	//
	// 	- For more information about the supported source and destination databases, see [Overview of data synchronization scenarios](https://help.aliyun.com/document_detail/130744.html) and [Overview of data migration scenarios](https://help.aliyun.com/document_detail/26618.html).
	//
	// 	- You must specify one of this parameter and the **JobId*	- parameter.
	//
	// example:
	//
	// MYSQL
	SourceEndpointEngineName *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	// The ID of the region in which the source instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// >  You must specify one of this parameter and the **JobId*	- parameter.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The synchronization topology. Valid values:
	//
	// 	- **oneway**: one-way synchronization. This is the default value.
	//
	// 	- **bidirectional**: two-way synchronization.
	//
	// example:
	//
	// oneway
	SyncArchitecture *string `json:"SyncArchitecture,omitempty" xml:"SyncArchitecture,omitempty"`
	// The type of the DTS instance. Valid values:
	//
	// 	- **MIGRATION**: data migration instance
	//
	// 	- **SYNC**: data synchronization instance
	//
	// 	- **SUBSCRIBE**: change tracking instance
	//
	// > You must specify one of this parameter and the **JobId*	- parameter.
	//
	// example:
	//
	// SYNC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The subscription duration.
	//
	// 	- Valid values if **Period*	- is set to **Month**: 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// 	- Valid values if **Period*	- is set to **Year**: 1, 2, 3, and 5.
	//
	// >
	//
	// 	- This parameter is valid and required only if **PayType*	- is set to **PrePaid**.
	//
	// 	- You can configure **Period*	- to specify the unit of the subscription duration.
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
