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
	SetTemplateId(v string) *DistillationTemplate
	GetTemplateId() *string
	SetTemplateName(v string) *DistillationTemplate
	GetTemplateName() *string
	SetTrainingOptions(v []*DistillationTemplateTrainingOptions) *DistillationTemplate
	GetTrainingOptions() []*DistillationTemplateTrainingOptions
}

type DistillationTemplate struct {
	// The algorithm name.
	//
	// example:
	//
	// easydistill
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// The algorithm provider.
	//
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// The algorithm version.
	//
	// example:
	//
	// v2.0.0
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	// The list of capability tags used for displaying scenario cards.
	CapabilityTags []*string `json:"CapabilityTags,omitempty" xml:"CapabilityTags,omitempty" type:"Repeated"`
	// The template category. The frontend uses this value to filter scenario cards.
	//
	// example:
	//
	// reasoning
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The raw YAML content of the EasyDistill default configurations. The frontend uses this content for rendering the configuration form and supports recovering to default configurations. The model and credential fields are intentionally left empty and are populated by the user in the form upon commit.
	//
	// example:
	//
	// job_type: advanced_cot_distill
	DefaultConfig *string `json:"DefaultConfig,omitempty" xml:"DefaultConfig,omitempty"`
	// The template description, localized based on the requested language. The description specifies the applicable scenarios and outputs.
	//
	// example:
	//
	// Designed for scenarios that require multi-step reasoning such as math, logic, and code. Produces an SFT dataset
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the input data must be an entire directory. If this parameter is set to true, only a directory can be selected on the form, not a single file. If this parameter is absent or set to false, either a file or a directory can be selected. This value is true when seed files reference other files in the same directory by relative path.
	InputDatasetMustBeDirectory *bool `json:"InputDatasetMustBeDirectory,omitempty" xml:"InputDatasetMustBeDirectory,omitempty"`
	// The OSS address of the sample input data, rendered based on the region. Users can download the sample and prepare their own data in the same format. An empty value indicates that the template does not provide a sample.
	//
	// example:
	//
	// oss://pai-quickstart-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/public_datasets/distillation_examples/advanced_cot_distill/input.jsonl
	InputExampleUri *string `json:"InputExampleUri,omitempty" xml:"InputExampleUri,omitempty"`
	// The algorithm job type. The value is the same as TemplateId.
	//
	// example:
	//
	// advanced_cot_distill
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The list of model slots that require user selection. The frontend uses this list to render the model selection form.
	ModelSlots []*DistillationTemplateModelSlots `json:"ModelSlots,omitempty" xml:"ModelSlots,omitempty" type:"Repeated"`
	// The display order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 10
	OrderNumber *int32 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// The list of pipeline stages. The order of the stages represents the execution order.
	PipelineStages []*DistillationTemplatePipelineStages `json:"PipelineStages,omitempty" xml:"PipelineStages,omitempty" type:"Repeated"`
	// The content of the preset configuration card, displayed in order to show the key default configurations of the template.
	PresetConfig []*DistillationTemplatePresetConfig `json:"PresetConfig,omitempty" xml:"PresetConfig,omitempty" type:"Repeated"`
	// The distillation template ID, which is the same as the algorithm job_type. Pass this value as TemplateId when creating a task plan.
	//
	// example:
	//
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template display name, localized based on the requested language.
	//
	// example:
	//
	// Chain-of-thought reasoning distillation
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The capability declaration for the second stage (training the student model with the distilled data). An empty value indicates that the template supports only the distillation stage.
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

func (s *DistillationTemplate) GetTemplateId() *string {
	return s.TemplateId
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

func (s *DistillationTemplate) SetTemplateId(v string) *DistillationTemplate {
	s.TemplateId = &v
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
	// The list of model access methods supported by this slot.
	Backends []*DistillationTemplateModelSlotsBackends `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The slot description, localized based on the requested language.
	//
	// example:
	//
	// The teacher model used to generate distillation data. We recommend selecting a model with strong reasoning capabilities
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The slot identifier, which corresponds to the backend section name in the submitted configuration.
	//
	// example:
	//
	// backend
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The slot display name, localized based on the requested language.
	//
	// example:
	//
	// Teacher model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether this slot is required. If this parameter is set to false, the user can skip the selection, and the algorithm falls back to other slots.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
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
	// The channel name of the PAI-Token gateway. The frontend uses this value to retrieve the list of available models for the channel. This value must be passed back as-is upon submission. This parameter is returned only when Type is pai_token.
	//
	// example:
	//
	// distillation
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The model access method. pai_token indicates the PAI-Token gateway, where the user selects from the list of available models for the channel. pai_eas indicates the user\\"s own PAI-EAS service instance, which requires the service address and token.
	//
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
	// The stage description, localized based on the requested language.
	//
	// example:
	//
	// The teacher model generates responses with reasoning processes for each question
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The stage identifier, which corresponds to the value of pipeline[].stage in the algorithm configuration.
	//
	// example:
	//
	// cot_distill
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The stage display name, localized based on the requested language.
	//
	// example:
	//
	// Generate chain of thought
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
	// The configuration item name, localized based on the requested language. Names are matched by position across languages, so the same row can have different names in different languages.
	//
	// example:
	//
	// Task type
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The configuration item value, localized based on the requested language.
	//
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
	// The range of Model Gallery tasks available for the student model.
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
