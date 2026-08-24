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
	CatalogUuid  *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Expr         *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// This parameter is required.
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	SqlParams     *string `json:"SqlParams,omitempty" xml:"SqlParams,omitempty"`
	Tag           *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
