// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetworkExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassicExpiredDays(v int32) *ModifyDBInstanceNetworkExpireTimeRequest
	GetClassicExpiredDays() *int32
	SetConnectionString(v string) *ModifyDBInstanceNetworkExpireTimeRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceNetworkExpireTimeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceNetworkExpireTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceNetworkExpireTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceNetworkExpireTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceNetworkExpireTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceNetworkExpireTimeRequest struct {
	// The retention days of the classic network endpoint. Valid values: **1 to 120**. Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	ClassicExpiredDays *int32 `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// The classic network endpoint whose expiration time you want to extend. Two types of classic network endpoints are supported:
	//
	// 	- The internal endpoint of the classic network.
	//
	// 	- The read/write splitting endpoint of the classic network.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxx.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceNetworkExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetworkExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) GetClassicExpiredDays() *int32 {
	return s.ClassicExpiredDays
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) SetClassicExpiredDays(v int32) *ModifyDBInstanceNetworkExpireTimeRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) SetConnectionString(v string) *ModifyDBInstanceNetworkExpireTimeRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkExpireTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) SetOwnerAccount(v string) *ModifyDBInstanceNetworkExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) SetOwnerId(v int64) *ModifyDBInstanceNetworkExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceNetworkExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceNetworkExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetworkExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
