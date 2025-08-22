// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRDDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRDDomainsRequest
	GetPageSize() *int32
}

type DescribeRDDomainsRequest struct {
	// The number of the page to return. Valid values: 1 to 100000.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The default value is 20. Valid values: an integer between 1 and 500. Default value: 20.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRDDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRDDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRDDomainsRequest) SetPageNumber(v int32) *DescribeRDDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRDDomainsRequest) SetPageSize(v int32) *DescribeRDDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRDDomainsRequest) Validate() error {
	return dara.Validate(s)
}
