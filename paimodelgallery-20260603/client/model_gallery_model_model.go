// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelGalleryModel interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *ModelGalleryModel
	GetCollection() *string
	SetCompressible(v bool) *ModelGalleryModel
	GetCompressible() *bool
	SetDeepThink(v bool) *ModelGalleryModel
	GetDeepThink() *bool
	SetDemonstrable(v bool) *ModelGalleryModel
	GetDemonstrable() *bool
	SetDeployable(v bool) *ModelGalleryModel
	GetDeployable() *bool
	SetDistillable(v bool) *ModelGalleryModel
	GetDistillable() *bool
	SetDomain(v string) *ModelGalleryModel
	GetDomain() *string
	SetEvaluable(v bool) *ModelGalleryModel
	GetEvaluable() *bool
	SetExtraInfo(v map[string]interface{}) *ModelGalleryModel
	GetExtraInfo() map[string]interface{}
	SetFunctionCall(v bool) *ModelGalleryModel
	GetFunctionCall() *bool
	SetGmtCreateTime(v string) *ModelGalleryModel
	GetGmtCreateTime() *string
	SetGmtLatestVersionModified(v string) *ModelGalleryModel
	GetGmtLatestVersionModified() *string
	SetGmtModifiedTime(v string) *ModelGalleryModel
	GetGmtModifiedTime() *string
	SetLatestVersionName(v string) *ModelGalleryModel
	GetLatestVersionName() *string
	SetModelDescription(v string) *ModelGalleryModel
	GetModelDescription() *string
	SetModelDoc(v string) *ModelGalleryModel
	GetModelDoc() *string
	SetModelId(v string) *ModelGalleryModel
	GetModelId() *string
	SetModelName(v string) *ModelGalleryModel
	GetModelName() *string
	SetModelSeries(v string) *ModelGalleryModel
	GetModelSeries() *string
	SetModelType(v string) *ModelGalleryModel
	GetModelType() *string
	SetOrderNumber(v int64) *ModelGalleryModel
	GetOrderNumber() *int64
	SetOrigin(v string) *ModelGalleryModel
	GetOrigin() *string
	SetParameterSize(v int64) *ModelGalleryModel
	GetParameterSize() *int64
	SetSearchWords(v string) *ModelGalleryModel
	GetSearchWords() *string
	SetSupportedCompressionMethods(v map[string]interface{}) *ModelGalleryModel
	GetSupportedCompressionMethods() map[string]interface{}
	SetSupportedCompressionResources(v string) *ModelGalleryModel
	GetSupportedCompressionResources() *string
	SetSupportedDistillationMethods(v map[string]interface{}) *ModelGalleryModel
	GetSupportedDistillationMethods() map[string]interface{}
	SetSupportedDistillationResources(v string) *ModelGalleryModel
	GetSupportedDistillationResources() *string
	SetSupportedEvaluationMethods(v map[string]interface{}) *ModelGalleryModel
	GetSupportedEvaluationMethods() map[string]interface{}
	SetSupportedEvaluationResources(v string) *ModelGalleryModel
	GetSupportedEvaluationResources() *string
	SetSupportedInferenceMethods(v map[string]interface{}) *ModelGalleryModel
	GetSupportedInferenceMethods() map[string]interface{}
	SetSupportedInferenceResources(v string) *ModelGalleryModel
	GetSupportedInferenceResources() *string
	SetSupportedTrainingMethods(v map[string]interface{}) *ModelGalleryModel
	GetSupportedTrainingMethods() map[string]interface{}
	SetSupportedTrainingResources(v string) *ModelGalleryModel
	GetSupportedTrainingResources() *string
	SetTags(v *ModelGalleryModelTags) *ModelGalleryModel
	GetTags() *ModelGalleryModelTags
	SetTask(v string) *ModelGalleryModel
	GetTask() *string
	SetTrainable(v bool) *ModelGalleryModel
	GetTrainable() *bool
}

