// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDBsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeBackupDBsRequest
	GetBackupId() *string
	SetOwnerAccount(v string) *DescribeBackupDBsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupDBsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBackupDBsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupDBsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeBackupDBsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeBackupDBsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupDBsRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *DescribeBackupDBsRequest
	GetRestoreTime() *string
	SetSourceDBInstance(v string) *DescribeBackupDBsRequest
	GetSourceDBInstance() *string
}

type DescribeBackupDBsRequest struct {
	// The backup ID.
	//
	// >
	//
	// 	- You can call the [DescribeBackups](https://help.aliyun.com/document_detail/62172.html) operation to query the backup ID.
	//
	// 	- You must specify one of the **RestoreTime*	- and BackupId parameters.
	//
	// example:
	//
	// 5664****
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: 30. Valid values: **30**, **50**, and **100**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-ax68****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which the instance is restored. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >
	//
	// 	- The time can be a point in time within the previous seven days. The time must be earlier than the current time, but later than the time when the instance was created.
	//
	// 	- You must specify one of the RestoreTime and **BackupId*	- parameters.
	//
	// example:
	//
	// 2019-08-22T12:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// dds-bp2286****
	SourceDBInstance *string `json:"SourceDBInstance,omitempty" xml:"SourceDBInstance,omitempty"`
}

func (s DescribeBackupDBsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDBsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupDBsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupDBsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupDBsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupDBsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupDBsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBackupDBsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupDBsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupDBsRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *DescribeBackupDBsRequest) GetSourceDBInstance() *string {
	return s.SourceDBInstance
}

func (s *DescribeBackupDBsRequest) SetBackupId(v string) *DescribeBackupDBsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetOwnerAccount(v string) *DescribeBackupDBsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetOwnerId(v int64) *DescribeBackupDBsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetPageNumber(v int32) *DescribeBackupDBsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetPageSize(v int32) *DescribeBackupDBsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetResourceGroupId(v string) *DescribeBackupDBsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetResourceOwnerAccount(v string) *DescribeBackupDBsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetResourceOwnerId(v int64) *DescribeBackupDBsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetRestoreTime(v string) *DescribeBackupDBsRequest {
	s.RestoreTime = &v
	return s
}

func (s *DescribeBackupDBsRequest) SetSourceDBInstance(v string) *DescribeBackupDBsRequest {
	s.SourceDBInstance = &v
	return s
}

func (s *DescribeBackupDBsRequest) Validate() error {
	return dara.Validate(s)
}
