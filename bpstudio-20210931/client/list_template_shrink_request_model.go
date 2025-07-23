// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListTemplateShrinkRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListTemplateShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v int32) *ListTemplateShrinkRequest
	GetNextToken() *int32
	SetOrderType(v int64) *ListTemplateShrinkRequest
	GetOrderType() *int64
	SetResourceGroupId(v string) *ListTemplateShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *ListTemplateShrinkRequest
	GetTagShrink() *string
	SetTagList(v int32) *ListTemplateShrinkRequest
	GetTagList() *int32
	SetType(v string) *ListTemplateShrinkRequest
	GetType() *string
}

type ListTemplateShrinkRequest struct {
	// The keyword that is used to search for templates.
	//
	// example:
	//
	// cadt
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The criterion by which the returned templates are sorted. Valid values:
	//
	// 	- 1: The templates are sorted by the time when they are updated.
	//
	// 	- 2: The templates are sorted by the time when they are created.
	//
	// 	- 3: The templates are sorted by the system.
	//
	// 	- 4: The templates are sorted by the number of times that they are used.
	//
	// 	- If you specify an integer other than 1, 2, 3, and 4 or do not specify any value, the templates are sorted by the system.
	//
	// example:
	//
	// 1
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagShrink       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// Deprecated
	//
	// The tag that you want to use to query templates.
	//
	// example:
	//
	// 1
	TagList *int32 `json:"TagList,omitempty" xml:"TagList,omitempty"`
	// The type of the templates to be returned. Valid values: public and private
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTemplateShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateShrinkRequest) GetNextToken() *int32 {
	return s.NextToken
}

func (s *ListTemplateShrinkRequest) GetOrderType() *int64 {
	return s.OrderType
}

func (s *ListTemplateShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplateShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListTemplateShrinkRequest) GetTagList() *int32 {
	return s.TagList
}

func (s *ListTemplateShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListTemplateShrinkRequest) SetKeyword(v string) *ListTemplateShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetMaxResults(v int32) *ListTemplateShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetNextToken(v int32) *ListTemplateShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetOrderType(v int64) *ListTemplateShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetResourceGroupId(v string) *ListTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetTagShrink(v string) *ListTemplateShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetTagList(v int32) *ListTemplateShrinkRequest {
	s.TagList = &v
	return s
}

func (s *ListTemplateShrinkRequest) SetType(v string) *ListTemplateShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
