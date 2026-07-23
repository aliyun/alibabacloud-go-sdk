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
	// Add column
	//
	// example:
	//
	// {"Name":"id","Type":"bigint","Comment":"主键"}
	AddColumnShrink *string `json:"AddColumn,omitempty" xml:"AddColumn,omitempty"`
	// Data catalog to which it belongs
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Idempotency token
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Delete column
	//
	// example:
	//
	// {"Name":"old_column"}
	DeleteColumnShrink *string `json:"DeleteColumn,omitempty" xml:"DeleteColumn,omitempty"`
	// Table name
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Namespace to which it belongs
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Rename column
	//
	// example:
	//
	// {"Name":"old_name","NewName":"new_name"}
	RenameColumnShrink *string `json:"RenameColumn,omitempty" xml:"RenameColumn,omitempty"`
	// Update column comment
	//
	// example:
	//
	// {"Name":"id","Comment":"主键ID"}
	UpdateColumnCommentShrink *string `json:"UpdateColumnComment,omitempty" xml:"UpdateColumnComment,omitempty"`
	// Update column type
	//
	// example:
	//
	// {"Name":"id","Type":"bigint"}
	UpdateColumnTypeShrink *string `json:"UpdateColumnType,omitempty" xml:"UpdateColumnType,omitempty"`
	// Update table comment
	//
	// example:
	//
	// 更新后的备注
	UpdateComment *string `json:"UpdateComment,omitempty" xml:"UpdateComment,omitempty"`
	// Update retention policy
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
