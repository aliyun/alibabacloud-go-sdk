// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *ListDatasetsRequest
	GetDatasetName() *string
	SetMaxResults(v int32) *ListDatasetsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetsRequest
	GetNextToken() *string
}

type ListDatasetsRequest struct {
	// example:
	//
	// test_dataset
	DatasetName *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// d9a48d977f45aa6fcf6981ed13b885b3fab0b124c12dcbbe70edce5d81ba****************
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetsRequest) SetDatasetName(v string) *ListDatasetsRequest {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetsRequest) SetMaxResults(v int32) *ListDatasetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetsRequest) SetNextToken(v string) *ListDatasetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsRequest) Validate() error {
	return dara.Validate(s)
}
