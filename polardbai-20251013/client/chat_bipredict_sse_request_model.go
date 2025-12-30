// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPredictSseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIPredictSseRequest
	GetDbName() *string
	SetGenerateChart(v bool) *ChatBIPredictSseRequest
	GetGenerateChart() *bool
	SetGenerateSummary(v bool) *ChatBIPredictSseRequest
	GetGenerateSummary() *bool
	SetInstanceName(v string) *ChatBIPredictSseRequest
	GetInstanceName() *string
	SetParameters(v *ChatBIPredictSseRequestParameters) *ChatBIPredictSseRequest
	GetParameters() *ChatBIPredictSseRequestParameters
	SetPatternIndexTableName(v string) *ChatBIPredictSseRequest
	GetPatternIndexTableName() *string
	SetQuestion(v string) *ChatBIPredictSseRequest
	GetQuestion() *string
	SetSchemaIndexTableName(v string) *ChatBIPredictSseRequest
	GetSchemaIndexTableName() *string
	SetSelectData(v bool) *ChatBIPredictSseRequest
	GetSelectData() *bool
}

type ChatBIPredictSseRequest struct {
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
	InstanceName *string                            `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Parameters   *ChatBIPredictSseRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
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
}

func (s ChatBIPredictSseRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPredictSseRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPredictSseRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPredictSseRequest) GetGenerateChart() *bool {
	return s.GenerateChart
}

func (s *ChatBIPredictSseRequest) GetGenerateSummary() *bool {
	return s.GenerateSummary
}

func (s *ChatBIPredictSseRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPredictSseRequest) GetParameters() *ChatBIPredictSseRequestParameters {
	return s.Parameters
}

func (s *ChatBIPredictSseRequest) GetPatternIndexTableName() *string {
	return s.PatternIndexTableName
}

func (s *ChatBIPredictSseRequest) GetQuestion() *string {
	return s.Question
}

func (s *ChatBIPredictSseRequest) GetSchemaIndexTableName() *string {
	return s.SchemaIndexTableName
}

func (s *ChatBIPredictSseRequest) GetSelectData() *bool {
	return s.SelectData
}

func (s *ChatBIPredictSseRequest) SetDbName(v string) *ChatBIPredictSseRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPredictSseRequest) SetGenerateChart(v bool) *ChatBIPredictSseRequest {
	s.GenerateChart = &v
	return s
}

func (s *ChatBIPredictSseRequest) SetGenerateSummary(v bool) *ChatBIPredictSseRequest {
	s.GenerateSummary = &v
	return s
}

func (s *ChatBIPredictSseRequest) SetInstanceName(v string) *ChatBIPredictSseRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPredictSseRequest) SetParameters(v *ChatBIPredictSseRequestParameters) *ChatBIPredictSseRequest {
	s.Parameters = v
	return s
}

func (s *ChatBIPredictSseRequest) SetPatternIndexTableName(v string) *ChatBIPredictSseRequest {
	s.PatternIndexTableName = &v
	return s
}

func (s *ChatBIPredictSseRequest) SetQuestion(v string) *ChatBIPredictSseRequest {
	s.Question = &v
	return s
}

func (s *ChatBIPredictSseRequest) SetSchemaIndexTableName(v string) *ChatBIPredictSseRequest {
	s.SchemaIndexTableName = &v
	return s
}

func (s *ChatBIPredictSseRequest) SetSelectData(v bool) *ChatBIPredictSseRequest {
	s.SelectData = &v
	return s
}

func (s *ChatBIPredictSseRequest) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatBIPredictSseRequestParameters struct {
	// example:
	//
	// 0.85
	PatternIndexThreshold *float64 `json:"PatternIndexThreshold,omitempty" xml:"PatternIndexThreshold,omitempty"`
	// example:
	//
	// 2
	PatternIndexTop *int32 `json:"PatternIndexTop,omitempty" xml:"PatternIndexTop,omitempty"`
	// example:
	//
	// IMAGE
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s ChatBIPredictSseRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPredictSseRequestParameters) GoString() string {
	return s.String()
}

func (s *ChatBIPredictSseRequestParameters) GetPatternIndexThreshold() *float64 {
	return s.PatternIndexThreshold
}

func (s *ChatBIPredictSseRequestParameters) GetPatternIndexTop() *int32 {
	return s.PatternIndexTop
}

func (s *ChatBIPredictSseRequestParameters) GetResultType() *string {
	return s.ResultType
}

func (s *ChatBIPredictSseRequestParameters) SetPatternIndexThreshold(v float64) *ChatBIPredictSseRequestParameters {
	s.PatternIndexThreshold = &v
	return s
}

func (s *ChatBIPredictSseRequestParameters) SetPatternIndexTop(v int32) *ChatBIPredictSseRequestParameters {
	s.PatternIndexTop = &v
	return s
}

func (s *ChatBIPredictSseRequestParameters) SetResultType(v string) *ChatBIPredictSseRequestParameters {
	s.ResultType = &v
	return s
}

func (s *ChatBIPredictSseRequestParameters) Validate() error {
	return dara.Validate(s)
}
