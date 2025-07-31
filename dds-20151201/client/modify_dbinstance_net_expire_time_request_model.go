// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceNetExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassicExpendExpiredDays(v int32) *ModifyDBInstanceNetExpireTimeRequest
	GetClassicExpendExpiredDays() *int32
	SetConnectionString(v string) *ModifyDBInstanceNetExpireTimeRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceNetExpireTimeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceNetExpireTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceNetExpireTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceNetExpireTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceNetExpireTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceNetExpireTimeRequest struct {
	// The retention period of the classic network endpoint of the instance. Valid values: **14**, **30**, **60**, and **120**. Unit: day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	ClassicExpendExpiredDays *int32 `json:"ClassicExpendExpiredDays,omitempty" xml:"ClassicExpendExpiredDays,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// dds-bpxxxxxxxx.mongodb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceNetExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceNetExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetExpireTimeRequest) GetClassicExpendExpiredDays() *int32 {
	return s.ClassicExpendExpiredDays
}

func (s *ModifyDBInstanceNetExpireTimeRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceNetExpireTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceNetExpireTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceNetExpireTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceNetExpireTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceNetExpireTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetClassicExpendExpiredDays(v int32) *ModifyDBInstanceNetExpireTimeRequest {
	s.ClassicExpendExpiredDays = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetConnectionString(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetOwnerAccount(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetOwnerId(v int64) *ModifyDBInstanceNetExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceNetExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceNetExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceNetExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
