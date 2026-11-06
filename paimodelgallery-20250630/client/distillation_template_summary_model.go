// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistillationTemplateSummary interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilityTags(v []*string) *DistillationTemplateSummary
	GetCapabilityTags() []*string
	SetCategory(v string) *DistillationTemplateSummary
	GetCategory() *string
	SetDescription(v string) *DistillationTemplateSummary
	GetDescription() *string
	SetJobType(v string) *DistillationTemplateSummary
	GetJobType() *string
	SetOrderNumber(v int32) *DistillationTemplateSummary
	GetOrderNumber() *int32
	SetPipelineStages(v []*DistillationTemplateSummaryPipelineStages) *DistillationTemplateSummary
	GetPipelineStages() []*DistillationTemplateSummaryPipelineStages
	SetTemplateId(v string) *DistillationTemplateSummary
	GetTemplateId() *string
	SetTemplateKind(v string) *DistillationTemplateSummary
	GetTemplateKind() *string
	SetTemplateName(v string) *DistillationTemplateSummary
	GetTemplateName() *string
	SetTrainingOptions(v []*DistillationTemplateSummaryTrainingOptions) *DistillationTemplateSummary
	GetTrainingOptions() []*DistillationTemplateSummaryTrainingOptions
}

type DistillationTemplateSummary struct {
	CapabilityTags []*string `json:"CapabilityTags,omitempty" xml:"CapabilityTags,omitempty" type:"Repeated"`
	// example:
	//
	// reasoning
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 面向数学、逻辑、代码等需要多步推理的场景，产出 SFT 数据集。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// advanced_cot_distill
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// 10
	OrderNumber    *int32                                       `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	PipelineStages []*DistillationTemplateSummaryPipelineStages `json:"PipelineStages,omitempty" xml:"PipelineStages,omitempty" type:"Repeated"`
	// example:
	//
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// scenario
	TemplateKind *string `json:"TemplateKind,omitempty" xml:"TemplateKind,omitempty"`
	// example:
	//
	// 思维链推理蒸馏
	TemplateName    *string                                       `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TrainingOptions []*DistillationTemplateSummaryTrainingOptions `json:"TrainingOptions,omitempty" xml:"TrainingOptions,omitempty" type:"Repeated"`
}

func (s DistillationTemplateSummary) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplateSummary) GoString() string {
	return s.String()
}

func (s *DistillationTemplateSummary) GetCapabilityTags() []*string {
	return s.CapabilityTags
}

func (s *DistillationTemplateSummary) GetCategory() *string {
	return s.Category
}

func (s *DistillationTemplateSummary) GetDescription() *string {
	return s.Description
}

func (s *DistillationTemplateSummary) GetJobType() *string {
	return s.JobType
}

func (s *DistillationTemplateSummary) GetOrderNumber() *int32 {
	return s.OrderNumber
}

func (s *DistillationTemplateSummary) GetPipelineStages() []*DistillationTemplateSummaryPipelineStages {
	return s.PipelineStages
}

func (s *DistillationTemplateSummary) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DistillationTemplateSummary) GetTemplateKind() *string {
	return s.TemplateKind
}

func (s *DistillationTemplateSummary) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DistillationTemplateSummary) GetTrainingOptions() []*DistillationTemplateSummaryTrainingOptions {
	return s.TrainingOptions
}

func (s *DistillationTemplateSummary) SetCapabilityTags(v []*string) *DistillationTemplateSummary {
	s.CapabilityTags = v
	return s
}

func (s *DistillationTemplateSummary) SetCategory(v string) *DistillationTemplateSummary {
	s.Category = &v
	return s
}

func (s *DistillationTemplateSummary) SetDescription(v string) *DistillationTemplateSummary {
	s.Description = &v
	return s
}

func (s *DistillationTemplateSummary) SetJobType(v string) *DistillationTemplateSummary {
	s.JobType = &v
	return s
}

func (s *DistillationTemplateSummary) SetOrderNumber(v int32) *DistillationTemplateSummary {
	s.OrderNumber = &v
	return s
}

