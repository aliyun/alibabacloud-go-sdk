// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPolicyMode(v string) *DescribeBackupPolicyRequest
	GetBackupPolicyMode() *string
	SetCompressType(v string) *DescribeBackupPolicyRequest
	GetCompressType() *string
	SetDBInstanceId(v string) *DescribeBackupPolicyRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupPolicyRequest
	GetOwnerId() *int64
	SetReleasedKeepPolicy(v string) *DescribeBackupPolicyRequest
	GetReleasedKeepPolicy() *string
	SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeBackupPolicyRequest struct {
	BackupPolicyMode *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	CompressType     *string `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReleasedKeepPolicy   *string `json:"ReleasedKeepPolicy,omitempty" xml:"ReleasedKeepPolicy,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) GetBackupPolicyMode() *string {
	return s.BackupPolicyMode
}

func (s *DescribeBackupPolicyRequest) GetCompressType() *string {
	return s.CompressType
}

func (s *DescribeBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupPolicyRequest) GetReleasedKeepPolicy() *string {
	return s.ReleasedKeepPolicy
}

func (s *DescribeBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupPolicyRequest) SetBackupPolicyMode(v string) *DescribeBackupPolicyRequest {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetCompressType(v string) *DescribeBackupPolicyRequest {
	s.CompressType = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetDBInstanceId(v string) *DescribeBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetReleasedKeepPolicy(v string) *DescribeBackupPolicyRequest {
	s.ReleasedKeepPolicy = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
