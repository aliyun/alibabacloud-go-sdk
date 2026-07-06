// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStoragesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *ListAgentStoragesRequest
	GetAgentStorageName() *string
	SetAgentStorageNameList(v []*string) *ListAgentStoragesRequest
	GetAgentStorageNameList() []*string
	SetMaxResults(v int32) *ListAgentStoragesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentStoragesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListAgentStoragesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListAgentStoragesRequest
	GetStatus() *string
	SetTag(v []*ListAgentStoragesRequestTag) *ListAgentStoragesRequest
	GetTag() []*ListAgentStoragesRequestTag
}

type ListAgentStoragesRequest struct {
	// The name of the agent storage.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The list of agent storage names. Use this parameter to query multiple specified agent storages in a batch.
	AgentStorageNameList []*string `json:"AgentStorageNameList,omitempty" xml:"AgentStorageNameList,omitempty" type:"Repeated"`
	// The maximum number of tag resources to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for querying subsequent pages. This parameter has a value only when not all tag resources are returned. If the total number of expected tag resources exceeds the MaxResults value, use this token to retrieve the next page.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group. You can query the ID in the Resource Group console.
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
	Tag []*ListAgentStoragesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAgentStoragesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStoragesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentStoragesRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *ListAgentStoragesRequest) GetAgentStorageNameList() []*string {
	return s.AgentStorageNameList
}

func (s *ListAgentStoragesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentStoragesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentStoragesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAgentStoragesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAgentStoragesRequest) GetTag() []*ListAgentStoragesRequestTag {
	return s.Tag
}

func (s *ListAgentStoragesRequest) SetAgentStorageName(v string) *ListAgentStoragesRequest {
	s.AgentStorageName = &v
	return s
}

func (s *ListAgentStoragesRequest) SetAgentStorageNameList(v []*string) *ListAgentStoragesRequest {
	s.AgentStorageNameList = v
	return s
}

func (s *ListAgentStoragesRequest) SetMaxResults(v int32) *ListAgentStoragesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentStoragesRequest) SetNextToken(v string) *ListAgentStoragesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentStoragesRequest) SetResourceGroupId(v string) *ListAgentStoragesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAgentStoragesRequest) SetStatus(v string) *ListAgentStoragesRequest {
	s.Status = &v
	return s
}

func (s *ListAgentStoragesRequest) SetTag(v []*ListAgentStoragesRequestTag) *ListAgentStoragesRequest {
	s.Tag = v
	return s
}

func (s *ListAgentStoragesRequest) Validate() error {
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

type ListAgentStoragesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Tester
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAgentStoragesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStoragesRequestTag) GoString() string {
	return s.String()
}

func (s *ListAgentStoragesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAgentStoragesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAgentStoragesRequestTag) SetKey(v string) *ListAgentStoragesRequestTag {
	s.Key = &v
	return s
}

func (s *ListAgentStoragesRequestTag) SetValue(v string) *ListAgentStoragesRequestTag {
	s.Value = &v
	return s
}

func (s *ListAgentStoragesRequestTag) Validate() error {
	return dara.Validate(s)
}
