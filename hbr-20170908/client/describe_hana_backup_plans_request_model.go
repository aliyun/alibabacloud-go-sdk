// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHanaBackupPlansRequest
	GetClusterId() *string
	SetDatabaseName(v string) *DescribeHanaBackupPlansRequest
	GetDatabaseName() *string
	SetPageNumber(v int32) *DescribeHanaBackupPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaBackupPlansRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeHanaBackupPlansRequest
	GetResourceGroupId() *string
	SetVaultId(v string) *DescribeHanaBackupPlansRequest
	GetVaultId() *string
}

type DescribeHanaBackupPlansRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000chxz******lz7bk
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
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvywqfey5njq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000i4lg4kz******ahl
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaBackupPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaBackupPlansRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaBackupPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaBackupPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaBackupPlansRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHanaBackupPlansRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaBackupPlansRequest) SetClusterId(v string) *DescribeHanaBackupPlansRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetDatabaseName(v string) *DescribeHanaBackupPlansRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetPageNumber(v int32) *DescribeHanaBackupPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetPageSize(v int32) *DescribeHanaBackupPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetResourceGroupId(v string) *DescribeHanaBackupPlansRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetVaultId(v string) *DescribeHanaBackupPlansRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) Validate() error {
	return dara.Validate(s)
}
