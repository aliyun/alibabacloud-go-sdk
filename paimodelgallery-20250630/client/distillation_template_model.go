// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistillationTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmName(v string) *DistillationTemplate
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *DistillationTemplate
	GetAlgorithmProvider() *string
	SetAlgorithmVersion(v string) *DistillationTemplate
	GetAlgorithmVersion() *string
	SetCapabilityTags(v []*string) *DistillationTemplate
	GetCapabilityTags() []*string
	SetCategory(v string) *DistillationTemplate
	GetCategory() *string
	SetDefaultConfig(v string) *DistillationTemplate
	GetDefaultConfig() *string
	SetDescription(v string) *DistillationTemplate
	GetDescription() *string
	SetInputDatasetMustBeDirectory(v bool) *DistillationTemplate
	GetInputDatasetMustBeDirectory() *bool
	SetInputExampleUri(v string) *DistillationTemplate
	GetInputExampleUri() *string
	SetJobType(v string) *DistillationTemplate
	GetJobType() *string
	SetModelSlots(v []*DistillationTemplateModelSlots) *DistillationTemplate
	GetModelSlots() []*DistillationTemplateModelSlots
	SetOrderNumber(v int32) *DistillationTemplate
	GetOrderNumber() *int32
	SetPipelineStages(v []*DistillationTemplatePipelineStages) *DistillationTemplate
	GetPipelineStages() []*DistillationTemplatePipelineStages
	SetPresetConfig(v []*DistillationTemplatePresetConfig) *DistillationTemplate
	GetPresetConfig() []*DistillationTemplatePresetConfig
	SetProcessorExampleUri(v string) *DistillationTemplate
	GetProcessorExampleUri() *string
	SetTemplateId(v string) *DistillationTemplate
	GetTemplateId() *string
	SetTemplateKind(v string) *DistillationTemplate
	GetTemplateKind() *string
	SetTemplateName(v string) *DistillationTemplate
	GetTemplateName() *string
	SetTrainingOptions(v []*DistillationTemplateTrainingOptions) *DistillationTemplate
	GetTrainingOptions() []*DistillationTemplateTrainingOptions
}

type DistillationTemplate struct {
	// example:
	//
	// easydistill
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// v2.0.0
	AlgorithmVersion *string   `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	CapabilityTags   []*string `json:"CapabilityTags,omitempty" xml:"CapabilityTags,omitempty" type:"Repeated"`
	// example:
	//
	// reasoning
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// job_type: advanced_cot_distill
	DefaultConfig *string `json:"DefaultConfig,omitempty" xml:"DefaultConfig,omitempty"`
	// example:
	//
	// 面向数学、逻辑、代码等需要多步推理的场景，产出 SFT 数据集。
	Description                 *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InputDatasetMustBeDirectory *bool   `json:"InputDatasetMustBeDirectory,omitempty" xml:"InputDatasetMustBeDirectory,omitempty"`
	// example:
	//
	// oss://pai-quickstart-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/public_datasets/distillation_examples/advanced_cot_distill/input.jsonl
	InputExampleUri *string `json:"InputExampleUri,omitempty" xml:"InputExampleUri,omitempty"`
	// example:
	//
	// advanced_cot_distill
	JobType    *string                           `json:"JobType,omitempty" xml:"JobType,omitempty"`
	ModelSlots []*DistillationTemplateModelSlots `json:"ModelSlots,omitempty" xml:"ModelSlots,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	OrderNumber    *int32                                `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	PipelineStages []*DistillationTemplatePipelineStages `json:"PipelineStages,omitempty" xml:"PipelineStages,omitempty" type:"Repeated"`
	PresetConfig   []*DistillationTemplatePresetConfig   `json:"PresetConfig,omitempty" xml:"PresetConfig,omitempty" type:"Repeated"`
	// example:
	//
	// oss://pai-quickstart-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/public_datasets/distillation_examples/custom_distill/processor_example.py
	ProcessorExampleUri *string `json:"ProcessorExampleUri,omitempty" xml:"ProcessorExampleUri,omitempty"`
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
	TemplateName    *string                                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TrainingOptions []*DistillationTemplateTrainingOptions `json:"TrainingOptions,omitempty" xml:"TrainingOptions,omitempty" type:"Repeated"`
}

func (s DistillationTemplate) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplate) GoString() string {
	return s.String()
}

func (s *DistillationTemplate) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *DistillationTemplate) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *DistillationTemplate) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *DistillationTemplate) GetCapabilityTags() []*string {
	return s.CapabilityTags
}

func (s *DistillationTemplate) GetCategory() *string {
	return s.Category
}

func (s *DistillationTemplate) GetDefaultConfig() *string {
	return s.DefaultConfig
}

func (s *DistillationTemplate) GetDescription() *string {
	return s.Description
}

func (s *DistillationTemplate) GetInputDatasetMustBeDirectory() *bool {
	return s.InputDatasetMustBeDirectory
}

func (s *DistillationTemplate) GetInputExampleUri() *string {
	return s.InputExampleUri
}

