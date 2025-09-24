// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySkuPriceListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *QuerySkuPriceListShrinkRequest
	GetCommodityCode() *string
	SetLang(v string) *QuerySkuPriceListShrinkRequest
	GetLang() *string
	SetNextPageToken(v string) *QuerySkuPriceListShrinkRequest
	GetNextPageToken() *string
	SetPageSize(v int32) *QuerySkuPriceListShrinkRequest
	GetPageSize() *int32
	SetPriceEntityCode(v string) *QuerySkuPriceListShrinkRequest
	GetPriceEntityCode() *string
	SetPriceFactorConditionMapShrink(v string) *QuerySkuPriceListShrinkRequest
	GetPriceFactorConditionMapShrink() *string
}

type QuerySkuPriceListShrinkRequest struct {
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The token that is used to retrieve the next page. You do not need to set this parameter if you query coverage details for the first time. The response returns a token that you can use to query coverage details of the next page. If a null value is returned for the NextPageToken parameter, no more coverage details can be queried.
	//
	// example:
	//
	// 080112060a0422020800180022490a470342000000315333303332363436363336333433393636333136333338333733373333333133373336363336323634363336363337333836333636333636313336363433363332
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The number of entries to be returned on each page. Maximum value: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the pricing object.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance_type
	PriceEntityCode *string `json:"PriceEntityCode,omitempty" xml:"PriceEntityCode,omitempty"`
	// The conditions of the pricing factors.
	PriceFactorConditionMapShrink *string `json:"PriceFactorConditionMap,omitempty" xml:"PriceFactorConditionMap,omitempty"`
}

func (s QuerySkuPriceListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QuerySkuPriceListShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *QuerySkuPriceListShrinkRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QuerySkuPriceListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySkuPriceListShrinkRequest) GetPriceEntityCode() *string {
	return s.PriceEntityCode
}

func (s *QuerySkuPriceListShrinkRequest) GetPriceFactorConditionMapShrink() *string {
	return s.PriceFactorConditionMapShrink
}

func (s *QuerySkuPriceListShrinkRequest) SetCommodityCode(v string) *QuerySkuPriceListShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *QuerySkuPriceListShrinkRequest) SetLang(v string) *QuerySkuPriceListShrinkRequest {
	s.Lang = &v
	return s
}

func (s *QuerySkuPriceListShrinkRequest) SetNextPageToken(v string) *QuerySkuPriceListShrinkRequest {
	s.NextPageToken = &v
	return s
}

func (s *QuerySkuPriceListShrinkRequest) SetPageSize(v int32) *QuerySkuPriceListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySkuPriceListShrinkRequest) SetPriceEntityCode(v string) *QuerySkuPriceListShrinkRequest {
	s.PriceEntityCode = &v
	return s
}

func (s *QuerySkuPriceListShrinkRequest) SetPriceFactorConditionMapShrink(v string) *QuerySkuPriceListShrinkRequest {
	s.PriceFactorConditionMapShrink = &v
	return s
}

func (s *QuerySkuPriceListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
