// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *ListPoolsRequestFilter) *ListPoolsRequest
	GetFilter() *ListPoolsRequestFilter
	SetPageNumber(v int32) *ListPoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPoolsRequest
	GetPageSize() *int32
}

type ListPoolsRequest struct {
	// The filter conditions for querying resource pools.
	Filter *ListPoolsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListPoolsRequest) GetFilter() *ListPoolsRequestFilter {
	return s.Filter
}

func (s *ListPoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoolsRequest) SetFilter(v *ListPoolsRequestFilter) *ListPoolsRequest {
	s.Filter = v
	return s
}

func (s *ListPoolsRequest) SetPageNumber(v int32) *ListPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoolsRequest) SetPageSize(v int32) *ListPoolsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPoolsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoolsRequestFilter struct {
	// The list of resource pool names to query.
	PoolName []*string `json:"PoolName,omitempty" xml:"PoolName,omitempty" type:"Repeated"`
	// The list of resource pool statuses to query.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// Returns only resource pools created after the specified time. The time must be a Unix timestamp in UTC+8.
	//
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// Returns only resource pools created before the specified time. The time must be a Unix timestamp in UTC+8.
	//
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32 `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
}

func (s ListPoolsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListPoolsRequestFilter) GetPoolName() []*string {
	return s.PoolName
}

func (s *ListPoolsRequestFilter) GetStatus() []*string {
	return s.Status
}

func (s *ListPoolsRequestFilter) GetTimeCreatedAfter() *int32 {
	return s.TimeCreatedAfter
}

func (s *ListPoolsRequestFilter) GetTimeCreatedBefore() *int32 {
	return s.TimeCreatedBefore
}

func (s *ListPoolsRequestFilter) SetPoolName(v []*string) *ListPoolsRequestFilter {
	s.PoolName = v
	return s
}

func (s *ListPoolsRequestFilter) SetStatus(v []*string) *ListPoolsRequestFilter {
	s.Status = v
	return s
}

func (s *ListPoolsRequestFilter) SetTimeCreatedAfter(v int32) *ListPoolsRequestFilter {
	s.TimeCreatedAfter = &v
	return s
}

func (s *ListPoolsRequestFilter) SetTimeCreatedBefore(v int32) *ListPoolsRequestFilter {
	s.TimeCreatedBefore = &v
	return s
}

func (s *ListPoolsRequestFilter) Validate() error {
	return dara.Validate(s)
}
