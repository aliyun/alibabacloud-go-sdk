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
}

type ModifyBackupSourceEndpointRequest struct {
	// The backup gateway ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the ID.
	//
	// >  If you set **SourceEndpointInstanceType*	- to Agent, this parameter is required.
	//
	// example:
	//
	// 21321323213
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The backup object. If you set SourceEndpointInstanceType to Agent, this parameter is optional. Otherwise, it is required. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the backup object.
	//
	// example:
	//
	// [{ "DBName":"Name of the database to be backed up", "SchemaName":"Name of the schema to be backed up", "TableIncludes":[{ "TableName":"Name of the table to be backed up" }], "TableExcludes":[{"TableName":"Name of the table to be excluded during the backup" }] } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1h****usfa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique ID (UID) of the Alibaba Cloud account to which the backup schedule belongs. You can call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) operation to obtain the UID.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The name of the RAM role that is used to perform backups across Alibaba Cloud accounts. You can call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) operation to obtain the RAM role.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the database.
	//
	// 	- This parameter is required if the database is a PostgreSQL or MongoDB database.
	//
	// 	- This parameter is required if the database is an SQL Server database that is connected to DBS over a DBS backup gateway.
	//
	// example:
	//
	// test
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The endpoint of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the endpoint.
	//
	// >  If you set **SourceEndpointInstanceType*	- to Express, Agent, or Other, this parameter is required.
	//
	// example:
	//
	// 100.*.*.10:3306
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The database instance ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the ID.
	//
	// >  If you set **SourceEndpointInstanceType*	- to RDS, ECS, DDS, or Express, this parameter is required.
	//
	// example:
	//
	// rm-bp1p8c29479jv****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The location of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it. Valid values:
	//
	// 	- **RDS**
	//
	// 	- **ECS**
	//
	// 	- **Express**: The database is connected to DBS via Express Connect, VPN Gateway, or Smart Access Gateway.
	//
	// 	- **Agent**: The database is connected to DBS over a DBS backup gateway.
	//
	// 	- **DDS**: The database is an ApsaraDB for MongoDB database.
	//
	// 	- **Other**: The database is connected to DBS using the IP address and port of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The SID of the Oracle source database. If the database is an Oracle database, this parameter is required.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password of the account that is used to connect to the database.
	//
	// This parameter is required except that the database is an SQL Server database that is connected to DBS over a DBS backup gateway or a Redis database.
	//
	// example:
	//
	// test
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The port that is used to connect to the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the port.
	//
	// >  If you set **SourceEndpointInstanceType*	- to Express, Agent, Other, or ECS, this parameter is required.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The region in which the database you want to back up resides. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the region.
	//
	// >  If you set **SourceEndpointInstanceType*	- to RDS, ECS, DDS, Express, or Agent, this parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The account that is used to log on to the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the account.
	//
	// If the database is an SQL Server database connected to DBS over a DBS backup gateway or a Redis database, this parameter is optional. Otherwise, it is required.
	//
	// example:
	//
	// test
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
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

func (s *ModifyBackupSourceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
