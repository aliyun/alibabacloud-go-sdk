// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserBackupFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ListUserBackupFilesRequest
	GetBackupId() *string
	SetComment(v string) *ListUserBackupFilesRequest
	GetComment() *string
	SetOssUrl(v string) *ListUserBackupFilesRequest
	GetOssUrl() *string
	SetOwnerId(v int64) *ListUserBackupFilesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListUserBackupFilesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListUserBackupFilesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListUserBackupFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListUserBackupFilesRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListUserBackupFilesRequest
	GetStatus() *string
	SetTags(v string) *ListUserBackupFilesRequest
	GetTags() *string
}

type ListUserBackupFilesRequest struct {
	// The user backup ID.
	//
	// example:
	//
	// b-kwwvr7v8t7of****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The comment of the user backup to query.
	//
	// >You can enter part of the comment for fuzzy matching.
	//
	// example:
	//
	// BackupTest
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The OSS download URL of the user backup file. For information about how to obtain the OSS download URL of a user backup file, see [How do I obtain the URL of an uploaded object?](https://help.aliyun.com/document_detail/39607.html).
	//
	// example:
	//
	// https://****.oss-ap-****.aliyuncs.com/backup_qp.xb
	OssUrl  *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call DescribeRegions to query the available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call DescribeDBInstanceAttribute to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the user backup file. Valid values:
	//
	// 	- **Importing**: The backup is being imported.
	//
	// 	- **Failed**: The import failed.
	//
	// 	- **CheckSuccess**: The verification passed.
	//
	// 	- **BackupSuccess**: The import succeeded.
	//
	// 	- **Deleted**: The backup is deleted.
	//
	// example:
	//
	// CheckSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information used to query the user backup.
	//
	// example:
	//
	// key1:value1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListUserBackupFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *ListUserBackupFilesRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ListUserBackupFilesRequest) GetComment() *string {
	return s.Comment
}

func (s *ListUserBackupFilesRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *ListUserBackupFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListUserBackupFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserBackupFilesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListUserBackupFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListUserBackupFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListUserBackupFilesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUserBackupFilesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListUserBackupFilesRequest) SetBackupId(v string) *ListUserBackupFilesRequest {
	s.BackupId = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetComment(v string) *ListUserBackupFilesRequest {
	s.Comment = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetOssUrl(v string) *ListUserBackupFilesRequest {
	s.OssUrl = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetOwnerId(v int64) *ListUserBackupFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetRegionId(v string) *ListUserBackupFilesRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetResourceGroupId(v string) *ListUserBackupFilesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetResourceOwnerAccount(v string) *ListUserBackupFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetResourceOwnerId(v int64) *ListUserBackupFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetStatus(v string) *ListUserBackupFilesRequest {
	s.Status = &v
	return s
}

func (s *ListUserBackupFilesRequest) SetTags(v string) *ListUserBackupFilesRequest {
	s.Tags = &v
	return s
}

func (s *ListUserBackupFilesRequest) Validate() error {
	return dara.Validate(s)
}
