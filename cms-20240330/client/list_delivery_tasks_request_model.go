// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWords(v string) *ListDeliveryTasksRequest
	GetKeyWords() *string
	SetMaxResults(v int32) *ListDeliveryTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeliveryTasksRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListDeliveryTasksRequest
	GetResourceGroupId() *string
	SetTag(v []*ListDeliveryTasksRequestTag) *ListDeliveryTasksRequest
	GetTag() []*ListDeliveryTasksRequestTag
}

type ListDeliveryTasksRequest struct {
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
	Tag []*ListDeliveryTasksRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListDeliveryTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksRequest) GetKeyWords() *string {
	return s.KeyWords
}

func (s *ListDeliveryTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeliveryTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeliveryTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDeliveryTasksRequest) GetTag() []*ListDeliveryTasksRequestTag {
	return s.Tag
}

func (s *ListDeliveryTasksRequest) SetKeyWords(v string) *ListDeliveryTasksRequest {
	s.KeyWords = &v
	return s
}

func (s *ListDeliveryTasksRequest) SetMaxResults(v int32) *ListDeliveryTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeliveryTasksRequest) SetNextToken(v string) *ListDeliveryTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeliveryTasksRequest) SetResourceGroupId(v string) *ListDeliveryTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDeliveryTasksRequest) SetTag(v []*ListDeliveryTasksRequestTag) *ListDeliveryTasksRequest {
	s.Tag = v
	return s
}

func (s *ListDeliveryTasksRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeliveryTasksRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListDeliveryTasksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksRequestTag) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListDeliveryTasksRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListDeliveryTasksRequestTag) SetKey(v string) *ListDeliveryTasksRequestTag {
	s.Key = &v
	return s
}

func (s *ListDeliveryTasksRequestTag) SetValue(v string) *ListDeliveryTasksRequestTag {
	s.Value = &v
	return s
}

func (s *ListDeliveryTasksRequestTag) Validate() error {
	return dara.Validate(s)
}
