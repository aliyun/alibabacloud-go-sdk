// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactPath(v string) *ListArtifactsRequest
	GetArtifactPath() *string
	SetMaxResults(v int32) *ListArtifactsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactsRequest
	GetNextToken() *string
}

type ListArtifactsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// missions/mission-xxx/artifacts/2026-05/05-01/
	ArtifactPath *string `json:"artifactPath,omitempty" xml:"artifactPath,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListArtifactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequest) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *ListArtifactsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactsRequest) SetArtifactPath(v string) *ListArtifactsRequest {
	s.ArtifactPath = &v
	return s
}

func (s *ListArtifactsRequest) SetMaxResults(v int32) *ListArtifactsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsRequest) SetNextToken(v string) *ListArtifactsRequest {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsRequest) Validate() error {
	return dara.Validate(s)
}
