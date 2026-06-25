// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollections(v string) *ListModelsRequest
	GetCollections() *string
	SetConditions(v []*ListModelsRequestConditions) *ListModelsRequest
	GetConditions() []*ListModelsRequestConditions
	SetDomain(v string) *ListModelsRequest
	GetDomain() *string
	SetLabel(v string) *ListModelsRequest
	GetLabel() *string
	SetModelName(v string) *ListModelsRequest
	GetModelName() *string
	SetModelType(v string) *ListModelsRequest
	GetModelType() *string
	SetOrder(v string) *ListModelsRequest
	GetOrder() *string
	SetOrigin(v string) *ListModelsRequest
	GetOrigin() *string
	SetPageNumber(v int32) *ListModelsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelsRequest
	GetPageSize() *int32
	SetProvider(v string) *ListModelsRequest
	GetProvider() *string
	SetQuery(v string) *ListModelsRequest
	GetQuery() *string
	SetSortBy(v string) *ListModelsRequest
	GetSortBy() *string
	SetTag(v []*ListModelsRequestTag) *ListModelsRequest
	GetTag() []*ListModelsRequestTag
	SetTask(v string) *ListModelsRequest
	GetTask() *string
	SetWorkspaceId(v string) *ListModelsRequest
	GetWorkspaceId() *string
}

type ListModelsRequest struct {
	// The collections to which the model belongs. You can specify multiple collections. Separate them with commas (,).
	//
	// example:
	//
	// AI4D,QuickStart
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	// The conditions.
	Conditions []*ListModelsRequestConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
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
	Tag []*ListModelsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s ListModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) GetCollections() *string {
	return s.Collections
}

func (s *ListModelsRequest) GetConditions() []*ListModelsRequestConditions {
	return s.Conditions
}

func (s *ListModelsRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListModelsRequest) GetLabel() *string {
	return s.Label
}

func (s *ListModelsRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelsRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelsRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListModelsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelsRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListModelsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListModelsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelsRequest) GetTag() []*ListModelsRequestTag {
	return s.Tag
}

func (s *ListModelsRequest) GetTask() *string {
	return s.Task
}

func (s *ListModelsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListModelsRequest) SetCollections(v string) *ListModelsRequest {
	s.Collections = &v
	return s
}

func (s *ListModelsRequest) SetConditions(v []*ListModelsRequestConditions) *ListModelsRequest {
	s.Conditions = v
	return s
}

func (s *ListModelsRequest) SetDomain(v string) *ListModelsRequest {
	s.Domain = &v
	return s
}

func (s *ListModelsRequest) SetLabel(v string) *ListModelsRequest {
	s.Label = &v
	return s
}

func (s *ListModelsRequest) SetModelName(v string) *ListModelsRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelsRequest) SetModelType(v string) *ListModelsRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelsRequest) SetOrder(v string) *ListModelsRequest {
	s.Order = &v
	return s
}

func (s *ListModelsRequest) SetOrigin(v string) *ListModelsRequest {
	s.Origin = &v
	return s
}

func (s *ListModelsRequest) SetPageNumber(v int32) *ListModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsRequest) SetPageSize(v int32) *ListModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsRequest) SetProvider(v string) *ListModelsRequest {
	s.Provider = &v
	return s
}

func (s *ListModelsRequest) SetQuery(v string) *ListModelsRequest {
	s.Query = &v
	return s
}

func (s *ListModelsRequest) SetSortBy(v string) *ListModelsRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelsRequest) SetTag(v []*ListModelsRequestTag) *ListModelsRequest {
	s.Tag = v
	return s
}

func (s *ListModelsRequest) SetTask(v string) *ListModelsRequest {
	s.Task = &v
	return s
}

func (s *ListModelsRequest) SetWorkspaceId(v string) *ListModelsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelsRequest) Validate() error {
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

type ListModelsRequestConditions struct {
	// The parameter name. Example: ParameterSize.
	//
	// example:
	//
	// ParameterSize
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The operator. Example: LessThan.
	//
	// example:
	//
	// LessThan
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value. Example: 3000.
	//
	// example:
	//
	// 3000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListModelsRequestConditions) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequestConditions) GoString() string {
	return s.String()
}

func (s *ListModelsRequestConditions) GetColumn() *string {
	return s.Column
}

func (s *ListModelsRequestConditions) GetOperator() *string {
	return s.Operator
}

func (s *ListModelsRequestConditions) GetValue() *string {
	return s.Value
}

func (s *ListModelsRequestConditions) SetColumn(v string) *ListModelsRequestConditions {
	s.Column = &v
	return s
}

func (s *ListModelsRequestConditions) SetOperator(v string) *ListModelsRequestConditions {
	s.Operator = &v
	return s
}

func (s *ListModelsRequestConditions) SetValue(v string) *ListModelsRequestConditions {
	s.Value = &v
	return s
}

func (s *ListModelsRequestConditions) Validate() error {
	return dara.Validate(s)
}

type ListModelsRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListModelsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequestTag) GoString() string {
	return s.String()
}

func (s *ListModelsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListModelsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListModelsRequestTag) SetKey(v string) *ListModelsRequestTag {
	s.Key = &v
	return s
}

func (s *ListModelsRequestTag) SetValue(v string) *ListModelsRequestTag {
	s.Value = &v
	return s
}

func (s *ListModelsRequestTag) Validate() error {
	return dara.Validate(s)
}
