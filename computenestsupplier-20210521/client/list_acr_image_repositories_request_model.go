// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcrImageRepositoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *ListAcrImageRepositoriesRequest
	GetArtifactType() *string
	SetMaxResults(v int32) *ListAcrImageRepositoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAcrImageRepositoriesRequest
	GetNextToken() *string
	SetRepoName(v string) *ListAcrImageRepositoriesRequest
	GetRepoName() *string
}

type ListAcrImageRepositoriesRequest struct {
	// The type of the artifact. The default value is AcrImage. Valid values:
	//
	// - HelmChart: a Helm chart image.
	//
	// - AcrImage: a container image.
	//
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s ListAcrImageRepositoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListAcrImageRepositoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAcrImageRepositoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAcrImageRepositoriesRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListAcrImageRepositoriesRequest) SetArtifactType(v string) *ListAcrImageRepositoriesRequest {
	s.ArtifactType = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetMaxResults(v int32) *ListAcrImageRepositoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetNextToken(v string) *ListAcrImageRepositoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetRepoName(v string) *ListAcrImageRepositoriesRequest {
	s.RepoName = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) Validate() error {
	return dara.Validate(s)
}
