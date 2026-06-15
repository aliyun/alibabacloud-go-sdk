// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureEntityIds(v []*string) *ListFeatureEntitiesRequest
	GetFeatureEntityIds() []*string
	SetName(v string) *ListFeatureEntitiesRequest
	GetName() *string
	SetOrder(v string) *ListFeatureEntitiesRequest
	GetOrder() *string
	SetOwner(v string) *ListFeatureEntitiesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListFeatureEntitiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFeatureEntitiesRequest
	GetPageSize() *int32
	SetParentFeatureEntityId(v string) *ListFeatureEntitiesRequest
	GetParentFeatureEntityId() *string
	SetProjectId(v string) *ListFeatureEntitiesRequest
	GetProjectId() *string
	SetSortBy(v string) *ListFeatureEntitiesRequest
	GetSortBy() *string
}

type ListFeatureEntitiesRequest struct {
	// Filters the results by a list of feature entity IDs.
	FeatureEntityIds []*string `json:"FeatureEntityIds,omitempty" xml:"FeatureEntityIds,omitempty" type:"Repeated"`
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

func (s ListFeatureEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesRequest) GetFeatureEntityIds() []*string {
	return s.FeatureEntityIds
}

func (s *ListFeatureEntitiesRequest) GetName() *string {
	return s.Name
}

func (s *ListFeatureEntitiesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListFeatureEntitiesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListFeatureEntitiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFeatureEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFeatureEntitiesRequest) GetParentFeatureEntityId() *string {
	return s.ParentFeatureEntityId
}

func (s *ListFeatureEntitiesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListFeatureEntitiesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListFeatureEntitiesRequest) SetFeatureEntityIds(v []*string) *ListFeatureEntitiesRequest {
	s.FeatureEntityIds = v
	return s
}

func (s *ListFeatureEntitiesRequest) SetName(v string) *ListFeatureEntitiesRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetOrder(v string) *ListFeatureEntitiesRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetOwner(v string) *ListFeatureEntitiesRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetPageNumber(v int32) *ListFeatureEntitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetPageSize(v int32) *ListFeatureEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetParentFeatureEntityId(v string) *ListFeatureEntitiesRequest {
	s.ParentFeatureEntityId = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetProjectId(v string) *ListFeatureEntitiesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureEntitiesRequest) SetSortBy(v string) *ListFeatureEntitiesRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
