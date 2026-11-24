// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributionProductsLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDistributionProductsLinkResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeDistributionProductsLinkResponseBodyResult) *DescribeDistributionProductsLinkResponseBody
	GetResult() []*DescribeDistributionProductsLinkResponseBodyResult
	SetSuccess(v bool) *DescribeDistributionProductsLinkResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeDistributionProductsLinkResponseBody
	GetTotalCount() *int64
}

type DescribeDistributionProductsLinkResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5BD09171-BF4D-18D8-890E-C70C067527BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The promotion link information for the promoted products.
	Result []*DescribeDistributionProductsLinkResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The success status indicator.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of promoted products.
	//
	// example:
	//
	// 51
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDistributionProductsLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDistributionProductsLinkResponseBody) GetResult() []*DescribeDistributionProductsLinkResponseBodyResult {
	return s.Result
}

func (s *DescribeDistributionProductsLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDistributionProductsLinkResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDistributionProductsLinkResponseBody) SetRequestId(v string) *DescribeDistributionProductsLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBody) SetResult(v []*DescribeDistributionProductsLinkResponseBodyResult) *DescribeDistributionProductsLinkResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBody) SetSuccess(v bool) *DescribeDistributionProductsLinkResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBody) SetTotalCount(v int64) *DescribeDistributionProductsLinkResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDistributionProductsLinkResponseBodyResult struct {
	// The code of the promoted product.
	//
	// example:
	//
	// cmapi000****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the promoted product.
	//
	// example:
	//
	// Test product
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The promotion link for the product. Note: The sample link is a test product link.
	//
	// example:
	//
	// https://market.aliyun.com/ad/adClick?param=QN%2BgAHA8mY%2BV1vSDtkNqyWG9G5uUX2Q8%2BJrC8wTzBo5YK7tTmm4Zdiz6nmmyLHMC5REAsq2viyeJYjBjVHno1HydPvHrVarNXQTMEzEmR8TRZaG2FsnzpszXKK2cMA69CUmF9qQpHqU%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDistributionProductsLinkResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsLinkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) GetCode() *string {
	return s.Code
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) GetUrl() *string {
	return s.Url
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) SetCode(v string) *DescribeDistributionProductsLinkResponseBodyResult {
	s.Code = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) SetName(v string) *DescribeDistributionProductsLinkResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) SetUrl(v string) *DescribeDistributionProductsLinkResponseBodyResult {
	s.Url = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
