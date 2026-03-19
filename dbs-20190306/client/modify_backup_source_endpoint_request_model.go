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
	// The backup gateway ID. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// > You must specify this parameter when **SourceEndpointInstanceType*	- is Agent.
	//
	// example:
	//
	// 21321323213
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The backup objects. This parameter is optional when SourceEndpointInstanceType is Agent. For all other cases, you must specify it. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// example:
	//
	// [{  "DBName":"待备份库名", "SchemaName":"待备份 Schema 名", "TableIncludes":[{ "TableName":"待备份表表名" }],  "TableExcludes":[{"TableName":"待备份库名不需要备份表的表名" }] } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The backup plan ID. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1h****usfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// A unique string that ensures idempotence and prevents duplicate requests.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The UID of the Alibaba Cloud account used for cross-account backup. Call [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) to get this value.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The RAM role name used for cross-account backup. Call [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) to get this value.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// example:
	//
	// true
	EnableSourceEndpointSsl *string `json:"EnableSourceEndpointSsl,omitempty" xml:"EnableSourceEndpointSsl,omitempty"`
	OwnerId                 *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The database name.
	//
	// - You must specify this parameter for PostgreSQL or MongoDB databases.
	//
	// - You must specify this parameter for Microsoft SQL Server databases connected through a backup gateway.
	//
	// example:
	//
	// test
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The database endpoint. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// > You must specify this parameter when **SourceEndpointInstanceType*	- is Express, Agent, or Other.
	//
	// example:
	//
	// 100.*.*.10:3306
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The database instance ID. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// > You must specify this parameter when **SourceEndpointInstanceType*	- is RDS, ECS, DDS, or Express.
	//
	// example:
	//
	// rm-bp1p8c29479jv****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The location of the database. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value. Valid values:
	//
	// - **RDS**
	//
	// - **ECS**
	//
	// - **Express**: A database connected through a leased line, VPN Gateway, or Smart Access Gateway
	//
	// - **Agent**: A database connected through a backup gateway
	//
	// - **DDS**: A MongoDB instance on Alibaba Cloud
	//
	// - **Other**: A database connected directly using an IP address and port
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// example:
	//
	// /home/test
	SourceEndpointOracleHome *string `json:"SourceEndpointOracleHome,omitempty" xml:"SourceEndpointOracleHome,omitempty"`
	// The Oracle system ID (SID). You must specify this parameter for Oracle databases.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password.
	//
	// This parameter is optional for Redis databases or for Microsoft SQL Server databases connected through a backup gateway. For all other cases, you must specify it.
	//
	// example:
	//
	// test
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The database port. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// > You must specify this parameter when **SourceEndpointInstanceType*	- is Express, Agent, Other, or ECS.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The region where the database is located. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// > You must specify this parameter when **SourceEndpointInstanceType*	- is RDS, ECS, DDS, Express, or Agent.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The database account. Call [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) to get this value.
	//
	// This parameter is optional for Redis databases or for Microsoft SQL Server databases connected through a backup gateway. For all other cases, you must specify it.
	//
	// example:
	//
	// test
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
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
