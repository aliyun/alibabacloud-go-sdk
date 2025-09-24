// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ListRunMetricsRequest
	GetKey() *string
	SetMaxResults(v int64) *ListRunMetricsRequest
	GetMaxResults() *int64
	SetPageToken(v int64) *ListRunMetricsRequest
	GetPageToken() *int64
}

type ListRunMetricsRequest struct {
	// The metric key of the run.
	//
	// This parameter is required.
	//
	// example:
	//
	// loss
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The maximum number of entries in the request. Default value: 10.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token, which starts from 0. Default value: 0.
	//
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
}

func (s ListRunMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRunMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListRunMetricsRequest) GetKey() *string {
	return s.Key
}

func (s *ListRunMetricsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListRunMetricsRequest) GetPageToken() *int64 {
	return s.PageToken
}

func (s *ListRunMetricsRequest) SetKey(v string) *ListRunMetricsRequest {
	s.Key = &v
	return s
}

func (s *ListRunMetricsRequest) SetMaxResults(v int64) *ListRunMetricsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRunMetricsRequest) SetPageToken(v int64) *ListRunMetricsRequest {
	s.PageToken = &v
	return s
}

func (s *ListRunMetricsRequest) Validate() error {
	return dara.Validate(s)
}
