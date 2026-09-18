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
	SetTemplateName(v string) *DistillationTemplateSummary
	GetTemplateName() *string
	SetTrainingOptions(v []*DistillationTemplateSummaryTrainingOptions) *DistillationTemplateSummary
	GetTrainingOptions() []*DistillationTemplateSummaryTrainingOptions
}

type DistillationTemplateSummary struct {
	// The list of capability tags, used for scenario card display.
	CapabilityTags []*string `json:"CapabilityTags,omitempty" xml:"CapabilityTags,omitempty" type:"Repeated"`
	// The template category. The frontend uses this value to filter scenario cards.
	//
	// example:
	//
	// reasoning
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The template description, localized based on the language specified in the request. The description specifies applicable scenarios and outputs.
	//
	// example:
	//
	// Designed for scenarios that require multi-step reasoning, such as math, logic, and code. Produces an SFT dataset
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The algorithm job type. The value is the same as TemplateId.
	//
	// example:
	//
	// advanced_cot_distill
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The display order. A smaller value indicates a higher position.
	//
	// example:
	//
	// 10
	OrderNumber *int32 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// The list of pipeline stages. The order of the stages represents the execution order.
	PipelineStages []*DistillationTemplateSummaryPipelineStages `json:"PipelineStages,omitempty" xml:"PipelineStages,omitempty" type:"Repeated"`
	// The distillation template ID, which is the same as the algorithm job_type. Pass this value as TemplateId when creating a task plan.
	//
	// example:
	//
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template display name, localized based on the language specified in the request.
	//
	// example:
	//
	// Chain-of-Thought Reasoning Distillation
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The capability declaration for the second stage, in which the distilled data is used to train the student model. An empty value indicates that the template supports only the distillation stage.
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
	// The stage description, localized based on the language specified in the request.
	//
	// example:
	//
	// The teacher model generates reasoning-augmented responses for each question
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The stage identifier, which corresponds to the pipeline[].stage value in the algorithm configuration.
	//
	// example:
	//
	// cot_distill
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The stage display name, localized based on the language specified in the request.
	//
	// example:
	//
	// Generate Chain-of-Thought
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
	// The available Model Gallery Task values for the student model.
	ModelTasks []*string `json:"ModelTasks,omitempty" xml:"ModelTasks,omitempty" type:"Repeated"`
	// The list of supported training method families.
	TrainingMethods []*string `json:"TrainingMethods,omitempty" xml:"TrainingMethods,omitempty" type:"Repeated"`
	// The training type. The frontend uses this value to select the training workflow and display text.
	//
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
