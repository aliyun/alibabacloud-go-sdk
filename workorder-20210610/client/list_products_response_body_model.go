// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListProductsResponseBody
	GetCode() *int32
	SetData(v []*ListProductsResponseBodyData) *ListProductsResponseBody
	GetData() []*ListProductsResponseBodyData
	SetMessage(v string) *ListProductsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProductsResponseBody
	GetSuccess() *bool
}

type ListProductsResponseBody struct {
	// The return code of the request result.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return value, that is, product list
	Data []*ListProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message. If success is set to false, the message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the API request. The requestID is unique for each call.
	//
	// example:
	//
	// AC0AB2EC-AFBC-44BA-AE77-132A5A1EC0AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is normal.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListProductsResponseBody) GetData() []*ListProductsResponseBodyData {
	return s.Data
}

func (s *ListProductsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProductsResponseBody) SetCode(v int32) *ListProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductsResponseBody) SetData(v []*ListProductsResponseBodyData) *ListProductsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductsResponseBody) SetMessage(v string) *ListProductsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetSuccess(v bool) *ListProductsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyData struct {
	// The ID of the product catalog, which represents the product category, such as elastic computing
	//
	// example:
	//
	// 1
	DirectoryId *int64 `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the product catalog, which represents the product category, such as elastic computing
	//
	// example:
	//
	// ECS
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// List of Alibaba Cloud products
	ProductList []*ListProductsResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyData) GetDirectoryId() *int64 {
	return s.DirectoryId
}

func (s *ListProductsResponseBodyData) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *ListProductsResponseBodyData) GetProductList() []*ListProductsResponseBodyDataProductList {
	return s.ProductList
}

func (s *ListProductsResponseBodyData) SetDirectoryId(v int64) *ListProductsResponseBodyData {
	s.DirectoryId = &v
	return s
}

func (s *ListProductsResponseBodyData) SetDirectoryName(v string) *ListProductsResponseBodyData {
	s.DirectoryName = &v
	return s
}

func (s *ListProductsResponseBodyData) SetProductList(v []*ListProductsResponseBodyDataProductList) *ListProductsResponseBodyData {
	s.ProductList = v
	return s
}

func (s *ListProductsResponseBodyData) Validate() error {
	if s.ProductList != nil {
		for _, item := range s.ProductList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyDataProductList struct {
	// Alibaba Cloud product ID
	//
	// example:
	//
	// 7160
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Alibaba Cloud product name
	//
	// example:
	//
	// ECS
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s ListProductsResponseBodyDataProductList) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataProductList) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListProductsResponseBodyDataProductList) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsResponseBodyDataProductList) SetProductId(v int64) *ListProductsResponseBodyDataProductList {
	s.ProductId = &v
	return s
}

func (s *ListProductsResponseBodyDataProductList) SetProductName(v string) *ListProductsResponseBodyDataProductList {
	s.ProductName = &v
	return s
}

func (s *ListProductsResponseBodyDataProductList) Validate() error {
	return dara.Validate(s)
}
