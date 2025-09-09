// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ListArtifactVersionsRequest
	GetArtifactId() *string
	SetFilters(v []*ListArtifactVersionsRequestFilters) *ListArtifactVersionsRequest
	GetFilters() []*ListArtifactVersionsRequestFilters
	SetMaxResults(v int32) *ListArtifactVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactVersionsRequest
	GetNextToken() *string
}

type ListArtifactVersionsRequest struct {
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The filter.
	Filters []*ListArtifactVersionsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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

func (s ListArtifactVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListArtifactVersionsRequest) GetFilters() []*ListArtifactVersionsRequestFilters {
	return s.Filters
}

func (s *ListArtifactVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactVersionsRequest) SetArtifactId(v string) *ListArtifactVersionsRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetFilters(v []*ListArtifactVersionsRequestFilters) *ListArtifactVersionsRequest {
	s.Filters = v
	return s
}

func (s *ListArtifactVersionsRequest) SetMaxResults(v int32) *ListArtifactVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetNextToken(v string) *ListArtifactVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListArtifactVersionsRequest) Validate() error {
	return dara.Validate(s)
}

type ListArtifactVersionsRequestFilters struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// **Status**：The artifact status
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListArtifactVersionsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactVersionsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsRequestFilters) GetName() *string {
	return s.Name
}

func (s *ListArtifactVersionsRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *ListArtifactVersionsRequestFilters) SetName(v string) *ListArtifactVersionsRequestFilters {
	s.Name = &v
	return s
}

func (s *ListArtifactVersionsRequestFilters) SetValues(v []*string) *ListArtifactVersionsRequestFilters {
	s.Values = v
	return s
}

func (s *ListArtifactVersionsRequestFilters) Validate() error {
	return dara.Validate(s)
}
