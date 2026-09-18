// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelGalleryModelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollections(v string) *ListModelGalleryModelsShrinkRequest
	GetCollections() *string
	SetCompressible(v bool) *ListModelGalleryModelsShrinkRequest
	GetCompressible() *bool
	SetConditionsShrink(v string) *ListModelGalleryModelsShrinkRequest
	GetConditionsShrink() *string
	SetDeepThink(v bool) *ListModelGalleryModelsShrinkRequest
	GetDeepThink() *bool
	SetDemonstrable(v bool) *ListModelGalleryModelsShrinkRequest
	GetDemonstrable() *bool
	SetDeployable(v bool) *ListModelGalleryModelsShrinkRequest
	GetDeployable() *bool
	SetDistillable(v bool) *ListModelGalleryModelsShrinkRequest
	GetDistillable() *bool
	SetDomain(v string) *ListModelGalleryModelsShrinkRequest
	GetDomain() *string
	SetEvaluable(v bool) *ListModelGalleryModelsShrinkRequest
	GetEvaluable() *bool
	SetFunctionCall(v bool) *ListModelGalleryModelsShrinkRequest
	GetFunctionCall() *bool
	SetModelName(v string) *ListModelGalleryModelsShrinkRequest
	GetModelName() *string
	SetModelSeries(v string) *ListModelGalleryModelsShrinkRequest
	GetModelSeries() *string
	SetModelType(v string) *ListModelGalleryModelsShrinkRequest
	GetModelType() *string
	SetOrder(v string) *ListModelGalleryModelsShrinkRequest
	GetOrder() *string
	SetOrigin(v string) *ListModelGalleryModelsShrinkRequest
	GetOrigin() *string
	SetPageNumber(v int32) *ListModelGalleryModelsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelGalleryModelsShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *ListModelGalleryModelsShrinkRequest
	GetQuery() *string
	SetSortBy(v string) *ListModelGalleryModelsShrinkRequest
	GetSortBy() *string
	SetSupportedCompressionResource(v string) *ListModelGalleryModelsShrinkRequest
	GetSupportedCompressionResource() *string
	SetSupportedDistillationResource(v string) *ListModelGalleryModelsShrinkRequest
	GetSupportedDistillationResource() *string
	SetSupportedEvaluationResource(v string) *ListModelGalleryModelsShrinkRequest
	GetSupportedEvaluationResource() *string
	SetSupportedInferenceResource(v string) *ListModelGalleryModelsShrinkRequest
	GetSupportedInferenceResource() *string
	SetSupportedTrainingResource(v string) *ListModelGalleryModelsShrinkRequest
	GetSupportedTrainingResource() *string
	SetTagShrink(v string) *ListModelGalleryModelsShrinkRequest
	GetTagShrink() *string
	SetTask(v string) *ListModelGalleryModelsShrinkRequest
	GetTask() *string
	SetTrainable(v bool) *ListModelGalleryModelsShrinkRequest
	GetTrainable() *bool
}

