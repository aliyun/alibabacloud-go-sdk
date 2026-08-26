// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOneMetaSqlTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *UpdateOneMetaSqlTemplateRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *UpdateOneMetaSqlTemplateRequest
	GetDatabaseUuid() *string
	SetDescription(v string) *UpdateOneMetaSqlTemplateRequest
	GetDescription() *string
	SetExpr(v string) *UpdateOneMetaSqlTemplateRequest
	GetExpr() *string
	SetKnowledgeUuid(v string) *UpdateOneMetaSqlTemplateRequest
	GetKnowledgeUuid() *string
	SetSqlParams(v string) *UpdateOneMetaSqlTemplateRequest
	GetSqlParams() *string
	SetTag(v string) *UpdateOneMetaSqlTemplateRequest
	GetTag() *string
	SetTitle(v string) *UpdateOneMetaSqlTemplateRequest
	GetTitle() *string
}

type UpdateOneMetaSqlTemplateRequest struct {
	// The UUID of the associated folder.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The UUID of the associated database.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The description of the SQL template.
	//
	// example:
	//
	// sales version 2
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The content of the SQL template.
	//
	// example:
	//
	// select count(1) from sales where dt = \\"2026-08-01\\"
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// The UUID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86c5c290052147c***
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	// The custom template parameters.
	//
	// example:
	//
	// {"dt": "2026-08-01"}
	SqlParams *string `json:"SqlParams,omitempty" xml:"SqlParams,omitempty"`
	// The tag of the SQL template.
	//
	// example:
	//
	// new_sales
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The title of the SQL template.
	//
	// example:
	//
	// sales_v2
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateOneMetaSqlTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOneMetaSqlTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateOneMetaSqlTemplateRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *UpdateOneMetaSqlTemplateRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *UpdateOneMetaSqlTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOneMetaSqlTemplateRequest) GetExpr() *string {
	return s.Expr
}

func (s *UpdateOneMetaSqlTemplateRequest) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *UpdateOneMetaSqlTemplateRequest) GetSqlParams() *string {
	return s.SqlParams
}

func (s *UpdateOneMetaSqlTemplateRequest) GetTag() *string {
	return s.Tag
}

func (s *UpdateOneMetaSqlTemplateRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateOneMetaSqlTemplateRequest) SetCatalogUuid(v string) *UpdateOneMetaSqlTemplateRequest {
	s.CatalogUuid = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) SetDatabaseUuid(v string) *UpdateOneMetaSqlTemplateRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) SetDescription(v string) *UpdateOneMetaSqlTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) SetExpr(v string) *UpdateOneMetaSqlTemplateRequest {
	s.Expr = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) SetKnowledgeUuid(v string) *UpdateOneMetaSqlTemplateRequest {
	s.KnowledgeUuid = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) SetSqlParams(v string) *UpdateOneMetaSqlTemplateRequest {
	s.SqlParams = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) SetTag(v string) *UpdateOneMetaSqlTemplateRequest {
	s.Tag = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) SetTitle(v string) *UpdateOneMetaSqlTemplateRequest {
	s.Title = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateRequest) Validate() error {
	return dara.Validate(s)
}
