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
	// The collection where the model is located. You can specify multiple collections and separate them with commas (,).
	//
	// example:
	//
	// AI4D,QuickStart
	Collections      *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	ConditionsShrink *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The domain. Only models in the domain are returned. Valid values: nlp (Natural Language Processing) and cv (Computer Vision).
	//
	// example:
	//
	// nlp
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The label. Models whose label key or label value contains a specific label are filtered.
	//
	// example:
	//
	// key1
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The model name used to filter the returned models.
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type.
	//
	// example:
	//
	// Endpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The order in which the entries are sorted by the specific field on the returned page. Default value: ASC.
	//
	// 	- ASC
	//
	// 	- DESC
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The model source used to filter the models that belong to a community or organization, such as ModelScope and Hugging Face.
	//
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The provider. If you configure this parameter, only the models exposed by the provider are returned. If you leave this parameter empty, only models owned by the user are returned.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The query condition. For example, if you set the value to nlp, all models that match ModelName, Domain, Task, LabelKey, and LabelValue are returned.
	//
	// example:
	//
	// nlp
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The field used to sort the results. The GmtCreateTime field is used for sorting.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The tags of the model.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The task used to filter the models that belong to the task type. Example: text-classification.
	//
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// The workspace ID. Only models in this workspace are queried. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
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
