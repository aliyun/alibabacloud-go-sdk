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
	// The artifact ID.
	//
	// To obtain the artifact ID, call the [ListArtifacts](https://help.aliyun.com/document_detail/469993.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43ae****
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The filter.
	Filters []*ListArtifactVersionsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. Set this to the NextToken value from the previous call.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hR****
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
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListArtifactVersionsRequestFilters struct {
	// The filter name. This parameter supports querying by one or more filter names. Valid values:
	//
	// **Status**: Filters by artifact status.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filter values.
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
