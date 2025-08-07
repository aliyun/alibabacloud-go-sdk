// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStrings(v string) *ModifyDBClusterMigrationRequest
	GetConnectionStrings() *string
	SetDBClusterId(v string) *ModifyDBClusterMigrationRequest
	GetDBClusterId() *string
	SetNewMasterInstanceId(v string) *ModifyDBClusterMigrationRequest
	GetNewMasterInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBClusterMigrationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterMigrationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterMigrationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterMigrationRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyDBClusterMigrationRequest
	GetSecurityToken() *string
	SetSourceRDSDBInstanceId(v string) *ModifyDBClusterMigrationRequest
	GetSourceRDSDBInstanceId() *string
	SetSwapConnectionString(v string) *ModifyDBClusterMigrationRequest
	GetSwapConnectionString() *string
}

type ModifyDBClusterMigrationRequest struct {
	// The endpoints to be switched. The endpoints are in the JSON format.
	//
	// > This parameter is valid when the SwapConnectionString parameter is set to true.
	//
	// example:
	//
	// {"rm-2ze73el581cs*****.mysql.pre.rds.aliyuncs.com":"pc-2ze8200s298e*****.mysql.polardb.pre.rds.aliyuncs.com","rm-2ze73el581cs86*****.mysql.pre.rds.aliyuncs.com":"test-p*****.mysql.polardb.pre.rds.aliyuncs.com"}
	ConnectionStrings *string `json:"ConnectionStrings,omitempty" xml:"ConnectionStrings,omitempty"`
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the new instance or new cluster. Valid values:
	//
	// 	- To perform a data migration, enter the ID of the PolarDB cluster.
	//
	// 	- To perform a migration rollback, enter the ID of the ApsaraDB for RDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	NewMasterInstanceId  *string `json:"NewMasterInstanceId,omitempty" xml:"NewMasterInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the source ApsaraDB RDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-************
	SourceRDSDBInstanceId *string `json:"SourceRDSDBInstanceId,omitempty" xml:"SourceRDSDBInstanceId,omitempty"`
	// Specifies whether to switch the endpoints. Valid values:
	//
	// 	- **true**: switches the endpoints. If you select this option, you do not need the change the endpoint in your applications.
	//
	// 	- **false**: does not switch the endpoints. If you select this option, you must specify the endpoint of the PolarDB cluster in your applications.
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	SwapConnectionString *string `json:"SwapConnectionString,omitempty" xml:"SwapConnectionString,omitempty"`
}

func (s ModifyDBClusterMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationRequest) GetConnectionStrings() *string {
	return s.ConnectionStrings
}

func (s *ModifyDBClusterMigrationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterMigrationRequest) GetNewMasterInstanceId() *string {
	return s.NewMasterInstanceId
}

func (s *ModifyDBClusterMigrationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterMigrationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterMigrationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterMigrationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterMigrationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDBClusterMigrationRequest) GetSourceRDSDBInstanceId() *string {
	return s.SourceRDSDBInstanceId
}

func (s *ModifyDBClusterMigrationRequest) GetSwapConnectionString() *string {
	return s.SwapConnectionString
}

func (s *ModifyDBClusterMigrationRequest) SetConnectionStrings(v string) *ModifyDBClusterMigrationRequest {
	s.ConnectionStrings = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetDBClusterId(v string) *ModifyDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetNewMasterInstanceId(v string) *ModifyDBClusterMigrationRequest {
	s.NewMasterInstanceId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetOwnerAccount(v string) *ModifyDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetOwnerId(v int64) *ModifyDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetSecurityToken(v string) *ModifyDBClusterMigrationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetSourceRDSDBInstanceId(v string) *ModifyDBClusterMigrationRequest {
	s.SourceRDSDBInstanceId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetSwapConnectionString(v string) *ModifyDBClusterMigrationRequest {
	s.SwapConnectionString = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) Validate() error {
	return dara.Validate(s)
}
