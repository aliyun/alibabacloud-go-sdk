// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBabelfishPort(v string) *ModifyDBInstanceConnectionStringRequest
	GetBabelfishPort() *string
	SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest
	GetConnectionStringPrefix() *string
	SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetCurrentConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest
	GetDBInstanceId() *string
	SetGeneralGroupName(v string) *ModifyDBInstanceConnectionStringRequest
	GetGeneralGroupName() *string
	SetOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest
	GetOwnerId() *int64
	SetPGBouncerPort(v string) *ModifyDBInstanceConnectionStringRequest
	GetPGBouncerPort() *string
	SetPort(v string) *ModifyDBInstanceConnectionStringRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerId() *int64
	SetRetainVip(v bool) *ModifyDBInstanceConnectionStringRequest
	GetRetainVip() *bool
	SetTargetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest
	GetTargetDBInstanceId() *string
}

type ModifyDBInstanceConnectionStringRequest struct {
	// The Tabular Data Stream (TDS) port of the instance for which Babelfish is enabled.
	//
	// > This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).
	//
	// example:
	//
	// 1433
	BabelfishPort *string `json:"BabelfishPort,omitempty" xml:"BabelfishPort,omitempty"`
	// The prefix of the endpoint after the change. Only the prefix of the value of **CurrentConnectionString*	- can be changed.
	//
	// > The value must be 8 to 64 characters in length and can contain letters, digits, and hyphens (-). The value cannot contain any of the following special characters: ! # % ^ & \\	- = + | {} ; : \\" " ,<> / ?
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The endpoint of the instance. It can be an internal endpoint, a public endpoint, or a classic network endpoint in hybrid access mode.
	//
	// > The read/write splitting endpoint cannot be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5x****.mysql.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the dedicated cluster to which the instance belongs. This parameter is returned only when the instance is created in an ApsaraDB MyBase cluster that runs MySQL on Standard Edition.
	//
	// example:
	//
	// rgc-bp1tkv8****
	GeneralGroupName *string `json:"GeneralGroupName,omitempty" xml:"GeneralGroupName,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The PgBouncer port.
	//
	// > This parameter is suitable only for ApsaraDB RDS for PostgreSQL instances. If you enable PgBouncer for your instance, you can change the PgBouncer port of the instance.
	//
	// example:
	//
	// 6432
	PGBouncerPort *string `json:"PGBouncerPort,omitempty" xml:"PGBouncerPort,omitempty"`
	// The port number after the change.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RetainVip            *bool   `json:"RetainVip,omitempty" xml:"RetainVip,omitempty"`
	TargetDBInstanceId   *string `json:"TargetDBInstanceId,omitempty" xml:"TargetDBInstanceId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) GetBabelfishPort() *string {
	return s.BabelfishPort
}

func (s *ModifyDBInstanceConnectionStringRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyDBInstanceConnectionStringRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetGeneralGroupName() *string {
	return s.GeneralGroupName
}

func (s *ModifyDBInstanceConnectionStringRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceConnectionStringRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetPGBouncerPort() *string {
	return s.PGBouncerPort
}

func (s *ModifyDBInstanceConnectionStringRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyDBInstanceConnectionStringRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceConnectionStringRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetRetainVip() *bool {
	return s.RetainVip
}

func (s *ModifyDBInstanceConnectionStringRequest) GetTargetDBInstanceId() *string {
	return s.TargetDBInstanceId
}

func (s *ModifyDBInstanceConnectionStringRequest) SetBabelfishPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.BabelfishPort = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetGeneralGroupName(v string) *ModifyDBInstanceConnectionStringRequest {
	s.GeneralGroupName = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPGBouncerPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.PGBouncerPort = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetRetainVip(v bool) *ModifyDBInstanceConnectionStringRequest {
	s.RetainVip = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetTargetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.TargetDBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
