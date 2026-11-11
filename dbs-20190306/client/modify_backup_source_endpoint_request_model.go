// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSourceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupGatewayId(v int64) *ModifyBackupSourceEndpointRequest
	GetBackupGatewayId() *int64
	SetBackupObjects(v string) *ModifyBackupSourceEndpointRequest
	GetBackupObjects() *string
	SetBackupPlanId(v string) *ModifyBackupSourceEndpointRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *ModifyBackupSourceEndpointRequest
	GetClientToken() *string
	SetCrossAliyunId(v string) *ModifyBackupSourceEndpointRequest
	GetCrossAliyunId() *string
	SetCrossRoleName(v string) *ModifyBackupSourceEndpointRequest
	GetCrossRoleName() *string
	SetEnableSourceEndpointSsl(v string) *ModifyBackupSourceEndpointRequest
	GetEnableSourceEndpointSsl() *string
	SetOwnerId(v string) *ModifyBackupSourceEndpointRequest
	GetOwnerId() *string
	SetSourceEndpointDatabaseName(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointDatabaseName() *string
	SetSourceEndpointIP(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointInstanceID(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointInstanceType(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointInstanceType() *string
	SetSourceEndpointOracleHome(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointOracleHome() *string
	SetSourceEndpointOracleSID(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointOracleSID() *string
	SetSourceEndpointPassword(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointPort(v int32) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointPort() *int32
	SetSourceEndpointRegion(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointUserName(v string) *ModifyBackupSourceEndpointRequest
	GetSourceEndpointUserName() *string
	SetSslCaPem(v string) *ModifyBackupSourceEndpointRequest
	GetSslCaPem() *string
}

type ModifyBackupSourceEndpointRequest struct {
	// The backup gateway ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to `Agent`.
	//
	// example:
	//
	// 21321323213
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// Backup objects: This parameter is optional if the database is located on an agent, but required for all other scenarios. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the value of this parameter.
	//
	// example:
	//
	// [{  "DBName":"待备份库名", "SchemaName":"待备份 Schema 名", "TableIncludes":[{ "TableName":"待备份表表名" }],  "TableExcludes":[{"TableName":"待备份库名不需要备份表的表名" }] } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The backup plan ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1h****usfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// A client-generated token to ensure request idempotence.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The UID for cross-account backup.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The RAM role name for cross-account backup.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// Specifies whether to use an SSL connection to back up the source database.
	//
	// example:
	//
	// true
	EnableSourceEndpointSsl *string `json:"EnableSourceEndpointSsl,omitempty" xml:"EnableSourceEndpointSsl,omitempty"`
	OwnerId                 *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The database name.
	//
	// - This parameter is required for PostgreSQL and MongoDB databases.
	//
	// - This parameter is also required for Microsoft SQL Server databases when the **SourceEndpointInstanceType*	- parameter is set to \\"Agent\\".
	//
	// example:
	//
	// test
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The IP address of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to `Express`, `Agent`, or `Other`.
	//
	// example:
	//
	// 100.*.*.10:3306
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The database instance ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to `RDS`, `ECS`, `DDS`, or `Express`.
	//
	// example:
	//
	// rm-bp1p8c29479jv****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The type of the source instance. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value. Valid values:
	//
	// - **RDS**: an RDS instance
	//
	// - **ECS**: a self-managed database on an ECS instance
	//
	// - **Express**: a database connected over a leased line, VPN Gateway, or Smart Access Gateway
	//
	// - **Agent**: a database connected over a backup gateway
	//
	// - **DDS**: an ApsaraDB for MongoDB instance
	//
	// - **Other**: a database connected over an IP address and port
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The Oracle Home path.
	//
	// example:
	//
	// /home/test
	SourceEndpointOracleHome *string `json:"SourceEndpointOracleHome,omitempty" xml:"SourceEndpointOracleHome,omitempty"`
	// The Oracle SID. This parameter is required for Oracle databases.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password for the database account.
	//
	// This parameter is required for all database engines except for Redis and Microsoft SQL Server (when the **SourceEndpointInstanceType*	- parameter is set to \\"Agent\\").
	//
	// example:
	//
	// test
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The database port. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to `Express`, `Agent`, `Other`, or `ECS`.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The database region. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to `RDS`, `ECS`, `DDS`, `Express`, or `Agent`.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The username of the database account. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to get this value.
	//
	// This parameter is required for all database engines except for Redis and Microsoft SQL Server (when the **SourceEndpointInstanceType*	- parameter is set to \\"Agent\\").
	//
	// example:
	//
	// test
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The PEM certificate for the SSL connection.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- ... -----END CERTIFICATE-----
	SslCaPem *string `json:"SslCaPem,omitempty" xml:"SslCaPem,omitempty"`
}

func (s ModifyBackupSourceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSourceEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupSourceEndpointRequest) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *ModifyBackupSourceEndpointRequest) GetBackupObjects() *string {
	return s.BackupObjects
}

func (s *ModifyBackupSourceEndpointRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyBackupSourceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBackupSourceEndpointRequest) GetCrossAliyunId() *string {
	return s.CrossAliyunId
}

func (s *ModifyBackupSourceEndpointRequest) GetCrossRoleName() *string {
	return s.CrossRoleName
}

func (s *ModifyBackupSourceEndpointRequest) GetEnableSourceEndpointSsl() *string {
	return s.EnableSourceEndpointSsl
}

func (s *ModifyBackupSourceEndpointRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointOracleHome() *string {
	return s.SourceEndpointOracleHome
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointPort() *int32 {
	return s.SourceEndpointPort
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *ModifyBackupSourceEndpointRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *ModifyBackupSourceEndpointRequest) GetSslCaPem() *string {
	return s.SslCaPem
}

func (s *ModifyBackupSourceEndpointRequest) SetBackupGatewayId(v int64) *ModifyBackupSourceEndpointRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetBackupObjects(v string) *ModifyBackupSourceEndpointRequest {
	s.BackupObjects = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetBackupPlanId(v string) *ModifyBackupSourceEndpointRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetClientToken(v string) *ModifyBackupSourceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetCrossAliyunId(v string) *ModifyBackupSourceEndpointRequest {
	s.CrossAliyunId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetCrossRoleName(v string) *ModifyBackupSourceEndpointRequest {
	s.CrossRoleName = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetEnableSourceEndpointSsl(v string) *ModifyBackupSourceEndpointRequest {
	s.EnableSourceEndpointSsl = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetOwnerId(v string) *ModifyBackupSourceEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointDatabaseName(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointIP(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointInstanceID(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointInstanceType(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointOracleHome(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointOracleHome = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointOracleSID(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointPassword(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointPort(v int32) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointRegion(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointUserName(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSslCaPem(v string) *ModifyBackupSourceEndpointRequest {
	s.SslCaPem = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
