// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeResolverEndpointsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeResolverEndpointsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeResolverEndpointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResolverEndpointsRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeResolverEndpointsRequest
	GetStatus() *string
	SetVpcRegionId(v string) *DescribeResolverEndpointsRequest
	GetVpcRegionId() *string
}

type DescribeResolverEndpointsRequest struct {
	// The keyword of the endpoint name, which is used for fuzzy searches.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The state of the endpoint that you want to query. Valid values:
	//
	// 	- SUCCESS: The endpoint works as expected.
	//
	// 	- INIT: The endpoint is being created.
	//
	// 	- FAILED: The endpoint failed to be created.
	//
	// 	- CHANGE_INIT: The endpoint is being modified.
	//
	// 	- CHANGE_FAILED: The endpoint failed to be modified.
	//
	// 	- EXCEPTION: The endpoint encountered an exception.
	//
	// >  If you do not specify this parameter, endpoints in all states are returned.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The region ID of the outbound virtual private cloud (VPC).
	//
	// example:
	//
	// cn-zhangjiakou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s DescribeResolverEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeResolverEndpointsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeResolverEndpointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResolverEndpointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResolverEndpointsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeResolverEndpointsRequest) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *DescribeResolverEndpointsRequest) SetKeyword(v string) *DescribeResolverEndpointsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetLang(v string) *DescribeResolverEndpointsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetPageNumber(v int32) *DescribeResolverEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetPageSize(v int32) *DescribeResolverEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetStatus(v string) *DescribeResolverEndpointsRequest {
	s.Status = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetVpcRegionId(v string) *DescribeResolverEndpointsRequest {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
