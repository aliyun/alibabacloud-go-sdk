// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatabaseSummary interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *DatabaseSummary
	GetCreatedAt() *int64
	SetDatabaseName(v string) *DatabaseSummary
	GetDatabaseName() *string
	SetGeneratedDate(v string) *DatabaseSummary
	GetGeneratedDate() *string
	SetLocation(v string) *DatabaseSummary
	GetLocation() *string
	SetObjTypeArchiveSize(v int64) *DatabaseSummary
	GetObjTypeArchiveSize() *int64
	SetObjTypeColdArchiveSize(v int64) *DatabaseSummary
	GetObjTypeColdArchiveSize() *int64
	SetObjTypeIaSize(v int64) *DatabaseSummary
	GetObjTypeIaSize() *int64
	SetObjTypeStandardSize(v int64) *DatabaseSummary
	GetObjTypeStandardSize() *int64
	SetPartitionCount(v int64) *DatabaseSummary
	GetPartitionCount() *int64
	SetTableCount(v int64) *DatabaseSummary
	GetTableCount() *int64
	SetTotalFileCount(v int64) *DatabaseSummary
	GetTotalFileCount() *int64
	SetTotalFileSizeInBytes(v int64) *DatabaseSummary
	GetTotalFileSizeInBytes() *int64
	SetTotalMetaSizeInBytes(v int64) *DatabaseSummary
	GetTotalMetaSizeInBytes() *int64
}

type DatabaseSummary struct {
	// Creation timestamp in milliseconds
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 库名 - Database name
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// Last profile update date in format yyyyMMdd
	GeneratedDate *string `json:"generatedDate,omitempty" xml:"generatedDate,omitempty"`
	// Storage location URI
	Location               *string `json:"location,omitempty" xml:"location,omitempty"`
	ObjTypeArchiveSize     *int64  `json:"objTypeArchiveSize,omitempty" xml:"objTypeArchiveSize,omitempty"`
	ObjTypeColdArchiveSize *int64  `json:"objTypeColdArchiveSize,omitempty" xml:"objTypeColdArchiveSize,omitempty"`
	ObjTypeIaSize          *int64  `json:"objTypeIaSize,omitempty" xml:"objTypeIaSize,omitempty"`
	ObjTypeStandardSize    *int64  `json:"objTypeStandardSize,omitempty" xml:"objTypeStandardSize,omitempty"`
	PartitionCount         *int64  `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// Total storage in bytes
	TableCount     *int64 `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
	TotalFileCount *int64 `json:"totalFileCount,omitempty" xml:"totalFileCount,omitempty"`
	// Total file count
	TotalFileSizeInBytes *int64 `json:"totalFileSizeInBytes,omitempty" xml:"totalFileSizeInBytes,omitempty"`
	TotalMetaSizeInBytes *int64 `json:"totalMetaSizeInBytes,omitempty" xml:"totalMetaSizeInBytes,omitempty"`
}

func (s DatabaseSummary) String() string {
	return dara.Prettify(s)
}

func (s DatabaseSummary) GoString() string {
	return s.String()
}

func (s *DatabaseSummary) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *DatabaseSummary) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DatabaseSummary) GetGeneratedDate() *string {
	return s.GeneratedDate
}

func (s *DatabaseSummary) GetLocation() *string {
	return s.Location
}

func (s *DatabaseSummary) GetObjTypeArchiveSize() *int64 {
	return s.ObjTypeArchiveSize
}

func (s *DatabaseSummary) GetObjTypeColdArchiveSize() *int64 {
	return s.ObjTypeColdArchiveSize
}

func (s *DatabaseSummary) GetObjTypeIaSize() *int64 {
	return s.ObjTypeIaSize
}

func (s *DatabaseSummary) GetObjTypeStandardSize() *int64 {
	return s.ObjTypeStandardSize
}

func (s *DatabaseSummary) GetPartitionCount() *int64 {
	return s.PartitionCount
}

func (s *DatabaseSummary) GetTableCount() *int64 {
	return s.TableCount
}

func (s *DatabaseSummary) GetTotalFileCount() *int64 {
	return s.TotalFileCount
}

func (s *DatabaseSummary) GetTotalFileSizeInBytes() *int64 {
	return s.TotalFileSizeInBytes
}

func (s *DatabaseSummary) GetTotalMetaSizeInBytes() *int64 {
	return s.TotalMetaSizeInBytes
}

func (s *DatabaseSummary) SetCreatedAt(v int64) *DatabaseSummary {
	s.CreatedAt = &v
	return s
}

func (s *DatabaseSummary) SetDatabaseName(v string) *DatabaseSummary {
	s.DatabaseName = &v
	return s
}

func (s *DatabaseSummary) SetGeneratedDate(v string) *DatabaseSummary {
	s.GeneratedDate = &v
	return s
}

func (s *DatabaseSummary) SetLocation(v string) *DatabaseSummary {
	s.Location = &v
	return s
}

func (s *DatabaseSummary) SetObjTypeArchiveSize(v int64) *DatabaseSummary {
	s.ObjTypeArchiveSize = &v
	return s
}

func (s *DatabaseSummary) SetObjTypeColdArchiveSize(v int64) *DatabaseSummary {
	s.ObjTypeColdArchiveSize = &v
	return s
}

func (s *DatabaseSummary) SetObjTypeIaSize(v int64) *DatabaseSummary {
	s.ObjTypeIaSize = &v
	return s
}

func (s *DatabaseSummary) SetObjTypeStandardSize(v int64) *DatabaseSummary {
	s.ObjTypeStandardSize = &v
	return s
}

func (s *DatabaseSummary) SetPartitionCount(v int64) *DatabaseSummary {
	s.PartitionCount = &v
	return s
}

func (s *DatabaseSummary) SetTableCount(v int64) *DatabaseSummary {
	s.TableCount = &v
	return s
}

func (s *DatabaseSummary) SetTotalFileCount(v int64) *DatabaseSummary {
	s.TotalFileCount = &v
	return s
}

func (s *DatabaseSummary) SetTotalFileSizeInBytes(v int64) *DatabaseSummary {
	s.TotalFileSizeInBytes = &v
	return s
}

func (s *DatabaseSummary) SetTotalMetaSizeInBytes(v int64) *DatabaseSummary {
	s.TotalMetaSizeInBytes = &v
	return s
}

func (s *DatabaseSummary) Validate() error {
	return dara.Validate(s)
}
