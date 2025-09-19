// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixSqlByMetaAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *FixSqlByMetaAgentRequest
	GetDbId() *string
	SetError(v string) *FixSqlByMetaAgentRequest
	GetError() *string
	SetModel(v string) *FixSqlByMetaAgentRequest
	GetModel() *string
	SetQuery(v string) *FixSqlByMetaAgentRequest
	GetQuery() *string
	SetSql(v string) *FixSqlByMetaAgentRequest
	GetSql() *string
}

type FixSqlByMetaAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// unknown column col
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 正确字段是啥？
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// select 	- from table where col = 1
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s FixSqlByMetaAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s FixSqlByMetaAgentRequest) GoString() string {
	return s.String()
}

func (s *FixSqlByMetaAgentRequest) GetDbId() *string {
	return s.DbId
}

func (s *FixSqlByMetaAgentRequest) GetError() *string {
	return s.Error
}

func (s *FixSqlByMetaAgentRequest) GetModel() *string {
	return s.Model
}

func (s *FixSqlByMetaAgentRequest) GetQuery() *string {
	return s.Query
}

func (s *FixSqlByMetaAgentRequest) GetSql() *string {
	return s.Sql
}

func (s *FixSqlByMetaAgentRequest) SetDbId(v string) *FixSqlByMetaAgentRequest {
	s.DbId = &v
	return s
}

func (s *FixSqlByMetaAgentRequest) SetError(v string) *FixSqlByMetaAgentRequest {
	s.Error = &v
	return s
}

func (s *FixSqlByMetaAgentRequest) SetModel(v string) *FixSqlByMetaAgentRequest {
	s.Model = &v
	return s
}

func (s *FixSqlByMetaAgentRequest) SetQuery(v string) *FixSqlByMetaAgentRequest {
	s.Query = &v
	return s
}

func (s *FixSqlByMetaAgentRequest) SetSql(v string) *FixSqlByMetaAgentRequest {
	s.Sql = &v
	return s
}

func (s *FixSqlByMetaAgentRequest) Validate() error {
	return dara.Validate(s)
}
