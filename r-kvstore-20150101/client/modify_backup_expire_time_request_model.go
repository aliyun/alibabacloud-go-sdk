// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ModifyBackupExpireTimeRequest
	GetBackupId() *string
	SetExpectExpireTime(v string) *ModifyBackupExpireTimeRequest
	GetExpectExpireTime() *string
	SetInstanceId(v string) *ModifyBackupExpireTimeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyBackupExpireTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupExpireTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyBackupExpireTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupExpireTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyBackupExpireTimeRequest struct {
	// The ID of the backup file. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query the IDs of backup files.
	//
	// example:
	//
	// 521****66
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The point in time to which you want to extend the expiration time. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The time cannot be earlier than the current expiration time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-07-06T07:25:57Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupExpireTimeRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ModifyBackupExpireTimeRequest) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *ModifyBackupExpireTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBackupExpireTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBackupExpireTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupExpireTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupExpireTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupExpireTimeRequest) SetBackupId(v string) *ModifyBackupExpireTimeRequest {
	s.BackupId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetExpectExpireTime(v string) *ModifyBackupExpireTimeRequest {
	s.ExpectExpireTime = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetInstanceId(v string) *ModifyBackupExpireTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetOwnerAccount(v string) *ModifyBackupExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetOwnerId(v int64) *ModifyBackupExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetResourceOwnerAccount(v string) *ModifyBackupExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyBackupExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
