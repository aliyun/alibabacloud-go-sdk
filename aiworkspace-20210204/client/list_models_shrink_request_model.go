// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollections(v string) *ListModelsShrinkRequest
	GetCollections() *string
	SetConditionsShrink(v string) *ListModelsShrinkRequest
	GetConditionsShrink() *string
	SetDomain(v string) *ListModelsShrinkRequest
	GetDomain() *string
	SetLabel(v string) *ListModelsShrinkRequest
	GetLabel() *string
	SetModelName(v string) *ListModelsShrinkRequest
	GetModelName() *string
	SetModelType(v string) *ListModelsShrinkRequest
	GetModelType() *string
	SetOrder(v string) *ListModelsShrinkRequest
	GetOrder() *string
	SetOrigin(v string) *ListModelsShrinkRequest
	GetOrigin() *string
	SetPageNumber(v int32) *ListModelsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelsShrinkRequest
	GetPageSize() *int32
	SetProvider(v string) *ListModelsShrinkRequest
	GetProvider() *string
	SetQuery(v string) *ListModelsShrinkRequest
	GetQuery() *string
	SetSortBy(v string) *ListModelsShrinkRequest
	GetSortBy() *string
	SetTagShrink(v string) *ListModelsShrinkRequest
	GetTagShrink() *string
	SetTask(v string) *ListModelsShrinkRequest
	GetTask() *string
	SetWorkspaceId(v string) *ListModelsShrinkRequest
	GetWorkspaceId() *string
}

type ListModelsShrinkRequest struct {
	// The collections to which the model belongs. You can specify multiple collections. Separate them with commas (,).
	//
	// example:
	//
	// AI4D,QuickStart
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	// The conditions.
	ConditionsShrink *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The domain. This parameter is used to filter the model list by domain. Examples: nlp (natural language processing) and cv (computer vision).
	//
	// example:
	//
	// nlp
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The label string. This parameter is used to filter the list. Models are returned if their label keys or values contain the specified string.
	//
	// example:
	//
	// key1
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The model name. This parameter is used to filter the model list.
	//
	// example:
	//
	// Sentiment analysis
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type.
	//
	// example:
	//
	// Endpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The order in which to sort the results of a paged query. The default value is ASC.
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The model source. This parameter is used to filter the model list by community or organization. Examples: ModelScope and HuggingFace.
	//
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The page number of the model list. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of models to display on each page in a paged query. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The provider. If you specify a provider, only the public models from that provider are returned. If you leave this parameter empty, your own models are returned.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The query condition. This parameter performs a fuzzy match on ModelName, Domain, Task, LabelKey, and LabelValue. For example, if you enter nlp, models that match in any of these fields are returned.
	//
	// example:
	//
	// nlp
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The field to use for sorting in a paged query. Currently, only the GmtCreateTime field is supported.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The list of tags.
	//
	// example:
	//
	// Endpoint
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The task. This parameter is used to filter the model list by task type. Example: text-classification.
	//
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// The workspace ID. The returned list contains only the models in the specified workspace. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 324**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListModelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelsShrinkRequest) GetCollections() *string {
	return s.Collections
}

func (s *ListModelsShrinkRequest) GetConditionsShrink() *string {
	return s.ConditionsShrink
}

func (s *ListModelsShrinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListModelsShrinkRequest) GetLabel() *string {
	return s.Label
}

func (s *ListModelsShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelsShrinkRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelsShrinkRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListModelsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelsShrinkRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListModelsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListModelsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelsShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListModelsShrinkRequest) GetTask() *string {
	return s.Task
}

func (s *ListModelsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListModelsShrinkRequest) SetCollections(v string) *ListModelsShrinkRequest {
	s.Collections = &v
	return s
}

func (s *ListModelsShrinkRequest) SetConditionsShrink(v string) *ListModelsShrinkRequest {
	s.ConditionsShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) SetDomain(v string) *ListModelsShrinkRequest {
	s.Domain = &v
	return s
}

func (s *ListModelsShrinkRequest) SetLabel(v string) *ListModelsShrinkRequest {
	s.Label = &v
	return s
}

func (s *ListModelsShrinkRequest) SetModelName(v string) *ListModelsShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelsShrinkRequest) SetModelType(v string) *ListModelsShrinkRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelsShrinkRequest) SetOrder(v string) *ListModelsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListModelsShrinkRequest) SetOrigin(v string) *ListModelsShrinkRequest {
	s.Origin = &v
	return s
}

func (s *ListModelsShrinkRequest) SetPageNumber(v int32) *ListModelsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsShrinkRequest) SetPageSize(v int32) *ListModelsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsShrinkRequest) SetProvider(v string) *ListModelsShrinkRequest {
	s.Provider = &v
	return s
}

func (s *ListModelsShrinkRequest) SetQuery(v string) *ListModelsShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListModelsShrinkRequest) SetSortBy(v string) *ListModelsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelsShrinkRequest) SetTagShrink(v string) *ListModelsShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListModelsShrinkRequest) SetTask(v string) *ListModelsShrinkRequest {
	s.Task = &v
	return s
}

func (s *ListModelsShrinkRequest) SetWorkspaceId(v string) *ListModelsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
