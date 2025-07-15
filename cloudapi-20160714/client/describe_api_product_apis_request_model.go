// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiProductApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductId(v string) *DescribeApiProductApisRequest
	GetApiProductId() *string
	SetPageNumber(v int32) *DescribeApiProductApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiProductApisRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApiProductApisRequest
	GetSecurityToken() *string
}

type DescribeApiProductApisRequest struct {
	// The ID of the API product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 117b7a64a8b3f064eaa4a47ac62aac5e
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiProductApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiProductApisRequest) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *DescribeApiProductApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiProductApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiProductApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiProductApisRequest) SetApiProductId(v string) *DescribeApiProductApisRequest {
	s.ApiProductId = &v
	return s
}

func (s *DescribeApiProductApisRequest) SetPageNumber(v int32) *DescribeApiProductApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiProductApisRequest) SetPageSize(v int32) *DescribeApiProductApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiProductApisRequest) SetSecurityToken(v string) *DescribeApiProductApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiProductApisRequest) Validate() error {
	return dara.Validate(s)
}
