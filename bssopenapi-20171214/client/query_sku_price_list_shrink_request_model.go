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
	// The commodity code. You can call the **QueryCommodityList*	- operation to query the commodity code.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The language. Default value: zh (Chinese).
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The pagination token. Leave this parameter empty for the first page query. The query result returns a token for the next page. An empty token in the result indicates that all data has been retrieved. You can obtain the next page token from the **NextPageToken*	- field in the query result.
	//
	// example:
	//
	// 080112060a0422020800180022490a470342000000315333303332363436363336333433393636333136333338333733373333333133373336363336323634363336363337333836333636333636313336363433363332
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The number of entries per page. Maximum value: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pricing entity code. You can call the **QueryPriceEntityList*	- operation to query the pricing entity code.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance_type
	PriceEntityCode *string `json:"PriceEntityCode,omitempty" xml:"PriceEntityCode,omitempty"`
	// The collection of pricing factor conditions. You can call the **QueryPriceEntityList*	- operation to query pricing factor codes and pricing factor values.
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
