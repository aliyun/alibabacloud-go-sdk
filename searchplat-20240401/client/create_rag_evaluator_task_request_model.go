// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRagEvaluatorTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateRagEvaluatorTaskRequest
	GetAppName() *string
	SetData(v []*CreateRagEvaluatorTaskRequestData) *CreateRagEvaluatorTaskRequest
	GetData() []*CreateRagEvaluatorTaskRequestData
	SetDataSourceConfig(v interface{}) *CreateRagEvaluatorTaskRequest
	GetDataSourceConfig() interface{}
	SetEmails(v []*string) *CreateRagEvaluatorTaskRequest
	GetEmails() []*string
	SetEvaluateConfig(v *CreateRagEvaluatorTaskRequestEvaluateConfig) *CreateRagEvaluatorTaskRequest
	GetEvaluateConfig() *CreateRagEvaluatorTaskRequestEvaluateConfig
	SetHasDataSource(v bool) *CreateRagEvaluatorTaskRequest
	GetHasDataSource() *bool
	SetMetrics(v []interface{}) *CreateRagEvaluatorTaskRequest
	GetMetrics() []interface{}
	SetTaskName(v string) *CreateRagEvaluatorTaskRequest
	GetTaskName() *string
}

type CreateRagEvaluatorTaskRequest struct {
	// app_name
	//
	// example:
	//
	// 空
	AppName *string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// The list of evaluation data.
	Data []*CreateRagEvaluatorTaskRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The datasource config.
	//
	// example:
	//
	// {
	//
	// "data_source_type": "oss",
	//
	// "file_path": "oss://xxx.",
	//
	// "file_name": "04837719-default-zz.xlsx"
	//
	// }
	DataSourceConfig interface{} `json:"data_source_config,omitempty" xml:"data_source_config,omitempty"`
	// emails
	Emails []*string `json:"emails,omitempty" xml:"emails,omitempty" type:"Repeated"`
	// The evaluation configuration.
	EvaluateConfig *CreateRagEvaluatorTaskRequestEvaluateConfig `json:"evaluate_config,omitempty" xml:"evaluate_config,omitempty" type:"Struct"`
	// has_data_source
	//
	// example:
	//
	// 空
	HasDataSource *bool `json:"has_data_source,omitempty" xml:"has_data_source,omitempty"`
	// The metric values. Valid values:
	//
	// - context_recall
	//
	// - context_precision
	//
	// - faithfulness
	//
	// - satisfaction
	//
	// - comprehensive_score.
	Metrics []interface{} `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// The evaluation task name.
	//
	// example:
	//
	// taskName
	TaskName *string `json:"task_name,omitempty" xml:"task_name,omitempty"`
}

func (s CreateRagEvaluatorTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRagEvaluatorTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRagEvaluatorTaskRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateRagEvaluatorTaskRequest) GetData() []*CreateRagEvaluatorTaskRequestData {
	return s.Data
}

func (s *CreateRagEvaluatorTaskRequest) GetDataSourceConfig() interface{} {
	return s.DataSourceConfig
}

func (s *CreateRagEvaluatorTaskRequest) GetEmails() []*string {
	return s.Emails
}

func (s *CreateRagEvaluatorTaskRequest) GetEvaluateConfig() *CreateRagEvaluatorTaskRequestEvaluateConfig {
	return s.EvaluateConfig
}

func (s *CreateRagEvaluatorTaskRequest) GetHasDataSource() *bool {
	return s.HasDataSource
}

func (s *CreateRagEvaluatorTaskRequest) GetMetrics() []interface{} {
	return s.Metrics
}

func (s *CreateRagEvaluatorTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateRagEvaluatorTaskRequest) SetAppName(v string) *CreateRagEvaluatorTaskRequest {
	s.AppName = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) SetData(v []*CreateRagEvaluatorTaskRequestData) *CreateRagEvaluatorTaskRequest {
	s.Data = v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) SetDataSourceConfig(v interface{}) *CreateRagEvaluatorTaskRequest {
	s.DataSourceConfig = v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) SetEmails(v []*string) *CreateRagEvaluatorTaskRequest {
	s.Emails = v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) SetEvaluateConfig(v *CreateRagEvaluatorTaskRequestEvaluateConfig) *CreateRagEvaluatorTaskRequest {
	s.EvaluateConfig = v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) SetHasDataSource(v bool) *CreateRagEvaluatorTaskRequest {
	s.HasDataSource = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) SetMetrics(v []interface{}) *CreateRagEvaluatorTaskRequest {
	s.Metrics = v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) SetTaskName(v string) *CreateRagEvaluatorTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequest) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EvaluateConfig != nil {
		if err := s.EvaluateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRagEvaluatorTaskRequestData struct {
	// model_answer
	//
	// example:
	//
	// 空
	ModelAnswer *string `json:"model_answer,omitempty" xml:"model_answer,omitempty"`
	// question
	//
	// example:
	//
	// 空
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// recall_docs
	RecallDocs []*string `json:"recall_docs,omitempty" xml:"recall_docs,omitempty" type:"Repeated"`
	// standard_answer
	//
	// example:
	//
	// 空
	StandardAnswer *string `json:"standard_answer,omitempty" xml:"standard_answer,omitempty"`
}

func (s CreateRagEvaluatorTaskRequestData) String() string {
	return dara.Prettify(s)
}

func (s CreateRagEvaluatorTaskRequestData) GoString() string {
	return s.String()
}

func (s *CreateRagEvaluatorTaskRequestData) GetModelAnswer() *string {
	return s.ModelAnswer
}

func (s *CreateRagEvaluatorTaskRequestData) GetQuestion() *string {
	return s.Question
}

func (s *CreateRagEvaluatorTaskRequestData) GetRecallDocs() []*string {
	return s.RecallDocs
}

func (s *CreateRagEvaluatorTaskRequestData) GetStandardAnswer() *string {
	return s.StandardAnswer
}

func (s *CreateRagEvaluatorTaskRequestData) SetModelAnswer(v string) *CreateRagEvaluatorTaskRequestData {
	s.ModelAnswer = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequestData) SetQuestion(v string) *CreateRagEvaluatorTaskRequestData {
	s.Question = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequestData) SetRecallDocs(v []*string) *CreateRagEvaluatorTaskRequestData {
	s.RecallDocs = v
	return s
}

func (s *CreateRagEvaluatorTaskRequestData) SetStandardAnswer(v string) *CreateRagEvaluatorTaskRequestData {
	s.StandardAnswer = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequestData) Validate() error {
	return dara.Validate(s)
}

type CreateRagEvaluatorTaskRequestEvaluateConfig struct {
	// The model to use.
	//
	// example:
	//
	// qwen-72b
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// prompt
	//
	// example:
	//
	// 空
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// run_all_step
	//
	// example:
	//
	// false
	RunAllStep *bool `json:"run_all_step,omitempty" xml:"run_all_step,omitempty"`
}

func (s CreateRagEvaluatorTaskRequestEvaluateConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRagEvaluatorTaskRequestEvaluateConfig) GoString() string {
	return s.String()
}

func (s *CreateRagEvaluatorTaskRequestEvaluateConfig) GetModel() *string {
	return s.Model
}

func (s *CreateRagEvaluatorTaskRequestEvaluateConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateRagEvaluatorTaskRequestEvaluateConfig) GetRunAllStep() *bool {
	return s.RunAllStep
}

func (s *CreateRagEvaluatorTaskRequestEvaluateConfig) SetModel(v string) *CreateRagEvaluatorTaskRequestEvaluateConfig {
	s.Model = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequestEvaluateConfig) SetPrompt(v string) *CreateRagEvaluatorTaskRequestEvaluateConfig {
	s.Prompt = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequestEvaluateConfig) SetRunAllStep(v bool) *CreateRagEvaluatorTaskRequestEvaluateConfig {
	s.RunAllStep = &v
	return s
}

func (s *CreateRagEvaluatorTaskRequestEvaluateConfig) Validate() error {
	return dara.Validate(s)
}
