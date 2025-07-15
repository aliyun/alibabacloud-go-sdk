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
	Filter *ListPoolsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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
	return dara.Validate(s)
}

type ListPoolsRequestFilter struct {
	PoolName []*string `json:"PoolName,omitempty" xml:"PoolName,omitempty" type:"Repeated"`
	Status   []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
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
