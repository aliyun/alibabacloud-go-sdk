// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeAdvisorResourcesRequest
	GetKeyword() *string
	SetLanguage(v string) *DescribeAdvisorResourcesRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeAdvisorResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAdvisorResourcesRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAdvisorResourcesRequest
	GetProduct() *string
	SetResourceId(v string) *DescribeAdvisorResourcesRequest
	GetResourceId() *string
}

type DescribeAdvisorResourcesRequest struct {
	// example:
	//
	// abcd
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvisorResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeAdvisorResourcesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAdvisorResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAdvisorResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAdvisorResourcesRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvisorResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvisorResourcesRequest) SetKeyword(v string) *DescribeAdvisorResourcesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetLanguage(v string) *DescribeAdvisorResourcesRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetPageNumber(v int32) *DescribeAdvisorResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetPageSize(v int32) *DescribeAdvisorResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetProduct(v string) *DescribeAdvisorResourcesRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) SetResourceId(v string) *DescribeAdvisorResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvisorResourcesRequest) Validate() error {
	return dara.Validate(s)
}
