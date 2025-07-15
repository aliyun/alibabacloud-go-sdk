// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiProductsByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductInfoList(v *DescribeApiProductsByAppResponseBodyApiProductInfoList) *DescribeApiProductsByAppResponseBody
	GetApiProductInfoList() *DescribeApiProductsByAppResponseBodyApiProductInfoList
	SetPageNumber(v int32) *DescribeApiProductsByAppResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiProductsByAppResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApiProductsByAppResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApiProductsByAppResponseBody
	GetTotalCount() *int32
}

type DescribeApiProductsByAppResponseBody struct {
	// The information about API products.
	ApiProductInfoList *DescribeApiProductsByAppResponseBodyApiProductInfoList `json:"ApiProductInfoList,omitempty" xml:"ApiProductInfoList,omitempty" type:"Struct"`
	// The page number. Default value: 1.
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B805201-AF4C-5788-AC9E-C3EEC83DC82A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiProductsByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductsByAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiProductsByAppResponseBody) GetApiProductInfoList() *DescribeApiProductsByAppResponseBodyApiProductInfoList {
	return s.ApiProductInfoList
}

func (s *DescribeApiProductsByAppResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiProductsByAppResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiProductsByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiProductsByAppResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApiProductsByAppResponseBody) SetApiProductInfoList(v *DescribeApiProductsByAppResponseBodyApiProductInfoList) *DescribeApiProductsByAppResponseBody {
	s.ApiProductInfoList = v
	return s
}

func (s *DescribeApiProductsByAppResponseBody) SetPageNumber(v int32) *DescribeApiProductsByAppResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiProductsByAppResponseBody) SetPageSize(v int32) *DescribeApiProductsByAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiProductsByAppResponseBody) SetRequestId(v string) *DescribeApiProductsByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiProductsByAppResponseBody) SetTotalCount(v int32) *DescribeApiProductsByAppResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApiProductsByAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiProductsByAppResponseBodyApiProductInfoList struct {
	ApiProductInfo []*DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo `json:"ApiProductInfo,omitempty" xml:"ApiProductInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiProductsByAppResponseBodyApiProductInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductsByAppResponseBodyApiProductInfoList) GoString() string {
	return s.String()
}

func (s *DescribeApiProductsByAppResponseBodyApiProductInfoList) GetApiProductInfo() []*DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo {
	return s.ApiProductInfo
}

func (s *DescribeApiProductsByAppResponseBodyApiProductInfoList) SetApiProductInfo(v []*DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo) *DescribeApiProductsByAppResponseBodyApiProductInfoList {
	s.ApiProductInfo = v
	return s
}

func (s *DescribeApiProductsByAppResponseBodyApiProductInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo struct {
	// The ID of the API product.
	//
	// example:
	//
	// 117b7a64a8b3f064eaa4a47ac62aac5e
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
}

func (s DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo) SetApiProductId(v string) *DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo {
	s.ApiProductId = &v
	return s
}

func (s *DescribeApiProductsByAppResponseBodyApiProductInfoListApiProductInfo) Validate() error {
	return dara.Validate(s)
}
