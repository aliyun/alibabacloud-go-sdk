// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddColumn(v *UpdateTableRequestAddColumn) *UpdateTableRequest
	GetAddColumn() *UpdateTableRequestAddColumn
	SetCatalog(v string) *UpdateTableRequest
	GetCatalog() *string
	SetClientToken(v string) *UpdateTableRequest
	GetClientToken() *string
	SetDeleteColumn(v *UpdateTableRequestDeleteColumn) *UpdateTableRequest
	GetDeleteColumn() *UpdateTableRequestDeleteColumn
	SetName(v string) *UpdateTableRequest
	GetName() *string
	SetNamespace(v string) *UpdateTableRequest
	GetNamespace() *string
	SetRenameColumn(v *UpdateTableRequestRenameColumn) *UpdateTableRequest
	GetRenameColumn() *UpdateTableRequestRenameColumn
	SetUpdateColumnComment(v *UpdateTableRequestUpdateColumnComment) *UpdateTableRequest
	GetUpdateColumnComment() *UpdateTableRequestUpdateColumnComment
	SetUpdateColumnType(v *UpdateTableRequestUpdateColumnType) *UpdateTableRequest
	GetUpdateColumnType() *UpdateTableRequestUpdateColumnType
	SetUpdateComment(v string) *UpdateTableRequest
	GetUpdateComment() *string
	SetUpdateRetentionPolicy(v *UpdateTableRequestUpdateRetentionPolicy) *UpdateTableRequest
	GetUpdateRetentionPolicy() *UpdateTableRequestUpdateRetentionPolicy
}

type UpdateTableRequest struct {
	// 新增列定义（JSON 对象）。包含 Name（列名，必填）、Type（数据类型，必填，如 STRING、INT32、INT64、FLOAT、DOUBLE、BOOLEAN、TIMESTAMP）、Comment（列备注，选填）。每次调用只能新增一列
	//
	// example:
	//
	// {"Name":"id","Type":"bigint","Comment":"主键"}
	AddColumn *UpdateTableRequestAddColumn `json:"AddColumn,omitempty" xml:"AddColumn,omitempty" type:"Struct"`
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
	DeleteColumn *UpdateTableRequestDeleteColumn `json:"DeleteColumn,omitempty" xml:"DeleteColumn,omitempty" type:"Struct"`
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
	RenameColumn *UpdateTableRequestRenameColumn `json:"RenameColumn,omitempty" xml:"RenameColumn,omitempty" type:"Struct"`
	// 修改列的备注信息（JSON 对象）。包含 Name（目标列名，必填）、Comment（新备注内容，必填，传空字符串可清除备注）。每次调用只能修改一列
	//
	// example:
	//
	// {"Name":"id","Comment":"主键ID"}
	UpdateColumnComment *UpdateTableRequestUpdateColumnComment `json:"UpdateColumnComment,omitempty" xml:"UpdateColumnComment,omitempty" type:"Struct"`
	// 修改列的数据类型（JSON 对象）。包含 Name（目标列名，必填）、Type（新数据类型，必填）。仅支持兼容类型转换，每次调用只能修改一列
	//
	// example:
	//
	// {"Name":"id","Type":"bigint"}
	UpdateColumnType *UpdateTableRequestUpdateColumnType `json:"UpdateColumnType,omitempty" xml:"UpdateColumnType,omitempty" type:"Struct"`
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
	UpdateRetentionPolicy *UpdateTableRequestUpdateRetentionPolicy `json:"UpdateRetentionPolicy,omitempty" xml:"UpdateRetentionPolicy,omitempty" type:"Struct"`
}

func (s UpdateTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableRequest) GetAddColumn() *UpdateTableRequestAddColumn {
	return s.AddColumn
}

func (s *UpdateTableRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *UpdateTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTableRequest) GetDeleteColumn() *UpdateTableRequestDeleteColumn {
	return s.DeleteColumn
}

func (s *UpdateTableRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTableRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateTableRequest) GetRenameColumn() *UpdateTableRequestRenameColumn {
	return s.RenameColumn
}

func (s *UpdateTableRequest) GetUpdateColumnComment() *UpdateTableRequestUpdateColumnComment {
	return s.UpdateColumnComment
}

