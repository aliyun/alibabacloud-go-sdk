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
	// The list of the filters.
	Filter []*ListRestoreTasksRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of rows displayed per page in paginated queries. Maximum: 100 rows per page. Default: 20 rows.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// NextToken
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
	return dara.Validate(s)
}

type ListRestoreTasksRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- RestoreTaskId: the ID of the restore task.
	//
	// 	- ServiceInstanceId: The ID of the service instance.
	//
	// 	- Status
	//
	// 	- StartTime
	//
	// 	- EndTime
	//
	// example:
	//
	// ServiceInstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
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
