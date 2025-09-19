// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerSqlSyntaxByMetaAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *AnswerSqlSyntaxByMetaAgentRequest
	GetDbId() *string
	SetModel(v string) *AnswerSqlSyntaxByMetaAgentRequest
	GetModel() *string
	SetQuery(v string) *AnswerSqlSyntaxByMetaAgentRequest
	GetQuery() *string
}

type AnswerSqlSyntaxByMetaAgentRequest struct {
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
	// 怎么获取当前时间的字符串
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s AnswerSqlSyntaxByMetaAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s AnswerSqlSyntaxByMetaAgentRequest) GoString() string {
	return s.String()
}

func (s *AnswerSqlSyntaxByMetaAgentRequest) GetDbId() *string {
	return s.DbId
}

func (s *AnswerSqlSyntaxByMetaAgentRequest) GetModel() *string {
	return s.Model
}

func (s *AnswerSqlSyntaxByMetaAgentRequest) GetQuery() *string {
	return s.Query
}

func (s *AnswerSqlSyntaxByMetaAgentRequest) SetDbId(v string) *AnswerSqlSyntaxByMetaAgentRequest {
	s.DbId = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentRequest) SetModel(v string) *AnswerSqlSyntaxByMetaAgentRequest {
	s.Model = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentRequest) SetQuery(v string) *AnswerSqlSyntaxByMetaAgentRequest {
	s.Query = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentRequest) Validate() error {
	return dara.Validate(s)
}
