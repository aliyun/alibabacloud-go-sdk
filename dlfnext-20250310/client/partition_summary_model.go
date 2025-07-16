// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPartitionSummary interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *PartitionSummary
	GetCreatedAt() *int64
	SetDatabaseName(v string) *PartitionSummary
	GetDatabaseName() *string
	SetLastAccessTime(v int64) *PartitionSummary
	GetLastAccessTime() *int64
	SetPartitionName(v string) *PartitionSummary
	GetPartitionName() *string
	SetTableName(v string) *PartitionSummary
	GetTableName() *string
	SetTotalFileCount(v int64) *PartitionSummary
	GetTotalFileCount() *int64
	SetTotalFileSizeInBytes(v int64) *PartitionSummary
	GetTotalFileSizeInBytes() *int64
	SetUpdatedAt(v int64) *PartitionSummary
	GetUpdatedAt() *int64
}

type PartitionSummary struct {
	// Partition creation timestamp in milliseconds
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Database name
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// Total files in partition
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// Partition identifier
	PartitionName *string `json:"partitionName,omitempty" xml:"partitionName,omitempty"`
	// Table name
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// 24h access count
	TotalFileCount *int64 `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	// Last data access timestamp in milliseconds
	TotalFileSizeInBytes *int64 `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
	UpdatedAt            *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s PartitionSummary) String() string {
	return dara.Prettify(s)
}

func (s PartitionSummary) GoString() string {
	return s.String()
}

func (s *PartitionSummary) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *PartitionSummary) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *PartitionSummary) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *PartitionSummary) GetPartitionName() *string {
	return s.PartitionName
}

func (s *PartitionSummary) GetTableName() *string {
	return s.TableName
}

func (s *PartitionSummary) GetTotalFileCount() *int64 {
	return s.TotalFileCount
}

func (s *PartitionSummary) GetTotalFileSizeInBytes() *int64 {
	return s.TotalFileSizeInBytes
}

func (s *PartitionSummary) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *PartitionSummary) SetCreatedAt(v int64) *PartitionSummary {
	s.CreatedAt = &v
	return s
}

func (s *PartitionSummary) SetDatabaseName(v string) *PartitionSummary {
	s.DatabaseName = &v
	return s
}

func (s *PartitionSummary) SetLastAccessTime(v int64) *PartitionSummary {
	s.LastAccessTime = &v
	return s
}

func (s *PartitionSummary) SetPartitionName(v string) *PartitionSummary {
	s.PartitionName = &v
	return s
}

func (s *PartitionSummary) SetTableName(v string) *PartitionSummary {
	s.TableName = &v
	return s
}

func (s *PartitionSummary) SetTotalFileCount(v int64) *PartitionSummary {
	s.TotalFileCount = &v
	return s
}

func (s *PartitionSummary) SetTotalFileSizeInBytes(v int64) *PartitionSummary {
	s.TotalFileSizeInBytes = &v
	return s
}

func (s *PartitionSummary) SetUpdatedAt(v int64) *PartitionSummary {
	s.UpdatedAt = &v
	return s
}

func (s *PartitionSummary) Validate() error {
	return dara.Validate(s)
}
