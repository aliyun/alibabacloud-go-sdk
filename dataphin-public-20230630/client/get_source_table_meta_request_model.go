// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceTableMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *GetSourceTableMetaRequestContext) *GetSourceTableMetaRequest
	GetContext() *GetSourceTableMetaRequestContext
	SetOpTenantId(v int64) *GetSourceTableMetaRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetSourceTableMetaRequest
	GetOpUserId() *string
	SetQuery(v *GetSourceTableMetaRequestQuery) *GetSourceTableMetaRequest
	GetQuery() *GetSourceTableMetaRequestQuery
}

type GetSourceTableMetaRequest struct {
	// This parameter is required.
	Context *GetSourceTableMetaRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// This parameter is required.
	Query *GetSourceTableMetaRequestQuery `json:"Query,omitempty" xml:"Query,omitempty" type:"Struct"`
}

func (s GetSourceTableMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaRequest) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaRequest) GetContext() *GetSourceTableMetaRequestContext {
	return s.Context
}

func (s *GetSourceTableMetaRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSourceTableMetaRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetSourceTableMetaRequest) GetQuery() *GetSourceTableMetaRequestQuery {
	return s.Query
}

func (s *GetSourceTableMetaRequest) SetContext(v *GetSourceTableMetaRequestContext) *GetSourceTableMetaRequest {
	s.Context = v
	return s
}

func (s *GetSourceTableMetaRequest) SetOpTenantId(v int64) *GetSourceTableMetaRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSourceTableMetaRequest) SetOpUserId(v string) *GetSourceTableMetaRequest {
	s.OpUserId = &v
	return s
}

func (s *GetSourceTableMetaRequest) SetQuery(v *GetSourceTableMetaRequestQuery) *GetSourceTableMetaRequest {
	s.Query = v
	return s
}

func (s *GetSourceTableMetaRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.Query != nil {
		if err := s.Query.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSourceTableMetaRequestContext struct {
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetSourceTableMetaRequestContext) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaRequestContext) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaRequestContext) GetEnv() *string {
	return s.Env
}

func (s *GetSourceTableMetaRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSourceTableMetaRequestContext) SetEnv(v string) *GetSourceTableMetaRequestContext {
	s.Env = &v
	return s
}

func (s *GetSourceTableMetaRequestContext) SetProjectId(v int64) *GetSourceTableMetaRequestContext {
	s.ProjectId = &v
	return s
}

func (s *GetSourceTableMetaRequestContext) Validate() error {
	return dara.Validate(s)
}

type GetSourceTableMetaRequestQuery struct {
	// example:
	//
	// hive_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// DATA_SOURCE
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// example:
	//
	// default
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// ods_user_info
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetSourceTableMetaRequestQuery) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaRequestQuery) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaRequestQuery) GetCatalog() *string {
	return s.Catalog
}

func (s *GetSourceTableMetaRequestQuery) GetId() *string {
	return s.Id
}

func (s *GetSourceTableMetaRequestQuery) GetQueryMode() *string {
	return s.QueryMode
}

func (s *GetSourceTableMetaRequestQuery) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetSourceTableMetaRequestQuery) GetTableName() *string {
	return s.TableName
}

func (s *GetSourceTableMetaRequestQuery) SetCatalog(v string) *GetSourceTableMetaRequestQuery {
	s.Catalog = &v
	return s
}

func (s *GetSourceTableMetaRequestQuery) SetId(v string) *GetSourceTableMetaRequestQuery {
	s.Id = &v
	return s
}

func (s *GetSourceTableMetaRequestQuery) SetQueryMode(v string) *GetSourceTableMetaRequestQuery {
	s.QueryMode = &v
	return s
}

func (s *GetSourceTableMetaRequestQuery) SetSchemaName(v string) *GetSourceTableMetaRequestQuery {
	s.SchemaName = &v
	return s
}

func (s *GetSourceTableMetaRequestQuery) SetTableName(v string) *GetSourceTableMetaRequestQuery {
	s.TableName = &v
	return s
}

func (s *GetSourceTableMetaRequestQuery) Validate() error {
	return dara.Validate(s)
}
