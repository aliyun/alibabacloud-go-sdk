// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *ListDatasetsShrinkRequest
	GetDatasetName() *string
	SetLabelsShrink(v string) *ListDatasetsShrinkRequest
	GetLabelsShrink() *string
	SetMaxResults(v int32) *ListDatasetsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetsShrinkRequest
	GetNextToken() *string
}

type ListDatasetsShrinkRequest struct {
	// The dataset name.
	//
	// example:
	//
	// product_faq_dataset
	DatasetName  *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	LabelsShrink *string `json:"labels,omitempty" xml:"labels,omitempty"`
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

func (s ListDatasetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetsShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *ListDatasetsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetsShrinkRequest) SetDatasetName(v string) *ListDatasetsShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetLabelsShrink(v string) *ListDatasetsShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetMaxResults(v int32) *ListDatasetsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetNextToken(v string) *ListDatasetsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
