// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesRequest
	GetMaxResults() *int32
	SetName(v string) *ListWorkspacesRequest
	GetName() *string
	SetNextToken(v string) *ListWorkspacesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListWorkspacesRequest
	GetRegionId() *string
	SetState(v string) *ListWorkspacesRequest
	GetState() *string
	SetTag(v []*ListWorkspacesRequestTag) *ListWorkspacesRequest
	GetTag() []*ListWorkspacesRequestTag
}

type ListWorkspacesRequest struct {
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
	State *string                     `json:"state,omitempty" xml:"state,omitempty"`
	Tag   []*ListWorkspacesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkspacesRequest) GetState() *string {
	return s.State
}

func (s *ListWorkspacesRequest) GetTag() []*ListWorkspacesRequestTag {
	return s.Tag
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetName(v string) *ListWorkspacesRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetRegionId(v string) *ListWorkspacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesRequest) SetState(v string) *ListWorkspacesRequest {
	s.State = &v
	return s
}

func (s *ListWorkspacesRequest) SetTag(v []*ListWorkspacesRequestTag) *ListWorkspacesRequest {
	s.Tag = v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}

type ListWorkspacesRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListWorkspacesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequestTag) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListWorkspacesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListWorkspacesRequestTag) SetKey(v string) *ListWorkspacesRequestTag {
	s.Key = &v
	return s
}

func (s *ListWorkspacesRequestTag) SetValue(v string) *ListWorkspacesRequestTag {
	s.Value = &v
	return s
}

func (s *ListWorkspacesRequestTag) Validate() error {
	return dara.Validate(s)
}
