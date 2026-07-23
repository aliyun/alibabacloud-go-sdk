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
	// The data catalog to which the table belongs.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The column definitions.
	//
	// example:
	//
	// [{"Name":"id","Type":"bigint","Comment":"主键"}]
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// The description.
	//
	// example:
	//
	// 测试事件表
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the table belongs.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The data retention policy.
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
