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
	SetDBClusterId(v string) *DeleteBackupRequest
	GetDBClusterId() *string
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
	// The backup ID. If you need to specify multiple backup IDs, separate the backup IDs with commas (,).
	//
	// >  You can call the [DescribeBackups](https://help.aliyun.com/document_detail/98102.html) operation to query the backup IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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

func (s *DeleteBackupRequest) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DeleteBackupRequest) SetDBClusterId(v string) *DeleteBackupRequest {
	s.DBClusterId = &v
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
