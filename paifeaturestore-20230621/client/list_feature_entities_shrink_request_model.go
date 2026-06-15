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
	SetParentFeatureEntityId(v string) *ListFeatureEntitiesShrinkRequest
	GetParentFeatureEntityId() *string
	SetProjectId(v string) *ListFeatureEntitiesShrinkRequest
	GetProjectId() *string
	SetSortBy(v string) *ListFeatureEntitiesShrinkRequest
	GetSortBy() *string
}

type ListFeatureEntitiesShrinkRequest struct {
	// Filters the results by a list of feature entity IDs.
	FeatureEntityIdsShrink *string `json:"FeatureEntityIds,omitempty" xml:"FeatureEntityIds,omitempty"`
	// Filters the results by feature entity name.
	//
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Valid values:
	//
	// - `Asc`: Ascending order.
	//
	// - `Desc`: Descending order.
	//
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The Alibaba Cloud account ID of the creator.
	//
	// example:
	//
	// 1231432*****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Values start at 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters the results by parent feature entity ID. Omit this parameter to return all feature entities. Set it to `0` to return all root feature entities. Set it to a non-zero value to return all child feature entities of the specified parent.
	//
	// example:
	//
	// 1
	ParentFeatureEntityId *string `json:"ParentFeatureEntityId,omitempty" xml:"ParentFeatureEntityId,omitempty"`
	// The project ID. You can obtain this ID by calling the `ListProjects` operation.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The field to sort the results by.
	//
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

func (s *ListFeatureEntitiesShrinkRequest) GetParentFeatureEntityId() *string {
	return s.ParentFeatureEntityId
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

func (s *ListFeatureEntitiesShrinkRequest) SetParentFeatureEntityId(v string) *ListFeatureEntitiesShrinkRequest {
	s.ParentFeatureEntityId = &v
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
