// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSDGRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSDGRequest
	GetPageSize() *int32
	SetSDGIds(v []*string) *DescribeSDGRequest
	GetSDGIds() []*string
}

type DescribeSDGRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of SDGs that you want to query. By default, all SDGs are queried.
	SDGIds []*string `json:"SDGIds,omitempty" xml:"SDGIds,omitempty" type:"Repeated"`
}

func (s DescribeSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDGRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSDGRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSDGRequest) GetSDGIds() []*string {
	return s.SDGIds
}

func (s *DescribeSDGRequest) SetPageNumber(v int32) *DescribeSDGRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGRequest) SetPageSize(v int32) *DescribeSDGRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGRequest) SetSDGIds(v []*string) *DescribeSDGRequest {
	s.SDGIds = v
	return s
}

func (s *DescribeSDGRequest) Validate() error {
	return dara.Validate(s)
}
