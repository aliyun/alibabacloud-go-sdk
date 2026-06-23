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
	// 注释。
	//
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 创建时间（毫秒级时间戳）。
	//
	// example:
	//
	// 1736852168000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// ID，可参考[元数据实体相关概念说明](https://help.aliyun.com/document_detail/2880092.html)。
	//
	// 格式为`${EntityType}:${实例ID或转义后的URL}:${数据目录名称}:${数据库名称}:${模式名称}`，对于不存在的层级置空。
	//
	// > 对于MaxCompute类型，此处的实例ID即为主账号ID，数据库名称即为MaxCompute项目名称。
	//
	// example:
	//
	// maxcompute-schema:123456XXX::test_project:default
	//
	// holo-schema:h-abc123xxx::test_db:test_schema
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 更新时间（毫秒级时间戳）。
	//
	// example:
	//
	// 1736852168000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 名称。
	//
	// example:
	//
	// test_db
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 父层级元数据实体ID，父层级实体类型取值参考ListCrawlerTypes接口。
	//
	// 格式为`${EntityType}:${实例ID或转义后的URL}:${数据目录名称}:${数据库名称}`，对于不存在的层级置空。
	//
	// > 对于MaxCompute类型，此处的实例ID即为主账号ID，数据库名称即为MaxCompute项目名称。
	//
	// example:
	//
	// maxcompute-project:123456XXX::test_project
	//
	// holo-database:h-abc123xxx::test_db
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// 类型。
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
