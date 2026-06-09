// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *CreateTableShrinkRequest
	GetCatalog() *string
	SetClientToken(v string) *CreateTableShrinkRequest
	GetClientToken() *string
	SetColumnsShrink(v string) *CreateTableShrinkRequest
	GetColumnsShrink() *string
	SetComment(v string) *CreateTableShrinkRequest
	GetComment() *string
	SetName(v string) *CreateTableShrinkRequest
	GetName() *string
	SetNamespace(v string) *CreateTableShrinkRequest
	GetNamespace() *string
	SetRetentionPolicyShrink(v string) *CreateTableShrinkRequest
	GetRetentionPolicyShrink() *string
}

type CreateTableShrinkRequest struct {
	// 表所属的数据目录名称。可通过 ListCatalogs 获取已有目录列表
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// 用于保证请求幂等性的Token，防止因网络重试导致重复创建。建议使用 UUID
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 表的列定义（JSON 数组）。每列包含 Name（列名，必填）、Type（数据类型，必填，如 STRING、INT32、INT64、FLOAT、DOUBLE、BOOLEAN、TIMESTAMP）、Comment（列备注，选填）
	//
	// example:
	//
	// [{"Name":"id","Type":"bigint","Comment":"主键"}]
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// 表的备注描述信息，无格式限制
	//
	// example:
	//
	// 测试事件表
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 事件表名称。以字母或数字开头，支持字母、数字、下划线和短横线，长度1~127。在同一命名空间下唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 表所属的命名空间名称。可通过 ListNamespaces 获取已有命名空间列表
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 数据保留策略（JSON 对象）。包含 HotTTL（热数据保留天数，高性能查询）和 ColdTTL（冷数据保留天数，低成本存储）。不传则使用系统默认值
	//
	// example:
	//
	// {"HotTTL":7,"ColdTTL":30}
	RetentionPolicyShrink *string `json:"RetentionPolicy,omitempty" xml:"RetentionPolicy,omitempty"`
}

func (s CreateTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTableShrinkRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateTableShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTableShrinkRequest) GetColumnsShrink() *string {
	return s.ColumnsShrink
}

func (s *CreateTableShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateTableShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateTableShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateTableShrinkRequest) GetRetentionPolicyShrink() *string {
	return s.RetentionPolicyShrink
}

func (s *CreateTableShrinkRequest) SetCatalog(v string) *CreateTableShrinkRequest {
	s.Catalog = &v
	return s
}

func (s *CreateTableShrinkRequest) SetClientToken(v string) *CreateTableShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTableShrinkRequest) SetColumnsShrink(v string) *CreateTableShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *CreateTableShrinkRequest) SetComment(v string) *CreateTableShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateTableShrinkRequest) SetName(v string) *CreateTableShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateTableShrinkRequest) SetNamespace(v string) *CreateTableShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateTableShrinkRequest) SetRetentionPolicyShrink(v string) *CreateTableShrinkRequest {
	s.RetentionPolicyShrink = &v
	return s
}

func (s *CreateTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
