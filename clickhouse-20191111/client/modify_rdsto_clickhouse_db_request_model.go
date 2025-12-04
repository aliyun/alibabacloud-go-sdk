// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRDSToClickhouseDbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCkPassword(v string) *ModifyRDSToClickhouseDbRequest
	GetCkPassword() *string
	SetCkUserName(v string) *ModifyRDSToClickhouseDbRequest
	GetCkUserName() *string
	SetClickhousePort(v int64) *ModifyRDSToClickhouseDbRequest
	GetClickhousePort() *int64
	SetDbClusterId(v string) *ModifyRDSToClickhouseDbRequest
	GetDbClusterId() *string
	SetLimitUpper(v int64) *ModifyRDSToClickhouseDbRequest
	GetLimitUpper() *int64
	SetOwnerAccount(v string) *ModifyRDSToClickhouseDbRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyRDSToClickhouseDbRequest
	GetOwnerId() *int64
	SetRdsId(v string) *ModifyRDSToClickhouseDbRequest
	GetRdsId() *string
	SetRdsPassword(v string) *ModifyRDSToClickhouseDbRequest
	GetRdsPassword() *string
	SetRdsPort(v int64) *ModifyRDSToClickhouseDbRequest
	GetRdsPort() *int64
	SetRdsSynDb(v string) *ModifyRDSToClickhouseDbRequest
	GetRdsSynDb() *string
	SetRdsSynTables(v string) *ModifyRDSToClickhouseDbRequest
	GetRdsSynTables() *string
	SetRdsUserName(v string) *ModifyRDSToClickhouseDbRequest
	GetRdsUserName() *string
	SetRdsVpcId(v string) *ModifyRDSToClickhouseDbRequest
	GetRdsVpcId() *string
	SetResourceOwnerAccount(v string) *ModifyRDSToClickhouseDbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRDSToClickhouseDbRequest
	GetResourceOwnerId() *int64
	SetSkipUnsupported(v bool) *ModifyRDSToClickhouseDbRequest
	GetSkipUnsupported() *bool
}

type ModifyRDSToClickhouseDbRequest struct {
	// The password of the account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	CkPassword *string `json:"CkPassword,omitempty" xml:"CkPassword,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// user1
	CkUserName *string `json:"CkUserName,omitempty" xml:"CkUserName,omitempty"`
	// The port number of the ApsaraDB for ClickHouse cluster.
	//
	// example:
	//
	// 8123
	ClickhousePort *int64 `json:"ClickhousePort,omitempty" xml:"ClickhousePort,omitempty"`
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp158i5wvj436****
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The maximum number of rows that can be synchronized per second.
	//
	// example:
	//
	// 50000
	LimitUpper   *int64  `json:"LimitUpper,omitempty" xml:"LimitUpper,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6x3qq4t90ok****
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The password of the account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Rr
	RdsPassword *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	// The port number of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// 3306
	RdsPort *int64 `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	// The database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	RdsSynDb *string `json:"RdsSynDb,omitempty" xml:"RdsSynDb,omitempty"`
	// The table in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// table
	RdsSynTables *string `json:"RdsSynTables,omitempty" xml:"RdsSynTables,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// user2
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the ApsaraDB RDS for MySQL instance belongs.
	//
	// example:
	//
	// vpc-bp1v9dtwmqqjhwwg****
	RdsVpcId             *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to ignore databases that do not support synchronization. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SkipUnsupported *bool `json:"SkipUnsupported,omitempty" xml:"SkipUnsupported,omitempty"`
}

func (s ModifyRDSToClickhouseDbRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRDSToClickhouseDbRequest) GoString() string {
	return s.String()
}

func (s *ModifyRDSToClickhouseDbRequest) GetCkPassword() *string {
	return s.CkPassword
}

func (s *ModifyRDSToClickhouseDbRequest) GetCkUserName() *string {
	return s.CkUserName
}

func (s *ModifyRDSToClickhouseDbRequest) GetClickhousePort() *int64 {
	return s.ClickhousePort
}

func (s *ModifyRDSToClickhouseDbRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *ModifyRDSToClickhouseDbRequest) GetLimitUpper() *int64 {
	return s.LimitUpper
}

func (s *ModifyRDSToClickhouseDbRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyRDSToClickhouseDbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRDSToClickhouseDbRequest) GetRdsId() *string {
	return s.RdsId
}

func (s *ModifyRDSToClickhouseDbRequest) GetRdsPassword() *string {
	return s.RdsPassword
}

func (s *ModifyRDSToClickhouseDbRequest) GetRdsPort() *int64 {
	return s.RdsPort
}

func (s *ModifyRDSToClickhouseDbRequest) GetRdsSynDb() *string {
	return s.RdsSynDb
}

func (s *ModifyRDSToClickhouseDbRequest) GetRdsSynTables() *string {
	return s.RdsSynTables
}

func (s *ModifyRDSToClickhouseDbRequest) GetRdsUserName() *string {
	return s.RdsUserName
}

func (s *ModifyRDSToClickhouseDbRequest) GetRdsVpcId() *string {
	return s.RdsVpcId
}

func (s *ModifyRDSToClickhouseDbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRDSToClickhouseDbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRDSToClickhouseDbRequest) GetSkipUnsupported() *bool {
	return s.SkipUnsupported
}

func (s *ModifyRDSToClickhouseDbRequest) SetCkPassword(v string) *ModifyRDSToClickhouseDbRequest {
	s.CkPassword = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetCkUserName(v string) *ModifyRDSToClickhouseDbRequest {
	s.CkUserName = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetClickhousePort(v int64) *ModifyRDSToClickhouseDbRequest {
	s.ClickhousePort = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetDbClusterId(v string) *ModifyRDSToClickhouseDbRequest {
	s.DbClusterId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetLimitUpper(v int64) *ModifyRDSToClickhouseDbRequest {
	s.LimitUpper = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetOwnerAccount(v string) *ModifyRDSToClickhouseDbRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetOwnerId(v int64) *ModifyRDSToClickhouseDbRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsId(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsPassword(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsPassword = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsPort(v int64) *ModifyRDSToClickhouseDbRequest {
	s.RdsPort = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsSynDb(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsSynDb = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsSynTables(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsSynTables = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsUserName(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsUserName = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsVpcId(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsVpcId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetResourceOwnerAccount(v string) *ModifyRDSToClickhouseDbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetResourceOwnerId(v int64) *ModifyRDSToClickhouseDbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetSkipUnsupported(v bool) *ModifyRDSToClickhouseDbRequest {
	s.SkipUnsupported = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) Validate() error {
	return dara.Validate(s)
}
