// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGetDBListFromAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupGatewayId(v int64) *CreateGetDBListFromAgentTaskRequest
	GetBackupGatewayId() *int64
	SetBackupGatewayIdentifier(v string) *CreateGetDBListFromAgentTaskRequest
	GetBackupGatewayIdentifier() *string
	SetClientToken(v string) *CreateGetDBListFromAgentTaskRequest
	GetClientToken() *string
	SetDatabaseType(v string) *CreateGetDBListFromAgentTaskRequest
	GetDatabaseType() *string
	SetOwnerId(v string) *CreateGetDBListFromAgentTaskRequest
	GetOwnerId() *string
	SetSourceEndpointIP(v string) *CreateGetDBListFromAgentTaskRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointPort(v int32) *CreateGetDBListFromAgentTaskRequest
	GetSourceEndpointPort() *int32
	SetSourceEndpointRegion(v string) *CreateGetDBListFromAgentTaskRequest
	GetSourceEndpointRegion() *string
}

type CreateGetDBListFromAgentTaskRequest struct {
	// The ID of the backup gateway. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the ID.
	//
	// > This parameter is required.
	//
	// example:
	//
	// 160813
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The unique identifier of the backup gateway.
	//
	// example:
	//
	// 4DF900B7082BC3****489D765C8D2ABE
	BackupGatewayIdentifier *string `json:"BackupGatewayIdentifier,omitempty" xml:"BackupGatewayIdentifier,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The type of the database. Valid values:
	//
	// - **MySQL**
	//
	// - **MSSQL**
	//
	// - **Oracle**
	//
	// - **MariaDB**
	//
	// - **PostgreSQL**
	//
	// - **DRDS**
	//
	// - **MongoDB**
	//
	// - **Redis**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The database connection address.
	//
	// example:
	//
	// 123.0.XX.XX
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The database connection port.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The region ID of the backup gateway.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
}

func (s CreateGetDBListFromAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGetDBListFromAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateGetDBListFromAgentTaskRequest) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *CreateGetDBListFromAgentTaskRequest) GetBackupGatewayIdentifier() *string {
	return s.BackupGatewayIdentifier
}

func (s *CreateGetDBListFromAgentTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGetDBListFromAgentTaskRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateGetDBListFromAgentTaskRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateGetDBListFromAgentTaskRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *CreateGetDBListFromAgentTaskRequest) GetSourceEndpointPort() *int32 {
	return s.SourceEndpointPort
}

func (s *CreateGetDBListFromAgentTaskRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *CreateGetDBListFromAgentTaskRequest) SetBackupGatewayId(v int64) *CreateGetDBListFromAgentTaskRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetBackupGatewayIdentifier(v string) *CreateGetDBListFromAgentTaskRequest {
	s.BackupGatewayIdentifier = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetClientToken(v string) *CreateGetDBListFromAgentTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetDatabaseType(v string) *CreateGetDBListFromAgentTaskRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetOwnerId(v string) *CreateGetDBListFromAgentTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetSourceEndpointIP(v string) *CreateGetDBListFromAgentTaskRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetSourceEndpointPort(v int32) *CreateGetDBListFromAgentTaskRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetSourceEndpointRegion(v string) *CreateGetDBListFromAgentTaskRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
