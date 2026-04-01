// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserBackupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DeleteUserBackupFileRequest
	GetBackupId() *string
	SetOwnerId(v int64) *DeleteUserBackupFileRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteUserBackupFileRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteUserBackupFileRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteUserBackupFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteUserBackupFileRequest
	GetResourceOwnerId() *int64
}

type DeleteUserBackupFileRequest struct {
	// The ID of the full backup file. You can call the ListUserBackupFiles operation to query the information about all full backup files in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-w1haya7e4i25********
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the DescribeRegions operation to query the most recent region list.
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
}

func (s DeleteUserBackupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserBackupFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserBackupFileRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DeleteUserBackupFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteUserBackupFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserBackupFileRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteUserBackupFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteUserBackupFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteUserBackupFileRequest) SetBackupId(v string) *DeleteUserBackupFileRequest {
	s.BackupId = &v
	return s
}

func (s *DeleteUserBackupFileRequest) SetOwnerId(v int64) *DeleteUserBackupFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUserBackupFileRequest) SetRegionId(v string) *DeleteUserBackupFileRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserBackupFileRequest) SetResourceGroupId(v string) *DeleteUserBackupFileRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteUserBackupFileRequest) SetResourceOwnerAccount(v string) *DeleteUserBackupFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUserBackupFileRequest) SetResourceOwnerId(v int64) *DeleteUserBackupFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUserBackupFileRequest) Validate() error {
	return dara.Validate(s)
}
