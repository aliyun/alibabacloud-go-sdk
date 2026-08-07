// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelGalleryModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollections(v string) *ListModelGalleryModelsRequest
	GetCollections() *string
	SetCompressible(v bool) *ListModelGalleryModelsRequest
	GetCompressible() *bool
	SetConditions(v []*ListModelGalleryModelsRequestConditions) *ListModelGalleryModelsRequest
	GetConditions() []*ListModelGalleryModelsRequestConditions
	SetDeepThink(v bool) *ListModelGalleryModelsRequest
	GetDeepThink() *bool
	SetDemonstrable(v bool) *ListModelGalleryModelsRequest
	GetDemonstrable() *bool
	SetDeployable(v bool) *ListModelGalleryModelsRequest
	GetDeployable() *bool
	SetDistillable(v bool) *ListModelGalleryModelsRequest
	GetDistillable() *bool
	SetDomain(v string) *ListModelGalleryModelsRequest
	GetDomain() *string
	SetEvaluable(v bool) *ListModelGalleryModelsRequest
	GetEvaluable() *bool
	SetFunctionCall(v bool) *ListModelGalleryModelsRequest
	GetFunctionCall() *bool
	SetModelName(v string) *ListModelGalleryModelsRequest
	GetModelName() *string
	SetModelSeries(v string) *ListModelGalleryModelsRequest
	GetModelSeries() *string
	SetModelType(v string) *ListModelGalleryModelsRequest
	GetModelType() *string
	SetOrder(v string) *ListModelGalleryModelsRequest
	GetOrder() *string
	SetOrigin(v string) *ListModelGalleryModelsRequest
	GetOrigin() *string
	SetPageNumber(v int32) *ListModelGalleryModelsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelGalleryModelsRequest
	GetPageSize() *int32
	SetQuery(v string) *ListModelGalleryModelsRequest
	GetQuery() *string
	SetSortBy(v string) *ListModelGalleryModelsRequest
	GetSortBy() *string
	SetSupportedCompressionResource(v string) *ListModelGalleryModelsRequest
	GetSupportedCompressionResource() *string
	SetSupportedDistillationResource(v string) *ListModelGalleryModelsRequest
	GetSupportedDistillationResource() *string
	SetSupportedEvaluationResource(v string) *ListModelGalleryModelsRequest
	GetSupportedEvaluationResource() *string
	SetSupportedInferenceResource(v string) *ListModelGalleryModelsRequest
	GetSupportedInferenceResource() *string
	SetSupportedTrainingResource(v string) *ListModelGalleryModelsRequest
	GetSupportedTrainingResource() *string
	SetTag(v []*ListModelGalleryModelsRequestTag) *ListModelGalleryModelsRequest
	GetTag() []*ListModelGalleryModelsRequestTag
	SetTask(v string) *ListModelGalleryModelsRequest
	GetTask() *string
	SetTrainable(v bool) *ListModelGalleryModelsRequest
	GetTrainable() *bool
}

type ListModelGalleryModelsRequest struct {
	// example:
	//
	// QuickStart
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	// example:
	//
	// true
	Compressible *bool                                      `json:"Compressible,omitempty" xml:"Compressible,omitempty"`
	Conditions   []*ListModelGalleryModelsRequestConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
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
	// true
	FunctionCall *bool `json:"FunctionCall,omitempty" xml:"FunctionCall,omitempty"`
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
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Qwen
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// GP7V
	SupportedCompressionResource *string `json:"SupportedCompressionResource,omitempty" xml:"SupportedCompressionResource,omitempty"`
	// example:
	//
	// GP7V
	SupportedDistillationResource *string `json:"SupportedDistillationResource,omitempty" xml:"SupportedDistillationResource,omitempty"`
	// example:
	//
	// GP7V
	SupportedEvaluationResource *string `json:"SupportedEvaluationResource,omitempty" xml:"SupportedEvaluationResource,omitempty"`
	// example:
	//
	// GP7V
	SupportedInferenceResource *string `json:"SupportedInferenceResource,omitempty" xml:"SupportedInferenceResource,omitempty"`
	// example:
	//
	// GP7V
	SupportedTrainingResource *string                             `json:"SupportedTrainingResource,omitempty" xml:"SupportedTrainingResource,omitempty"`
	Tag                       []*ListModelGalleryModelsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// large-language-model
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// example:
	//
	// true
	Trainable *bool `json:"Trainable,omitempty" xml:"Trainable,omitempty"`
}

func (s ListModelGalleryModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelGalleryModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelGalleryModelsRequest) GetCollections() *string {
	return s.Collections
}

func (s *ListModelGalleryModelsRequest) GetCompressible() *bool {
	return s.Compressible
}

func (s *ListModelGalleryModelsRequest) GetConditions() []*ListModelGalleryModelsRequestConditions {
	return s.Conditions
}

func (s *ListModelGalleryModelsRequest) GetDeepThink() *bool {
	return s.DeepThink
}

func (s *ListModelGalleryModelsRequest) GetDemonstrable() *bool {
	return s.Demonstrable
}

func (s *ListModelGalleryModelsRequest) GetDeployable() *bool {
	return s.Deployable
}

func (s *ListModelGalleryModelsRequest) GetDistillable() *bool {
	return s.Distillable
}

func (s *ListModelGalleryModelsRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListModelGalleryModelsRequest) GetEvaluable() *bool {
	return s.Evaluable
}

func (s *ListModelGalleryModelsRequest) GetFunctionCall() *bool {
	return s.FunctionCall
}

func (s *ListModelGalleryModelsRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelGalleryModelsRequest) GetModelSeries() *string {
	return s.ModelSeries
}

