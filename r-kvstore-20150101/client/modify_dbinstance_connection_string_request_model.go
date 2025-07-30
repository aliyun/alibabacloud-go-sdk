// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetCurrentConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest
	GetDBInstanceId() *string
	SetIPType(v string) *ModifyDBInstanceConnectionStringRequest
	GetIPType() *string
	SetNewConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetNewConnectionString() *string
	SetOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest
	GetOwnerId() *int64
	SetPort(v string) *ModifyDBInstanceConnectionStringRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyDBInstanceConnectionStringRequest
	GetSecurityToken() *string
}

type ModifyDBInstanceConnectionStringRequest struct {
	// The current endpoint of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- **Private**: internal network
	//
	// 	- **Public**: Internet
	//
	// example:
	//
	// Public
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The prefix of the new endpoint. Specify the endpoint in the `<prefix>.redis.rds.aliyuncs.com` format. The prefix must be 8 to 40 characters in length and can contain lowercase letters and digits. It must start with a lowercase letter.
	//
	// >  You must specify one of the **NewConnectionString*	- and **Port*	- parameters.
	//
	// example:
	//
	// standardredis
	NewConnectionString *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number of the instance. Valid values: **1024*	- to **65535**.
	//
	// >  You must specify one of the **NewConnectionString*	- and **Port*	- parameters.
	//
	// example:
	//
	// 6379
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetIPType() *string {
	return s.IPType
}

func (s *ModifyDBInstanceConnectionStringRequest) GetNewConnectionString() *string {
	return s.NewConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceConnectionStringRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *ModifyDBInstanceConnectionStringRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetIPType(v string) *ModifyDBInstanceConnectionStringRequest {
	s.IPType = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewConnectionString = &v
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

func (s *ModifyDBInstanceConnectionStringRequest) SetSecurityToken(v string) *ModifyDBInstanceConnectionStringRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
