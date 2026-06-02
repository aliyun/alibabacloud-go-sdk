// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaDatabaseEngineMeta interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *OneMetaDatabaseEngineMeta
	GetCatalogName() *string
	SetEncoding(v string) *OneMetaDatabaseEngineMeta
	GetEncoding() *string
	SetSchemaName(v string) *OneMetaDatabaseEngineMeta
	GetSchemaName() *string
	SetStorageCapacity(v int64) *OneMetaDatabaseEngineMeta
	GetStorageCapacity() *int64
}

type OneMetaDatabaseEngineMeta struct {
	CatalogName     *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	Encoding        *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	SchemaName      *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	StorageCapacity *int64  `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
}

func (s OneMetaDatabaseEngineMeta) String() string {
	return dara.Prettify(s)
}

func (s OneMetaDatabaseEngineMeta) GoString() string {
	return s.String()
}

func (s *OneMetaDatabaseEngineMeta) GetCatalogName() *string {
	return s.CatalogName
}

func (s *OneMetaDatabaseEngineMeta) GetEncoding() *string {
	return s.Encoding
}

func (s *OneMetaDatabaseEngineMeta) GetSchemaName() *string {
	return s.SchemaName
}

func (s *OneMetaDatabaseEngineMeta) GetStorageCapacity() *int64 {
	return s.StorageCapacity
}

func (s *OneMetaDatabaseEngineMeta) SetCatalogName(v string) *OneMetaDatabaseEngineMeta {
	s.CatalogName = &v
	return s
}

func (s *OneMetaDatabaseEngineMeta) SetEncoding(v string) *OneMetaDatabaseEngineMeta {
	s.Encoding = &v
	return s
}

func (s *OneMetaDatabaseEngineMeta) SetSchemaName(v string) *OneMetaDatabaseEngineMeta {
	s.SchemaName = &v
	return s
}

func (s *OneMetaDatabaseEngineMeta) SetStorageCapacity(v int64) *OneMetaDatabaseEngineMeta {
	s.StorageCapacity = &v
	return s
}

func (s *OneMetaDatabaseEngineMeta) Validate() error {
	return dara.Validate(s)
}
