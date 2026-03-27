// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchema interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Schema
	GetComment() *string
	SetCreateTime(v int64) *Schema
	GetCreateTime() *int64
	SetId(v string) *Schema
	GetId() *string
	SetModifyTime(v int64) *Schema
	GetModifyTime() *int64
	SetName(v string) *Schema
	GetName() *string
	SetParentMetaEntityId(v string) *Schema
	GetParentMetaEntityId() *string
	SetType(v string) *Schema
	GetType() *string
}

type Schema struct {
	// The comment.
	//
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1736852168000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The schema ID. For more information, see [Concepts related to metadata entities.](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format is `${EntityType}:${Instance ID or escaped URL}:${Catalog name}:${Database name}`. Use empty strings as placeholders for levels that do not exist.
	//
	// >  For the MaxCompute type, the instance ID level is represented by an empty string. The database name is the name of the MaxCompute project, which has enabled the schema feature.
	//
	// Examples of common ID formats
	//
	// `maxcompute-project:::project_name` (For MaxCompute projects schema enabled)
	//
	// `holo-schema:instance_id::database_name:schema_name`
	//
	// > \\
	//
	// `instance_id`: The Hologres instance ID\\
	//
	// . `database_name`: The database name\\
	//
	// . `project_name`: The MaxCompute project name\\
	//
	// . `schema_name`: The schema name.
	//
	// example:
	//
	// maxcompute-schema:123456XXX::test_project:default
	//
	// holo-schema:h-abc123xxx::test_db:test_schema
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1736852168000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name.
	//
	// example:
	//
	// test_db
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent entity ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format: `${EntityType}:${Instance ID or escaped URL}:${Catalog name}:${Database name}`. Use empty strings as placeholders for levels that do not exist.
	//
	// >  For the MaxCompute type, the instance ID level is represented by an empty string. The database name is the name of the MaxCompute project with schema enabled.
	//
	// Examples of common ParentMetaEntityId formats
	//
	// `maxcompute-project:::project_name` (For MaxCompute projects with schema enabled)
	//
	// `holo-database:instance_id::database_name`
	//
	// > \\
	//
	// `instance_id`: The Hologres instance ID\\
	//
	// . `database_name`: The database name\\
	//
	// . `project_name`: The MaxCompute project name.
	//
	// example:
	//
	// maxcompute-project:123456XXX::test_project
	//
	// holo-database:h-abc123xxx::test_db
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The type.
	//
	// example:
	//
	// MANAGED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Schema) String() string {
	return dara.Prettify(s)
}

func (s Schema) GoString() string {
	return s.String()
}

func (s *Schema) GetComment() *string {
	return s.Comment
}

func (s *Schema) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Schema) GetId() *string {
	return s.Id
}

func (s *Schema) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *Schema) GetName() *string {
	return s.Name
}

func (s *Schema) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *Schema) GetType() *string {
	return s.Type
}

func (s *Schema) SetComment(v string) *Schema {
	s.Comment = &v
	return s
}

func (s *Schema) SetCreateTime(v int64) *Schema {
	s.CreateTime = &v
	return s
}

func (s *Schema) SetId(v string) *Schema {
	s.Id = &v
	return s
}

func (s *Schema) SetModifyTime(v int64) *Schema {
	s.ModifyTime = &v
	return s
}

func (s *Schema) SetName(v string) *Schema {
	s.Name = &v
	return s
}

func (s *Schema) SetParentMetaEntityId(v string) *Schema {
	s.ParentMetaEntityId = &v
	return s
}

func (s *Schema) SetType(v string) *Schema {
	s.Type = &v
	return s
}

func (s *Schema) Validate() error {
	return dara.Validate(s)
}