func (s *UpdateTableRequest) GetUpdateColumnType() *UpdateTableRequestUpdateColumnType {
	return s.UpdateColumnType
}

func (s *UpdateTableRequest) GetUpdateComment() *string {
	return s.UpdateComment
}

func (s *UpdateTableRequest) GetUpdateRetentionPolicy() *UpdateTableRequestUpdateRetentionPolicy {
	return s.UpdateRetentionPolicy
}

func (s *UpdateTableRequest) SetAddColumn(v *UpdateTableRequestAddColumn) *UpdateTableRequest {
	s.AddColumn = v
	return s
}

func (s *UpdateTableRequest) SetCatalog(v string) *UpdateTableRequest {
	s.Catalog = &v
	return s
}

func (s *UpdateTableRequest) SetClientToken(v string) *UpdateTableRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTableRequest) SetDeleteColumn(v *UpdateTableRequestDeleteColumn) *UpdateTableRequest {
	s.DeleteColumn = v
	return s
}

func (s *UpdateTableRequest) SetName(v string) *UpdateTableRequest {
	s.Name = &v
	return s
}

func (s *UpdateTableRequest) SetNamespace(v string) *UpdateTableRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateTableRequest) SetRenameColumn(v *UpdateTableRequestRenameColumn) *UpdateTableRequest {
	s.RenameColumn = v
	return s
}

func (s *UpdateTableRequest) SetUpdateColumnComment(v *UpdateTableRequestUpdateColumnComment) *UpdateTableRequest {
	s.UpdateColumnComment = v
	return s
}

func (s *UpdateTableRequest) SetUpdateColumnType(v *UpdateTableRequestUpdateColumnType) *UpdateTableRequest {
	s.UpdateColumnType = v
	return s
}

func (s *UpdateTableRequest) SetUpdateComment(v string) *UpdateTableRequest {
	s.UpdateComment = &v
	return s
}

func (s *UpdateTableRequest) SetUpdateRetentionPolicy(v *UpdateTableRequestUpdateRetentionPolicy) *UpdateTableRequest {
	s.UpdateRetentionPolicy = v
	return s
}

func (s *UpdateTableRequest) Validate() error {
	if s.AddColumn != nil {
		if err := s.AddColumn.Validate(); err != nil {
			return err
		}
	}
	if s.DeleteColumn != nil {
		if err := s.DeleteColumn.Validate(); err != nil {
			return err
		}
	}
	if s.RenameColumn != nil {
		if err := s.RenameColumn.Validate(); err != nil {
			return err
		}
	}
	if s.UpdateColumnComment != nil {
		if err := s.UpdateColumnComment.Validate(); err != nil {
			return err
		}
	}
	if s.UpdateColumnType != nil {
		if err := s.UpdateColumnType.Validate(); err != nil {
			return err
		}
	}
	if s.UpdateRetentionPolicy != nil {
		if err := s.UpdateRetentionPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTableRequestAddColumn struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// kafka-default-agent-alikafka_pre-cn-28t3sfzno003
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTableRequestAddColumn) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestAddColumn) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestAddColumn) GetComment() *string {
	return s.Comment
}

func (s *UpdateTableRequestAddColumn) GetName() *string {
	return s.Name
}

func (s *UpdateTableRequestAddColumn) GetType() *string {
	return s.Type
}

func (s *UpdateTableRequestAddColumn) SetComment(v string) *UpdateTableRequestAddColumn {
	s.Comment = &v
	return s
}

func (s *UpdateTableRequestAddColumn) SetName(v string) *UpdateTableRequestAddColumn {
	s.Name = &v
	return s
}

func (s *UpdateTableRequestAddColumn) SetType(v string) *UpdateTableRequestAddColumn {
	s.Type = &v
	return s
}

func (s *UpdateTableRequestAddColumn) Validate() error {
	return dara.Validate(s)
}

type UpdateTableRequestDeleteColumn struct {
	// example:
	//
	// kafka-default-agent-alikafka_pre-cn-28t3sfzno003
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTableRequestDeleteColumn) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestDeleteColumn) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestDeleteColumn) GetName() *string {
	return s.Name
}

