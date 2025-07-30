// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationEndpointArchitecture(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointArchitecture() *string
	SetDestinationEndpointDatabaseName(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointDatabaseName() *string
	SetDestinationEndpointEngineName(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointEngineName() *string
	SetDestinationEndpointIP(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointIP() *string
	SetDestinationEndpointInstanceID(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointInstanceID() *string
	SetDestinationEndpointInstanceType(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointInstanceType() *string
	SetDestinationEndpointOracleSID(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointOracleSID() *string
	SetDestinationEndpointPassword(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointPassword() *string
	SetDestinationEndpointPort(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointPort() *string
	SetDestinationEndpointRegion(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointRegion() *string
	SetDestinationEndpointUserName(v string) *DescribeConnectionStatusRequest
	GetDestinationEndpointUserName() *string
	SetRegionId(v string) *DescribeConnectionStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeConnectionStatusRequest
	GetResourceGroupId() *string
	SetSourceEndpointArchitecture(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointArchitecture() *string
	SetSourceEndpointDatabaseName(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointDatabaseName() *string
	SetSourceEndpointEngineName(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointEngineName() *string
	SetSourceEndpointIP(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointInstanceID(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointInstanceType(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointInstanceType() *string
	SetSourceEndpointOracleSID(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointOracleSID() *string
	SetSourceEndpointPassword(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointPort(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointPort() *string
	SetSourceEndpointRegion(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointUserName(v string) *DescribeConnectionStatusRequest
	GetSourceEndpointUserName() *string
}

type DescribeConnectionStatusRequest struct {
	// You must specify this parameter only if the **SourceEndpointEngineName*	- parameter is set to **Oracle**. Valid values:
	//
	// 	- **SID**: non-RAC architecture
	//
	// 	- **RAC**: Real Application Cluster (RAC) architecture
	//
	// >  This parameter is optional. The data type of this parameter is String.
	//
	// example:
	//
	// SID
	DestinationEndpointArchitecture *string `json:"DestinationEndpointArchitecture,omitempty" xml:"DestinationEndpointArchitecture,omitempty"`
	// The name of the destination database or the authentication database.
	//
	// >
	//
	// 	- You must specify this parameter if the **DestinationEndpointEngineName*	- parameter is set to **PostgreSQL**, **DRDS**, or **MongoDB**. You must also specify this parameter if the **DestinationEndpointInstanceType*	- parameter is set to **PolarDB_o**.
	//
	// 	- If the **DestinationEndpointEngineName*	- parameter is set to **PostgreSQL*	- or **DRDS**, specify the name of the destination database. If the DestinationEndpointEngineName parameter is set to **MongoDB**, specify the name of the authentication database.
	//
	// 	- If the **DestinationEndpointInstanceType*	- parameter is set to **PolarDB_o**, specify the name of the destination database.
	//
	// example:
	//
	// dtstestdata
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	// The engine type of the destination database. Valid values: **MySQL**, **DRDS**, **SQLServer**, **PostgreSQL**, **PPAS**, **MongoDB**, and **Redis**.
	//
	// >  You must specify this parameter only if the **DestinationEndpointInstanceType*	- parameter is set to **RDS**, **DRDS**, **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// MySQL
	DestinationEndpointEngineName *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	// The endpoint of the destination database.
	//
	// >  You must specify this parameter only if the **DestinationEndpointInstanceType*	- parameter is set to **LocalInstance*	- or **Express**.
	//
	// example:
	//
	// 172.16.88.***
	DestinationEndpointIP *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// testsid
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The instance type of the destination database. Valid values:
	//
	// >
	//
	// 	- **ECS**: self-managed database that is hosted on Elastic Compute Service (ECS)
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **RDS**: ApsaraDB RDS instance
	//
	// 	- **DRDS**: PolarDB-X instance
	//
	// 	- **MongoDB**: ApsaraDB for MongoDB instance
	//
	// 	- **Redis**: ApsaraDB for Redis instance
	//
	// 	- **PetaData**: HybridDB for MySQL instance
	//
	// 	- **POLARDB**: PolarDB for MySQL cluster
	//
	// 	- **PolarDB_o**: PolarDB for Oracle cluster
	//
	// 	- **AnalyticDB**: AnalyticDB for MySQL cluster V3.0 or V2.0
	//
	// 	- **Greenplum**: AnalyticDB for PostgreSQL instance
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDB_o
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// You must specify this parameter only if the **DestinationEndpointEngineName*	- parameter is set to **Oracle**. Valid values:
	//
	// 	- **SID**: non-RAC architecture
	//
	// 	- **RAC**: RAC architecture
	//
	// >  This parameter is optional. The data type of this parameter is String.
	//
	// example:
	//
	// SID
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// The password of the destination database account.
	//
	// example:
	//
	// Test123456
	DestinationEndpointPassword *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	// The service port number of the source database.
	//
	// >  You must specify this parameter only if the **SourceEndpointInstanceType*	- parameter is set to **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// 3306
	DestinationEndpointPort *string `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	// The ID of the region where the destination instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The database account of the destination database.
	//
	// >  The permissions that are required for database accounts vary with the migration or synchronization scenario. For more information, see [Overview of data migration scenarios](https://help.aliyun.com/document_detail/26618.html) and [Overview of data synchronization scenarios](https://help.aliyun.com/document_detail/130744.html).
	//
	// example:
	//
	// dtstest
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// The ID of the region where the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// You must specify this parameter only if the **SourceEndpointEngineName*	- parameter is set to **Oracle**. Valid values:
	//
	// 	- **SID**: non-RAC architecture
	//
	// 	- **RAC**: RAC architecture
	//
	// >  This parameter is optional.
	//
	// example:
	//
	// SID
	SourceEndpointArchitecture *string `json:"SourceEndpointArchitecture,omitempty" xml:"SourceEndpointArchitecture,omitempty"`
	// The name of the source database or the authentication database.
	//
	// >
	//
	// 	- You must specify this parameter if the **SourceEndpointEngineName*	- parameter is set to **PostgreSQL*	- or **MongoDB**. You must also specify this parameter if the **SourceEndpointInstanceType*	- parameter is set to **PolarDB_o**.
	//
	// 	- If the **SourceEndpointEngineName*	- parameter is set to **PostgreSQL*	- or **DRDS**, specify the name of the source database. If the SourceEndpointEngineName parameter is set to **MongoDB**, specify the name of the authentication database.
	//
	// 	- If the **SourceEndpointInstanceType*	- parameter is set to **PolarDB_o**, specify the name of the source database.
	//
	// example:
	//
	// dtstestdata
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The engine type of the source database. Valid values: **MySQL**, **TiDB**, **SQLServer**, **PostgreSQL**, **Oracle**, **MongoDB**, and **Redis**.
	//
	// >  Default value: **MySQL**.
	//
	// example:
	//
	// MySQL
	SourceEndpointEngineName *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	// The endpoint of the source database.
	//
	// >  You must specify this parameter only if the **SourceEndpointInstanceType*	- parameter is set to **LocalInstance*	- or **Express**.
	//
	// example:
	//
	// 172.16.88.***
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// rm-bp1imrtn6fq7h****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The type of the source instance. Valid values:
	//
	// 	- **RDS**: ApsaraDB RDS instance
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **ECS**: self-managed database that is hosted on ECS
	//
	// 	- **Express**: self-managed database that is connected over Express Connect
	//
	// 	- **dg**: self-managed database that is connected over Database Gateway
	//
	// 	- **MongoDB**: ApsaraDB for MongoDB instance
	//
	// 	- **POLARDB**: PolarDB for MySQL cluster
	//
	// 	- **PolarDB_o**: PolarDB for Oracle cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The SID of the Oracle database.
	//
	// >  You must specify this parameter only if the **SourceEndpointEngineName*	- parameter is set to **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password of the source database account.
	//
	// example:
	//
	// Test123456
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The service port number of the source database.
	//
	// >  You must specify this parameter only if the **SourceEndpointInstanceType*	- parameter is set to **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The ID of the region where the source instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The database account of the source database.
	//
	// >  The permissions that are required for database accounts vary with the migration or synchronization scenario. For more information, see [Overview of data migration scenarios](https://help.aliyun.com/document_detail/26618.html) and [Overview of data synchronization scenarios](https://help.aliyun.com/document_detail/130744.html).
	//
	// example:
	//
	// dtstest
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
}

func (s DescribeConnectionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointArchitecture() *string {
	return s.DestinationEndpointArchitecture
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointDatabaseName() *string {
	return s.DestinationEndpointDatabaseName
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointEngineName() *string {
	return s.DestinationEndpointEngineName
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointIP() *string {
	return s.DestinationEndpointIP
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointInstanceID() *string {
	return s.DestinationEndpointInstanceID
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointInstanceType() *string {
	return s.DestinationEndpointInstanceType
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointOracleSID() *string {
	return s.DestinationEndpointOracleSID
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointPassword() *string {
	return s.DestinationEndpointPassword
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointPort() *string {
	return s.DestinationEndpointPort
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointRegion() *string {
	return s.DestinationEndpointRegion
}

func (s *DescribeConnectionStatusRequest) GetDestinationEndpointUserName() *string {
	return s.DestinationEndpointUserName
}

func (s *DescribeConnectionStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeConnectionStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointArchitecture() *string {
	return s.SourceEndpointArchitecture
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointEngineName() *string {
	return s.SourceEndpointEngineName
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointPort() *string {
	return s.SourceEndpointPort
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *DescribeConnectionStatusRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointArchitecture(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointArchitecture = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointDatabaseName(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointDatabaseName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointEngineName(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointEngineName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointIP(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointIP = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointInstanceID(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointInstanceType(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointOracleSID(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointPassword(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointPort(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointPort = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointRegion(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointUserName(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetRegionId(v string) *DescribeConnectionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetResourceGroupId(v string) *DescribeConnectionStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointArchitecture(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointArchitecture = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointDatabaseName(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointEngineName(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointIP(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointInstanceID(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointInstanceType(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointOracleSID(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointPassword(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointPort(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointRegion(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointUserName(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) Validate() error {
	return dara.Validate(s)
}
