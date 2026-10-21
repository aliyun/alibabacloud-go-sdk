// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticTableBaseInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogType(v string) *AgenticTableBaseInfo
	GetCatalogType() *string
	SetDatabaseQualifiedName(v string) *AgenticTableBaseInfo
	GetDatabaseQualifiedName() *string
	SetDatabaseUuid(v string) *AgenticTableBaseInfo
	GetDatabaseUuid() *string
	SetDescription(v string) *AgenticTableBaseInfo
	GetDescription() *string
	SetEngineMeta(v *AgenticTableEngineMeta) *AgenticTableBaseInfo
	GetEngineMeta() *AgenticTableEngineMeta
	SetName(v string) *AgenticTableBaseInfo
	GetName() *string
	SetQualifiedName(v string) *AgenticTableBaseInfo
	GetQualifiedName() *string
	SetTableType(v string) *AgenticTableBaseInfo
	GetTableType() *string
}

type AgenticTableBaseInfo struct {
	// The type of the catalog to which the table belongs. Valid values:
	//
	// - **JDBC**: relational database catalog
	//
	// - **DataLake**: data lake catalog
	//
	// - **Dataset**: knowledge base catalog
	//
	// example:
	//
	// JDBC
	CatalogType *string `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	// The fully qualified name of the database to which the table belongs.
	//
	// example:
	//
	// order_db
	DatabaseQualifiedName *string `json:"DatabaseQualifiedName,omitempty" xml:"DatabaseQualifiedName,omitempty"`
	// The UUID of the database to which the table belongs.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The description of the table.
	//
	// example:
	//
	// Order table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The engine metadata of the table.
	EngineMeta *AgenticTableEngineMeta `json:"EngineMeta,omitempty" xml:"EngineMeta,omitempty"`
	// The table name.
	//
	// example:
	//
	// orders
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The fully qualified name of the table. You can use this name as the TableQualifiedName parameter for GetAgenticTableColumn and GetAgenticTableIndex. For SQL Server, the format is schema name.table name.
	//
	// example:
	//
	// orders
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The table type. Valid values:
	//
	// - **NORMAL**: regular table
	//
	// - **VIEW**: view
	//
	// - **FOREIGN**: foreign table
	//
	// - **SYNONYM**: synonym, supported only for Oracle data sources
	//
	// example:
	//
	// NORMAL
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s AgenticTableBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s AgenticTableBaseInfo) GoString() string {
	return s.String()
}

func (s *AgenticTableBaseInfo) GetCatalogType() *string {
	return s.CatalogType
}

func (s *AgenticTableBaseInfo) GetDatabaseQualifiedName() *string {
	return s.DatabaseQualifiedName
}

func (s *AgenticTableBaseInfo) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *AgenticTableBaseInfo) GetDescription() *string {
	return s.Description
}

func (s *AgenticTableBaseInfo) GetEngineMeta() *AgenticTableEngineMeta {
	return s.EngineMeta
}

func (s *AgenticTableBaseInfo) GetName() *string {
	return s.Name
}

func (s *AgenticTableBaseInfo) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *AgenticTableBaseInfo) GetTableType() *string {
	return s.TableType
}

func (s *AgenticTableBaseInfo) SetCatalogType(v string) *AgenticTableBaseInfo {
	s.CatalogType = &v
	return s
}

func (s *AgenticTableBaseInfo) SetDatabaseQualifiedName(v string) *AgenticTableBaseInfo {
	s.DatabaseQualifiedName = &v
	return s
}

func (s *AgenticTableBaseInfo) SetDatabaseUuid(v string) *AgenticTableBaseInfo {
	s.DatabaseUuid = &v
	return s
}

func (s *AgenticTableBaseInfo) SetDescription(v string) *AgenticTableBaseInfo {
	s.Description = &v
	return s
}

func (s *AgenticTableBaseInfo) SetEngineMeta(v *AgenticTableEngineMeta) *AgenticTableBaseInfo {
	s.EngineMeta = v
	return s
}

func (s *AgenticTableBaseInfo) SetName(v string) *AgenticTableBaseInfo {
	s.Name = &v
	return s
}

func (s *AgenticTableBaseInfo) SetQualifiedName(v string) *AgenticTableBaseInfo {
	s.QualifiedName = &v
	return s
}

func (s *AgenticTableBaseInfo) SetTableType(v string) *AgenticTableBaseInfo {
	s.TableType = &v
	return s
}

func (s *AgenticTableBaseInfo) Validate() error {
	if s.EngineMeta != nil {
		if err := s.EngineMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