func (s *UpdateTableRequestDeleteColumn) SetName(v string) *UpdateTableRequestDeleteColumn {
	s.Name = &v
	return s
}

func (s *UpdateTableRequestDeleteColumn) Validate() error {
	return dara.Validate(s)
}

type UpdateTableRequestRenameColumn struct {
	// example:
	//
	// kafka-default-agent-alikafka_pre-cn-28t3sfzno003
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// fvt-oos-application-group-56ca74b000
	NewName *string `json:"NewName,omitempty" xml:"NewName,omitempty"`
}

func (s UpdateTableRequestRenameColumn) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestRenameColumn) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestRenameColumn) GetName() *string {
	return s.Name
}

func (s *UpdateTableRequestRenameColumn) GetNewName() *string {
	return s.NewName
}

func (s *UpdateTableRequestRenameColumn) SetName(v string) *UpdateTableRequestRenameColumn {
	s.Name = &v
	return s
}

func (s *UpdateTableRequestRenameColumn) SetNewName(v string) *UpdateTableRequestRenameColumn {
	s.NewName = &v
	return s
}

func (s *UpdateTableRequestRenameColumn) Validate() error {
	return dara.Validate(s)
}

type UpdateTableRequestUpdateColumnComment struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// kafka-default-agent-alikafka_pre-cn-28t3sfzno003
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateTableRequestUpdateColumnComment) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestUpdateColumnComment) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestUpdateColumnComment) GetComment() *string {
	return s.Comment
}

func (s *UpdateTableRequestUpdateColumnComment) GetName() *string {
	return s.Name
}

func (s *UpdateTableRequestUpdateColumnComment) SetComment(v string) *UpdateTableRequestUpdateColumnComment {
	s.Comment = &v
	return s
}

func (s *UpdateTableRequestUpdateColumnComment) SetName(v string) *UpdateTableRequestUpdateColumnComment {
	s.Name = &v
	return s
}

func (s *UpdateTableRequestUpdateColumnComment) Validate() error {
	return dara.Validate(s)
}

type UpdateTableRequestUpdateColumnType struct {
	// example:
	//
	// kafka-default-agent-alikafka_pre-cn-28t3sfzno003
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PRIVATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTableRequestUpdateColumnType) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestUpdateColumnType) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestUpdateColumnType) GetName() *string {
	return s.Name
}

func (s *UpdateTableRequestUpdateColumnType) GetType() *string {
	return s.Type
}

func (s *UpdateTableRequestUpdateColumnType) SetName(v string) *UpdateTableRequestUpdateColumnType {
	s.Name = &v
	return s
}

func (s *UpdateTableRequestUpdateColumnType) SetType(v string) *UpdateTableRequestUpdateColumnType {
	s.Type = &v
	return s
}

func (s *UpdateTableRequestUpdateColumnType) Validate() error {
	return dara.Validate(s)
}

type UpdateTableRequestUpdateRetentionPolicy struct {
	// example:
	//
	// 17
	ColdTTL *int32 `json:"ColdTTL,omitempty" xml:"ColdTTL,omitempty"`
	// example:
	//
	// 7
	HotTTL *int32 `json:"HotTTL,omitempty" xml:"HotTTL,omitempty"`
}

func (s UpdateTableRequestUpdateRetentionPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestUpdateRetentionPolicy) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestUpdateRetentionPolicy) GetColdTTL() *int32 {
	return s.ColdTTL
}

func (s *UpdateTableRequestUpdateRetentionPolicy) GetHotTTL() *int32 {
	return s.HotTTL
}

func (s *UpdateTableRequestUpdateRetentionPolicy) SetColdTTL(v int32) *UpdateTableRequestUpdateRetentionPolicy {
	s.ColdTTL = &v
	return s
}

func (s *UpdateTableRequestUpdateRetentionPolicy) SetHotTTL(v int32) *UpdateTableRequestUpdateRetentionPolicy {
	s.HotTTL = &v
	return s
}

func (s *UpdateTableRequestUpdateRetentionPolicy) Validate() error {
	return dara.Validate(s)
}