func (s *DistillationTemplate) GetJobType() *string {
	return s.JobType
}

func (s *DistillationTemplate) GetModelSlots() []*DistillationTemplateModelSlots {
	return s.ModelSlots
}

func (s *DistillationTemplate) GetOrderNumber() *int32 {
	return s.OrderNumber
}

func (s *DistillationTemplate) GetPipelineStages() []*DistillationTemplatePipelineStages {
	return s.PipelineStages
}

func (s *DistillationTemplate) GetPresetConfig() []*DistillationTemplatePresetConfig {
	return s.PresetConfig
}

func (s *DistillationTemplate) GetProcessorExampleUri() *string {
	return s.ProcessorExampleUri
}

func (s *DistillationTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DistillationTemplate) GetTemplateKind() *string {
	return s.TemplateKind
}

func (s *DistillationTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DistillationTemplate) GetTrainingOptions() []*DistillationTemplateTrainingOptions {
	return s.TrainingOptions
}

func (s *DistillationTemplate) SetAlgorithmName(v string) *DistillationTemplate {
	s.AlgorithmName = &v
	return s
}

func (s *DistillationTemplate) SetAlgorithmProvider(v string) *DistillationTemplate {
	s.AlgorithmProvider = &v
	return s
}

func (s *DistillationTemplate) SetAlgorithmVersion(v string) *DistillationTemplate {
	s.AlgorithmVersion = &v
	return s
}

func (s *DistillationTemplate) SetCapabilityTags(v []*string) *DistillationTemplate {
	s.CapabilityTags = v
	return s
}

func (s *DistillationTemplate) SetCategory(v string) *DistillationTemplate {
	s.Category = &v
	return s
}

func (s *DistillationTemplate) SetDefaultConfig(v string) *DistillationTemplate {
	s.DefaultConfig = &v
	return s
}

func (s *DistillationTemplate) SetDescription(v string) *DistillationTemplate {
	s.Description = &v
	return s
}

func (s *DistillationTemplate) SetInputDatasetMustBeDirectory(v bool) *DistillationTemplate {
	s.InputDatasetMustBeDirectory = &v
	return s
}

func (s *DistillationTemplate) SetInputExampleUri(v string) *DistillationTemplate {
	s.InputExampleUri = &v
	return s
}

func (s *DistillationTemplate) SetJobType(v string) *DistillationTemplate {
	s.JobType = &v
	return s
}

func (s *DistillationTemplate) SetModelSlots(v []*DistillationTemplateModelSlots) *DistillationTemplate {
	s.ModelSlots = v
	return s
}

func (s *DistillationTemplate) SetOrderNumber(v int32) *DistillationTemplate {
	s.OrderNumber = &v
	return s
}

func (s *DistillationTemplate) SetPipelineStages(v []*DistillationTemplatePipelineStages) *DistillationTemplate {
	s.PipelineStages = v
	return s
}

func (s *DistillationTemplate) SetPresetConfig(v []*DistillationTemplatePresetConfig) *DistillationTemplate {
	s.PresetConfig = v
	return s
}

func (s *DistillationTemplate) SetProcessorExampleUri(v string) *DistillationTemplate {
	s.ProcessorExampleUri = &v
	return s
}

func (s *DistillationTemplate) SetTemplateId(v string) *DistillationTemplate {
	s.TemplateId = &v
	return s
}

func (s *DistillationTemplate) SetTemplateKind(v string) *DistillationTemplate {
	s.TemplateKind = &v
	return s
}

func (s *DistillationTemplate) SetTemplateName(v string) *DistillationTemplate {
	s.TemplateName = &v
	return s
}

func (s *DistillationTemplate) SetTrainingOptions(v []*DistillationTemplateTrainingOptions) *DistillationTemplate {
	s.TrainingOptions = v
	return s
}

