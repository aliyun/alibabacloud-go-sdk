// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListTemplateRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListTemplateRequest
	GetMaxResults() *int32
	SetNextToken(v int32) *ListTemplateRequest
	GetNextToken() *int32
	SetOrderType(v int64) *ListTemplateRequest
	GetOrderType() *int64
	SetResourceGroupId(v string) *ListTemplateRequest
	GetResourceGroupId() *string
	SetTag(v []*ListTemplateRequestTag) *ListTemplateRequest
	GetTag() []*ListTemplateRequestTag
	SetTagList(v int32) *ListTemplateRequest
	GetTagList() *int32
	SetType(v string) *ListTemplateRequest
	GetType() *string
}

type ListTemplateRequest struct {
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
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListTemplateRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s ListTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTemplateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateRequest) GetNextToken() *int32 {
	return s.NextToken
}

func (s *ListTemplateRequest) GetOrderType() *int64 {
	return s.OrderType
}

func (s *ListTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplateRequest) GetTag() []*ListTemplateRequestTag {
	return s.Tag
}

func (s *ListTemplateRequest) GetTagList() *int32 {
	return s.TagList
}

func (s *ListTemplateRequest) GetType() *string {
	return s.Type
}

func (s *ListTemplateRequest) SetKeyword(v string) *ListTemplateRequest {
	s.Keyword = &v
	return s
}

func (s *ListTemplateRequest) SetMaxResults(v int32) *ListTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateRequest) SetNextToken(v int32) *ListTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateRequest) SetOrderType(v int64) *ListTemplateRequest {
	s.OrderType = &v
	return s
}

func (s *ListTemplateRequest) SetResourceGroupId(v string) *ListTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateRequest) SetTag(v []*ListTemplateRequestTag) *ListTemplateRequest {
	s.Tag = v
	return s
}

func (s *ListTemplateRequest) SetTagList(v int32) *ListTemplateRequest {
	s.TagList = &v
	return s
}

func (s *ListTemplateRequest) SetType(v string) *ListTemplateRequest {
	s.Type = &v
	return s
}

func (s *ListTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type ListTemplateRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplateRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateRequestTag) GoString() string {
	return s.String()
}

func (s *ListTemplateRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTemplateRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTemplateRequestTag) SetKey(v string) *ListTemplateRequestTag {
	s.Key = &v
	return s
}

func (s *ListTemplateRequestTag) SetValue(v string) *ListTemplateRequestTag {
	s.Value = &v
	return s
}

func (s *ListTemplateRequestTag) Validate() error {
	return dara.Validate(s)
}
