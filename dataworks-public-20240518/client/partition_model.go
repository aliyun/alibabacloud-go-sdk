// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPartition interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *Partition
	GetCreateTime() *int64
	SetDataSize(v int64) *Partition
	GetDataSize() *int64
	SetModifyTime(v int64) *Partition
	GetModifyTime() *int64
	SetName(v string) *Partition
	GetName() *string
	SetRecordCount(v int64) *Partition
	GetRecordCount() *int64
	SetTableId(v string) *Partition
	GetTableId() *string
}

type Partition struct {
	// The creation time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1700192563000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The storage size of the partition. Unit: bytes.
	//
	// example:
	//
	// 4096
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The modification time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1700192563000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The partition name.
	//
	// example:
	//
	// ds=20250101
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of records in the partition.
	//
	// example:
	//
	// 1000000
	RecordCount *int64 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The ID of the table to which the partition belongs. You can configure this parameter by referring to the `Table` topic.
	//
	// example:
	//
	// maxcompute-table:accountId::project::table
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s Partition) String() string {
	return dara.Prettify(s)
}

func (s Partition) GoString() string {
	return s.String()
}

func (s *Partition) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Partition) GetDataSize() *int64 {
	return s.DataSize
}

func (s *Partition) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *Partition) GetName() *string {
	return s.Name
}

func (s *Partition) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *Partition) GetTableId() *string {
	return s.TableId
}

func (s *Partition) SetCreateTime(v int64) *Partition {
	s.CreateTime = &v
	return s
}

func (s *Partition) SetDataSize(v int64) *Partition {
	s.DataSize = &v
	return s
}

func (s *Partition) SetModifyTime(v int64) *Partition {
	s.ModifyTime = &v
	return s
}

func (s *Partition) SetName(v string) *Partition {
	s.Name = &v
	return s
}

func (s *Partition) SetRecordCount(v int64) *Partition {
	s.RecordCount = &v
	return s
}

func (s *Partition) SetTableId(v string) *Partition {
	s.TableId = &v
	return s
}

func (s *Partition) Validate() error {
	return dara.Validate(s)
}
