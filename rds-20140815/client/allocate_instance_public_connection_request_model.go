// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBabelfishPort(v string) *AllocateInstancePublicConnectionRequest
	GetBabelfishPort() *string
	SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest
	GetDBInstanceId() *string
	SetGeneralGroupName(v string) *AllocateInstancePublicConnectionRequest
	GetGeneralGroupName() *string
	SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetOwnerId() *int64
	SetPGBouncerPort(v string) *AllocateInstancePublicConnectionRequest
	GetPGBouncerPort() *string
	SetPort(v string) *AllocateInstancePublicConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
}

type AllocateInstancePublicConnectionRequest struct {
	// The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
	//
	// > This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).
	//
	// example:
	//
	// 1433
	BabelfishPort *string `json:"BabelfishPort,omitempty" xml:"BabelfishPort,omitempty"`
	// The prefix of the public endpoint. A valid public endpoint is in the following format: `Prefix.Database engine.rds.aliyuncs.com`. Example: `test1234.mysql.rds.aliyuncs.com`.
	//
	// > The value can be 5 to 40 characters in length and can contain letters, digits, and hyphens (-). The value cannot contain any of the following characters: ~ ! # % ^ & \\	- = + | {} ; : \\" " , <> / ?
	//
	// This parameter is required.
	//
	// example:
	//
	// test1234
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the dedicated cluster to which the instance belongs. This parameter is available only when the instance is created in an ApsaraDB MyBase cluster that runs MySQL on Standard Edition.
	//
	// example:
	//
	// rgc-bp1tkv8*****
	GeneralGroupName *string `json:"GeneralGroupName,omitempty" xml:"GeneralGroupName,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The PgBouncer port.
	//
	// > This parameter is available only for instances that run PostgreSQL.
	//
	// example:
	//
	// 6432
	PGBouncerPort *string `json:"PGBouncerPort,omitempty" xml:"PGBouncerPort,omitempty"`
	// The public port of the instance. Valid values: **1000 to 5999**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) GetBabelfishPort() *string {
	return s.BabelfishPort
}

func (s *AllocateInstancePublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateInstancePublicConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateInstancePublicConnectionRequest) GetGeneralGroupName() *string {
	return s.GeneralGroupName
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateInstancePublicConnectionRequest) GetPGBouncerPort() *string {
	return s.PGBouncerPort
}

func (s *AllocateInstancePublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateInstancePublicConnectionRequest) SetBabelfishPort(v string) *AllocateInstancePublicConnectionRequest {
	s.BabelfishPort = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetGeneralGroupName(v string) *AllocateInstancePublicConnectionRequest {
	s.GeneralGroupName = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPGBouncerPort(v string) *AllocateInstancePublicConnectionRequest {
	s.PGBouncerPort = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
