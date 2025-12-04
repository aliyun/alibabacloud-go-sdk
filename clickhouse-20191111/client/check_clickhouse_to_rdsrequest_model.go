// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckClickhouseToRDSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCkPassword(v string) *CheckClickhouseToRDSRequest
	GetCkPassword() *string
	SetCkUserName(v string) *CheckClickhouseToRDSRequest
	GetCkUserName() *string
	SetClickhousePort(v int64) *CheckClickhouseToRDSRequest
	GetClickhousePort() *int64
	SetDbClusterId(v string) *CheckClickhouseToRDSRequest
	GetDbClusterId() *string
	SetOwnerAccount(v string) *CheckClickhouseToRDSRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckClickhouseToRDSRequest
	GetOwnerId() *int64
	SetRdsId(v string) *CheckClickhouseToRDSRequest
	GetRdsId() *string
	SetRdsPassword(v string) *CheckClickhouseToRDSRequest
	GetRdsPassword() *string
	SetRdsPort(v int64) *CheckClickhouseToRDSRequest
	GetRdsPort() *int64
	SetRdsUserName(v string) *CheckClickhouseToRDSRequest
	GetRdsUserName() *string
	SetRdsVpcId(v string) *CheckClickhouseToRDSRequest
	GetRdsVpcId() *string
	SetRdsVpcUrl(v string) *CheckClickhouseToRDSRequest
	GetRdsVpcUrl() *string
	SetResourceOwnerAccount(v string) *CheckClickhouseToRDSRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckClickhouseToRDSRequest
	GetResourceOwnerId() *int64
}

type CheckClickhouseToRDSRequest struct {
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
	// cc-2zeyy362b5sbm****
	DbClusterId  *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp13v4bnwlu8j****
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
	// The account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// user2
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	// The ID of the VPC in which the ApsaraDB RDS for MySQL instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz9mm0qka0winfl47****
	RdsVpcId *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	// The internal endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp16t9h3999xb0a7****.mysql.rds.aliyuncs.com
	RdsVpcUrl            *string `json:"RdsVpcUrl,omitempty" xml:"RdsVpcUrl,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckClickhouseToRDSRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckClickhouseToRDSRequest) GoString() string {
	return s.String()
}

func (s *CheckClickhouseToRDSRequest) GetCkPassword() *string {
	return s.CkPassword
}

func (s *CheckClickhouseToRDSRequest) GetCkUserName() *string {
	return s.CkUserName
}

func (s *CheckClickhouseToRDSRequest) GetClickhousePort() *int64 {
	return s.ClickhousePort
}

func (s *CheckClickhouseToRDSRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *CheckClickhouseToRDSRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckClickhouseToRDSRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckClickhouseToRDSRequest) GetRdsId() *string {
	return s.RdsId
}

func (s *CheckClickhouseToRDSRequest) GetRdsPassword() *string {
	return s.RdsPassword
}

func (s *CheckClickhouseToRDSRequest) GetRdsPort() *int64 {
	return s.RdsPort
}

func (s *CheckClickhouseToRDSRequest) GetRdsUserName() *string {
	return s.RdsUserName
}

func (s *CheckClickhouseToRDSRequest) GetRdsVpcId() *string {
	return s.RdsVpcId
}

func (s *CheckClickhouseToRDSRequest) GetRdsVpcUrl() *string {
	return s.RdsVpcUrl
}

func (s *CheckClickhouseToRDSRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckClickhouseToRDSRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckClickhouseToRDSRequest) SetCkPassword(v string) *CheckClickhouseToRDSRequest {
	s.CkPassword = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetCkUserName(v string) *CheckClickhouseToRDSRequest {
	s.CkUserName = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetClickhousePort(v int64) *CheckClickhouseToRDSRequest {
	s.ClickhousePort = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetDbClusterId(v string) *CheckClickhouseToRDSRequest {
	s.DbClusterId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetOwnerAccount(v string) *CheckClickhouseToRDSRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetOwnerId(v int64) *CheckClickhouseToRDSRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsId(v string) *CheckClickhouseToRDSRequest {
	s.RdsId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsPassword(v string) *CheckClickhouseToRDSRequest {
	s.RdsPassword = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsPort(v int64) *CheckClickhouseToRDSRequest {
	s.RdsPort = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsUserName(v string) *CheckClickhouseToRDSRequest {
	s.RdsUserName = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsVpcId(v string) *CheckClickhouseToRDSRequest {
	s.RdsVpcId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsVpcUrl(v string) *CheckClickhouseToRDSRequest {
	s.RdsVpcUrl = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetResourceOwnerAccount(v string) *CheckClickhouseToRDSRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetResourceOwnerId(v int64) *CheckClickhouseToRDSRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) Validate() error {
	return dara.Validate(s)
}