func (s *DistillationTemplate) Validate() error {
	if s.ModelSlots != nil {
		for _, item := range s.ModelSlots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PipelineStages != nil {
		for _, item := range s.PipelineStages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PresetConfig != nil {
		for _, item := range s.PresetConfig {
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

type DistillationTemplateModelSlots struct {
	Backends []*DistillationTemplateModelSlotsBackends `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// example:
	//
	// 用于生成蒸馏数据的教师模型，建议选择推理能力强的模型。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// backend
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 教师模型
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DistillationTemplateModelSlots) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplateModelSlots) GoString() string {
	return s.String()
}

func (s *DistillationTemplateModelSlots) GetBackends() []*DistillationTemplateModelSlotsBackends {
	return s.Backends
}

func (s *DistillationTemplateModelSlots) GetDescription() *string {
	return s.Description
}

func (s *DistillationTemplateModelSlots) GetKey() *string {
	return s.Key
}

func (s *DistillationTemplateModelSlots) GetName() *string {
	return s.Name
}

func (s *DistillationTemplateModelSlots) GetRequired() *bool {
	return s.Required
}

func (s *DistillationTemplateModelSlots) SetBackends(v []*DistillationTemplateModelSlotsBackends) *DistillationTemplateModelSlots {
	s.Backends = v
	return s
}

func (s *DistillationTemplateModelSlots) SetDescription(v string) *DistillationTemplateModelSlots {
	s.Description = &v
	return s
}

func (s *DistillationTemplateModelSlots) SetKey(v string) *DistillationTemplateModelSlots {
	s.Key = &v
	return s
}

func (s *DistillationTemplateModelSlots) SetName(v string) *DistillationTemplateModelSlots {
	s.Name = &v
	return s
}

func (s *DistillationTemplateModelSlots) SetRequired(v bool) *DistillationTemplateModelSlots {
	s.Required = &v
	return s
}

func (s *DistillationTemplateModelSlots) Validate() error {
	if s.Backends != nil {
		for _, item := range s.Backends {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DistillationTemplateModelSlotsBackends struct {
	// example:
	//
	// distillation
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// pai_token
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DistillationTemplateModelSlotsBackends) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplateModelSlotsBackends) GoString() string {
	return s.String()
}

func (s *DistillationTemplateModelSlotsBackends) GetChannel() *string {
	return s.Channel
}

func (s *DistillationTemplateModelSlotsBackends) GetType() *string {
	return s.Type
}

func (s *DistillationTemplateModelSlotsBackends) SetChannel(v string) *DistillationTemplateModelSlotsBackends {
	s.Channel = &v
	return s
}

func (s *DistillationTemplateModelSlotsBackends) SetType(v string) *DistillationTemplateModelSlotsBackends {
	s.Type = &v
	return s
}

func (s *DistillationTemplateModelSlotsBackends) Validate() error {
	return dara.Validate(s)
}

type DistillationTemplatePipelineStages struct {
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

func (s DistillationTemplatePipelineStages) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplatePipelineStages) GoString() string {
	return s.String()
}

func (s *DistillationTemplatePipelineStages) GetDescription() *string {
	return s.Description
}

func (s *DistillationTemplatePipelineStages) GetKey() *string {
	return s.Key
}

func (s *DistillationTemplatePipelineStages) GetName() *string {
	return s.Name
}

func (s *DistillationTemplatePipelineStages) SetDescription(v string) *DistillationTemplatePipelineStages {
	s.Description = &v
	return s
}

func (s *DistillationTemplatePipelineStages) SetKey(v string) *DistillationTemplatePipelineStages {
	s.Key = &v
	return s
}

func (s *DistillationTemplatePipelineStages) SetName(v string) *DistillationTemplatePipelineStages {
	s.Name = &v
	return s
}

func (s *DistillationTemplatePipelineStages) Validate() error {
	return dara.Validate(s)
}

type DistillationTemplatePresetConfig struct {
	// example:
	//
	// 任务类型
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// advanced_cot_distill
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DistillationTemplatePresetConfig) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplatePresetConfig) GoString() string {
	return s.String()
}

func (s *DistillationTemplatePresetConfig) GetLabel() *string {
	return s.Label
}

func (s *DistillationTemplatePresetConfig) GetValue() *string {
	return s.Value
}

func (s *DistillationTemplatePresetConfig) SetLabel(v string) *DistillationTemplatePresetConfig {
	s.Label = &v
	return s
}

func (s *DistillationTemplatePresetConfig) SetValue(v string) *DistillationTemplatePresetConfig {
	s.Value = &v
	return s
}

func (s *DistillationTemplatePresetConfig) Validate() error {
	return dara.Validate(s)
}

type DistillationTemplateTrainingOptions struct {
	ModelTasks      []*string `json:"ModelTasks,omitempty" xml:"ModelTasks,omitempty" type:"Repeated"`
	TrainingMethods []*string `json:"TrainingMethods,omitempty" xml:"TrainingMethods,omitempty" type:"Repeated"`
	// example:
	//
	// sft
	TrainingType *string `json:"TrainingType,omitempty" xml:"TrainingType,omitempty"`
}

func (s DistillationTemplateTrainingOptions) String() string {
	return dara.Prettify(s)
}

func (s DistillationTemplateTrainingOptions) GoString() string {
	return s.String()
}

func (s *DistillationTemplateTrainingOptions) GetModelTasks() []*string {
	return s.ModelTasks
}

func (s *DistillationTemplateTrainingOptions) GetTrainingMethods() []*string {
	return s.TrainingMethods
}

func (s *DistillationTemplateTrainingOptions) GetTrainingType() *string {
	return s.TrainingType
}

func (s *DistillationTemplateTrainingOptions) SetModelTasks(v []*string) *DistillationTemplateTrainingOptions {
	s.ModelTasks = v
	return s
}

func (s *DistillationTemplateTrainingOptions) SetTrainingMethods(v []*string) *DistillationTemplateTrainingOptions {
	s.TrainingMethods = v
	return s
}

func (s *DistillationTemplateTrainingOptions) SetTrainingType(v string) *DistillationTemplateTrainingOptions {
	s.TrainingType = &v
	return s
}

func (s *DistillationTemplateTrainingOptions) Validate() error {
	return dara.Validate(s)
}
