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
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeDistributionProductsLinkResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 50
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
	return dara.Validate(s)
}

type DescribeDistributionProductsLinkResponseBodyResult struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
