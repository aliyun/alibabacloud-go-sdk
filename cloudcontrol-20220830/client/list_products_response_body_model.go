// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsResponseBody
	GetNextToken() *string
	SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody
	GetProducts() []*ListProductsResponseBodyProducts
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductsResponseBody
	GetTotalCount() *int32
}

type ListProductsResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// Redis
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The cloud services.
	Products []*ListProductsResponseBodyProducts `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// The ID of a request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsResponseBody) GetProducts() []*ListProductsResponseBodyProducts {
	return s.Products
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int32) *ListProductsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyProducts struct {
	// The code of the service.
	//
	// example:
	//
	// ECS
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// The service name.
	//
	// example:
	//
	// Elastic Compute Service
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductsResponseBodyProducts) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsResponseBodyProducts) SetProductCode(v string) *ListProductsResponseBodyProducts {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductName(v string) *ListProductsResponseBodyProducts {
	s.ProductName = &v
	return s
}

func (s *ListProductsResponseBodyProducts) Validate() error {
	return dara.Validate(s)
}
