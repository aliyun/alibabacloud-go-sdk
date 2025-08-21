// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *ListBucketsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListBucketsRequest
	GetPageSize() *string
	SetPrefix(v string) *ListBucketsRequest
	GetPrefix() *string
}

type ListBucketsRequest struct {
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of returned buckets. You can leave this parameter empty. The default value is 10. The value cannot be greater than 100.
	//
	// example:
	//
	// 5
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The prefix that returned bucket names must contain. If this parameter is not specified, prefix information will not be used as a filter.
	//
	// example:
	//
	// image
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListBucketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListBucketsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListBucketsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListBucketsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListBucketsRequest) SetPageNumber(v string) *ListBucketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBucketsRequest) SetPageSize(v string) *ListBucketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBucketsRequest) SetPrefix(v string) *ListBucketsRequest {
	s.Prefix = &v
	return s
}

func (s *ListBucketsRequest) Validate() error {
	return dara.Validate(s)
}
