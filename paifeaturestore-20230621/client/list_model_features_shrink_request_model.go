// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelFeaturesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelFeatureIdsShrink(v string) *ListModelFeaturesShrinkRequest
	GetModelFeatureIdsShrink() *string
	SetName(v string) *ListModelFeaturesShrinkRequest
	GetName() *string
	SetOrder(v string) *ListModelFeaturesShrinkRequest
	GetOrder() *string
	SetOwner(v string) *ListModelFeaturesShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListModelFeaturesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelFeaturesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListModelFeaturesShrinkRequest
	GetProjectId() *string
	SetSortBy(v string) *ListModelFeaturesShrinkRequest
	GetSortBy() *string
}

type ListModelFeaturesShrinkRequest struct {
	ModelFeatureIdsShrink *string `json:"ModelFeatureIds,omitempty" xml:"ModelFeatureIds,omitempty"`
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

func (s ListModelFeaturesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelFeaturesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelFeaturesShrinkRequest) GetModelFeatureIdsShrink() *string {
	return s.ModelFeatureIdsShrink
}

func (s *ListModelFeaturesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListModelFeaturesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelFeaturesShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListModelFeaturesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelFeaturesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelFeaturesShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListModelFeaturesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelFeaturesShrinkRequest) SetModelFeatureIdsShrink(v string) *ListModelFeaturesShrinkRequest {
	s.ModelFeatureIdsShrink = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetName(v string) *ListModelFeaturesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetOrder(v string) *ListModelFeaturesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetOwner(v string) *ListModelFeaturesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetPageNumber(v int32) *ListModelFeaturesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetPageSize(v int32) *ListModelFeaturesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetProjectId(v string) *ListModelFeaturesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) SetSortBy(v string) *ListModelFeaturesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelFeaturesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
