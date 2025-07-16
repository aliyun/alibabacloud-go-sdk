// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureName(v string) *ListFeatureViewsRequest
	GetFeatureName() *string
	SetFeatureViewIds(v []*string) *ListFeatureViewsRequest
	GetFeatureViewIds() []*string
	SetName(v string) *ListFeatureViewsRequest
	GetName() *string
	SetOrder(v string) *ListFeatureViewsRequest
	GetOrder() *string
	SetOwner(v string) *ListFeatureViewsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListFeatureViewsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFeatureViewsRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListFeatureViewsRequest
	GetProjectId() *string
	SetSortBy(v string) *ListFeatureViewsRequest
	GetSortBy() *string
	SetTag(v string) *ListFeatureViewsRequest
	GetTag() *string
	SetType(v string) *ListFeatureViewsRequest
	GetType() *string
}

type ListFeatureViewsRequest struct {
	// example:
	//
	// feature1
	FeatureName    *string   `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	FeatureViewIds []*string `json:"FeatureViewIds,omitempty" xml:"FeatureViewIds,omitempty" type:"Repeated"`
	// example:
	//
	// fv1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1232143243242****
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
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFeatureViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListFeatureViewsRequest) GetFeatureViewIds() []*string {
	return s.FeatureViewIds
}

func (s *ListFeatureViewsRequest) GetName() *string {
	return s.Name
}

func (s *ListFeatureViewsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListFeatureViewsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListFeatureViewsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFeatureViewsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFeatureViewsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListFeatureViewsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListFeatureViewsRequest) GetTag() *string {
	return s.Tag
}

func (s *ListFeatureViewsRequest) GetType() *string {
	return s.Type
}

func (s *ListFeatureViewsRequest) SetFeatureName(v string) *ListFeatureViewsRequest {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureViewsRequest) SetFeatureViewIds(v []*string) *ListFeatureViewsRequest {
	s.FeatureViewIds = v
	return s
}

func (s *ListFeatureViewsRequest) SetName(v string) *ListFeatureViewsRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureViewsRequest) SetOrder(v string) *ListFeatureViewsRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureViewsRequest) SetOwner(v string) *ListFeatureViewsRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureViewsRequest) SetPageNumber(v int32) *ListFeatureViewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureViewsRequest) SetPageSize(v int32) *ListFeatureViewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureViewsRequest) SetProjectId(v string) *ListFeatureViewsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureViewsRequest) SetSortBy(v string) *ListFeatureViewsRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureViewsRequest) SetTag(v string) *ListFeatureViewsRequest {
	s.Tag = &v
	return s
}

func (s *ListFeatureViewsRequest) SetType(v string) *ListFeatureViewsRequest {
	s.Type = &v
	return s
}

func (s *ListFeatureViewsRequest) Validate() error {
	return dara.Validate(s)
}
