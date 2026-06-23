// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeFilters(v []*ListMetaEntitiesRequestAttributeFilters) *ListMetaEntitiesRequest
	GetAttributeFilters() []*ListMetaEntitiesRequestAttributeFilters
	SetComment(v string) *ListMetaEntitiesRequest
	GetComment() *string
	SetCustomAttributeFilters(v []*ListMetaEntitiesRequestCustomAttributeFilters) *ListMetaEntitiesRequest
	GetCustomAttributeFilters() []*ListMetaEntitiesRequestCustomAttributeFilters
	SetEntityType(v string) *ListMetaEntitiesRequest
	GetEntityType() *string
	SetMaxResults(v int32) *ListMetaEntitiesRequest
	GetMaxResults() *int32
	SetName(v string) *ListMetaEntitiesRequest
	GetName() *string
	SetNextToken(v string) *ListMetaEntitiesRequest
	GetNextToken() *string
	SetOrder(v string) *ListMetaEntitiesRequest
	GetOrder() *string
	SetSortBy(v string) *ListMetaEntitiesRequest
	GetSortBy() *string
}

type ListMetaEntitiesRequest struct {
	// Conditions for filtering entities by entity attributes. The `AND` operator is used between different filters, and the `OR` operator is used for multiple values within a single filter.
	//
	// example:
	//
	// []
	AttributeFilters []*ListMetaEntitiesRequestAttributeFilters `json:"AttributeFilters,omitempty" xml:"AttributeFilters,omitempty" type:"Repeated"`
	// Filters entities by comment. This is a token-based match.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Conditions for filtering entities by custom attributes. The `AND` operator is used between different filters, and the `OR` operator is used for multiple values within a single filter. This parameter supports only `ENUM` custom attributes.
	//
	// example:
	//
	// []
	CustomAttributeFilters []*ListMetaEntitiesRequestCustomAttributeFilters `json:"CustomAttributeFilters,omitempty" xml:"CustomAttributeFilters,omitempty" type:"Repeated"`
	// The type of the entity to list.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The maximum number of results to return per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Filters entities by name. This is a containment match.
	//
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that specifies the next page of results. To retrieve the first page, do not specify this parameter. To retrieve subsequent pages, set this parameter to the `NextToken` value from the previous response.
	//
	// example:
	//
	// AAAAAaUpAxoCTD/+sbOf3f+uxvnYyILMeAjoTFQSX64R12GN
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Valid values: `Asc` and `Desc`.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field to use for sorting the results.
	//
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListMetaEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListMetaEntitiesRequest) GetAttributeFilters() []*ListMetaEntitiesRequestAttributeFilters {
	return s.AttributeFilters
}

func (s *ListMetaEntitiesRequest) GetComment() *string {
	return s.Comment
}

func (s *ListMetaEntitiesRequest) GetCustomAttributeFilters() []*ListMetaEntitiesRequestCustomAttributeFilters {
	return s.CustomAttributeFilters
}

func (s *ListMetaEntitiesRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListMetaEntitiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMetaEntitiesRequest) GetName() *string {
	return s.Name
}

func (s *ListMetaEntitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMetaEntitiesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListMetaEntitiesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListMetaEntitiesRequest) SetAttributeFilters(v []*ListMetaEntitiesRequestAttributeFilters) *ListMetaEntitiesRequest {
	s.AttributeFilters = v
	return s
}

func (s *ListMetaEntitiesRequest) SetComment(v string) *ListMetaEntitiesRequest {
	s.Comment = &v
	return s
}

func (s *ListMetaEntitiesRequest) SetCustomAttributeFilters(v []*ListMetaEntitiesRequestCustomAttributeFilters) *ListMetaEntitiesRequest {
	s.CustomAttributeFilters = v
	return s
}

func (s *ListMetaEntitiesRequest) SetEntityType(v string) *ListMetaEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *ListMetaEntitiesRequest) SetMaxResults(v int32) *ListMetaEntitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMetaEntitiesRequest) SetName(v string) *ListMetaEntitiesRequest {
	s.Name = &v
	return s
}

func (s *ListMetaEntitiesRequest) SetNextToken(v string) *ListMetaEntitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMetaEntitiesRequest) SetOrder(v string) *ListMetaEntitiesRequest {
	s.Order = &v
	return s
}

func (s *ListMetaEntitiesRequest) SetSortBy(v string) *ListMetaEntitiesRequest {
	s.SortBy = &v
	return s
}

func (s *ListMetaEntitiesRequest) Validate() error {
	if s.AttributeFilters != nil {
		for _, item := range s.AttributeFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomAttributeFilters != nil {
		for _, item := range s.CustomAttributeFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMetaEntitiesRequestAttributeFilters struct {
	// The key of the entity attribute to filter by.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A list of values for the specified entity attribute.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListMetaEntitiesRequestAttributeFilters) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntitiesRequestAttributeFilters) GoString() string {
	return s.String()
}

func (s *ListMetaEntitiesRequestAttributeFilters) GetKey() *string {
	return s.Key
}

func (s *ListMetaEntitiesRequestAttributeFilters) GetValues() []*string {
	return s.Values
}

func (s *ListMetaEntitiesRequestAttributeFilters) SetKey(v string) *ListMetaEntitiesRequestAttributeFilters {
	s.Key = &v
	return s
}

func (s *ListMetaEntitiesRequestAttributeFilters) SetValues(v []*string) *ListMetaEntitiesRequestAttributeFilters {
	s.Values = v
	return s
}

func (s *ListMetaEntitiesRequestAttributeFilters) Validate() error {
	return dara.Validate(s)
}

type ListMetaEntitiesRequestCustomAttributeFilters struct {
	// The key of the custom attribute to filter by.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom_attr_1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A list of values for the specified custom attribute.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListMetaEntitiesRequestCustomAttributeFilters) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntitiesRequestCustomAttributeFilters) GoString() string {
	return s.String()
}

func (s *ListMetaEntitiesRequestCustomAttributeFilters) GetKey() *string {
	return s.Key
}

func (s *ListMetaEntitiesRequestCustomAttributeFilters) GetValues() []*string {
	return s.Values
}

func (s *ListMetaEntitiesRequestCustomAttributeFilters) SetKey(v string) *ListMetaEntitiesRequestCustomAttributeFilters {
	s.Key = &v
	return s
}

func (s *ListMetaEntitiesRequestCustomAttributeFilters) SetValues(v []*string) *ListMetaEntitiesRequestCustomAttributeFilters {
	s.Values = v
	return s
}

func (s *ListMetaEntitiesRequestCustomAttributeFilters) Validate() error {
	return dara.Validate(s)
}
