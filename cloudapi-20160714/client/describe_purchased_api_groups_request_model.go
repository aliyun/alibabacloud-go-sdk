// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedApiGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePurchasedApiGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePurchasedApiGroupsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribePurchasedApiGroupsRequest
	GetSecurityToken() *string
}

type DescribePurchasedApiGroupsRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePurchasedApiGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApiGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePurchasedApiGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePurchasedApiGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePurchasedApiGroupsRequest) SetPageNumber(v int32) *DescribePurchasedApiGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApiGroupsRequest) SetPageSize(v int32) *DescribePurchasedApiGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApiGroupsRequest) SetSecurityToken(v string) *DescribePurchasedApiGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePurchasedApiGroupsRequest) Validate() error {
	return dara.Validate(s)
}
