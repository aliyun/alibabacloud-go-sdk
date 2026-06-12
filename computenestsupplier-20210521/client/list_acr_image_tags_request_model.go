// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcrImageTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *ListAcrImageTagsRequest
	GetArtifactType() *string
	SetMaxResults(v int32) *ListAcrImageTagsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAcrImageTagsRequest
	GetNextToken() *string
	SetRepoId(v string) *ListAcrImageTagsRequest
	GetRepoId() *string
}

type ListAcrImageTagsRequest struct {
	// The artifact type. The default value is AcrImage. Possible values:
	//
	// - HelmChart: A Helm Chart image.
	//
	// - AcrImage: A container image.
	//
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-3gqhkza0wbxxxxxx
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListAcrImageTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageTagsRequest) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListAcrImageTagsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAcrImageTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAcrImageTagsRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListAcrImageTagsRequest) SetArtifactType(v string) *ListAcrImageTagsRequest {
	s.ArtifactType = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetMaxResults(v int32) *ListAcrImageTagsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetNextToken(v string) *ListAcrImageTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetRepoId(v string) *ListAcrImageTagsRequest {
	s.RepoId = &v
	return s
}

func (s *ListAcrImageTagsRequest) Validate() error {
	return dara.Validate(s)
}
