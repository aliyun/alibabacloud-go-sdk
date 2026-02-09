// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPptArtifactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPptArtifactsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPptArtifactsRequest
	GetNextToken() *string
	SetQuery(v string) *ListPptArtifactsRequest
	GetQuery() *string
	SetWorkspaceId(v string) *ListPptArtifactsRequest
	GetWorkspaceId() *string
}

type ListPptArtifactsRequest struct {
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cEoBWREAXdxaOyjq/cqAbg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Query     *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// llm-az2xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPptArtifactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPptArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListPptArtifactsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPptArtifactsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPptArtifactsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListPptArtifactsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListPptArtifactsRequest) SetMaxResults(v int32) *ListPptArtifactsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPptArtifactsRequest) SetNextToken(v string) *ListPptArtifactsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPptArtifactsRequest) SetQuery(v string) *ListPptArtifactsRequest {
	s.Query = &v
	return s
}

func (s *ListPptArtifactsRequest) SetWorkspaceId(v string) *ListPptArtifactsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListPptArtifactsRequest) Validate() error {
	return dara.Validate(s)
}
