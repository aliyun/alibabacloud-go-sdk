// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRecordId(v string) *DescribeBackupTablesRequest
	GetBackupRecordId() *string
	SetClusterId(v string) *DescribeBackupTablesRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeBackupTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupTablesRequest
	GetPageSize() *int32
}

type DescribeBackupTablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2020110302xxxx
	BackupRecordId *string `json:"BackupRecordId,omitempty" xml:"BackupRecordId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackupTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesRequest) GetBackupRecordId() *string {
	return s.BackupRecordId
}

func (s *DescribeBackupTablesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupTablesRequest) SetBackupRecordId(v string) *DescribeBackupTablesRequest {
	s.BackupRecordId = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetClusterId(v string) *DescribeBackupTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetPageNumber(v int32) *DescribeBackupTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetPageSize(v int32) *DescribeBackupTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupTablesRequest) Validate() error {
	return dara.Validate(s)
}
