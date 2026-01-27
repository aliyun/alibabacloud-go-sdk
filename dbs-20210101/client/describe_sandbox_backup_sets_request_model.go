// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxBackupSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeSandboxBackupSetsRequest
	GetBackupPlanId() *string
	SetBackupSetId(v string) *DescribeSandboxBackupSetsRequest
	GetBackupSetId() *string
	SetPageNumber(v string) *DescribeSandboxBackupSetsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeSandboxBackupSetsRequest
	GetPageSize() *string
}

type DescribeSandboxBackupSetsRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/437215.html) operation to query the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](https://help.aliyun.com/document_detail/193091.html) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the backup set. If this parameter is specified, only the snapshot of the specified backup set is returned. If this parameter is not specified, all the snapshots of the backup schedule are returned.
	//
	// example:
	//
	// 1xxxx2xxxxx1e
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- 30: This is the default value.
	//
	// 	- 50\\.
	//
	// 	- 100\\.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSandboxBackupSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxBackupSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeSandboxBackupSetsRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeSandboxBackupSetsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSandboxBackupSetsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSandboxBackupSetsRequest) SetBackupPlanId(v string) *DescribeSandboxBackupSetsRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetBackupSetId(v string) *DescribeSandboxBackupSetsRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetPageNumber(v string) *DescribeSandboxBackupSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetPageSize(v string) *DescribeSandboxBackupSetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) Validate() error {
	return dara.Validate(s)
}
