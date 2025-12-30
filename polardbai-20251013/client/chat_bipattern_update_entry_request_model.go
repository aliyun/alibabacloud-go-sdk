// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternUpdateEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIPatternUpdateEntryRequest
	GetDbName() *string
	SetId(v int32) *ChatBIPatternUpdateEntryRequest
	GetId() *int32
	SetInstanceName(v string) *ChatBIPatternUpdateEntryRequest
	GetInstanceName() *string
	SetPatternDescription(v string) *ChatBIPatternUpdateEntryRequest
	GetPatternDescription() *string
	SetPatternParams(v string) *ChatBIPatternUpdateEntryRequest
	GetPatternParams() *string
	SetPatternQuestion(v string) *ChatBIPatternUpdateEntryRequest
	GetPatternQuestion() *string
	SetPatternSql(v string) *ChatBIPatternUpdateEntryRequest
	GetPatternSql() *string
	SetTableName(v string) *ChatBIPatternUpdateEntryRequest
	GetTableName() *string
}

type ChatBIPatternUpdateEntryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 1, 2, 3, ...
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 【课程名称】【授课状态】的课程有哪些？
	PatternDescription *string `json:"PatternDescription,omitempty" xml:"PatternDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// []
	PatternParams *string `json:"PatternParams,omitempty" xml:"PatternParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 查询课程名称为#{courseName}授课状态为#{status}的课程
	PatternQuestion *string `json:"PatternQuestion,omitempty" xml:"PatternQuestion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SELECT course_name, course_time, course_location
	//
	// FROM courses
	//
	// WHERE
	//
	// course_name=#{courseName}
	//
	// AND status=#{statusCode}
	PatternSql *string `json:"PatternSql,omitempty" xml:"PatternSql,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polar4ai_nl2sql_pattern, polar4ai_nl2sql_pattern_1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ChatBIPatternUpdateEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternUpdateEntryRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternUpdateEntryRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternUpdateEntryRequest) GetId() *int32 {
	return s.Id
}

func (s *ChatBIPatternUpdateEntryRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternUpdateEntryRequest) GetPatternDescription() *string {
	return s.PatternDescription
}

func (s *ChatBIPatternUpdateEntryRequest) GetPatternParams() *string {
	return s.PatternParams
}

func (s *ChatBIPatternUpdateEntryRequest) GetPatternQuestion() *string {
	return s.PatternQuestion
}

func (s *ChatBIPatternUpdateEntryRequest) GetPatternSql() *string {
	return s.PatternSql
}

func (s *ChatBIPatternUpdateEntryRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBIPatternUpdateEntryRequest) SetDbName(v string) *ChatBIPatternUpdateEntryRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) SetId(v int32) *ChatBIPatternUpdateEntryRequest {
	s.Id = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) SetInstanceName(v string) *ChatBIPatternUpdateEntryRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) SetPatternDescription(v string) *ChatBIPatternUpdateEntryRequest {
	s.PatternDescription = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) SetPatternParams(v string) *ChatBIPatternUpdateEntryRequest {
	s.PatternParams = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) SetPatternQuestion(v string) *ChatBIPatternUpdateEntryRequest {
	s.PatternQuestion = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) SetPatternSql(v string) *ChatBIPatternUpdateEntryRequest {
	s.PatternSql = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) SetTableName(v string) *ChatBIPatternUpdateEntryRequest {
	s.TableName = &v
	return s
}

func (s *ChatBIPatternUpdateEntryRequest) Validate() error {
	return dara.Validate(s)
}
