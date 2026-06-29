// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStoragesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *ListAgentStoragesShrinkRequest
	GetAgentStorageName() *string
	SetAgentStorageNameListShrink(v string) *ListAgentStoragesShrinkRequest
	GetAgentStorageNameListShrink() *string
	SetMaxResults(v int32) *ListAgentStoragesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentStoragesShrinkRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListAgentStoragesShrinkRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListAgentStoragesShrinkRequest
	GetStatus() *string
	SetTagShrink(v string) *ListAgentStoragesShrinkRequest
	GetTagShrink() *string
}

type ListAgentStoragesShrinkRequest struct {
	// agent storage name
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The list of agent storage names, used to query multiple specified agent storages in a batch.
	AgentStorageNameListShrink *string `json:"AgentStorageNameList,omitempty" xml:"AgentStorageNameList,omitempty"`
	// The maximum number of tag resources to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results when the total number of tag resources exceeds the value of MaxResults. This parameter has a value only when not all tag resources are returned.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID. You can query this ID in the resource group console.
	//
	// example:
	//
	// rg-acfmxh4em5jncda
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the agent storage.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAgentStoragesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStoragesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAgentStoragesShrinkRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *ListAgentStoragesShrinkRequest) GetAgentStorageNameListShrink() *string {
	return s.AgentStorageNameListShrink
}

func (s *ListAgentStoragesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentStoragesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentStoragesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAgentStoragesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAgentStoragesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListAgentStoragesShrinkRequest) SetAgentStorageName(v string) *ListAgentStoragesShrinkRequest {
	s.AgentStorageName = &v
	return s
}

func (s *ListAgentStoragesShrinkRequest) SetAgentStorageNameListShrink(v string) *ListAgentStoragesShrinkRequest {
	s.AgentStorageNameListShrink = &v
	return s
}

func (s *ListAgentStoragesShrinkRequest) SetMaxResults(v int32) *ListAgentStoragesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentStoragesShrinkRequest) SetNextToken(v string) *ListAgentStoragesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentStoragesShrinkRequest) SetResourceGroupId(v string) *ListAgentStoragesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAgentStoragesShrinkRequest) SetStatus(v string) *ListAgentStoragesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListAgentStoragesShrinkRequest) SetTagShrink(v string) *ListAgentStoragesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListAgentStoragesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
