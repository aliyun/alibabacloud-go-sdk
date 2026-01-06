// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPredictSseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIPredictSseShrinkRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIPredictSseShrinkRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIPredictSseShrinkRequest
	GetDbName() *string
	SetGenerateChart(v bool) *ChatBIPredictSseShrinkRequest
	GetGenerateChart() *bool
	SetGenerateSummary(v bool) *ChatBIPredictSseShrinkRequest
	GetGenerateSummary() *bool
	SetInstanceName(v string) *ChatBIPredictSseShrinkRequest
	GetInstanceName() *string
	SetParametersShrink(v string) *ChatBIPredictSseShrinkRequest
	GetParametersShrink() *string
	SetPatternIndexTableName(v string) *ChatBIPredictSseShrinkRequest
	GetPatternIndexTableName() *string
	SetQuestion(v string) *ChatBIPredictSseShrinkRequest
	GetQuestion() *string
	SetSchemaIndexTableName(v string) *ChatBIPredictSseShrinkRequest
	GetSchemaIndexTableName() *string
	SetSelectData(v bool) *ChatBIPredictSseShrinkRequest
	GetSelectData() *bool
	SetThinkingMode(v bool) *ChatBIPredictSseShrinkRequest
	GetThinkingMode() *bool
}

type ChatBIPredictSseShrinkRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// false
	GenerateChart *bool `json:"GenerateChart,omitempty" xml:"GenerateChart,omitempty"`
	// example:
	//
	// true
	GenerateSummary *bool `json:"GenerateSummary,omitempty" xml:"GenerateSummary,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze7smdi2f*******
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 空字符串, pattern_index_1
	PatternIndexTableName *string `json:"PatternIndexTableName,omitempty" xml:"PatternIndexTableName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 查询2023年10月1日至2023年10月3日期间开课的课程名称和上课地点。
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// schema_index_1
	SchemaIndexTableName *string `json:"SchemaIndexTableName,omitempty" xml:"SchemaIndexTableName,omitempty"`
	// example:
	//
	// true
	SelectData *bool `json:"SelectData,omitempty" xml:"SelectData,omitempty"`
	// example:
	//
	// false
	ThinkingMode *bool `json:"ThinkingMode,omitempty" xml:"ThinkingMode,omitempty"`
}

func (s ChatBIPredictSseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPredictSseShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPredictSseShrinkRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIPredictSseShrinkRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIPredictSseShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPredictSseShrinkRequest) GetGenerateChart() *bool {
	return s.GenerateChart
}

func (s *ChatBIPredictSseShrinkRequest) GetGenerateSummary() *bool {
	return s.GenerateSummary
}

func (s *ChatBIPredictSseShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPredictSseShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *ChatBIPredictSseShrinkRequest) GetPatternIndexTableName() *string {
	return s.PatternIndexTableName
}

func (s *ChatBIPredictSseShrinkRequest) GetQuestion() *string {
	return s.Question
}

func (s *ChatBIPredictSseShrinkRequest) GetSchemaIndexTableName() *string {
	return s.SchemaIndexTableName
}

func (s *ChatBIPredictSseShrinkRequest) GetSelectData() *bool {
	return s.SelectData
}

func (s *ChatBIPredictSseShrinkRequest) GetThinkingMode() *bool {
	return s.ThinkingMode
}

func (s *ChatBIPredictSseShrinkRequest) SetAuthMessage(v string) *ChatBIPredictSseShrinkRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetAuthType(v string) *ChatBIPredictSseShrinkRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetDbName(v string) *ChatBIPredictSseShrinkRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetGenerateChart(v bool) *ChatBIPredictSseShrinkRequest {
	s.GenerateChart = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetGenerateSummary(v bool) *ChatBIPredictSseShrinkRequest {
	s.GenerateSummary = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetInstanceName(v string) *ChatBIPredictSseShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetParametersShrink(v string) *ChatBIPredictSseShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetPatternIndexTableName(v string) *ChatBIPredictSseShrinkRequest {
	s.PatternIndexTableName = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetQuestion(v string) *ChatBIPredictSseShrinkRequest {
	s.Question = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetSchemaIndexTableName(v string) *ChatBIPredictSseShrinkRequest {
	s.SchemaIndexTableName = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetSelectData(v bool) *ChatBIPredictSseShrinkRequest {
	s.SelectData = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) SetThinkingMode(v bool) *ChatBIPredictSseShrinkRequest {
	s.ThinkingMode = &v
	return s
}

func (s *ChatBIPredictSseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
