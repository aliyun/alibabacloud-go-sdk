// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListModelsRequest
	GetMaxResults() *int32
	SetModelType(v string) *ListModelsRequest
	GetModelType() *string
	SetNextToken(v string) *ListModelsRequest
	GetNextToken() *string
	SetPageNumber(v int64) *ListModelsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListModelsRequest
	GetPageSize() *int64
}

type ListModelsRequest struct {
	// The maximum number of models to return per page. Valid values: 1 to 100. Explicitly specify this parameter.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Set to `text` for text sessions. If omitted, all visible models are queried.
	//
	// example:
	//
	// text
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The pagination token returned from the previous page. Omit this parameter for the first page.
	//
	// example:
	//
	// next-token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The legacy pagination parameter for compatibility. The page number starts from 1. Use MaxResults and NextToken instead. Do not mix the two sets of pagination parameters.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The legacy pagination parameter for compatibility. The number of entries per page. Valid values: 1 to 100. Use MaxResults and NextToken instead. Do not mix the two sets of pagination parameters.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelsRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListModelsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListModelsRequest) SetMaxResults(v int32) *ListModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelsRequest) SetModelType(v string) *ListModelsRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelsRequest) SetNextToken(v string) *ListModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelsRequest) SetPageNumber(v int64) *ListModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsRequest) SetPageSize(v int64) *ListModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsRequest) Validate() error {
	return dara.Validate(s)
}
