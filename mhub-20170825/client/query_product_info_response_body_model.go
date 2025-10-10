// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductInfo(v *QueryProductInfoResponseBodyProductInfo) *QueryProductInfoResponseBody
	GetProductInfo() *QueryProductInfoResponseBodyProductInfo
	SetRequestId(v string) *QueryProductInfoResponseBody
	GetRequestId() *string
}

type QueryProductInfoResponseBody struct {
	ProductInfo *QueryProductInfoResponseBodyProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryProductInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProductInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductInfoResponseBody) GetProductInfo() *QueryProductInfoResponseBodyProductInfo {
	return s.ProductInfo
}

func (s *QueryProductInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProductInfoResponseBody) SetProductInfo(v *QueryProductInfoResponseBodyProductInfo) *QueryProductInfoResponseBody {
	s.ProductInfo = v
	return s
}

func (s *QueryProductInfoResponseBody) SetRequestId(v string) *QueryProductInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProductInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryProductInfoResponseBodyProductInfo struct {
	// example:
	//
	// 2022-12-02T09:50:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// endpoint for ecs: kangaroo-xuanji-cn-hangzhou-ecs-console-mirror-0(i-bp1fs84ua5zw4aljdlh1)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// icon图片base64字符串
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// icon图片名称
	IconName *string `json:"IconName,omitempty" xml:"IconName,omitempty"`
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
	// linux/amd64
	Platforms *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// example:
	//
	// 3910360
	ProductId *int32 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
}

func (s QueryProductInfoResponseBodyProductInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryProductInfoResponseBodyProductInfo) GoString() string {
	return s.String()
}

func (s *QueryProductInfoResponseBodyProductInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProductInfoResponseBodyProductInfo) GetDescription() *string {
	return s.Description
}

func (s *QueryProductInfoResponseBodyProductInfo) GetEncodedIcon() *string {
	return s.EncodedIcon
}

func (s *QueryProductInfoResponseBodyProductInfo) GetIconName() *string {
	return s.IconName
}

func (s *QueryProductInfoResponseBodyProductInfo) GetIndustryId() *int32 {
	return s.IndustryId
}

func (s *QueryProductInfoResponseBodyProductInfo) GetName() *string {
	return s.Name
}

func (s *QueryProductInfoResponseBodyProductInfo) GetPlatforms() *string {
	return s.Platforms
}

func (s *QueryProductInfoResponseBodyProductInfo) GetProductId() *int32 {
	return s.ProductId
}

func (s *QueryProductInfoResponseBodyProductInfo) GetReadonly() *bool {
	return s.Readonly
}

func (s *QueryProductInfoResponseBodyProductInfo) SetCreateTime(v string) *QueryProductInfoResponseBodyProductInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetDescription(v string) *QueryProductInfoResponseBodyProductInfo {
	s.Description = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetEncodedIcon(v string) *QueryProductInfoResponseBodyProductInfo {
	s.EncodedIcon = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetIconName(v string) *QueryProductInfoResponseBodyProductInfo {
	s.IconName = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetIndustryId(v int32) *QueryProductInfoResponseBodyProductInfo {
	s.IndustryId = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetName(v string) *QueryProductInfoResponseBodyProductInfo {
	s.Name = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetPlatforms(v string) *QueryProductInfoResponseBodyProductInfo {
	s.Platforms = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetProductId(v int32) *QueryProductInfoResponseBodyProductInfo {
	s.ProductId = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetReadonly(v bool) *QueryProductInfoResponseBodyProductInfo {
	s.Readonly = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) Validate() error {
	return dara.Validate(s)
}
