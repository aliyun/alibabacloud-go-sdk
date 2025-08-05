// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaRestoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v int64) *DescribeHanaRestoresRequest
	GetBackupId() *int64
	SetClusterId(v string) *DescribeHanaRestoresRequest
	GetClusterId() *string
	SetDatabaseName(v string) *DescribeHanaRestoresRequest
	GetDatabaseName() *string
	SetPageNumber(v int32) *DescribeHanaRestoresRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaRestoresRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeHanaRestoresRequest
	GetResourceGroupId() *string
	SetRestoreId(v string) *DescribeHanaRestoresRequest
	GetRestoreId() *string
	SetRestoreStatus(v string) *DescribeHanaRestoresRequest
	GetRestoreStatus() *string
	SetVaultId(v string) *DescribeHanaRestoresRequest
	GetVaultId() *string
}

type DescribeHanaRestoresRequest struct {
	// The backup ID.
	//
	// example:
	//
	// 1632754800158
	BackupId *int64 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000b******soejg
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.\\`
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4ebtpkzx7zy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the restore job.
	//
	// example:
	//
	// r-0007o3vqfukfe92hvf13
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	// The status of the restore job. Valid values:
	//
	// 	- **RUNNING**: The job is running.
	//
	// 	- **COMPLETE**: The job is completed.
	//
	// 	- **PARTIAL_COMPLETE**: The job is partially completed.
	//
	// 	- **FAILED**: The job failed.
	//
	// 	- **CANCELED**: The job is canceled.
	//
	// 	- **EXPIRED**: The job timed out.
	//
	// example:
	//
	// COMPLETE
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000au6bq******mpu
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaRestoresRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRestoresRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresRequest) GetBackupId() *int64 {
	return s.BackupId
}

func (s *DescribeHanaRestoresRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaRestoresRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaRestoresRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaRestoresRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaRestoresRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHanaRestoresRequest) GetRestoreId() *string {
	return s.RestoreId
}

func (s *DescribeHanaRestoresRequest) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *DescribeHanaRestoresRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaRestoresRequest) SetBackupId(v int64) *DescribeHanaRestoresRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetClusterId(v string) *DescribeHanaRestoresRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetDatabaseName(v string) *DescribeHanaRestoresRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetPageNumber(v int32) *DescribeHanaRestoresRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetPageSize(v int32) *DescribeHanaRestoresRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetResourceGroupId(v string) *DescribeHanaRestoresRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetRestoreId(v string) *DescribeHanaRestoresRequest {
	s.RestoreId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetRestoreStatus(v string) *DescribeHanaRestoresRequest {
	s.RestoreStatus = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetVaultId(v string) *DescribeHanaRestoresRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) Validate() error {
	return dara.Validate(s)
}
