// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDrdsDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateDrdsDBRequest
	GetAccountName() *string
	SetDbInstType(v string) *CreateDrdsDBRequest
	GetDbInstType() *string
	SetDbInstanceIsCreating(v bool) *CreateDrdsDBRequest
	GetDbInstanceIsCreating() *bool
	SetDbName(v string) *CreateDrdsDBRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *CreateDrdsDBRequest
	GetDrdsInstanceId() *string
	SetEncode(v string) *CreateDrdsDBRequest
	GetEncode() *string
	SetInstDbName(v []*CreateDrdsDBRequestInstDbName) *CreateDrdsDBRequest
	GetInstDbName() []*CreateDrdsDBRequestInstDbName
	SetPassword(v string) *CreateDrdsDBRequest
	GetPassword() *string
	SetRdsInstance(v []*string) *CreateDrdsDBRequest
	GetRdsInstance() []*string
	SetRdsSuperAccount(v []*CreateDrdsDBRequestRdsSuperAccount) *CreateDrdsDBRequest
	GetRdsSuperAccount() []*CreateDrdsDBRequestRdsSuperAccount
	SetType(v string) *CreateDrdsDBRequest
	GetType() *string
}

type CreateDrdsDBRequest struct {
	// The name of the account that has permissions to access all databases on the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required only when the Type parameter is set to VERTICAL.
	//
	// example:
	//
	// drds_sample_account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the storage instances that are used by the PolarDB-X 1.0 database. Set the value to RDS.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Specifies whether the required ApsaraDB RDS for MySQL instance is being created.
	//
	// example:
	//
	// false
	DbInstanceIsCreating *bool `json:"DbInstanceIsCreating,omitempty" xml:"DbInstanceIsCreating,omitempty"`
	// The name of the PolarDB-X 1.0 database you want to create.
	//
	// example:
	//
	// testdb
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance on which you want to create the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbgal154****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The encoding method that is used by the database.
	//
	// example:
	//
	// utf8
	Encode     *string                          `json:"Encode,omitempty" xml:"Encode,omitempty"`
	InstDbName []*CreateDrdsDBRequestInstDbName `json:"InstDbName,omitempty" xml:"InstDbName,omitempty" type:"Repeated"`
	// The password that is used to log on to the database.
	//
	// example:
	//
	// drds_sample_password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// ["drds_sample_rds_id1", "drds_sample_rds_id2"]
	RdsInstance     []*string                             `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
	RdsSuperAccount []*CreateDrdsDBRequestRdsSuperAccount `json:"RdsSuperAccount,omitempty" xml:"RdsSuperAccount,omitempty" type:"Repeated"`
	// The partitioning mode of the database. Valid values:
	//
	// 	- **HORIZONTAL**: The database is horizontally partitioned (sharded).
	//
	// 	- **VERTICAL**: The database is vertically partitioned.
	//
	// example:
	//
	// HORIZONTAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDrdsDBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateDrdsDBRequest) GetDbInstType() *string {
	return s.DbInstType
}

func (s *CreateDrdsDBRequest) GetDbInstanceIsCreating() *bool {
	return s.DbInstanceIsCreating
}

func (s *CreateDrdsDBRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDrdsDBRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *CreateDrdsDBRequest) GetEncode() *string {
	return s.Encode
}

func (s *CreateDrdsDBRequest) GetInstDbName() []*CreateDrdsDBRequestInstDbName {
	return s.InstDbName
}

func (s *CreateDrdsDBRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateDrdsDBRequest) GetRdsInstance() []*string {
	return s.RdsInstance
}

func (s *CreateDrdsDBRequest) GetRdsSuperAccount() []*CreateDrdsDBRequestRdsSuperAccount {
	return s.RdsSuperAccount
}

func (s *CreateDrdsDBRequest) GetType() *string {
	return s.Type
}

func (s *CreateDrdsDBRequest) SetAccountName(v string) *CreateDrdsDBRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbInstType(v string) *CreateDrdsDBRequest {
	s.DbInstType = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbInstanceIsCreating(v bool) *CreateDrdsDBRequest {
	s.DbInstanceIsCreating = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDbName(v string) *CreateDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *CreateDrdsDBRequest) SetDrdsInstanceId(v string) *CreateDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequest) SetEncode(v string) *CreateDrdsDBRequest {
	s.Encode = &v
	return s
}

func (s *CreateDrdsDBRequest) SetInstDbName(v []*CreateDrdsDBRequestInstDbName) *CreateDrdsDBRequest {
	s.InstDbName = v
	return s
}

func (s *CreateDrdsDBRequest) SetPassword(v string) *CreateDrdsDBRequest {
	s.Password = &v
	return s
}

func (s *CreateDrdsDBRequest) SetRdsInstance(v []*string) *CreateDrdsDBRequest {
	s.RdsInstance = v
	return s
}

func (s *CreateDrdsDBRequest) SetRdsSuperAccount(v []*CreateDrdsDBRequestRdsSuperAccount) *CreateDrdsDBRequest {
	s.RdsSuperAccount = v
	return s
}

func (s *CreateDrdsDBRequest) SetType(v string) *CreateDrdsDBRequest {
	s.Type = &v
	return s
}

func (s *CreateDrdsDBRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDrdsDBRequestInstDbName struct {
	// The ID of the ApsaraDB RDS for MySQL instance on which the databases need to be vertically partitioned. This parameter is required only when the Type parameter is set to VERTICAL.
	//
	// example:
	//
	// drds_sample_rds_id
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// ["drds_sample_db1", "drds_sample_db2"]
	ShardDbName []*string `json:"ShardDbName,omitempty" xml:"ShardDbName,omitempty" type:"Repeated"`
}

func (s CreateDrdsDBRequestInstDbName) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsDBRequestInstDbName) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequestInstDbName) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateDrdsDBRequestInstDbName) GetShardDbName() []*string {
	return s.ShardDbName
}

func (s *CreateDrdsDBRequestInstDbName) SetDbInstanceId(v string) *CreateDrdsDBRequestInstDbName {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequestInstDbName) SetShardDbName(v []*string) *CreateDrdsDBRequestInstDbName {
	s.ShardDbName = v
	return s
}

func (s *CreateDrdsDBRequestInstDbName) Validate() error {
	return dara.Validate(s)
}

type CreateDrdsDBRequestRdsSuperAccount struct {
	// The account name of the super administrator that is used to connect to the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// drds_sample_rds_super_account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of ApsaraDB RDS instance.
	//
	// example:
	//
	// drds_sample_rds_id
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The password of the super administrator account that is used to connect to the ApsaraDB RDS instance.
	//
	// example:
	//
	// drds_sample_rds_super_password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateDrdsDBRequestRdsSuperAccount) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsDBRequestRdsSuperAccount) GoString() string {
	return s.String()
}

func (s *CreateDrdsDBRequestRdsSuperAccount) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateDrdsDBRequestRdsSuperAccount) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateDrdsDBRequestRdsSuperAccount) GetPassword() *string {
	return s.Password
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetAccountName(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.AccountName = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetDbInstanceId(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) SetPassword(v string) *CreateDrdsDBRequestRdsSuperAccount {
	s.Password = &v
	return s
}

func (s *CreateDrdsDBRequestRdsSuperAccount) Validate() error {
	return dara.Validate(s)
}
