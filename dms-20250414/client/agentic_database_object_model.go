// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticDatabaseObject interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseQualifiedName(v string) *AgenticDatabaseObject
	GetDatabaseQualifiedName() *string
	SetDatabaseUuid(v string) *AgenticDatabaseObject
	GetDatabaseUuid() *string
	SetDdlSql(v string) *AgenticDatabaseObject
	GetDdlSql() *string
	SetObjectName(v string) *AgenticDatabaseObject
	GetObjectName() *string
	SetObjectQualifiedName(v string) *AgenticDatabaseObject
	GetObjectQualifiedName() *string
	SetObjectType(v string) *AgenticDatabaseObject
	GetObjectType() *string
}

type AgenticDatabaseObject struct {
	// The fully qualified name of the database. This name uniquely identifies the database within the system.
	DatabaseQualifiedName *string `json:"DatabaseQualifiedName,omitempty" xml:"DatabaseQualifiedName,omitempty"`
	// The unique identifier (UUID) of the database that contains the object.
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The Data Definition Language (DDL) SQL statement that defines the object\\"s structure.
	DdlSql *string `json:"DdlSql,omitempty" xml:"DdlSql,omitempty"`
	// The name of the database object, such as a table, view, or index.
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The fully qualified name that uniquely identifies the object, typically formatted as <database>.<schema>.<object>.
	ObjectQualifiedName *string `json:"ObjectQualifiedName,omitempty" xml:"ObjectQualifiedName,omitempty"`
	// The type of the database object. For example, `TABLE`, `VIEW`, or `INDEX`.
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s AgenticDatabaseObject) String() string {
	return dara.Prettify(s)
}

func (s AgenticDatabaseObject) GoString() string {
	return s.String()
}

func (s *AgenticDatabaseObject) GetDatabaseQualifiedName() *string {
	return s.DatabaseQualifiedName
}

func (s *AgenticDatabaseObject) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *AgenticDatabaseObject) GetDdlSql() *string {
	return s.DdlSql
}

func (s *AgenticDatabaseObject) GetObjectName() *string {
	return s.ObjectName
}

func (s *AgenticDatabaseObject) GetObjectQualifiedName() *string {
	return s.ObjectQualifiedName
}

func (s *AgenticDatabaseObject) GetObjectType() *string {
	return s.ObjectType
}

func (s *AgenticDatabaseObject) SetDatabaseQualifiedName(v string) *AgenticDatabaseObject {
	s.DatabaseQualifiedName = &v
	return s
}

func (s *AgenticDatabaseObject) SetDatabaseUuid(v string) *AgenticDatabaseObject {
	s.DatabaseUuid = &v
	return s
}

func (s *AgenticDatabaseObject) SetDdlSql(v string) *AgenticDatabaseObject {
	s.DdlSql = &v
	return s
}

func (s *AgenticDatabaseObject) SetObjectName(v string) *AgenticDatabaseObject {
	s.ObjectName = &v
	return s
}

func (s *AgenticDatabaseObject) SetObjectQualifiedName(v string) *AgenticDatabaseObject {
	s.ObjectQualifiedName = &v
	return s
}

func (s *AgenticDatabaseObject) SetObjectType(v string) *AgenticDatabaseObject {
	s.ObjectType = &v
	return s
}

func (s *AgenticDatabaseObject) Validate() error {
	return dara.Validate(s)
}
