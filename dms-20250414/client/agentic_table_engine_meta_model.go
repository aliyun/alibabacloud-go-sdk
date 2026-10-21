// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticTableEngineMeta interface {
	dara.Model
	String() string
	GoString() string
	SetChecksum(v string) *AgenticTableEngineMeta
	GetChecksum() *string
	SetCreateTime(v string) *AgenticTableEngineMeta
	GetCreateTime() *string
	SetDataBytes(v int64) *AgenticTableEngineMeta
	GetDataBytes() *int64
	SetEncoding(v string) *AgenticTableEngineMeta
	GetEncoding() *string
	SetEngine(v string) *AgenticTableEngineMeta
	GetEngine() *string
	SetFullChecksum(v string) *AgenticTableEngineMeta
	GetFullChecksum() *string
	SetIndexBytes(v int64) *AgenticTableEngineMeta
	GetIndexBytes() *int64
	SetLastDdlTime(v string) *AgenticTableEngineMeta
	GetLastDdlTime() *string
	SetNumRows(v int64) *AgenticTableEngineMeta
	GetNumRows() *int64
	SetRefInfo(v string) *AgenticTableEngineMeta
	GetRefInfo() *string
	SetStorageCapacity(v int64) *AgenticTableEngineMeta
	GetStorageCapacity() *int64
	SetTableSchemaName(v string) *AgenticTableEngineMeta
	GetTableSchemaName() *string
}

type AgenticTableEngineMeta struct {
	// The checksum of the table metadata.
	//
	// example:
	//
	// 9f86d081884c7d65****
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// The time when the table was created. The format is yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-01-01 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data size of the table. Unit: bytes.
	//
	// example:
	//
	// 786432
	DataBytes *int64 `json:"DataBytes,omitempty" xml:"DataBytes,omitempty"`
	// The character set encoding of the table.
	//
	// example:
	//
	// utf8mb4
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The storage engine of the table.
	//
	// example:
	//
	// InnoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The full checksum of the table metadata.
	//
	// example:
	//
	// 2c26b46b68ffc68f****
	FullChecksum *string `json:"FullChecksum,omitempty" xml:"FullChecksum,omitempty"`
	// The index size of the table. Unit: bytes.
	//
	// example:
	//
	// 262144
	IndexBytes *int64 `json:"IndexBytes,omitempty" xml:"IndexBytes,omitempty"`
	// The time when the last DDL operation was executed on the table. The format is yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-01-02 10:00:00
	LastDdlTime *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 10000
	NumRows *int64 `json:"NumRows,omitempty" xml:"NumRows,omitempty"`
	// The reference information of the table.
	//
	// example:
	//
	// order_db.orders_view
	RefInfo *string `json:"RefInfo,omitempty" xml:"RefInfo,omitempty"`
	// The storage capacity of the table. Unit: MB.
	//
	// example:
	//
	// 16
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
	// The name of the schema to which the table belongs.
	//
	// example:
	//
	// public
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
}

func (s AgenticTableEngineMeta) String() string {
	return dara.Prettify(s)
}

func (s AgenticTableEngineMeta) GoString() string {
	return s.String()
}

func (s *AgenticTableEngineMeta) GetChecksum() *string {
	return s.Checksum
}

func (s *AgenticTableEngineMeta) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AgenticTableEngineMeta) GetDataBytes() *int64 {
	return s.DataBytes
}

func (s *AgenticTableEngineMeta) GetEncoding() *string {
	return s.Encoding
}

func (s *AgenticTableEngineMeta) GetEngine() *string {
	return s.Engine
}

func (s *AgenticTableEngineMeta) GetFullChecksum() *string {
	return s.FullChecksum
}

func (s *AgenticTableEngineMeta) GetIndexBytes() *int64 {
	return s.IndexBytes
}

func (s *AgenticTableEngineMeta) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *AgenticTableEngineMeta) GetNumRows() *int64 {
	return s.NumRows
}

func (s *AgenticTableEngineMeta) GetRefInfo() *string {
	return s.RefInfo
}

func (s *AgenticTableEngineMeta) GetStorageCapacity() *int64 {
	return s.StorageCapacity
}

func (s *AgenticTableEngineMeta) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *AgenticTableEngineMeta) SetChecksum(v string) *AgenticTableEngineMeta {
	s.Checksum = &v
	return s
}

func (s *AgenticTableEngineMeta) SetCreateTime(v string) *AgenticTableEngineMeta {
	s.CreateTime = &v
	return s
}

func (s *AgenticTableEngineMeta) SetDataBytes(v int64) *AgenticTableEngineMeta {
	s.DataBytes = &v
	return s
}

func (s *AgenticTableEngineMeta) SetEncoding(v string) *AgenticTableEngineMeta {
	s.Encoding = &v
	return s
}

func (s *AgenticTableEngineMeta) SetEngine(v string) *AgenticTableEngineMeta {
	s.Engine = &v
	return s
}

func (s *AgenticTableEngineMeta) SetFullChecksum(v string) *AgenticTableEngineMeta {
	s.FullChecksum = &v
	return s
}

func (s *AgenticTableEngineMeta) SetIndexBytes(v int64) *AgenticTableEngineMeta {
	s.IndexBytes = &v
	return s
}

func (s *AgenticTableEngineMeta) SetLastDdlTime(v string) *AgenticTableEngineMeta {
	s.LastDdlTime = &v
	return s
}

func (s *AgenticTableEngineMeta) SetNumRows(v int64) *AgenticTableEngineMeta {
	s.NumRows = &v
	return s
}

func (s *AgenticTableEngineMeta) SetRefInfo(v string) *AgenticTableEngineMeta {
	s.RefInfo = &v
	return s
}

func (s *AgenticTableEngineMeta) SetStorageCapacity(v int64) *AgenticTableEngineMeta {
	s.StorageCapacity = &v
	return s
}

func (s *AgenticTableEngineMeta) SetTableSchemaName(v string) *AgenticTableEngineMeta {
	s.TableSchemaName = &v
	return s
}

func (s *AgenticTableEngineMeta) Validate() error {
	return dara.Validate(s)
}