type ListModelGalleryModelsShrinkRequest struct {
	// The collection to which the model belongs. The collection for ModelGallery models is QuickStart.
	//
	// example:
	//
	// QuickStart
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	// Specifies whether model compression is supported.
	Compressible *bool `json:"Compressible,omitempty" xml:"Compressible,omitempty"`
	// The list of conditions.
	ConditionsShrink *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// Specifies whether deep thinking is supported.
	DeepThink *bool `json:"DeepThink,omitempty" xml:"DeepThink,omitempty"`
	// Specifies whether online experience is supported.
	Demonstrable *bool `json:"Demonstrable,omitempty" xml:"Demonstrable,omitempty"`
	// Specifies whether deployment is supported.
	Deployable *bool `json:"Deployable,omitempty" xml:"Deployable,omitempty"`
	// Specifies whether distillation is supported.
	Distillable *bool `json:"Distillable,omitempty" xml:"Distillable,omitempty"`
	// The domain used to filter the model list. For example, aigc (generative AI), nlp (natural language processing), or cv (computer vision).
	//
	// example:
	//
	// aigc
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether evaluation is supported.
	Evaluable *bool `json:"Evaluable,omitempty" xml:"Evaluable,omitempty"`
	// Specifies whether FunctionCall is supported.
	FunctionCall *bool `json:"FunctionCall,omitempty" xml:"FunctionCall,omitempty"`
	// The model name. By default, fuzzy match is used to filter the model list. Enclose the name in double quotation marks for exact match. For example, "DeepSeek-V3.2" exactly matches the model DeepSeek-V3.2.
	//
	// example:
	//
	// Qwen3-235B-A22B-Thinking-2507
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model series. For example, PAI Optimized Model, Qwen, Wan-AI, GLM, Baichuan, DeepSeek, Kimi, MiniMax, Yi, InternLM, InternVL, OpenAI, NVIDIA, Gemma, Phi, dots_vlm, Llama, Mistral, Stable Diffusion, FLUX, Byte Dance, StepFun AI, ERNIE Bot, Tencent Hunyuan, or YOLO.
	//
	// example:
	//
	// Qwen
	ModelSeries *string `json:"ModelSeries,omitempty" xml:"ModelSeries,omitempty"`
	// The model type.
	//
	// example:
	//
	// Endpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The sort order for the specified sort field in a paged query. Default value: ASC.
	//
	// Valid values:
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The model source used to filter models that belong to a specific community or organization. For example, ModelScope, PAI, or NIM.
	//
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The page number of the model list. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of models to display on each page in a paged query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition. Fuzzy match is performed across multiple fields such as ModelName, Domain, and Task.
	//
	// example:
	//
	// Qwen
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The sort field for a paged query. Currently, the GmtCreateTime field is used for sorting. Valid values:
	//
	// - GmtCreateTime: the model creation time.
	//
	// - GmtLatestVersionModified: the time when the latest model version was updated.
	//
	// - OrderNumber: the ordinal number.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The supported compression resources, such as CPU, GPU, or GP7V. For example, if you specify GP7V, only models that support compression on GP7V resources are returned.
	//
	// example:
	//
	// GP7V
	SupportedCompressionResource *string `json:"SupportedCompressionResource,omitempty" xml:"SupportedCompressionResource,omitempty"`
	// The supported distillation resources, such as CPU, GPU, or GP7V. For example, if you specify GP7V, only models that support distillation on GP7V resources are returned.
	//
	// example:
	//
	// GP7V
	SupportedDistillationResource *string `json:"SupportedDistillationResource,omitempty" xml:"SupportedDistillationResource,omitempty"`
	// The supported evaluation resources, such as CPU, GPU, or GP7V. For example, if you specify GP7V, only models that support evaluation on GP7V resources are returned.
	//
	// example:
	//
	// GP7V
	SupportedEvaluationResource *string `json:"SupportedEvaluationResource,omitempty" xml:"SupportedEvaluationResource,omitempty"`
	// The supported deployment resources, such as CPU, GPU, or GP7V. For example, if you specify GP7V, only models that support deployment on GP7V resources are returned.
	//
	// example:
	//
	// GP7V
	SupportedInferenceResource *string `json:"SupportedInferenceResource,omitempty" xml:"SupportedInferenceResource,omitempty"`
	// The supported training resources, such as CPU, GPU, or GP7V. For example, if you specify GP7V, only models that support training on GP7V resources are returned.
	//
	// example:
	//
	// GP7V
	SupportedTrainingResource *string `json:"SupportedTrainingResource,omitempty" xml:"SupportedTrainingResource,omitempty"`
	// The list of labels.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The task. For example, large-language-model (large language model), image-classification (image classification), or embedding.
	//
	// example:
	//
	// large-language-model
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// Specifies whether training is supported.
	Trainable *bool `json:"Trainable,omitempty" xml:"Trainable,omitempty"`
}

func (s ListModelGalleryModelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelGalleryModelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelGalleryModelsShrinkRequest) GetCollections() *string {
	return s.Collections
}

func (s *ListModelGalleryModelsShrinkRequest) GetCompressible() *bool {
	return s.Compressible
}

func (s *ListModelGalleryModelsShrinkRequest) GetConditionsShrink() *string {
	return s.ConditionsShrink
}

func (s *ListModelGalleryModelsShrinkRequest) GetDeepThink() *bool {
	return s.DeepThink
}

func (s *ListModelGalleryModelsShrinkRequest) GetDemonstrable() *bool {
	return s.Demonstrable
}

func (s *ListModelGalleryModelsShrinkRequest) GetDeployable() *bool {
	return s.Deployable
}

