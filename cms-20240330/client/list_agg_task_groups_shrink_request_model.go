// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggTaskGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterAggTaskGroupIds(v string) *ListAggTaskGroupsShrinkRequest
	GetFilterAggTaskGroupIds() *string
	SetFilterAggTaskGroupNames(v string) *ListAggTaskGroupsShrinkRequest
	GetFilterAggTaskGroupNames() *string
	SetMaxResults(v int32) *ListAggTaskGroupsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggTaskGroupsShrinkRequest
	GetNextToken() *string
	SetQuery(v string) *ListAggTaskGroupsShrinkRequest
	GetQuery() *string
	SetStatus(v string) *ListAggTaskGroupsShrinkRequest
	GetStatus() *string
	SetTagsShrink(v string) *ListAggTaskGroupsShrinkRequest
	GetTagsShrink() *string
	SetTargetPrometheusId(v string) *ListAggTaskGroupsShrinkRequest
	GetTargetPrometheusId() *string
}

type ListAggTaskGroupsShrinkRequest struct {
	// example:
	//
	// ["aggTaskGroup-xxx"]
	FilterAggTaskGroupIds *string `json:"filterAggTaskGroupIds,omitempty" xml:"filterAggTaskGroupIds,omitempty"`
	// example:
	//
	// ["apiserver_request_total"]
	FilterAggTaskGroupNames *string `json:"filterAggTaskGroupNames,omitempty" xml:"filterAggTaskGroupNames,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 28036394xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// Running
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// rw-pq4apob9jm
	TargetPrometheusId *string `json:"targetPrometheusId,omitempty" xml:"targetPrometheusId,omitempty"`
}

func (s ListAggTaskGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggTaskGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAggTaskGroupsShrinkRequest) GetFilterAggTaskGroupIds() *string {
	return s.FilterAggTaskGroupIds
}

func (s *ListAggTaskGroupsShrinkRequest) GetFilterAggTaskGroupNames() *string {
	return s.FilterAggTaskGroupNames
}

func (s *ListAggTaskGroupsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggTaskGroupsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggTaskGroupsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListAggTaskGroupsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAggTaskGroupsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListAggTaskGroupsShrinkRequest) GetTargetPrometheusId() *string {
	return s.TargetPrometheusId
}

func (s *ListAggTaskGroupsShrinkRequest) SetFilterAggTaskGroupIds(v string) *ListAggTaskGroupsShrinkRequest {
	s.FilterAggTaskGroupIds = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) SetFilterAggTaskGroupNames(v string) *ListAggTaskGroupsShrinkRequest {
	s.FilterAggTaskGroupNames = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) SetMaxResults(v int32) *ListAggTaskGroupsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) SetNextToken(v string) *ListAggTaskGroupsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) SetQuery(v string) *ListAggTaskGroupsShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) SetStatus(v string) *ListAggTaskGroupsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) SetTagsShrink(v string) *ListAggTaskGroupsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) SetTargetPrometheusId(v string) *ListAggTaskGroupsShrinkRequest {
	s.TargetPrometheusId = &v
	return s
}

func (s *ListAggTaskGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
