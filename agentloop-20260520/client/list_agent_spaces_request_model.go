// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *ListAgentSpacesRequest
	GetAgentSpace() *string
	SetMaxResults(v int32) *ListAgentSpacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentSpacesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListAgentSpacesRequest
	GetRegionId() *string
	SetTag(v []*ListAgentSpacesRequestTag) *ListAgentSpacesRequest
	GetTag() []*ListAgentSpacesRequestTag
}

type ListAgentSpacesRequest struct {
	// The name of the AgentSpace.
	//
	// example:
	//
	// test-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The maximum number of results.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// pEL20OGYeZQez8NdW7ve
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Filters only active AgentSpaces in the current region. Multiple conditions use the AND operator, and keys must be unique. Omitting the value matches the existence of the key, and an empty character string is an exact match for an empty value. An empty array is equivalent to not using label filtering and retains the original cross-region resources.
	//
	// example:
	//
	// [{\\"key\\": \\"env\\", \\"value\\": \\"test\\"}]
	Tag []*ListAgentSpacesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListAgentSpacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListAgentSpacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSpacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSpacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentSpacesRequest) GetTag() []*ListAgentSpacesRequestTag {
	return s.Tag
}

func (s *ListAgentSpacesRequest) SetAgentSpace(v string) *ListAgentSpacesRequest {
	s.AgentSpace = &v
	return s
}

func (s *ListAgentSpacesRequest) SetMaxResults(v int32) *ListAgentSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSpacesRequest) SetNextToken(v string) *ListAgentSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentSpacesRequest) SetRegionId(v string) *ListAgentSpacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAgentSpacesRequest) SetTag(v []*ListAgentSpacesRequestTag) *ListAgentSpacesRequest {
	s.Tag = v
	return s
}

func (s *ListAgentSpacesRequest) Validate() error {
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

type ListAgentSpacesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAgentSpacesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesRequestTag) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAgentSpacesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAgentSpacesRequestTag) SetKey(v string) *ListAgentSpacesRequestTag {
	s.Key = &v
	return s
}

func (s *ListAgentSpacesRequestTag) SetValue(v string) *ListAgentSpacesRequestTag {
	s.Value = &v
	return s
}

func (s *ListAgentSpacesRequestTag) Validate() error {
	return dara.Validate(s)
}
