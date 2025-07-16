// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelFeatureIds(v []*string) *ListModelFeaturesRequest
	GetModelFeatureIds() []*string
	SetName(v string) *ListModelFeaturesRequest
	GetName() *string
	SetOrder(v string) *ListModelFeaturesRequest
	GetOrder() *string
	SetOwner(v string) *ListModelFeaturesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListModelFeaturesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelFeaturesRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListModelFeaturesRequest
	GetProjectId() *string
	SetSortBy(v string) *ListModelFeaturesRequest
	GetSortBy() *string
}

type ListModelFeaturesRequest struct {
	ModelFeatureIds []*string `json:"ModelFeatureIds,omitempty" xml:"ModelFeatureIds,omitempty" type:"Repeated"`
	// example:
	//
	// model_feature1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 12323143****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// DESC
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListModelFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesRequest) GetModelFeatureIds() []*string {
	return s.ModelFeatureIds
}

func (s *ListModelFeaturesRequest) GetName() *string {
	return s.Name
}

func (s *ListModelFeaturesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelFeaturesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListModelFeaturesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelFeaturesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelFeaturesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListModelFeaturesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelFeaturesRequest) SetModelFeatureIds(v []*string) *ListModelFeaturesRequest {
	s.ModelFeatureIds = v
	return s
}

func (s *ListModelFeaturesRequest) SetName(v string) *ListModelFeaturesRequest {
	s.Name = &v
	return s
}

func (s *ListModelFeaturesRequest) SetOrder(v string) *ListModelFeaturesRequest {
	s.Order = &v
	return s
}

func (s *ListModelFeaturesRequest) SetOwner(v string) *ListModelFeaturesRequest {
	s.Owner = &v
	return s
}

func (s *ListModelFeaturesRequest) SetPageNumber(v int32) *ListModelFeaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelFeaturesRequest) SetPageSize(v int32) *ListModelFeaturesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelFeaturesRequest) SetProjectId(v string) *ListModelFeaturesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModelFeaturesRequest) SetSortBy(v string) *ListModelFeaturesRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
