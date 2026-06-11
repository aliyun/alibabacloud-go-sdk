// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticDatabaseEngineMeta interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *AgenticDatabaseEngineMeta
	GetCatalogName() *string
	SetEncoding(v string) *AgenticDatabaseEngineMeta
	GetEncoding() *string
	SetSchemaName(v string) *AgenticDatabaseEngineMeta
	GetSchemaName() *string
	SetStorageCapacity(v int64) *AgenticDatabaseEngineMeta
	GetStorageCapacity() *int64
}

type AgenticDatabaseEngineMeta struct {
	// The name of the database catalog.
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The character encoding for the database.
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The name of the database schema.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The storage capacity of the database, in GB.
	StorageCapacity *int64 `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
}

func (s AgenticDatabaseEngineMeta) String() string {
	return dara.Prettify(s)
}

func (s AgenticDatabaseEngineMeta) GoString() string {
	return s.String()
}

func (s *AgenticDatabaseEngineMeta) GetCatalogName() *string {
	return s.CatalogName
}

func (s *AgenticDatabaseEngineMeta) GetEncoding() *string {
	return s.Encoding
}

func (s *AgenticDatabaseEngineMeta) GetSchemaName() *string {
	return s.SchemaName
}

func (s *AgenticDatabaseEngineMeta) GetStorageCapacity() *int64 {
	return s.StorageCapacity
}

func (s *AgenticDatabaseEngineMeta) SetCatalogName(v string) *AgenticDatabaseEngineMeta {
	s.CatalogName = &v
	return s
}

func (s *AgenticDatabaseEngineMeta) SetEncoding(v string) *AgenticDatabaseEngineMeta {
	s.Encoding = &v
	return s
}

func (s *AgenticDatabaseEngineMeta) SetSchemaName(v string) *AgenticDatabaseEngineMeta {
	s.SchemaName = &v
	return s
}

func (s *AgenticDatabaseEngineMeta) SetStorageCapacity(v int64) *AgenticDatabaseEngineMeta {
	s.StorageCapacity = &v
	return s
}

func (s *AgenticDatabaseEngineMeta) Validate() error {
	return dara.Validate(s)
}
