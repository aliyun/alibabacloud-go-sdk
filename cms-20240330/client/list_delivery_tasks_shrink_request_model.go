// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWords(v string) *ListDeliveryTasksShrinkRequest
	GetKeyWords() *string
	SetMaxResults(v int32) *ListDeliveryTasksShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeliveryTasksShrinkRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListDeliveryTasksShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *ListDeliveryTasksShrinkRequest
	GetTagShrink() *string
}

type ListDeliveryTasksShrinkRequest struct {
	// The keyword used to perform a fuzzy match by task name or task ID.
	//
	// example:
	//
	// test-task
	KeyWords *string `json:"keyWords,omitempty" xml:"keyWords,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Leave this parameter empty for the first request. For subsequent requests, set this parameter to the nextToken value returned in the previous response.
	//
	// example:
	//
	// sjC5rekx93Ew7K7VcmI3wkBZBYQ-GphB2ilQu3zJCGxoZuicwyJznfo2riTjr-lq
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bhocin5e2na
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags.
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListDeliveryTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksShrinkRequest) GetKeyWords() *string {
	return s.KeyWords
}

func (s *ListDeliveryTasksShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeliveryTasksShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeliveryTasksShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDeliveryTasksShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListDeliveryTasksShrinkRequest) SetKeyWords(v string) *ListDeliveryTasksShrinkRequest {
	s.KeyWords = &v
	return s
}

func (s *ListDeliveryTasksShrinkRequest) SetMaxResults(v int32) *ListDeliveryTasksShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeliveryTasksShrinkRequest) SetNextToken(v string) *ListDeliveryTasksShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeliveryTasksShrinkRequest) SetResourceGroupId(v string) *ListDeliveryTasksShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDeliveryTasksShrinkRequest) SetTagShrink(v string) *ListDeliveryTasksShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListDeliveryTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
