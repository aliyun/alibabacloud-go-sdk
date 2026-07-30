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
	SetLabels(v map[string][]*string) *ListDatasetsRequest
	GetLabels() map[string][]*string
	SetMaxResults(v int32) *ListDatasetsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetsRequest
	GetNextToken() *string
}

type ListDatasetsRequest struct {
	// The dataset name.
	//
	// example:
	//
	// product_faq_dataset
	DatasetName *string              `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	Labels      map[string][]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. You do not need to set this parameter for the first request. For subsequent requests, set this parameter to the nextToken value returned in the previous response.
	//
	// example:
	//
	// RsfoUqpOJd5nd0F1e4OquY/7dKNGp1JMgsKtvCagmtY=
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

func (s *ListDatasetsRequest) GetLabels() map[string][]*string {
	return s.Labels
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

func (s *ListDatasetsRequest) SetLabels(v map[string][]*string) *ListDatasetsRequest {
	s.Labels = v
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
