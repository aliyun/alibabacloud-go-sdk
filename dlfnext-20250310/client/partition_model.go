// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPartition interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Partition
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Partition
	GetCreatedBy() *string
	SetDone(v bool) *Partition
	GetDone() *bool
	SetFileCount(v int64) *Partition
	GetFileCount() *int64
	SetFileSizeInBytes(v int64) *Partition
	GetFileSizeInBytes() *int64
	SetLastFileCreationTime(v int64) *Partition
	GetLastFileCreationTime() *int64
	SetRecordCount(v int64) *Partition
	GetRecordCount() *int64
	SetSpec(v map[string]interface{}) *Partition
	GetSpec() map[string]interface{}
	SetStorageAction(v string) *Partition
	GetStorageAction() *string
	SetStorageActionTimestamp(v int64) *Partition
	GetStorageActionTimestamp() *int64
	SetStorageClass(v string) *Partition
	GetStorageClass() *string
	SetUpdatedAt(v int64) *Partition
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Partition
	GetUpdatedBy() *string
}

type Partition struct {
	CreatedAt              *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy              *string                `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Done                   *bool                  `json:"done,omitempty" xml:"done,omitempty"`
	FileCount              *int64                 `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	FileSizeInBytes        *int64                 `json:"fileSizeInBytes,omitempty" xml:"fileSizeInBytes,omitempty"`
	LastFileCreationTime   *int64                 `json:"lastFileCreationTime,omitempty" xml:"lastFileCreationTime,omitempty"`
	RecordCount            *int64                 `json:"recordCount,omitempty" xml:"recordCount,omitempty"`
	Spec                   map[string]interface{} `json:"spec,omitempty" xml:"spec,omitempty"`
	StorageAction          *string                `json:"storageAction,omitempty" xml:"storageAction,omitempty"`
	StorageActionTimestamp *int64                 `json:"storageActionTimestamp,omitempty" xml:"storageActionTimestamp,omitempty"`
	StorageClass           *string                `json:"storageClass,omitempty" xml:"storageClass,omitempty"`
	UpdatedAt              *int64                 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy              *string                `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Partition) String() string {
	return dara.Prettify(s)
}

func (s Partition) GoString() string {
	return s.String()
}

func (s *Partition) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Partition) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Partition) GetDone() *bool {
	return s.Done
}

func (s *Partition) GetFileCount() *int64 {
	return s.FileCount
}

func (s *Partition) GetFileSizeInBytes() *int64 {
	return s.FileSizeInBytes
}

func (s *Partition) GetLastFileCreationTime() *int64 {
	return s.LastFileCreationTime
}

func (s *Partition) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *Partition) GetSpec() map[string]interface{} {
	return s.Spec
}

func (s *Partition) GetStorageAction() *string {
	return s.StorageAction
}

func (s *Partition) GetStorageActionTimestamp() *int64 {
	return s.StorageActionTimestamp
}

func (s *Partition) GetStorageClass() *string {
	return s.StorageClass
}

func (s *Partition) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Partition) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Partition) SetCreatedAt(v int64) *Partition {
	s.CreatedAt = &v
	return s
}

func (s *Partition) SetCreatedBy(v string) *Partition {
	s.CreatedBy = &v
	return s
}

func (s *Partition) SetDone(v bool) *Partition {
	s.Done = &v
	return s
}

func (s *Partition) SetFileCount(v int64) *Partition {
	s.FileCount = &v
	return s
}

func (s *Partition) SetFileSizeInBytes(v int64) *Partition {
	s.FileSizeInBytes = &v
	return s
}

func (s *Partition) SetLastFileCreationTime(v int64) *Partition {
	s.LastFileCreationTime = &v
	return s
}

func (s *Partition) SetRecordCount(v int64) *Partition {
	s.RecordCount = &v
	return s
}

func (s *Partition) SetSpec(v map[string]interface{}) *Partition {
	s.Spec = v
	return s
}

func (s *Partition) SetStorageAction(v string) *Partition {
	s.StorageAction = &v
	return s
}

func (s *Partition) SetStorageActionTimestamp(v int64) *Partition {
	s.StorageActionTimestamp = &v
	return s
}

func (s *Partition) SetStorageClass(v string) *Partition {
	s.StorageClass = &v
	return s
}

func (s *Partition) SetUpdatedAt(v int64) *Partition {
	s.UpdatedAt = &v
	return s
}

func (s *Partition) SetUpdatedBy(v string) *Partition {
	s.UpdatedBy = &v
	return s
}

func (s *Partition) Validate() error {
	return dara.Validate(s)
}
