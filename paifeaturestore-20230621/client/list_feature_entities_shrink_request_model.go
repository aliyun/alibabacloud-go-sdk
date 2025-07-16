// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureEntitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureEntityIdsShrink(v string) *ListFeatureEntitiesShrinkRequest
	GetFeatureEntityIdsShrink() *string
	SetName(v string) *ListFeatureEntitiesShrinkRequest
	GetName() *string
	SetOrder(v string) *ListFeatureEntitiesShrinkRequest
	GetOrder() *string
	SetOwner(v string) *ListFeatureEntitiesShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListFeatureEntitiesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFeatureEntitiesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListFeatureEntitiesShrinkRequest
	GetProjectId() *string
	SetSortBy(v string) *ListFeatureEntitiesShrinkRequest
	GetSortBy() *string
}

type ListFeatureEntitiesShrinkRequest struct {
	FeatureEntityIdsShrink *string `json:"FeatureEntityIds,omitempty" xml:"FeatureEntityIds,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1231432*****
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
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListFeatureEntitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureEntitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesShrinkRequest) GetFeatureEntityIdsShrink() *string {
	return s.FeatureEntityIdsShrink
}

func (s *ListFeatureEntitiesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListFeatureEntitiesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListFeatureEntitiesShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListFeatureEntitiesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFeatureEntitiesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFeatureEntitiesShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListFeatureEntitiesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListFeatureEntitiesShrinkRequest) SetFeatureEntityIdsShrink(v string) *ListFeatureEntitiesShrinkRequest {
	s.FeatureEntityIdsShrink = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetName(v string) *ListFeatureEntitiesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetOrder(v string) *ListFeatureEntitiesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetOwner(v string) *ListFeatureEntitiesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetPageNumber(v int32) *ListFeatureEntitiesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetPageSize(v int32) *ListFeatureEntitiesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetProjectId(v string) *ListFeatureEntitiesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) SetSortBy(v string) *ListFeatureEntitiesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureEntitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
