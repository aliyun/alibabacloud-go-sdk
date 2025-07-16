// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableSummary interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *TableSummary
	GetCreatedAt() *int64
	SetDatabaseName(v string) *TableSummary
	GetDatabaseName() *string
	SetGeneratedDate(v string) *TableSummary
	GetGeneratedDate() *string
	SetLastAccessTime(v int64) *TableSummary
	GetLastAccessTime() *int64
	SetPartitionCount(v int64) *TableSummary
	GetPartitionCount() *int64
	SetPath(v string) *TableSummary
	GetPath() *string
	SetTableName(v string) *TableSummary
	GetTableName() *string
	SetTotalFileCount(v int64) *TableSummary
	GetTotalFileCount() *int64
	SetTotalFileSizeInBytes(v int64) *TableSummary
	GetTotalFileSizeInBytes() *int64
}

type TableSummary struct {
	// Latest snapshot storage size
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Database name
	DatabaseName   *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	GeneratedDate  *string `json:"generatedDate,omitempty" xml:"generatedDate,omitempty"`
	LastAccessTime *int64  `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// Creation timestamp in milliseconds
	PartitionCount *int64  `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	// Table name
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// 30-day access count
	TotalFileCount       *int64 `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	TotalFileSizeInBytes *int64 `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
}

func (s TableSummary) String() string {
	return dara.Prettify(s)
}

func (s TableSummary) GoString() string {
	return s.String()
}

func (s *TableSummary) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *TableSummary) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *TableSummary) GetGeneratedDate() *string {
	return s.GeneratedDate
}

func (s *TableSummary) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *TableSummary) GetPartitionCount() *int64 {
	return s.PartitionCount
}

func (s *TableSummary) GetPath() *string {
	return s.Path
}

func (s *TableSummary) GetTableName() *string {
	return s.TableName
}

func (s *TableSummary) GetTotalFileCount() *int64 {
	return s.TotalFileCount
}

func (s *TableSummary) GetTotalFileSizeInBytes() *int64 {
	return s.TotalFileSizeInBytes
}

func (s *TableSummary) SetCreatedAt(v int64) *TableSummary {
	s.CreatedAt = &v
	return s
}

func (s *TableSummary) SetDatabaseName(v string) *TableSummary {
	s.DatabaseName = &v
	return s
}

func (s *TableSummary) SetGeneratedDate(v string) *TableSummary {
	s.GeneratedDate = &v
	return s
}

func (s *TableSummary) SetLastAccessTime(v int64) *TableSummary {
	s.LastAccessTime = &v
	return s
}

func (s *TableSummary) SetPartitionCount(v int64) *TableSummary {
	s.PartitionCount = &v
	return s
}

func (s *TableSummary) SetPath(v string) *TableSummary {
	s.Path = &v
	return s
}

func (s *TableSummary) SetTableName(v string) *TableSummary {
	s.TableName = &v
	return s
}

func (s *TableSummary) SetTotalFileCount(v int64) *TableSummary {
	s.TotalFileCount = &v
	return s
}

func (s *TableSummary) SetTotalFileSizeInBytes(v int64) *TableSummary {
	s.TotalFileSizeInBytes = &v
	return s
}

func (s *TableSummary) Validate() error {
	return dara.Validate(s)
}
