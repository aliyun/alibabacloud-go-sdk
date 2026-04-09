// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeDomainsRequest
	GetAccountId() *string
	SetPageNumber(v int64) *DescribeDomainsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainsRequest
	GetPageSize() *int64
}

type DescribeDomainsRequest struct {
	// example:
	//
	// 123456
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeDomainsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainsRequest) SetAccountId(v string) *DescribeDomainsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageNumber(v int64) *DescribeDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int64) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) Validate() error {
	return dara.Validate(s)
}
