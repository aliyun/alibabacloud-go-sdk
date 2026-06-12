// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRestoreTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListRestoreTasksRequestFilter) *ListRestoreTasksRequest
	GetFilter() []*ListRestoreTasksRequestFilter
	SetMaxResults(v int32) *ListRestoreTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRestoreTasksRequest
	GetNextToken() *string
}

type ListRestoreTasksRequest struct {
	// The filter.
	Filter []*ListRestoreTasksRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListRestoreTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRestoreTasksRequest) GoString() string {
	return s.String()
}

func (s *ListRestoreTasksRequest) GetFilter() []*ListRestoreTasksRequestFilter {
	return s.Filter
}

func (s *ListRestoreTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRestoreTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRestoreTasksRequest) SetFilter(v []*ListRestoreTasksRequestFilter) *ListRestoreTasksRequest {
	s.Filter = v
	return s
}

func (s *ListRestoreTasksRequest) SetMaxResults(v int32) *ListRestoreTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRestoreTasksRequest) SetNextToken(v string) *ListRestoreTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListRestoreTasksRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRestoreTasksRequestFilter struct {
	// The name of the filter. You can query by one or more filter names. Valid values:
	//
	// - RestoreTaskId: The restore job ID.
	//
	// - ServiceInstanceId: The service instance ID.
	//
	// - Status: The status.
	//
	// - StartTime: The start time.
	//
	// - EndTime: The end time.
	//
	// example:
	//
	// ServiceInstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListRestoreTasksRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListRestoreTasksRequestFilter) GoString() string {
	return s.String()
}

func (s *ListRestoreTasksRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListRestoreTasksRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListRestoreTasksRequestFilter) SetName(v string) *ListRestoreTasksRequestFilter {
	s.Name = &v
	return s
}

func (s *ListRestoreTasksRequestFilter) SetValue(v []*string) *ListRestoreTasksRequestFilter {
	s.Value = v
	return s
}

func (s *ListRestoreTasksRequestFilter) Validate() error {
	return dara.Validate(s)
}
