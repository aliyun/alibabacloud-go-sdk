// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DeleteBackupRequest
	GetBackupId() *string
	SetInstanceId(v string) *DeleteBackupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteBackupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteBackupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteBackupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBackupRequest
	GetResourceOwnerId() *int64
}

type DeleteBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 521****66
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of instance.
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

func (s DeleteBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DeleteBackupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBackupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteBackupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBackupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBackupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBackupRequest) SetBackupId(v string) *DeleteBackupRequest {
	s.BackupId = &v
	return s
}

func (s *DeleteBackupRequest) SetInstanceId(v string) *DeleteBackupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBackupRequest) SetOwnerAccount(v string) *DeleteBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetOwnerId(v int64) *DeleteBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerAccount(v string) *DeleteBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerId(v int64) *DeleteBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBackupRequest) Validate() error {
	return dara.Validate(s)
}
