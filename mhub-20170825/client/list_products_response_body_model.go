// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductInfos(v *ListProductsResponseBodyProductInfos) *ListProductsResponseBody
	GetProductInfos() *ListProductsResponseBodyProductInfos
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListProductsResponseBody
	GetTotal() *int32
	SetUbsmsStatus(v string) *ListProductsResponseBody
	GetUbsmsStatus() *string
}

type ListProductsResponseBody struct {
	ProductInfos *ListProductsResponseBodyProductInfos `json:"ProductInfos,omitempty" xml:"ProductInfos,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// enabled
	UbsmsStatus *string `json:"UbsmsStatus,omitempty" xml:"UbsmsStatus,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetProductInfos() *ListProductsResponseBodyProductInfos {
	return s.ProductInfos
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListProductsResponseBody) GetUbsmsStatus() *string {
	return s.UbsmsStatus
}

func (s *ListProductsResponseBody) SetProductInfos(v *ListProductsResponseBodyProductInfos) *ListProductsResponseBody {
	s.ProductInfos = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotal(v int32) *ListProductsResponseBody {
	s.Total = &v
	return s
}

func (s *ListProductsResponseBody) SetUbsmsStatus(v string) *ListProductsResponseBody {
	s.UbsmsStatus = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyProductInfos struct {
	ProductInfo []*ListProductsResponseBodyProductInfosProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyProductInfos) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProductInfos) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductInfos) GetProductInfo() []*ListProductsResponseBodyProductInfosProductInfo {
	return s.ProductInfo
}

func (s *ListProductsResponseBodyProductInfos) SetProductInfo(v []*ListProductsResponseBodyProductInfosProductInfo) *ListProductsResponseBodyProductInfos {
	s.ProductInfo = v
	return s
}

func (s *ListProductsResponseBodyProductInfos) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyProductInfosProductInfo struct {
	// example:
	//
	// 2020-12-16T06:25:52Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 图片 base64
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// example:
	//
	// 我的项目
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// iOS
	Platforms *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// example:
	//
	// 1234
	ProductId *int32 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
}

func (s ListProductsResponseBodyProductInfosProductInfo) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyProductInfosProductInfo) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetDescription() *string {
	return s.Description
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetEncodedIcon() *string {
	return s.EncodedIcon
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetIndustryId() *int32 {
	return s.IndustryId
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetName() *string {
	return s.Name
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetPlatforms() *string {
	return s.Platforms
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetProductId() *int32 {
	return s.ProductId
}

func (s *ListProductsResponseBodyProductInfosProductInfo) GetReadonly() *bool {
	return s.Readonly
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetCreateTime(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.CreateTime = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetDescription(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetEncodedIcon(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.EncodedIcon = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetIndustryId(v int32) *ListProductsResponseBodyProductInfosProductInfo {
	s.IndustryId = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetName(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetPlatforms(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.Platforms = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetProductId(v int32) *ListProductsResponseBodyProductInfosProductInfo {
	s.ProductId = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetReadonly(v bool) *ListProductsResponseBodyProductInfosProductInfo {
	s.Readonly = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) Validate() error {
	return dara.Validate(s)
}
