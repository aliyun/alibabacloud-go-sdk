// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListBackupsRequestFilter) *ListBackupsRequest
	GetFilter() []*ListBackupsRequestFilter
	SetMaxResults(v int32) *ListBackupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBackupsRequest
	GetNextToken() *string
}

type ListBackupsRequest struct {
	// The filter.
	Filter []*ListBackupsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLYoaeZA6xVdkCrmG9EmGshtmECUGpq9Qm7x5vQkpz9NXH0XzUc9t4Kxaf3UtuPY4a0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsRequest) GoString() string {
	return s.String()
}

func (s *ListBackupsRequest) GetFilter() []*ListBackupsRequestFilter {
	return s.Filter
}

func (s *ListBackupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBackupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBackupsRequest) SetFilter(v []*ListBackupsRequestFilter) *ListBackupsRequest {
	s.Filter = v
	return s
}

func (s *ListBackupsRequest) SetMaxResults(v int32) *ListBackupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBackupsRequest) SetNextToken(v string) *ListBackupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBackupsRequest) Validate() error {
	return dara.Validate(s)
}

type ListBackupsRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- BackupId: the ID of the backup.
	//
	// 	- ServiceInstanceId: The ID of the service instance.
	//
	// 	- Status: the state of the service.
	//
	// 	- StartTime
	//
	// 	- EndTime
	//
	// example:
	//
	// ServiceInstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filters.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListBackupsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListBackupsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListBackupsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListBackupsRequestFilter) SetName(v string) *ListBackupsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListBackupsRequestFilter) SetValue(v []*string) *ListBackupsRequestFilter {
	s.Value = v
	return s
}

func (s *ListBackupsRequestFilter) Validate() error {
	return dara.Validate(s)
}
