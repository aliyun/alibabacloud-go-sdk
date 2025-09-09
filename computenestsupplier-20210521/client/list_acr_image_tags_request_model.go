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
	// The type of the deployment package. Default value: AcrImage. Valid values:
	//
	// 	- HelmChart: Helm chart image.
	//
	// 	- AcrImage: container image.
	//
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The image ID.
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
