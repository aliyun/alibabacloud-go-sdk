// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetFileCount(v int64) *TableSnapshot
	GetFileCount() *int64
	SetFileSizeInBytes(v int64) *TableSnapshot
	GetFileSizeInBytes() *int64
	SetLastFileCreationTime(v int64) *TableSnapshot
	GetLastFileCreationTime() *int64
	SetRecordCount(v int64) *TableSnapshot
	GetRecordCount() *int64
	SetSnapshot(v *Snapshot) *TableSnapshot
	GetSnapshot() *Snapshot
}

type TableSnapshot struct {
	FileCount            *int64    `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	FileSizeInBytes      *int64    `json:"fileSizeInBytes,omitempty" xml:"fileSizeInBytes,omitempty"`
	LastFileCreationTime *int64    `json:"lastFileCreationTime,omitempty" xml:"lastFileCreationTime,omitempty"`
	RecordCount          *int64    `json:"recordCount,omitempty" xml:"recordCount,omitempty"`
	Snapshot             *Snapshot `json:"snapshot,omitempty" xml:"snapshot,omitempty"`
}

func (s TableSnapshot) String() string {
	return dara.Prettify(s)
}

func (s TableSnapshot) GoString() string {
	return s.String()
}

func (s *TableSnapshot) GetFileCount() *int64 {
	return s.FileCount
}

func (s *TableSnapshot) GetFileSizeInBytes() *int64 {
	return s.FileSizeInBytes
}

func (s *TableSnapshot) GetLastFileCreationTime() *int64 {
	return s.LastFileCreationTime
}

func (s *TableSnapshot) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *TableSnapshot) GetSnapshot() *Snapshot {
	return s.Snapshot
}

func (s *TableSnapshot) SetFileCount(v int64) *TableSnapshot {
	s.FileCount = &v
	return s
}

func (s *TableSnapshot) SetFileSizeInBytes(v int64) *TableSnapshot {
	s.FileSizeInBytes = &v
	return s
}

func (s *TableSnapshot) SetLastFileCreationTime(v int64) *TableSnapshot {
	s.LastFileCreationTime = &v
	return s
}

func (s *TableSnapshot) SetRecordCount(v int64) *TableSnapshot {
	s.RecordCount = &v
	return s
}

func (s *TableSnapshot) SetSnapshot(v *Snapshot) *TableSnapshot {
	s.Snapshot = v
	return s
}

func (s *TableSnapshot) Validate() error {
	if s.Snapshot != nil {
		if err := s.Snapshot.Validate(); err != nil {
			return err
		}
	}
	return nil
}
