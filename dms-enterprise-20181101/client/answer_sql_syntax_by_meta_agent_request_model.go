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
	// The ID of the database. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name of the selected model. You can use only Qwen series models.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The syntax question.
	//
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
