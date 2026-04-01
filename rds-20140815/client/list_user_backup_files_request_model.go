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
	// The ID of the full backup file.
	//
	// example:
	//
	// b-kwwvr7v8t7of********
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The description of the full backup file.
	//
	// > The system implements a fuzzy match based on the value of this parameter.
	//
	// example:
	//
	// BackupTest
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The URL from which you can download the full backup file that is stored as an object in an Object Storage Service (OSS) bucket. For more information about how to obtain the URL, see [Obtain the access URL after you upload objects](https://help.aliyun.com/document_detail/39607.html).
	//
	// example:
	//
	// https://******.oss-ap-********.aliyuncs.com/backup_qp.xb
	OssUrl  *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to obtain the ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the full backup file. Valid values:
	//
	// 	- **Importing**: The full backup file is being imported.
	//
	// 	- **Failed**: The full backup file fails to be imported.
	//
	// 	- **CheckSucccess**: The full backup file passes the check.
	//
	// 	- **BackupSuccess**: The full backup file is imported.
	//
	// 	- **Deleted**: The full backup file is deleted.
	//
	// example:
	//
	// CheckSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the full backup file.
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