func (s *ListModelGalleryModelsShrinkRequest) GetDistillable() *bool {
	return s.Distillable
}

func (s *ListModelGalleryModelsShrinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListModelGalleryModelsShrinkRequest) GetEvaluable() *bool {
	return s.Evaluable
}

func (s *ListModelGalleryModelsShrinkRequest) GetFunctionCall() *bool {
	return s.FunctionCall
}

func (s *ListModelGalleryModelsShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelGalleryModelsShrinkRequest) GetModelSeries() *string {
	return s.ModelSeries
}

func (s *ListModelGalleryModelsShrinkRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelGalleryModelsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelGalleryModelsShrinkRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListModelGalleryModelsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelGalleryModelsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelGalleryModelsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListModelGalleryModelsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelGalleryModelsShrinkRequest) GetSupportedCompressionResource() *string {
	return s.SupportedCompressionResource
}

func (s *ListModelGalleryModelsShrinkRequest) GetSupportedDistillationResource() *string {
	return s.SupportedDistillationResource
}

func (s *ListModelGalleryModelsShrinkRequest) GetSupportedEvaluationResource() *string {
	return s.SupportedEvaluationResource
}

func (s *ListModelGalleryModelsShrinkRequest) GetSupportedInferenceResource() *string {
	return s.SupportedInferenceResource
}

func (s *ListModelGalleryModelsShrinkRequest) GetSupportedTrainingResource() *string {
	return s.SupportedTrainingResource
}

func (s *ListModelGalleryModelsShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListModelGalleryModelsShrinkRequest) GetTask() *string {
	return s.Task
}

func (s *ListModelGalleryModelsShrinkRequest) GetTrainable() *bool {
	return s.Trainable
}

func (s *ListModelGalleryModelsShrinkRequest) SetCollections(v string) *ListModelGalleryModelsShrinkRequest {
	s.Collections = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetCompressible(v bool) *ListModelGalleryModelsShrinkRequest {
	s.Compressible = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetConditionsShrink(v string) *ListModelGalleryModelsShrinkRequest {
	s.ConditionsShrink = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetDeepThink(v bool) *ListModelGalleryModelsShrinkRequest {
	s.DeepThink = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetDemonstrable(v bool) *ListModelGalleryModelsShrinkRequest {
	s.Demonstrable = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetDeployable(v bool) *ListModelGalleryModelsShrinkRequest {
	s.Deployable = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetDistillable(v bool) *ListModelGalleryModelsShrinkRequest {
	s.Distillable = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetDomain(v string) *ListModelGalleryModelsShrinkRequest {
	s.Domain = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetEvaluable(v bool) *ListModelGalleryModelsShrinkRequest {
	s.Evaluable = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetFunctionCall(v bool) *ListModelGalleryModelsShrinkRequest {
	s.FunctionCall = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetModelName(v string) *ListModelGalleryModelsShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetModelSeries(v string) *ListModelGalleryModelsShrinkRequest {
	s.ModelSeries = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetModelType(v string) *ListModelGalleryModelsShrinkRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetOrder(v string) *ListModelGalleryModelsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetOrigin(v string) *ListModelGalleryModelsShrinkRequest {
	s.Origin = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetPageNumber(v int32) *ListModelGalleryModelsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetPageSize(v int32) *ListModelGalleryModelsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetQuery(v string) *ListModelGalleryModelsShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetSortBy(v string) *ListModelGalleryModelsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetSupportedCompressionResource(v string) *ListModelGalleryModelsShrinkRequest {
	s.SupportedCompressionResource = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetSupportedDistillationResource(v string) *ListModelGalleryModelsShrinkRequest {
	s.SupportedDistillationResource = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetSupportedEvaluationResource(v string) *ListModelGalleryModelsShrinkRequest {
	s.SupportedEvaluationResource = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetSupportedInferenceResource(v string) *ListModelGalleryModelsShrinkRequest {
	s.SupportedInferenceResource = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetSupportedTrainingResource(v string) *ListModelGalleryModelsShrinkRequest {
	s.SupportedTrainingResource = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetTagShrink(v string) *ListModelGalleryModelsShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetTask(v string) *ListModelGalleryModelsShrinkRequest {
	s.Task = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) SetTrainable(v bool) *ListModelGalleryModelsShrinkRequest {
	s.Trainable = &v
	return s
}

func (s *ListModelGalleryModelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
