// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaTableEngineMeta interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *OneMetaTableEngineMeta
	GetCreateTime() *string
	SetDataBytes(v int64) *OneMetaTableEngineMeta
	GetDataBytes() *int64
	SetEncoding(v string) *OneMetaTableEngineMeta
	GetEncoding() *string
	SetEngine(v string) *OneMetaTableEngineMeta
	GetEngine() *string
	SetIndexBytes(v int64) *OneMetaTableEngineMeta
	GetIndexBytes() *int64
	SetLastDdlTime(v string) *OneMetaTableEngineMeta
	GetLastDdlTime() *string
	SetNumRows(v int64) *OneMetaTableEngineMeta
	GetNumRows() *int64
	SetRefInfo(v string) *OneMetaTableEngineMeta
	GetRefInfo() *string
	SetStorageCapacity(v int64) *OneMetaTableEngineMeta
	GetStorageCapacity() *int64
	SetTableSchemaName(v string) *OneMetaTableEngineMeta
	GetTableSchemaName() *string
}

type OneMetaTableEngineMeta struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataBytes       *int64  `json:"DataBytes,omitempty" xml:"DataBytes,omitempty"`
	Encoding        *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	Engine          *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IndexBytes      *int64  `json:"IndexBytes,omitempty" xml:"IndexBytes,omitempty"`
	LastDdlTime     *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	NumRows         *int64  `json:"NumRows,omitempty" xml:"NumRows,omitempty"`
	RefInfo         *string `json:"RefInfo,omitempty" xml:"RefInfo,omitempty"`
	StorageCapacity *int64  `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
}

func (s OneMetaTableEngineMeta) String() string {
	return dara.Prettify(s)
}

func (s OneMetaTableEngineMeta) GoString() string {
	return s.String()
}

func (s *OneMetaTableEngineMeta) GetCreateTime() *string {
	return s.CreateTime
}

func (s *OneMetaTableEngineMeta) GetDataBytes() *int64 {
	return s.DataBytes
}

func (s *OneMetaTableEngineMeta) GetEncoding() *string {
	return s.Encoding
}

func (s *OneMetaTableEngineMeta) GetEngine() *string {
	return s.Engine
}

func (s *OneMetaTableEngineMeta) GetIndexBytes() *int64 {
	return s.IndexBytes
}

func (s *OneMetaTableEngineMeta) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *OneMetaTableEngineMeta) GetNumRows() *int64 {
	return s.NumRows
}

func (s *OneMetaTableEngineMeta) GetRefInfo() *string {
	return s.RefInfo
}

func (s *OneMetaTableEngineMeta) GetStorageCapacity() *int64 {
	return s.StorageCapacity
}

func (s *OneMetaTableEngineMeta) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *OneMetaTableEngineMeta) SetCreateTime(v string) *OneMetaTableEngineMeta {
	s.CreateTime = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetDataBytes(v int64) *OneMetaTableEngineMeta {
	s.DataBytes = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetEncoding(v string) *OneMetaTableEngineMeta {
	s.Encoding = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetEngine(v string) *OneMetaTableEngineMeta {
	s.Engine = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetIndexBytes(v int64) *OneMetaTableEngineMeta {
	s.IndexBytes = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetLastDdlTime(v string) *OneMetaTableEngineMeta {
	s.LastDdlTime = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetNumRows(v int64) *OneMetaTableEngineMeta {
	s.NumRows = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetRefInfo(v string) *OneMetaTableEngineMeta {
	s.RefInfo = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetStorageCapacity(v int64) *OneMetaTableEngineMeta {
	s.StorageCapacity = &v
	return s
}

func (s *OneMetaTableEngineMeta) SetTableSchemaName(v string) *OneMetaTableEngineMeta {
	s.TableSchemaName = &v
	return s
}

func (s *OneMetaTableEngineMeta) Validate() error {
	return dara.Validate(s)
}
