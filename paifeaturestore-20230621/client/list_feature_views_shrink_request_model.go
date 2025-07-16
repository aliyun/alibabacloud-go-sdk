// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureName(v string) *ListFeatureViewsShrinkRequest
	GetFeatureName() *string
	SetFeatureViewIdsShrink(v string) *ListFeatureViewsShrinkRequest
	GetFeatureViewIdsShrink() *string
	SetName(v string) *ListFeatureViewsShrinkRequest
	GetName() *string
	SetOrder(v string) *ListFeatureViewsShrinkRequest
	GetOrder() *string
	SetOwner(v string) *ListFeatureViewsShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListFeatureViewsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFeatureViewsShrinkRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListFeatureViewsShrinkRequest
	GetProjectId() *string
	SetSortBy(v string) *ListFeatureViewsShrinkRequest
	GetSortBy() *string
	SetTag(v string) *ListFeatureViewsShrinkRequest
	GetTag() *string
	SetType(v string) *ListFeatureViewsShrinkRequest
	GetType() *string
}

type ListFeatureViewsShrinkRequest struct {
	// example:
	//
	// feature1
	FeatureName          *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	FeatureViewIdsShrink *string `json:"FeatureViewIds,omitempty" xml:"FeatureViewIds,omitempty"`
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

func (s ListFeatureViewsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsShrinkRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListFeatureViewsShrinkRequest) GetFeatureViewIdsShrink() *string {
	return s.FeatureViewIdsShrink
}

func (s *ListFeatureViewsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListFeatureViewsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListFeatureViewsShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListFeatureViewsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFeatureViewsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFeatureViewsShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListFeatureViewsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListFeatureViewsShrinkRequest) GetTag() *string {
	return s.Tag
}

func (s *ListFeatureViewsShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListFeatureViewsShrinkRequest) SetFeatureName(v string) *ListFeatureViewsShrinkRequest {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetFeatureViewIdsShrink(v string) *ListFeatureViewsShrinkRequest {
	s.FeatureViewIdsShrink = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetName(v string) *ListFeatureViewsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetOrder(v string) *ListFeatureViewsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetOwner(v string) *ListFeatureViewsShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetPageNumber(v int32) *ListFeatureViewsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetPageSize(v int32) *ListFeatureViewsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetProjectId(v string) *ListFeatureViewsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetSortBy(v string) *ListFeatureViewsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetTag(v string) *ListFeatureViewsShrinkRequest {
	s.Tag = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) SetType(v string) *ListFeatureViewsShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListFeatureViewsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
