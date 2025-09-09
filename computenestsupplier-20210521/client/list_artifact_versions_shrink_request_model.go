// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactVersionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ListArtifactVersionsShrinkRequest
	GetArtifactId() *string
	SetFiltersShrink(v string) *ListArtifactVersionsShrinkRequest
	GetFiltersShrink() *string
	SetMaxResults(v int32) *ListArtifactVersionsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactVersionsShrinkRequest
	GetNextToken() *string
}

type ListArtifactVersionsShrinkRequest struct {
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The filter.
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListArtifactVersionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactVersionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsShrinkRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListArtifactVersionsShrinkRequest) GetFiltersShrink() *string {
	return s.FiltersShrink
}

func (s *ListArtifactVersionsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactVersionsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactVersionsShrinkRequest) SetArtifactId(v string) *ListArtifactVersionsShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsShrinkRequest) SetFiltersShrink(v string) *ListArtifactVersionsShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *ListArtifactVersionsShrinkRequest) SetMaxResults(v int32) *ListArtifactVersionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsShrinkRequest) SetNextToken(v string) *ListArtifactVersionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListArtifactVersionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
