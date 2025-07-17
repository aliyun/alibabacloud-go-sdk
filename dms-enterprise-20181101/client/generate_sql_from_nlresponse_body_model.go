// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSqlFromNLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateSqlFromNLResponseBodyData) *GenerateSqlFromNLResponseBody
	GetData() *GenerateSqlFromNLResponseBodyData
	SetErrorCode(v string) *GenerateSqlFromNLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GenerateSqlFromNLResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GenerateSqlFromNLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateSqlFromNLResponseBody
	GetSuccess() *bool
}

type GenerateSqlFromNLResponseBody struct {
	Data *GenerateSqlFromNLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateSqlFromNLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlFromNLResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateSqlFromNLResponseBody) GetData() *GenerateSqlFromNLResponseBodyData {
	return s.Data
}

func (s *GenerateSqlFromNLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GenerateSqlFromNLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GenerateSqlFromNLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateSqlFromNLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateSqlFromNLResponseBody) SetData(v *GenerateSqlFromNLResponseBodyData) *GenerateSqlFromNLResponseBody {
	s.Data = v
	return s
}

func (s *GenerateSqlFromNLResponseBody) SetErrorCode(v string) *GenerateSqlFromNLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GenerateSqlFromNLResponseBody) SetErrorMessage(v string) *GenerateSqlFromNLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenerateSqlFromNLResponseBody) SetRequestId(v string) *GenerateSqlFromNLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateSqlFromNLResponseBody) SetSuccess(v bool) *GenerateSqlFromNLResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateSqlFromNLResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateSqlFromNLResponseBodyData struct {
	KnowledgeReferences []*GenerateSqlFromNLResponseBodyDataKnowledgeReferences `json:"KnowledgeReferences,omitempty" xml:"KnowledgeReferences,omitempty" type:"Repeated"`
	Question            *string                                                 `json:"Question,omitempty" xml:"Question,omitempty"`
	SimilarSql          []*GenerateSqlFromNLResponseBodyDataSimilarSql          `json:"SimilarSql,omitempty" xml:"SimilarSql,omitempty" type:"Repeated"`
	// example:
	//
	// SELECT 	- FROM table WHERE condition;
	Sql    *string                                    `json:"Sql,omitempty" xml:"Sql,omitempty"`
	Tables []*GenerateSqlFromNLResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 通过分析用户的问题和提供的知识，生成了相应的SQL语句。
	Thought *string `json:"Thought,omitempty" xml:"Thought,omitempty"`
}

func (s GenerateSqlFromNLResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlFromNLResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateSqlFromNLResponseBodyData) GetKnowledgeReferences() []*GenerateSqlFromNLResponseBodyDataKnowledgeReferences {
	return s.KnowledgeReferences
}

func (s *GenerateSqlFromNLResponseBodyData) GetQuestion() *string {
	return s.Question
}

func (s *GenerateSqlFromNLResponseBodyData) GetSimilarSql() []*GenerateSqlFromNLResponseBodyDataSimilarSql {
	return s.SimilarSql
}

func (s *GenerateSqlFromNLResponseBodyData) GetSql() *string {
	return s.Sql
}

func (s *GenerateSqlFromNLResponseBodyData) GetTables() []*GenerateSqlFromNLResponseBodyDataTables {
	return s.Tables
}

func (s *GenerateSqlFromNLResponseBodyData) GetThought() *string {
	return s.Thought
}

func (s *GenerateSqlFromNLResponseBodyData) SetKnowledgeReferences(v []*GenerateSqlFromNLResponseBodyDataKnowledgeReferences) *GenerateSqlFromNLResponseBodyData {
	s.KnowledgeReferences = v
	return s
}

func (s *GenerateSqlFromNLResponseBodyData) SetQuestion(v string) *GenerateSqlFromNLResponseBodyData {
	s.Question = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyData) SetSimilarSql(v []*GenerateSqlFromNLResponseBodyDataSimilarSql) *GenerateSqlFromNLResponseBodyData {
	s.SimilarSql = v
	return s
}

func (s *GenerateSqlFromNLResponseBodyData) SetSql(v string) *GenerateSqlFromNLResponseBodyData {
	s.Sql = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyData) SetTables(v []*GenerateSqlFromNLResponseBodyDataTables) *GenerateSqlFromNLResponseBodyData {
	s.Tables = v
	return s
}

func (s *GenerateSqlFromNLResponseBodyData) SetThought(v string) *GenerateSqlFromNLResponseBodyData {
	s.Thought = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GenerateSqlFromNLResponseBodyDataKnowledgeReferences struct {
	// example:
	//
	// {\\"Status\\": \\"OK\\"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// verified
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// sample_tbl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GenerateSqlFromNLResponseBodyDataKnowledgeReferences) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlFromNLResponseBodyDataKnowledgeReferences) GoString() string {
	return s.String()
}

func (s *GenerateSqlFromNLResponseBodyDataKnowledgeReferences) GetContent() *string {
	return s.Content
}

func (s *GenerateSqlFromNLResponseBodyDataKnowledgeReferences) GetLevel() *string {
	return s.Level
}

func (s *GenerateSqlFromNLResponseBodyDataKnowledgeReferences) GetName() *string {
	return s.Name
}

func (s *GenerateSqlFromNLResponseBodyDataKnowledgeReferences) SetContent(v string) *GenerateSqlFromNLResponseBodyDataKnowledgeReferences {
	s.Content = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataKnowledgeReferences) SetLevel(v string) *GenerateSqlFromNLResponseBodyDataKnowledgeReferences {
	s.Level = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataKnowledgeReferences) SetName(v string) *GenerateSqlFromNLResponseBodyDataKnowledgeReferences {
	s.Name = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataKnowledgeReferences) Validate() error {
	return dara.Validate(s)
}

type GenerateSqlFromNLResponseBodyDataSimilarSql struct {
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 0.52
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// SELECT 	- WHERE ResourceType = \\"ACS::ECS::Instance\\" AND ResourceGroupId != \\"rg-xxx\\"
	Sql     *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	Thought *string `json:"Thought,omitempty" xml:"Thought,omitempty"`
}

func (s GenerateSqlFromNLResponseBodyDataSimilarSql) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlFromNLResponseBodyDataSimilarSql) GoString() string {
	return s.String()
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) GetQuestion() *string {
	return s.Question
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) GetScore() *string {
	return s.Score
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) GetSql() *string {
	return s.Sql
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) GetThought() *string {
	return s.Thought
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) SetQuestion(v string) *GenerateSqlFromNLResponseBodyDataSimilarSql {
	s.Question = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) SetScore(v string) *GenerateSqlFromNLResponseBodyDataSimilarSql {
	s.Score = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) SetSql(v string) *GenerateSqlFromNLResponseBodyDataSimilarSql {
	s.Sql = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) SetThought(v string) *GenerateSqlFromNLResponseBodyDataSimilarSql {
	s.Thought = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataSimilarSql) Validate() error {
	return dara.Validate(s)
}

type GenerateSqlFromNLResponseBodyDataTables struct {
	// example:
	//
	// ins_1.db1.table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GenerateSqlFromNLResponseBodyDataTables) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlFromNLResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *GenerateSqlFromNLResponseBodyDataTables) GetTableName() *string {
	return s.TableName
}

func (s *GenerateSqlFromNLResponseBodyDataTables) SetTableName(v string) *GenerateSqlFromNLResponseBodyDataTables {
	s.TableName = &v
	return s
}

func (s *GenerateSqlFromNLResponseBodyDataTables) Validate() error {
	return dara.Validate(s)
}
