// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsForRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeEventsForRegionRequest
	GetClusterId() *string
	SetPageNumber(v int64) *DescribeEventsForRegionRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEventsForRegionRequest
	GetPageSize() *int64
}

type DescribeEventsForRegionRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cf62854ac2130470897be7a27ed1f****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The number of pages.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of records on each page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s DescribeEventsForRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsForRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsForRegionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeEventsForRegionRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEventsForRegionRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventsForRegionRequest) SetClusterId(v string) *DescribeEventsForRegionRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeEventsForRegionRequest) SetPageNumber(v int64) *DescribeEventsForRegionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsForRegionRequest) SetPageSize(v int64) *DescribeEventsForRegionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsForRegionRequest) Validate() error {
	return dara.Validate(s)
}