func (s *ListModelGalleryModelsRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelGalleryModelsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelGalleryModelsRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListModelGalleryModelsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelGalleryModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelGalleryModelsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListModelGalleryModelsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelGalleryModelsRequest) GetSupportedCompressionResource() *string {
	return s.SupportedCompressionResource
}

func (s *ListModelGalleryModelsRequest) GetSupportedDistillationResource() *string {
	return s.SupportedDistillationResource
}

func (s *ListModelGalleryModelsRequest) GetSupportedEvaluationResource() *string {
	return s.SupportedEvaluationResource
}

func (s *ListModelGalleryModelsRequest) GetSupportedInferenceResource() *string {
	return s.SupportedInferenceResource
}

func (s *ListModelGalleryModelsRequest) GetSupportedTrainingResource() *string {
	return s.SupportedTrainingResource
}

func (s *ListModelGalleryModelsRequest) GetTag() []*ListModelGalleryModelsRequestTag {
	return s.Tag
}

func (s *ListModelGalleryModelsRequest) GetTask() *string {
	return s.Task
}

func (s *ListModelGalleryModelsRequest) GetTrainable() *bool {
	return s.Trainable
}

func (s *ListModelGalleryModelsRequest) SetCollections(v string) *ListModelGalleryModelsRequest {
	s.Collections = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetCompressible(v bool) *ListModelGalleryModelsRequest {
	s.Compressible = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetConditions(v []*ListModelGalleryModelsRequestConditions) *ListModelGalleryModelsRequest {
	s.Conditions = v
	return s
}

func (s *ListModelGalleryModelsRequest) SetDeepThink(v bool) *ListModelGalleryModelsRequest {
	s.DeepThink = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetDemonstrable(v bool) *ListModelGalleryModelsRequest {
	s.Demonstrable = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetDeployable(v bool) *ListModelGalleryModelsRequest {
	s.Deployable = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetDistillable(v bool) *ListModelGalleryModelsRequest {
	s.Distillable = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetDomain(v string) *ListModelGalleryModelsRequest {
	s.Domain = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetEvaluable(v bool) *ListModelGalleryModelsRequest {
	s.Evaluable = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetFunctionCall(v bool) *ListModelGalleryModelsRequest {
	s.FunctionCall = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetModelName(v string) *ListModelGalleryModelsRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetModelSeries(v string) *ListModelGalleryModelsRequest {
	s.ModelSeries = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetModelType(v string) *ListModelGalleryModelsRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetOrder(v string) *ListModelGalleryModelsRequest {
	s.Order = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetOrigin(v string) *ListModelGalleryModelsRequest {
	s.Origin = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetPageNumber(v int32) *ListModelGalleryModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetPageSize(v int32) *ListModelGalleryModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetQuery(v string) *ListModelGalleryModelsRequest {
	s.Query = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetSortBy(v string) *ListModelGalleryModelsRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetSupportedCompressionResource(v string) *ListModelGalleryModelsRequest {
	s.SupportedCompressionResource = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetSupportedDistillationResource(v string) *ListModelGalleryModelsRequest {
	s.SupportedDistillationResource = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetSupportedEvaluationResource(v string) *ListModelGalleryModelsRequest {
	s.SupportedEvaluationResource = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetSupportedInferenceResource(v string) *ListModelGalleryModelsRequest {
	s.SupportedInferenceResource = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetSupportedTrainingResource(v string) *ListModelGalleryModelsRequest {
	s.SupportedTrainingResource = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetTag(v []*ListModelGalleryModelsRequestTag) *ListModelGalleryModelsRequest {
	s.Tag = v
	return s
}

func (s *ListModelGalleryModelsRequest) SetTask(v string) *ListModelGalleryModelsRequest {
	s.Task = &v
	return s
}

func (s *ListModelGalleryModelsRequest) SetTrainable(v bool) *ListModelGalleryModelsRequest {
	s.Trainable = &v
	return s
}

func (s *ListModelGalleryModelsRequest) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelGalleryModelsRequestConditions struct {
	// example:
	//
	// ParameterSize
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// example:
	//
	// LessThan
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 1024
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListModelGalleryModelsRequestConditions) String() string {
	return dara.Prettify(s)
}

func (s ListModelGalleryModelsRequestConditions) GoString() string {
	return s.String()
}

func (s *ListModelGalleryModelsRequestConditions) GetColumn() *string {
	return s.Column
}

func (s *ListModelGalleryModelsRequestConditions) GetOperator() *string {
	return s.Operator
}

func (s *ListModelGalleryModelsRequestConditions) GetValue() *string {
	return s.Value
}

func (s *ListModelGalleryModelsRequestConditions) SetColumn(v string) *ListModelGalleryModelsRequestConditions {
	s.Column = &v
	return s
}

func (s *ListModelGalleryModelsRequestConditions) SetOperator(v string) *ListModelGalleryModelsRequestConditions {
	s.Operator = &v
	return s
}

func (s *ListModelGalleryModelsRequestConditions) SetValue(v string) *ListModelGalleryModelsRequestConditions {
	s.Value = &v
	return s
}

func (s *ListModelGalleryModelsRequestConditions) Validate() error {
	return dara.Validate(s)
}

type ListModelGalleryModelsRequestTag struct {
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListModelGalleryModelsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListModelGalleryModelsRequestTag) GoString() string {
	return s.String()
}

func (s *ListModelGalleryModelsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListModelGalleryModelsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListModelGalleryModelsRequestTag) SetKey(v string) *ListModelGalleryModelsRequestTag {
	s.Key = &v
	return s
}

func (s *ListModelGalleryModelsRequestTag) SetValue(v string) *ListModelGalleryModelsRequestTag {
	s.Value = &v
	return s
}

func (s *ListModelGalleryModelsRequestTag) Validate() error {
	return dara.Validate(s)
}