type ModelGalleryModel struct {
	// example:
	//
	// QuickStart
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// example:
	//
	// true
	Compressible *bool `json:"Compressible,omitempty" xml:"Compressible,omitempty"`
	// example:
	//
	// true
	DeepThink *bool `json:"DeepThink,omitempty" xml:"DeepThink,omitempty"`
	// example:
	//
	// true
	Demonstrable *bool `json:"Demonstrable,omitempty" xml:"Demonstrable,omitempty"`
	// example:
	//
	// true
	Deployable *bool `json:"Deployable,omitempty" xml:"Deployable,omitempty"`
	// example:
	//
	// true
	Distillable *bool `json:"Distillable,omitempty" xml:"Distillable,omitempty"`
	// example:
	//
	// aigc
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
	Evaluable *bool `json:"Evaluable,omitempty" xml:"Evaluable,omitempty"`
	// example:
	//
	// {}
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// true
	FunctionCall *bool `json:"FunctionCall,omitempty" xml:"FunctionCall,omitempty"`
	// example:
	//
	// 2026-04-03T05:54:02.000Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2025-12-03T07:21:28.000Z
	GmtLatestVersionModified *string `json:"GmtLatestVersionModified,omitempty" xml:"GmtLatestVersionModified,omitempty"`
	// example:
	//
	// 2026-04-03T05:54:02.000Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 1.0.0
	LatestVersionName *string `json:"LatestVersionName,omitempty" xml:"LatestVersionName,omitempty"`
	// example:
	//
	// 大语言模型。
	ModelDescription *string `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// example:
	//
	// http://pai-quickstart-test.com
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// example:
	//
	// model-gj5mifpeol92kx619y
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// Qwen3-235B-A22B-Thinking-2507
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Qwen
	ModelSeries *string `json:"ModelSeries,omitempty" xml:"ModelSeries,omitempty"`
	// example:
	//
	// LoRA
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// 1
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// PAI
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// 1024
	ParameterSize *int64 `json:"ParameterSize,omitempty" xml:"ParameterSize,omitempty"`
	// example:
	//
	// llm
	SearchWords *string `json:"SearchWords,omitempty" xml:"SearchWords,omitempty"`
	// example:
	//
	// {
	//
	//   "Methods": [{
	//
	//     "Method": "Quantization",
	//
	//     "SubMethod": "MinMax-8Bit",
	//
	//     "Resource": "GPU"
	//
	//   }]
	//
	// }
	SupportedCompressionMethods map[string]interface{} `json:"SupportedCompressionMethods,omitempty" xml:"SupportedCompressionMethods,omitempty"`
	// example:
	//
	// GPU,GP7V
	SupportedCompressionResources *string `json:"SupportedCompressionResources,omitempty" xml:"SupportedCompressionResources,omitempty"`
	// example:
	//
	// {
	//
	//   "Methods": [{
	//
	//     "Method": "EasyDistill",
	//
	//     "SubMethod": "DataSynthesis",
	//
	//     "Resource": "GPU"
	//
	//   }]
	//
	// }
	SupportedDistillationMethods map[string]interface{} `json:"SupportedDistillationMethods,omitempty" xml:"SupportedDistillationMethods,omitempty"`
	// example:
	//
	// GPU,GP7V
	SupportedDistillationResources *string `json:"SupportedDistillationResources,omitempty" xml:"SupportedDistillationResources,omitempty"`
	// example:
	//
	// {
	//
	//   "Methods": [{
	//
	//     "Method": "Single-Node-Standard",
	//
	//     "Resource": "GPU"
	//
	//   }]
	//
	// }
	SupportedEvaluationMethods map[string]interface{} `json:"SupportedEvaluationMethods,omitempty" xml:"SupportedEvaluationMethods,omitempty"`
	// example:
	//
	// GPU,GP7V
	SupportedEvaluationResources *string `json:"SupportedEvaluationResources,omitempty" xml:"SupportedEvaluationResources,omitempty"`
	// example:
	//
	// {
	//
	//   "Methods": [{
	//
	//     "framework": "blade",
	//
	//     "scenario": "nvidia-standard",
	//
	//     "Resource": "GPU"
	//
	//   }]
	//
	// }
	SupportedInferenceMethods map[string]interface{} `json:"SupportedInferenceMethods,omitempty" xml:"SupportedInferenceMethods,omitempty"`
	// example:
	//
	// GPU,GP7V
	SupportedInferenceResources *string `json:"SupportedInferenceResources,omitempty" xml:"SupportedInferenceResources,omitempty"`
	// example:
	//
	// {
	//
	//   "Methods": [{
	//
	//     "Method": "SFT",
	//
	//     "SubMethod": "LoRA_LLM",
	//
	//     "Resource": "GPU"
	//
	//   }]
	//
	// }
	SupportedTrainingMethods map[string]interface{} `json:"SupportedTrainingMethods,omitempty" xml:"SupportedTrainingMethods,omitempty"`
	// example:
	//
	// GPU,GP7V
	SupportedTrainingResources *string                `json:"SupportedTrainingResources,omitempty" xml:"SupportedTrainingResources,omitempty"`
	Tags                       *ModelGalleryModelTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// large-language-model
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// example:
	//
	// true
	Trainable *bool `json:"Trainable,omitempty" xml:"Trainable,omitempty"`
}

func (s ModelGalleryModel) String() string {
	return dara.Prettify(s)
}

func (s ModelGalleryModel) GoString() string {
	return s.String()
}

func (s *ModelGalleryModel) GetCollection() *string {
	return s.Collection
}

func (s *ModelGalleryModel) GetCompressible() *bool {
	return s.Compressible
}

func (s *ModelGalleryModel) GetDeepThink() *bool {
	return s.DeepThink
}

func (s *ModelGalleryModel) GetDemonstrable() *bool {
	return s.Demonstrable
}

func (s *ModelGalleryModel) GetDeployable() *bool {
	return s.Deployable
}

func (s *ModelGalleryModel) GetDistillable() *bool {
	return s.Distillable
}

func (s *ModelGalleryModel) GetDomain() *string {
	return s.Domain
}

func (s *ModelGalleryModel) GetEvaluable() *bool {
	return s.Evaluable
}

func (s *ModelGalleryModel) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *ModelGalleryModel) GetFunctionCall() *bool {
	return s.FunctionCall
}

func (s *ModelGalleryModel) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ModelGalleryModel) GetGmtLatestVersionModified() *string {
	return s.GmtLatestVersionModified
}

func (s *ModelGalleryModel) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ModelGalleryModel) GetLatestVersionName() *string {
	return s.LatestVersionName
}

func (s *ModelGalleryModel) GetModelDescription() *string {
	return s.ModelDescription
}

func (s *ModelGalleryModel) GetModelDoc() *string {
	return s.ModelDoc
}

func (s *ModelGalleryModel) GetModelId() *string {
	return s.ModelId
}

func (s *ModelGalleryModel) GetModelName() *string {
	return s.ModelName
}

func (s *ModelGalleryModel) GetModelSeries() *string {
	return s.ModelSeries
}

func (s *ModelGalleryModel) GetModelType() *string {
	return s.ModelType
}

func (s *ModelGalleryModel) GetOrderNumber() *int64 {
	return s.OrderNumber
}

func (s *ModelGalleryModel) GetOrigin() *string {
	return s.Origin
}

func (s *ModelGalleryModel) GetParameterSize() *int64 {
	return s.ParameterSize
}

func (s *ModelGalleryModel) GetSearchWords() *string {
	return s.SearchWords
}

func (s *ModelGalleryModel) GetSupportedCompressionMethods() map[string]interface{} {
	return s.SupportedCompressionMethods
}

func (s *ModelGalleryModel) GetSupportedCompressionResources() *string {
	return s.SupportedCompressionResources
}

func (s *ModelGalleryModel) GetSupportedDistillationMethods() map[string]interface{} {
	return s.SupportedDistillationMethods
}

func (s *ModelGalleryModel) GetSupportedDistillationResources() *string {
	return s.SupportedDistillationResources
}

func (s *ModelGalleryModel) GetSupportedEvaluationMethods() map[string]interface{} {
	return s.SupportedEvaluationMethods
}

func (s *ModelGalleryModel) GetSupportedEvaluationResources() *string {
	return s.SupportedEvaluationResources
}

func (s *ModelGalleryModel) GetSupportedInferenceMethods() map[string]interface{} {
	return s.SupportedInferenceMethods
}

func (s *ModelGalleryModel) GetSupportedInferenceResources() *string {
	return s.SupportedInferenceResources
}

func (s *ModelGalleryModel) GetSupportedTrainingMethods() map[string]interface{} {
	return s.SupportedTrainingMethods
}

func (s *ModelGalleryModel) GetSupportedTrainingResources() *string {
	return s.SupportedTrainingResources
}

func (s *ModelGalleryModel) GetTags() *ModelGalleryModelTags {
	return s.Tags
}

func (s *ModelGalleryModel) GetTask() *string {
	return s.Task
}

func (s *ModelGalleryModel) GetTrainable() *bool {
	return s.Trainable
}

func (s *ModelGalleryModel) SetCollection(v string) *ModelGalleryModel {
	s.Collection = &v
	return s
}

func (s *ModelGalleryModel) SetCompressible(v bool) *ModelGalleryModel {
	s.Compressible = &v
	return s
}

func (s *ModelGalleryModel) SetDeepThink(v bool) *ModelGalleryModel {
	s.DeepThink = &v
	return s
}

func (s *ModelGalleryModel) SetDemonstrable(v bool) *ModelGalleryModel {
	s.Demonstrable = &v
	return s
}

func (s *ModelGalleryModel) SetDeployable(v bool) *ModelGalleryModel {
	s.Deployable = &v
	return s
}

func (s *ModelGalleryModel) SetDistillable(v bool) *ModelGalleryModel {
	s.Distillable = &v
	return s
}

func (s *ModelGalleryModel) SetDomain(v string) *ModelGalleryModel {
	s.Domain = &v
	return s
}

func (s *ModelGalleryModel) SetEvaluable(v bool) *ModelGalleryModel {
	s.Evaluable = &v
	return s
}

func (s *ModelGalleryModel) SetExtraInfo(v map[string]interface{}) *ModelGalleryModel {
	s.ExtraInfo = v
	return s
}

func (s *ModelGalleryModel) SetFunctionCall(v bool) *ModelGalleryModel {
	s.FunctionCall = &v
	return s
}

func (s *ModelGalleryModel) SetGmtCreateTime(v string) *ModelGalleryModel {
	s.GmtCreateTime = &v
	return s
}

func (s *ModelGalleryModel) SetGmtLatestVersionModified(v string) *ModelGalleryModel {
	s.GmtLatestVersionModified = &v
	return s
}

func (s *ModelGalleryModel) SetGmtModifiedTime(v string) *ModelGalleryModel {
	s.GmtModifiedTime = &v
	return s
}

func (s *ModelGalleryModel) SetLatestVersionName(v string) *ModelGalleryModel {
	s.LatestVersionName = &v
	return s
}

func (s *ModelGalleryModel) SetModelDescription(v string) *ModelGalleryModel {
	s.ModelDescription = &v
	return s
}

func (s *ModelGalleryModel) SetModelDoc(v string) *ModelGalleryModel {
	s.ModelDoc = &v
	return s
}

func (s *ModelGalleryModel) SetModelId(v string) *ModelGalleryModel {
	s.ModelId = &v
	return s
}

func (s *ModelGalleryModel) SetModelName(v string) *ModelGalleryModel {
	s.ModelName = &v
	return s
}

func (s *ModelGalleryModel) SetModelSeries(v string) *ModelGalleryModel {
	s.ModelSeries = &v
	return s
}

func (s *ModelGalleryModel) SetModelType(v string) *ModelGalleryModel {
	s.ModelType = &v
	return s
}

func (s *ModelGalleryModel) SetOrderNumber(v int64) *ModelGalleryModel {
	s.OrderNumber = &v
	return s
}

func (s *ModelGalleryModel) SetOrigin(v string) *ModelGalleryModel {
	s.Origin = &v
	return s
}

func (s *ModelGalleryModel) SetParameterSize(v int64) *ModelGalleryModel {
	s.ParameterSize = &v
	return s
}

func (s *ModelGalleryModel) SetSearchWords(v string) *ModelGalleryModel {
	s.SearchWords = &v
	return s
}

func (s *ModelGalleryModel) SetSupportedCompressionMethods(v map[string]interface{}) *ModelGalleryModel {
	s.SupportedCompressionMethods = v
	return s
}

func (s *ModelGalleryModel) SetSupportedCompressionResources(v string) *ModelGalleryModel {
	s.SupportedCompressionResources = &v
	return s
}

func (s *ModelGalleryModel) SetSupportedDistillationMethods(v map[string]interface{}) *ModelGalleryModel {
	s.SupportedDistillationMethods = v
	return s
}

func (s *ModelGalleryModel) SetSupportedDistillationResources(v string) *ModelGalleryModel {
	s.SupportedDistillationResources = &v
	return s
}

func (s *ModelGalleryModel) SetSupportedEvaluationMethods(v map[string]interface{}) *ModelGalleryModel {
	s.SupportedEvaluationMethods = v
	return s
}

func (s *ModelGalleryModel) SetSupportedEvaluationResources(v string) *ModelGalleryModel {
	s.SupportedEvaluationResources = &v
	return s
}

func (s *ModelGalleryModel) SetSupportedInferenceMethods(v map[string]interface{}) *ModelGalleryModel {
	s.SupportedInferenceMethods = v
	return s
}

func (s *ModelGalleryModel) SetSupportedInferenceResources(v string) *ModelGalleryModel {
	s.SupportedInferenceResources = &v
	return s
}

func (s *ModelGalleryModel) SetSupportedTrainingMethods(v map[string]interface{}) *ModelGalleryModel {
	s.SupportedTrainingMethods = v
	return s
}

func (s *ModelGalleryModel) SetSupportedTrainingResources(v string) *ModelGalleryModel {
	s.SupportedTrainingResources = &v
	return s
}

func (s *ModelGalleryModel) SetTags(v *ModelGalleryModelTags) *ModelGalleryModel {
	s.Tags = v
	return s
}

func (s *ModelGalleryModel) SetTask(v string) *ModelGalleryModel {
	s.Task = &v
	return s
}

func (s *ModelGalleryModel) SetTrainable(v bool) *ModelGalleryModel {
	s.Trainable = &v
	return s
}

func (s *ModelGalleryModel) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelGalleryModelTags struct {
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModelGalleryModelTags) String() string {
	return dara.Prettify(s)
}

func (s ModelGalleryModelTags) GoString() string {
	return s.String()
}

func (s *ModelGalleryModelTags) GetKey() *string {
	return s.Key
}

func (s *ModelGalleryModelTags) GetValue() *string {
	return s.Value
}

func (s *ModelGalleryModelTags) SetKey(v string) *ModelGalleryModelTags {
	s.Key = &v
	return s
}

func (s *ModelGalleryModelTags) SetValue(v string) *ModelGalleryModelTags {
	s.Value = &v
	return s
}

func (s *ModelGalleryModelTags) Validate() error {
	return dara.Validate(s)
}
