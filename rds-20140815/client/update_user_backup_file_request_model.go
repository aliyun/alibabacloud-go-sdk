// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserBackupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *UpdateUserBackupFileRequest
	GetBackupId() *string
	SetComment(v string) *UpdateUserBackupFileRequest
	GetComment() *string
	SetOwnerId(v int64) *UpdateUserBackupFileRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateUserBackupFileRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateUserBackupFileRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *UpdateUserBackupFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateUserBackupFileRequest
	GetResourceOwnerId() *int64
	SetRetention(v int32) *UpdateUserBackupFileRequest
	GetRetention() *int32
}

type UpdateUserBackupFileRequest struct {
	// The backup ID. You can call the ListUserBackupFiles operation to query the backup ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-kwwvr7v8t7of********
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The new description of the full backup file.
	//
	// example:
	//
	// CommentTest
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new retention period of the full backup file. Unit: days. Valid values: any non-zero positive integer.
	//
	// example:
	//
	// 30
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s UpdateUserBackupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserBackupFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserBackupFileRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *UpdateUserBackupFileRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateUserBackupFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateUserBackupFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateUserBackupFileRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateUserBackupFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateUserBackupFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateUserBackupFileRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *UpdateUserBackupFileRequest) SetBackupId(v string) *UpdateUserBackupFileRequest {
	s.BackupId = &v
	return s
}

func (s *UpdateUserBackupFileRequest) SetComment(v string) *UpdateUserBackupFileRequest {
	s.Comment = &v
	return s
}

func (s *UpdateUserBackupFileRequest) SetOwnerId(v int64) *UpdateUserBackupFileRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateUserBackupFileRequest) SetRegionId(v string) *UpdateUserBackupFileRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateUserBackupFileRequest) SetResourceGroupId(v string) *UpdateUserBackupFileRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateUserBackupFileRequest) SetResourceOwnerAccount(v string) *UpdateUserBackupFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateUserBackupFileRequest) SetResourceOwnerId(v int64) *UpdateUserBackupFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateUserBackupFileRequest) SetRetention(v int32) *UpdateUserBackupFileRequest {
	s.Retention = &v
	return s
}

func (s *UpdateUserBackupFileRequest) Validate() error {
	return dara.Validate(s)
}
