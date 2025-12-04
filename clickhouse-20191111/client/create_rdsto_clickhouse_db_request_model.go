// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRDSToClickhouseDbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCkPassword(v string) *CreateRDSToClickhouseDbRequest
	GetCkPassword() *string
	SetCkUserName(v string) *CreateRDSToClickhouseDbRequest
	GetCkUserName() *string
	SetClickhousePort(v int64) *CreateRDSToClickhouseDbRequest
	GetClickhousePort() *int64
	SetDbClusterId(v string) *CreateRDSToClickhouseDbRequest
	GetDbClusterId() *string
	SetLimitUpper(v int64) *CreateRDSToClickhouseDbRequest
	GetLimitUpper() *int64
	SetOwnerAccount(v string) *CreateRDSToClickhouseDbRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRDSToClickhouseDbRequest
	GetOwnerId() *int64
	SetRdsId(v string) *CreateRDSToClickhouseDbRequest
	GetRdsId() *string
	SetRdsPassword(v string) *CreateRDSToClickhouseDbRequest
	GetRdsPassword() *string
	SetRdsPort(v int64) *CreateRDSToClickhouseDbRequest
	GetRdsPort() *int64
	SetRdsUserName(v string) *CreateRDSToClickhouseDbRequest
	GetRdsUserName() *string
	SetRdsVpcId(v string) *CreateRDSToClickhouseDbRequest
	GetRdsVpcId() *string
	SetRdsVpcUrl(v string) *CreateRDSToClickhouseDbRequest
	GetRdsVpcUrl() *string
	SetResourceOwnerAccount(v string) *CreateRDSToClickhouseDbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRDSToClickhouseDbRequest
	GetResourceOwnerId() *int64
	SetSkipUnsupported(v bool) *CreateRDSToClickhouseDbRequest
	GetSkipUnsupported() *bool
	SetSynDbTables(v string) *CreateRDSToClickhouseDbRequest
	GetSynDbTables() *string
}

type CreateRDSToClickhouseDbRequest struct {
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
	// cc-2ze5zeyl72188****
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
	// rm-8vb989qj9roh0****
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The password of the account that is used to log on to the ApsaraDB RDS for MySQL instance.
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
	// vpc-2zen93xryil99jsfy****
	RdsVpcId *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	// The private endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-bp16t9h3999xb0a7****.mysql.rds.aliyuncs.com
	RdsVpcUrl            *string `json:"RdsVpcUrl,omitempty" xml:"RdsVpcUrl,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to ignore the table schemas that do not support synchronization. Valid values:
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
	// The tables whose data you want to synchronize.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"Schema":"recommend","Tables":["mr_platform_cpm","mr_platform_ecpm","p_monitor_record"]}]
	SynDbTables *string `json:"SynDbTables,omitempty" xml:"SynDbTables,omitempty"`
}

func (s CreateRDSToClickhouseDbRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRDSToClickhouseDbRequest) GoString() string {
	return s.String()
}

func (s *CreateRDSToClickhouseDbRequest) GetCkPassword() *string {
	return s.CkPassword
}

func (s *CreateRDSToClickhouseDbRequest) GetCkUserName() *string {
	return s.CkUserName
}

func (s *CreateRDSToClickhouseDbRequest) GetClickhousePort() *int64 {
	return s.ClickhousePort
}

func (s *CreateRDSToClickhouseDbRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *CreateRDSToClickhouseDbRequest) GetLimitUpper() *int64 {
	return s.LimitUpper
}

func (s *CreateRDSToClickhouseDbRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRDSToClickhouseDbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRDSToClickhouseDbRequest) GetRdsId() *string {
	return s.RdsId
}

func (s *CreateRDSToClickhouseDbRequest) GetRdsPassword() *string {
	return s.RdsPassword
}

func (s *CreateRDSToClickhouseDbRequest) GetRdsPort() *int64 {
	return s.RdsPort
}

func (s *CreateRDSToClickhouseDbRequest) GetRdsUserName() *string {
	return s.RdsUserName
}

func (s *CreateRDSToClickhouseDbRequest) GetRdsVpcId() *string {
	return s.RdsVpcId
}

func (s *CreateRDSToClickhouseDbRequest) GetRdsVpcUrl() *string {
	return s.RdsVpcUrl
}

func (s *CreateRDSToClickhouseDbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRDSToClickhouseDbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRDSToClickhouseDbRequest) GetSkipUnsupported() *bool {
	return s.SkipUnsupported
}

func (s *CreateRDSToClickhouseDbRequest) GetSynDbTables() *string {
	return s.SynDbTables
}

func (s *CreateRDSToClickhouseDbRequest) SetCkPassword(v string) *CreateRDSToClickhouseDbRequest {
	s.CkPassword = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetCkUserName(v string) *CreateRDSToClickhouseDbRequest {
	s.CkUserName = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetClickhousePort(v int64) *CreateRDSToClickhouseDbRequest {
	s.ClickhousePort = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetDbClusterId(v string) *CreateRDSToClickhouseDbRequest {
	s.DbClusterId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetLimitUpper(v int64) *CreateRDSToClickhouseDbRequest {
	s.LimitUpper = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetOwnerAccount(v string) *CreateRDSToClickhouseDbRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetOwnerId(v int64) *CreateRDSToClickhouseDbRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsId(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsPassword(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsPassword = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsPort(v int64) *CreateRDSToClickhouseDbRequest {
	s.RdsPort = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsUserName(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsUserName = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsVpcId(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsVpcId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsVpcUrl(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsVpcUrl = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetResourceOwnerAccount(v string) *CreateRDSToClickhouseDbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetResourceOwnerId(v int64) *CreateRDSToClickhouseDbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetSkipUnsupported(v bool) *CreateRDSToClickhouseDbRequest {
	s.SkipUnsupported = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetSynDbTables(v string) *CreateRDSToClickhouseDbRequest {
	s.SynDbTables = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) Validate() error {
	return dara.Validate(s)
}
