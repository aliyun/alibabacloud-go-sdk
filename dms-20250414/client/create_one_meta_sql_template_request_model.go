// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOneMetaSqlTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *CreateOneMetaSqlTemplateRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *CreateOneMetaSqlTemplateRequest
	GetDatabaseUuid() *string
	SetDescription(v string) *CreateOneMetaSqlTemplateRequest
	GetDescription() *string
	SetExpr(v string) *CreateOneMetaSqlTemplateRequest
	GetExpr() *string
	SetSource(v string) *CreateOneMetaSqlTemplateRequest
	GetSource() *string
	SetSqlParams(v string) *CreateOneMetaSqlTemplateRequest
	GetSqlParams() *string
	SetTag(v string) *CreateOneMetaSqlTemplateRequest
	GetTag() *string
	SetTitle(v string) *CreateOneMetaSqlTemplateRequest
	GetTitle() *string
}

type CreateOneMetaSqlTemplateRequest struct {
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
	// sales template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The content of the SQL template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT SUM(amount) AS total_sales FROM store_daily_sales
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// The knowledge source of the SQL template.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_AGENT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The custom template parameters.
	//
	// example:
	//
	// {"start_day_id": "2026-08-01", "end_day_id": "2026-08-16"}
	SqlParams *string `json:"SqlParams,omitempty" xml:"SqlParams,omitempty"`
	// The tag of the SQL template.
	//
	// example:
	//
	// sales
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The title of the SQL template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecommerce_sales
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateOneMetaSqlTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOneMetaSqlTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateOneMetaSqlTemplateRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *CreateOneMetaSqlTemplateRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *CreateOneMetaSqlTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOneMetaSqlTemplateRequest) GetExpr() *string {
	return s.Expr
}

func (s *CreateOneMetaSqlTemplateRequest) GetSource() *string {
	return s.Source
}

func (s *CreateOneMetaSqlTemplateRequest) GetSqlParams() *string {
	return s.SqlParams
}

func (s *CreateOneMetaSqlTemplateRequest) GetTag() *string {
	return s.Tag
}

func (s *CreateOneMetaSqlTemplateRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateOneMetaSqlTemplateRequest) SetCatalogUuid(v string) *CreateOneMetaSqlTemplateRequest {
	s.CatalogUuid = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) SetDatabaseUuid(v string) *CreateOneMetaSqlTemplateRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) SetDescription(v string) *CreateOneMetaSqlTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) SetExpr(v string) *CreateOneMetaSqlTemplateRequest {
	s.Expr = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) SetSource(v string) *CreateOneMetaSqlTemplateRequest {
	s.Source = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) SetSqlParams(v string) *CreateOneMetaSqlTemplateRequest {
	s.SqlParams = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) SetTag(v string) *CreateOneMetaSqlTemplateRequest {
	s.Tag = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) SetTitle(v string) *CreateOneMetaSqlTemplateRequest {
	s.Title = &v
	return s
}

func (s *CreateOneMetaSqlTemplateRequest) Validate() error {
	return dara.Validate(s)
}
