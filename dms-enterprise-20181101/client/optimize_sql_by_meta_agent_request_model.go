// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptimizeSqlByMetaAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *OptimizeSqlByMetaAgentRequest
	GetDbId() *string
	SetModel(v string) *OptimizeSqlByMetaAgentRequest
	GetModel() *string
	SetQuery(v string) *OptimizeSqlByMetaAgentRequest
	GetQuery() *string
	SetSql(v string) *OptimizeSqlByMetaAgentRequest
	GetSql() *string
}

type OptimizeSqlByMetaAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 是否为慢SQL
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// select 	- from table where col = 1
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s OptimizeSqlByMetaAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s OptimizeSqlByMetaAgentRequest) GoString() string {
	return s.String()
}

func (s *OptimizeSqlByMetaAgentRequest) GetDbId() *string {
	return s.DbId
}

func (s *OptimizeSqlByMetaAgentRequest) GetModel() *string {
	return s.Model
}

func (s *OptimizeSqlByMetaAgentRequest) GetQuery() *string {
	return s.Query
}

func (s *OptimizeSqlByMetaAgentRequest) GetSql() *string {
	return s.Sql
}

func (s *OptimizeSqlByMetaAgentRequest) SetDbId(v string) *OptimizeSqlByMetaAgentRequest {
	s.DbId = &v
	return s
}

func (s *OptimizeSqlByMetaAgentRequest) SetModel(v string) *OptimizeSqlByMetaAgentRequest {
	s.Model = &v
	return s
}

func (s *OptimizeSqlByMetaAgentRequest) SetQuery(v string) *OptimizeSqlByMetaAgentRequest {
	s.Query = &v
	return s
}

func (s *OptimizeSqlByMetaAgentRequest) SetSql(v string) *OptimizeSqlByMetaAgentRequest {
	s.Sql = &v
	return s
}

func (s *OptimizeSqlByMetaAgentRequest) Validate() error {
	return dara.Validate(s)
}
