// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *ListAgentSpacesShrinkRequest
	GetAgentSpace() *string
	SetMaxResults(v int32) *ListAgentSpacesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentSpacesShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListAgentSpacesShrinkRequest
	GetRegionId() *string
	SetTagShrink(v string) *ListAgentSpacesShrinkRequest
	GetTagShrink() *string
}

type ListAgentSpacesShrinkRequest struct {
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
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListAgentSpacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesShrinkRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListAgentSpacesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSpacesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSpacesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentSpacesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListAgentSpacesShrinkRequest) SetAgentSpace(v string) *ListAgentSpacesShrinkRequest {
	s.AgentSpace = &v
	return s
}

func (s *ListAgentSpacesShrinkRequest) SetMaxResults(v int32) *ListAgentSpacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSpacesShrinkRequest) SetNextToken(v string) *ListAgentSpacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentSpacesShrinkRequest) SetRegionId(v string) *ListAgentSpacesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListAgentSpacesShrinkRequest) SetTagShrink(v string) *ListAgentSpacesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListAgentSpacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
