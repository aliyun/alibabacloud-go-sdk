// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaQualityAnalysisJobsResponseItem interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *ListMediaQualityAnalysisJobsResponseItem
	GetCreateTime() *string
	SetFinishTime(v string) *ListMediaQualityAnalysisJobsResponseItem
	GetFinishTime() *string
	SetInput(v *ListMediaQualityAnalysisJobsResponseItemInput) *ListMediaQualityAnalysisJobsResponseItem
	GetInput() *ListMediaQualityAnalysisJobsResponseItemInput
	SetJobId(v string) *ListMediaQualityAnalysisJobsResponseItem
	GetJobId() *string
	SetName(v string) *ListMediaQualityAnalysisJobsResponseItem
	GetName() *string
	SetScheduleConfig(v *ListMediaQualityAnalysisJobsResponseItemScheduleConfig) *ListMediaQualityAnalysisJobsResponseItem
	GetScheduleConfig() *ListMediaQualityAnalysisJobsResponseItemScheduleConfig
	SetState(v string) *ListMediaQualityAnalysisJobsResponseItem
	GetState() *string
	SetTemplateConfig(v *ListMediaQualityAnalysisJobsResponseItemTemplateConfig) *ListMediaQualityAnalysisJobsResponseItem
	GetTemplateConfig() *ListMediaQualityAnalysisJobsResponseItemTemplateConfig
	SetUserData(v string) *ListMediaQualityAnalysisJobsResponseItem
	GetUserData() *string
}

type ListMediaQualityAnalysisJobsResponseItem struct {
	CreateTime     *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime     *string                                                 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Input          *ListMediaQualityAnalysisJobsResponseItemInput          `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	JobId          *string                                                 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name           *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	ScheduleConfig *ListMediaQualityAnalysisJobsResponseItemScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	State          *string                                                 `json:"State,omitempty" xml:"State,omitempty"`
	TemplateConfig *ListMediaQualityAnalysisJobsResponseItemTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	UserData       *string                                                 `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListMediaQualityAnalysisJobsResponseItem) String() string {
	return dara.Prettify(s)
}

func (s ListMediaQualityAnalysisJobsResponseItem) GoString() string {
	return s.String()
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetInput() *ListMediaQualityAnalysisJobsResponseItemInput {
	return s.Input
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetName() *string {
	return s.Name
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetScheduleConfig() *ListMediaQualityAnalysisJobsResponseItemScheduleConfig {
	return s.ScheduleConfig
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetState() *string {
	return s.State
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetTemplateConfig() *ListMediaQualityAnalysisJobsResponseItemTemplateConfig {
	return s.TemplateConfig
}

func (s *ListMediaQualityAnalysisJobsResponseItem) GetUserData() *string {
	return s.UserData
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetCreateTime(v string) *ListMediaQualityAnalysisJobsResponseItem {
	s.CreateTime = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetFinishTime(v string) *ListMediaQualityAnalysisJobsResponseItem {
	s.FinishTime = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetInput(v *ListMediaQualityAnalysisJobsResponseItemInput) *ListMediaQualityAnalysisJobsResponseItem {
	s.Input = v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetJobId(v string) *ListMediaQualityAnalysisJobsResponseItem {
	s.JobId = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetName(v string) *ListMediaQualityAnalysisJobsResponseItem {
	s.Name = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetScheduleConfig(v *ListMediaQualityAnalysisJobsResponseItemScheduleConfig) *ListMediaQualityAnalysisJobsResponseItem {
	s.ScheduleConfig = v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetState(v string) *ListMediaQualityAnalysisJobsResponseItem {
	s.State = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetTemplateConfig(v *ListMediaQualityAnalysisJobsResponseItemTemplateConfig) *ListMediaQualityAnalysisJobsResponseItem {
	s.TemplateConfig = v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) SetUserData(v string) *ListMediaQualityAnalysisJobsResponseItem {
	s.UserData = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItem) Validate() error {
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
	return nil
}

type ListMediaQualityAnalysisJobsResponseItemInput struct {
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMediaQualityAnalysisJobsResponseItemInput) String() string {
	return dara.Prettify(s)
}

func (s ListMediaQualityAnalysisJobsResponseItemInput) GoString() string {
	return s.String()
}

func (s *ListMediaQualityAnalysisJobsResponseItemInput) GetMedia() *string {
	return s.Media
}

func (s *ListMediaQualityAnalysisJobsResponseItemInput) GetType() *string {
	return s.Type
}

func (s *ListMediaQualityAnalysisJobsResponseItemInput) SetMedia(v string) *ListMediaQualityAnalysisJobsResponseItemInput {
	s.Media = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItemInput) SetType(v string) *ListMediaQualityAnalysisJobsResponseItemInput {
	s.Type = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItemInput) Validate() error {
	return dara.Validate(s)
}

type ListMediaQualityAnalysisJobsResponseItemScheduleConfig struct {
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Priority   *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ListMediaQualityAnalysisJobsResponseItemScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMediaQualityAnalysisJobsResponseItemScheduleConfig) GoString() string {
	return s.String()
}

func (s *ListMediaQualityAnalysisJobsResponseItemScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListMediaQualityAnalysisJobsResponseItemScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *ListMediaQualityAnalysisJobsResponseItemScheduleConfig) SetPipelineId(v string) *ListMediaQualityAnalysisJobsResponseItemScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItemScheduleConfig) SetPriority(v int32) *ListMediaQualityAnalysisJobsResponseItemScheduleConfig {
	s.Priority = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItemScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type ListMediaQualityAnalysisJobsResponseItemTemplateConfig struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListMediaQualityAnalysisJobsResponseItemTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMediaQualityAnalysisJobsResponseItemTemplateConfig) GoString() string {
	return s.String()
}

func (s *ListMediaQualityAnalysisJobsResponseItemTemplateConfig) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListMediaQualityAnalysisJobsResponseItemTemplateConfig) SetTemplateId(v string) *ListMediaQualityAnalysisJobsResponseItemTemplateConfig {
	s.TemplateId = &v
	return s
}

func (s *ListMediaQualityAnalysisJobsResponseItemTemplateConfig) Validate() error {
	return dara.Validate(s)
}
