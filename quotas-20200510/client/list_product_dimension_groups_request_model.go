// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductDimensionGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductDimensionGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductDimensionGroupsRequest
	GetNextToken() *string
	SetProductCode(v string) *ListProductDimensionGroupsRequest
	GetProductCode() *string
}

type ListProductDimensionGroupsRequest struct {
	// The maximum number of records that can be returned for the query. Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service code.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductDimensionGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductDimensionGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductDimensionGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductDimensionGroupsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductDimensionGroupsRequest) SetMaxResults(v int32) *ListProductDimensionGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductDimensionGroupsRequest) SetNextToken(v string) *ListProductDimensionGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductDimensionGroupsRequest) SetProductCode(v string) *ListProductDimensionGroupsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListProductDimensionGroupsRequest) Validate() error {
	return dara.Validate(s)
}
