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
	Checksum        *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataBytes       *int64  `json:"DataBytes,omitempty" xml:"DataBytes,omitempty"`
	Encoding        *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	Engine          *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	FullChecksum    *string `json:"FullChecksum,omitempty" xml:"FullChecksum,omitempty"`
	IndexBytes      *int64  `json:"IndexBytes,omitempty" xml:"IndexBytes,omitempty"`
	LastDdlTime     *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	NumRows         *int64  `json:"NumRows,omitempty" xml:"NumRows,omitempty"`
	RefInfo         *string `json:"RefInfo,omitempty" xml:"RefInfo,omitempty"`
	StorageCapacity *int64  `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
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
