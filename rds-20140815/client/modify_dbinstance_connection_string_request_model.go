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
	// The TDS port number for Babelfish for RDS PostgreSQL.
	//
	// > This parameter is applicable only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for RDS PostgreSQL, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).
	//
	// example:
	//
	// 1433
	BabelfishPort *string `json:"BabelfishPort,omitempty" xml:"BabelfishPort,omitempty"`
	// The prefix of the endpoint. You can modify only the prefix of the value specified by the **CurrentConnectionString*	- parameter.
	//
	// >The prefix must be 8 to 64 characters in length and cannot contain Chinese characters or special characters (~!#%^&*=+\\|{};:\\"",<>/?). The prefix can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The current endpoint of the instance. The endpoint can be a public endpoint or internal endpoint, or a classic network connectivity endpoint in hybrid access mode.
	//
	// >Modification of read/write splitting connection endpoints is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5x****.mysql.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the group to which the dedicated cluster MySQL general-purpose instance belongs.
	//
	// example:
	//
	// rgc-bp1tkv8****
	GeneralGroupName *string `json:"GeneralGroupName,omitempty" xml:"GeneralGroupName,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The PgBouncer port number.
	//
	// > This parameter is applicable only to ApsaraDB RDS for PostgreSQL instances. If PgBouncer is enabled, you can modify the PgBouncer port number.
	//
	// example:
	//
	// 6432
	PGBouncerPort *string `json:"PGBouncerPort,omitempty" xml:"PGBouncerPort,omitempty"`
	// The target port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to retain the virtual IP address (VIP) when swapping the endpoint.
	//
	// - **true**: The VIP is retained.
	//
	// - **false*	- (default): The VIP is not retained.
	//
	// > This parameter is applicable only to ApsaraDB RDS for PostgreSQL instances.
	//
	// example:
	//
	// false
	RetainVip *bool `json:"RetainVip,omitempty" xml:"RetainVip,omitempty"`
	// The instance ID of the target ApsaraDB RDS for PostgreSQL instance with which you want to swap the endpoint.
	//
	// > This parameter is applicable only to ApsaraDB RDS for PostgreSQL instances.
	//
	// example:
	//
	// pgm-bp1206s14p3o****
	TargetDBInstanceId *string `json:"TargetDBInstanceId,omitempty" xml:"TargetDBInstanceId,omitempty"`
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
