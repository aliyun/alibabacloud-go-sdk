// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListWorkspacesShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListWorkspacesShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListWorkspacesShrinkRequest
	GetRegionId() *string
	SetState(v string) *ListWorkspacesShrinkRequest
	GetState() *string
	SetTagShrink(v string) *ListWorkspacesShrinkRequest
	GetTagShrink() *string
}

type ListWorkspacesShrinkRequest struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the workspace. Fuzzy match is supported.
	//
	// example:
	//
	// test_workspace
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The state of the workspace.
	//
	// example:
	//
	// running
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListWorkspacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkspacesShrinkRequest) GetState() *string {
	return s.State
}

func (s *ListWorkspacesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListWorkspacesShrinkRequest) SetMaxResults(v int32) *ListWorkspacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetName(v string) *ListWorkspacesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetNextToken(v string) *ListWorkspacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetRegionId(v string) *ListWorkspacesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetState(v string) *ListWorkspacesShrinkRequest {
	s.State = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetTagShrink(v string) *ListWorkspacesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