func (s *DistillationTemplateSummary) SetPipelineStages(v []*DistillationTemplateSummaryPipelineStages) *DistillationTemplateSummary {
	s.PipelineStages = v
	return s
}

func (s *DistillationTemplateSummary) SetTemplateId(v string) *DistillationTemplateSummary {
	s.TemplateId = &v
	return s
}

func (s *DistillationTemplateSummary) SetTemplateKind(v string) *DistillationTemplateSummary {
	s.TemplateKind = &v
	return s
}

func (s *DistillationTemplateSummary) SetTemplateName(v string) *DistillationTemplateSummary {
	s.TemplateName = &v
	return s
}

func (s *DistillationTemplateSummary) SetTrainingOptions(v []*DistillationTemplateSummaryTrainingOptions) *DistillationTemplateSummary {
	s.TrainingOptions = v
	return s
}

func (s *DistillationTemplateSummary) Validate() error {
	if s.PipelineStages != nil {
		for _, item := range s.PipelineStages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TrainingOptions != nil {
		for _, item := range s.TrainingOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DistillationTemplateSummaryPipelineStages struct {
	// example:
	//
	// 由教师模型对每个问题生成带推理过程的回答。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cot_distill
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 生成思维链
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DistillationTemplateSummaryPipelineStages) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplateSummaryPipelineStages) GoString() string {
	return s.String()
}

func (s *DistillationTemplateSummaryPipelineStages) GetDescription() *string {
	return s.Description
}

func (s *DistillationTemplateSummaryPipelineStages) GetKey() *string {
	return s.Key
}

func (s *DistillationTemplateSummaryPipelineStages) GetName() *string {
	return s.Name
}

func (s *DistillationTemplateSummaryPipelineStages) SetDescription(v string) *DistillationTemplateSummaryPipelineStages {
	s.Description = &v
	return s
}

func (s *DistillationTemplateSummaryPipelineStages) SetKey(v string) *DistillationTemplateSummaryPipelineStages {
	s.Key = &v
	return s
}

func (s *DistillationTemplateSummaryPipelineStages) SetName(v string) *DistillationTemplateSummaryPipelineStages {
	s.Name = &v
	return s
}

func (s *DistillationTemplateSummaryPipelineStages) Validate() error {
	return dara.Validate(s)
}

type DistillationTemplateSummaryTrainingOptions struct {
	ModelTasks      []*string `json:"ModelTasks,omitempty" xml:"ModelTasks,omitempty" type:"Repeated"`
	TrainingMethods []*string `json:"TrainingMethods,omitempty" xml:"TrainingMethods,omitempty" type:"Repeated"`
	// example:
	//
	// sft
	TrainingType *string `json:"TrainingType,omitempty" xml:"TrainingType,omitempty"`
}

func (s DistillationTemplateSummaryTrainingOptions) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplateSummaryTrainingOptions) GoString() string {
	return s.String()
}

func (s *DistillationTemplateSummaryTrainingOptions) GetModelTasks() []*string {
	return s.ModelTasks
}

func (s *DistillationTemplateSummaryTrainingOptions) GetTrainingMethods() []*string {
	return s.TrainingMethods
}

func (s *DistillationTemplateSummaryTrainingOptions) GetTrainingType() *string {
	return s.TrainingType
}

func (s *DistillationTemplateSummaryTrainingOptions) SetModelTasks(v []*string) *DistillationTemplateSummaryTrainingOptions {
	s.ModelTasks = v
	return s
}

func (s *DistillationTemplateSummaryTrainingOptions) SetTrainingMethods(v []*string) *DistillationTemplateSummaryTrainingOptions {
	s.TrainingMethods = v
	return s
}

func (s *DistillationTemplateSummaryTrainingOptions) SetTrainingType(v string) *DistillationTemplateSummaryTrainingOptions {
	s.TrainingType = &v
	return s
}

func (s *DistillationTemplateSummaryTrainingOptions) Validate() error {
	return dara.Validate(s)
}
