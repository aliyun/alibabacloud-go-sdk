// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListConsumerGroupsRequest
	GetFilter() *string
	SetPageNumber(v int32) *ListConsumerGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumerGroupsRequest
	GetPageSize() *int32
}

type ListConsumerGroupsRequest struct {
	// The filter condition for the query. If not provided, all consumer groups under the specified instance will be queried.
	//
	// example:
	//
	// CID-TEST
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// Page number, indicating which page of results to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, the maximum number of results to display per page.
	//
	// Value range: [10, 100].
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConsumerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListConsumerGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerGroupsRequest) SetFilter(v string) *ListConsumerGroupsRequest {
	s.Filter = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageNumber(v int32) *ListConsumerGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageSize(v int32) *ListConsumerGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupsRequest) Validate() error {
	return dara.Validate(s)
}
