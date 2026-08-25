// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *ListServerIdeImagesRequest
	GetLabels() *string
	SetMaxResults(v int32) *ListServerIdeImagesRequest
	GetMaxResults() *int32
	SetName(v string) *ListServerIdeImagesRequest
	GetName() *string
	SetNextToken(v string) *ListServerIdeImagesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListServerIdeImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServerIdeImagesRequest
	GetPageSize() *int32
}

type ListServerIdeImagesRequest struct {
	// The image label filter condition. Separate multiple Key=Value conditions with commas.
	//
	// example:
	//
	// system.chipType=GPU,system.official=true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of records to return in a single request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The image name filter condition. Supports fuzzy match.
	//
	// example:
	//
	// tensorflow_2.9
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token used to continue a query. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// CAESG****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. The value starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServerIdeImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeImagesRequest) GoString() string {
	return s.String()
}

func (s *ListServerIdeImagesRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListServerIdeImagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerIdeImagesRequest) GetName() *string {
	return s.Name
}

func (s *ListServerIdeImagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerIdeImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServerIdeImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServerIdeImagesRequest) SetLabels(v string) *ListServerIdeImagesRequest {
	s.Labels = &v
	return s
}

func (s *ListServerIdeImagesRequest) SetMaxResults(v int32) *ListServerIdeImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerIdeImagesRequest) SetName(v string) *ListServerIdeImagesRequest {
	s.Name = &v
	return s
}

func (s *ListServerIdeImagesRequest) SetNextToken(v string) *ListServerIdeImagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerIdeImagesRequest) SetPageNumber(v int32) *ListServerIdeImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServerIdeImagesRequest) SetPageSize(v int32) *ListServerIdeImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServerIdeImagesRequest) Validate() error {
	return dara.Validate(s)
}
