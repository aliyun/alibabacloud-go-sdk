// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaEntitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeFiltersShrink(v string) *ListMetaEntitiesShrinkRequest
	GetAttributeFiltersShrink() *string
	SetComment(v string) *ListMetaEntitiesShrinkRequest
	GetComment() *string
	SetCustomAttributeFiltersShrink(v string) *ListMetaEntitiesShrinkRequest
	GetCustomAttributeFiltersShrink() *string
	SetEntityType(v string) *ListMetaEntitiesShrinkRequest
	GetEntityType() *string
	SetMaxResults(v int32) *ListMetaEntitiesShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListMetaEntitiesShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListMetaEntitiesShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListMetaEntitiesShrinkRequest
	GetOrder() *string
	SetSortBy(v string) *ListMetaEntitiesShrinkRequest
	GetSortBy() *string
}

type ListMetaEntitiesShrinkRequest struct {
	// example:
	//
	// []
	AttributeFiltersShrink *string `json:"AttributeFilters,omitempty" xml:"AttributeFilters,omitempty"`
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// []
	CustomAttributeFiltersShrink *string `json:"CustomAttributeFilters,omitempty" xml:"CustomAttributeFilters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AAAAAaUpAxoCTD/+sbOf3f+uxvnYyILMeAjoTFQSX64R12GN
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListMetaEntitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMetaEntitiesShrinkRequest) GetAttributeFiltersShrink() *string {
	return s.AttributeFiltersShrink
}

func (s *ListMetaEntitiesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *ListMetaEntitiesShrinkRequest) GetCustomAttributeFiltersShrink() *string {
	return s.CustomAttributeFiltersShrink
}

func (s *ListMetaEntitiesShrinkRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListMetaEntitiesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMetaEntitiesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListMetaEntitiesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMetaEntitiesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListMetaEntitiesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListMetaEntitiesShrinkRequest) SetAttributeFiltersShrink(v string) *ListMetaEntitiesShrinkRequest {
	s.AttributeFiltersShrink = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetComment(v string) *ListMetaEntitiesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetCustomAttributeFiltersShrink(v string) *ListMetaEntitiesShrinkRequest {
	s.CustomAttributeFiltersShrink = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetEntityType(v string) *ListMetaEntitiesShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetMaxResults(v int32) *ListMetaEntitiesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetName(v string) *ListMetaEntitiesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetNextToken(v string) *ListMetaEntitiesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetOrder(v string) *ListMetaEntitiesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) SetSortBy(v string) *ListMetaEntitiesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListMetaEntitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
