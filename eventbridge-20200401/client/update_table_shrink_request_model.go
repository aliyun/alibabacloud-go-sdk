// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddColumnShrink(v string) *UpdateTableShrinkRequest
	GetAddColumnShrink() *string
	SetCatalog(v string) *UpdateTableShrinkRequest
	GetCatalog() *string
	SetClientToken(v string) *UpdateTableShrinkRequest
	GetClientToken() *string
	SetDeleteColumnShrink(v string) *UpdateTableShrinkRequest
	GetDeleteColumnShrink() *string
	SetName(v string) *UpdateTableShrinkRequest
	GetName() *string
	SetNamespace(v string) *UpdateTableShrinkRequest
	GetNamespace() *string
	SetRenameColumnShrink(v string) *UpdateTableShrinkRequest
	GetRenameColumnShrink() *string
	SetUpdateColumnCommentShrink(v string) *UpdateTableShrinkRequest
	GetUpdateColumnCommentShrink() *string
	SetUpdateColumnTypeShrink(v string) *UpdateTableShrinkRequest
	GetUpdateColumnTypeShrink() *string
	SetUpdateComment(v string) *UpdateTableShrinkRequest
	GetUpdateComment() *string
	SetUpdateRetentionPolicyShrink(v string) *UpdateTableShrinkRequest
	GetUpdateRetentionPolicyShrink() *string
}

type UpdateTableShrinkRequest struct {
	// 新增列定义（JSON 对象）。包含 Name（列名，必填）、Type（数据类型，必填，如 STRING、INT32、INT64、FLOAT、DOUBLE、BOOLEAN、TIMESTAMP）、Comment（列备注，选填）。每次调用只能新增一列
	//
	// example:
	//
	// {"Name":"id","Type":"bigint","Comment":"主键"}
	AddColumnShrink *string `json:"AddColumn,omitempty" xml:"AddColumn,omitempty"`
	// 表所属的数据目录名称。可通过 ListCatalogs 获取
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// 用于保证请求幂等性的Token。建议使用 UUID
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 删除列定义（JSON 对象）。包含 Name（要删除的列名，必填）。删除后不可恢复，已有数据中该列的值将丢失。每次调用只能删除一列
	//
	// example:
	//
	// {"Name":"old_column"}
	DeleteColumnShrink *string `json:"DeleteColumn,omitempty" xml:"DeleteColumn,omitempty"`
	// 要修改的事件表名称。名称本身不可修改，此处用于定位目标表。需同时指定所属 Catalog 和 Namespace。可通过 ListTables 获取
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 表所属的命名空间名称。可通过 ListNamespaces 获取
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 重命名列（JSON 对象）。包含 Name（原列名，必填）、NewName（新列名，必填）。每次调用只能重命名一列
	//
	// example:
	//
	// {"Name":"old_name","NewName":"new_name"}
	RenameColumnShrink *string `json:"RenameColumn,omitempty" xml:"RenameColumn,omitempty"`
	// 修改列的备注信息（JSON 对象）。包含 Name（目标列名，必填）、Comment（新备注内容，必填，传空字符串可清除备注）。每次调用只能修改一列
	//
	// example:
	//
	// {"Name":"id","Comment":"主键ID"}
	UpdateColumnCommentShrink *string `json:"UpdateColumnComment,omitempty" xml:"UpdateColumnComment,omitempty"`
	// 修改列的数据类型（JSON 对象）。包含 Name（目标列名，必填）、Type（新数据类型，必填）。仅支持兼容类型转换，每次调用只能修改一列
	//
	// example:
	//
	// {"Name":"id","Type":"bigint"}
	UpdateColumnTypeShrink *string `json:"UpdateColumnType,omitempty" xml:"UpdateColumnType,omitempty"`
	// 修改表的备注描述。传入新的备注内容替换原有备注，传空字符串可清除备注
	//
	// example:
	//
	// 更新后的备注
	UpdateComment *string `json:"UpdateComment,omitempty" xml:"UpdateComment,omitempty"`
	// 修改数据保留策略（JSON 对象）。包含 HotTTL（热数据保留天数）、ColdTTL（冷数据保留天数）。传入后会替换原有策略
	//
	// example:
	//
	// {"HotTTL":7,"ColdTTL":30}
	UpdateRetentionPolicyShrink *string `json:"UpdateRetentionPolicy,omitempty" xml:"UpdateRetentionPolicy,omitempty"`
}

func (s UpdateTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableShrinkRequest) GetAddColumnShrink() *string {
	return s.AddColumnShrink
}

func (s *UpdateTableShrinkRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *UpdateTableShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTableShrinkRequest) GetDeleteColumnShrink() *string {
	return s.DeleteColumnShrink
}

func (s *UpdateTableShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTableShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateTableShrinkRequest) GetRenameColumnShrink() *string {
	return s.RenameColumnShrink
}

func (s *UpdateTableShrinkRequest) GetUpdateColumnCommentShrink() *string {
	return s.UpdateColumnCommentShrink
}

func (s *UpdateTableShrinkRequest) GetUpdateColumnTypeShrink() *string {
	return s.UpdateColumnTypeShrink
}

func (s *UpdateTableShrinkRequest) GetUpdateComment() *string {
	return s.UpdateComment
}

func (s *UpdateTableShrinkRequest) GetUpdateRetentionPolicyShrink() *string {
	return s.UpdateRetentionPolicyShrink
}

func (s *UpdateTableShrinkRequest) SetAddColumnShrink(v string) *UpdateTableShrinkRequest {
	s.AddColumnShrink = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetCatalog(v string) *UpdateTableShrinkRequest {
	s.Catalog = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetClientToken(v string) *UpdateTableShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetDeleteColumnShrink(v string) *UpdateTableShrinkRequest {
	s.DeleteColumnShrink = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetName(v string) *UpdateTableShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetNamespace(v string) *UpdateTableShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetRenameColumnShrink(v string) *UpdateTableShrinkRequest {
	s.RenameColumnShrink = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetUpdateColumnCommentShrink(v string) *UpdateTableShrinkRequest {
	s.UpdateColumnCommentShrink = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetUpdateColumnTypeShrink(v string) *UpdateTableShrinkRequest {
	s.UpdateColumnTypeShrink = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetUpdateComment(v string) *UpdateTableShrinkRequest {
	s.UpdateComment = &v
	return s
}

func (s *UpdateTableShrinkRequest) SetUpdateRetentionPolicyShrink(v string) *UpdateTableShrinkRequest {
	s.UpdateRetentionPolicyShrink = &v
	return s
}

func (s *UpdateTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
