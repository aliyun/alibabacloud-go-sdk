// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggTaskGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterAggTaskGroupIds(v string) *ListAggTaskGroupsRequest
	GetFilterAggTaskGroupIds() *string
	SetFilterAggTaskGroupNames(v string) *ListAggTaskGroupsRequest
	GetFilterAggTaskGroupNames() *string
	SetMaxResults(v int32) *ListAggTaskGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggTaskGroupsRequest
	GetNextToken() *string
	SetQuery(v string) *ListAggTaskGroupsRequest
	GetQuery() *string
	SetStatus(v string) *ListAggTaskGroupsRequest
	GetStatus() *string
	SetTags(v []*ListAggTaskGroupsRequestTags) *ListAggTaskGroupsRequest
	GetTags() []*ListAggTaskGroupsRequestTags
	SetTargetPrometheusId(v string) *ListAggTaskGroupsRequest
	GetTargetPrometheusId() *string
}

type ListAggTaskGroupsRequest struct {
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
	Status *string                         `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListAggTaskGroupsRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// rw-pq4apob9jm
	TargetPrometheusId *string `json:"targetPrometheusId,omitempty" xml:"targetPrometheusId,omitempty"`
}

func (s ListAggTaskGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggTaskGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAggTaskGroupsRequest) GetFilterAggTaskGroupIds() *string {
	return s.FilterAggTaskGroupIds
}

func (s *ListAggTaskGroupsRequest) GetFilterAggTaskGroupNames() *string {
	return s.FilterAggTaskGroupNames
}

func (s *ListAggTaskGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggTaskGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggTaskGroupsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListAggTaskGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAggTaskGroupsRequest) GetTags() []*ListAggTaskGroupsRequestTags {
	return s.Tags
}

func (s *ListAggTaskGroupsRequest) GetTargetPrometheusId() *string {
	return s.TargetPrometheusId
}

func (s *ListAggTaskGroupsRequest) SetFilterAggTaskGroupIds(v string) *ListAggTaskGroupsRequest {
	s.FilterAggTaskGroupIds = &v
	return s
}

func (s *ListAggTaskGroupsRequest) SetFilterAggTaskGroupNames(v string) *ListAggTaskGroupsRequest {
	s.FilterAggTaskGroupNames = &v
	return s
}

func (s *ListAggTaskGroupsRequest) SetMaxResults(v int32) *ListAggTaskGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggTaskGroupsRequest) SetNextToken(v string) *ListAggTaskGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggTaskGroupsRequest) SetQuery(v string) *ListAggTaskGroupsRequest {
	s.Query = &v
	return s
}

func (s *ListAggTaskGroupsRequest) SetStatus(v string) *ListAggTaskGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListAggTaskGroupsRequest) SetTags(v []*ListAggTaskGroupsRequestTags) *ListAggTaskGroupsRequest {
	s.Tags = v
	return s
}

func (s *ListAggTaskGroupsRequest) SetTargetPrometheusId(v string) *ListAggTaskGroupsRequest {
	s.TargetPrometheusId = &v
	return s
}

func (s *ListAggTaskGroupsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggTaskGroupsRequestTags struct {
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAggTaskGroupsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListAggTaskGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *ListAggTaskGroupsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListAggTaskGroupsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListAggTaskGroupsRequestTags) SetKey(v string) *ListAggTaskGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *ListAggTaskGroupsRequestTags) SetValue(v string) *ListAggTaskGroupsRequestTags {
	s.Value = &v
	return s
}

func (s *ListAggTaskGroupsRequestTags) Validate() error {
	return dara.Validate(s)
}
