// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaQualityAnalysisJob interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *MediaQualityAnalysisJob
	GetCreateTime() *string
	SetFinishTime(v string) *MediaQualityAnalysisJob
	GetFinishTime() *string
	SetInput(v *MediaQualityAnalysisJobInput) *MediaQualityAnalysisJob
	GetInput() *MediaQualityAnalysisJobInput
	SetJobId(v string) *MediaQualityAnalysisJob
	GetJobId() *string
	SetName(v string) *MediaQualityAnalysisJob
	GetName() *string
	SetScheduleConfig(v *MediaQualityAnalysisJobScheduleConfig) *MediaQualityAnalysisJob
	GetScheduleConfig() *MediaQualityAnalysisJobScheduleConfig
	SetState(v string) *MediaQualityAnalysisJob
	GetState() *string
	SetTemplateConfig(v *MediaQualityAnalysisJobTemplateConfig) *MediaQualityAnalysisJob
	GetTemplateConfig() *MediaQualityAnalysisJobTemplateConfig
	SetUserData(v string) *MediaQualityAnalysisJob
	GetUserData() *string
	SetVqaResult(v *MediaQualityAnalysisJobVqaResult) *MediaQualityAnalysisJob
	GetVqaResult() *MediaQualityAnalysisJobVqaResult
}

type MediaQualityAnalysisJob struct {
	CreateTime     *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime     *string                                `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Input          *MediaQualityAnalysisJobInput          `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	JobId          *string                                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name           *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	ScheduleConfig *MediaQualityAnalysisJobScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	State          *string                                `json:"State,omitempty" xml:"State,omitempty"`
	TemplateConfig *MediaQualityAnalysisJobTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	UserData       *string                                `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VqaResult      *MediaQualityAnalysisJobVqaResult      `json:"VqaResult,omitempty" xml:"VqaResult,omitempty"`
}

func (s MediaQualityAnalysisJob) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJob) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MediaQualityAnalysisJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *MediaQualityAnalysisJob) GetInput() *MediaQualityAnalysisJobInput {
	return s.Input
}

func (s *MediaQualityAnalysisJob) GetJobId() *string {
	return s.JobId
}

func (s *MediaQualityAnalysisJob) GetName() *string {
	return s.Name
}

func (s *MediaQualityAnalysisJob) GetScheduleConfig() *MediaQualityAnalysisJobScheduleConfig {
	return s.ScheduleConfig
}

func (s *MediaQualityAnalysisJob) GetState() *string {
	return s.State
}

func (s *MediaQualityAnalysisJob) GetTemplateConfig() *MediaQualityAnalysisJobTemplateConfig {
	return s.TemplateConfig
}

func (s *MediaQualityAnalysisJob) GetUserData() *string {
	return s.UserData
}

func (s *MediaQualityAnalysisJob) GetVqaResult() *MediaQualityAnalysisJobVqaResult {
	return s.VqaResult
}

func (s *MediaQualityAnalysisJob) SetCreateTime(v string) *MediaQualityAnalysisJob {
	s.CreateTime = &v
	return s
}

func (s *MediaQualityAnalysisJob) SetFinishTime(v string) *MediaQualityAnalysisJob {
	s.FinishTime = &v
	return s
}

func (s *MediaQualityAnalysisJob) SetInput(v *MediaQualityAnalysisJobInput) *MediaQualityAnalysisJob {
	s.Input = v
	return s
}

func (s *MediaQualityAnalysisJob) SetJobId(v string) *MediaQualityAnalysisJob {
	s.JobId = &v
	return s
}

func (s *MediaQualityAnalysisJob) SetName(v string) *MediaQualityAnalysisJob {
	s.Name = &v
	return s
}

func (s *MediaQualityAnalysisJob) SetScheduleConfig(v *MediaQualityAnalysisJobScheduleConfig) *MediaQualityAnalysisJob {
	s.ScheduleConfig = v
	return s
}

func (s *MediaQualityAnalysisJob) SetState(v string) *MediaQualityAnalysisJob {
	s.State = &v
	return s
}

func (s *MediaQualityAnalysisJob) SetTemplateConfig(v *MediaQualityAnalysisJobTemplateConfig) *MediaQualityAnalysisJob {
	s.TemplateConfig = v
	return s
}

func (s *MediaQualityAnalysisJob) SetUserData(v string) *MediaQualityAnalysisJob {
	s.UserData = &v
	return s
}

func (s *MediaQualityAnalysisJob) SetVqaResult(v *MediaQualityAnalysisJobVqaResult) *MediaQualityAnalysisJob {
	s.VqaResult = v
	return s
}

func (s *MediaQualityAnalysisJob) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateConfig != nil {
		if err := s.TemplateConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VqaResult != nil {
		if err := s.VqaResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaQualityAnalysisJobInput struct {
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaQualityAnalysisJobInput) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobInput) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobInput) GetMedia() *string {
	return s.Media
}

func (s *MediaQualityAnalysisJobInput) GetType() *string {
	return s.Type
}

func (s *MediaQualityAnalysisJobInput) SetMedia(v string) *MediaQualityAnalysisJobInput {
	s.Media = &v
	return s
}

func (s *MediaQualityAnalysisJobInput) SetType(v string) *MediaQualityAnalysisJobInput {
	s.Type = &v
	return s
}

func (s *MediaQualityAnalysisJobInput) Validate() error {
	return dara.Validate(s)
}

type MediaQualityAnalysisJobScheduleConfig struct {
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Priority   *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s MediaQualityAnalysisJobScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobScheduleConfig) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *MediaQualityAnalysisJobScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *MediaQualityAnalysisJobScheduleConfig) SetPipelineId(v string) *MediaQualityAnalysisJobScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *MediaQualityAnalysisJobScheduleConfig) SetPriority(v int32) *MediaQualityAnalysisJobScheduleConfig {
	s.Priority = &v
	return s
}

func (s *MediaQualityAnalysisJobScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type MediaQualityAnalysisJobTemplateConfig struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s MediaQualityAnalysisJobTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobTemplateConfig) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobTemplateConfig) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MediaQualityAnalysisJobTemplateConfig) SetTemplateId(v string) *MediaQualityAnalysisJobTemplateConfig {
	s.TemplateId = &v
	return s
}

func (s *MediaQualityAnalysisJobTemplateConfig) Validate() error {
	return dara.Validate(s)
}
