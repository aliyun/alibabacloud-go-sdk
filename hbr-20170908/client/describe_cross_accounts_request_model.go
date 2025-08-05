// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeCrossAccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossAccountsRequest
	GetPageSize() *int32
}

type DescribeCrossAccountsRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCrossAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossAccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossAccountsRequest) SetPageNumber(v int32) *DescribeCrossAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossAccountsRequest) SetPageSize(v int32) *DescribeCrossAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossAccountsRequest) Validate() error {
	return dara.Validate(s)
}
