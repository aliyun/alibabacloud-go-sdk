// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySkuPriceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *QuerySkuPriceListRequest
	GetCommodityCode() *string
	SetLang(v string) *QuerySkuPriceListRequest
	GetLang() *string
	SetNextPageToken(v string) *QuerySkuPriceListRequest
	GetNextPageToken() *string
	SetPageSize(v int32) *QuerySkuPriceListRequest
	GetPageSize() *int32
	SetPriceEntityCode(v string) *QuerySkuPriceListRequest
	GetPriceEntityCode() *string
	SetPriceFactorConditionMap(v map[string][]*string) *QuerySkuPriceListRequest
	GetPriceFactorConditionMap() map[string][]*string
}

type QuerySkuPriceListRequest struct {
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
	PriceFactorConditionMap map[string][]*string `json:"PriceFactorConditionMap,omitempty" xml:"PriceFactorConditionMap,omitempty"`
}

func (s QuerySkuPriceListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListRequest) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QuerySkuPriceListRequest) GetLang() *string {
	return s.Lang
}

func (s *QuerySkuPriceListRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QuerySkuPriceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySkuPriceListRequest) GetPriceEntityCode() *string {
	return s.PriceEntityCode
}

func (s *QuerySkuPriceListRequest) GetPriceFactorConditionMap() map[string][]*string {
	return s.PriceFactorConditionMap
}

func (s *QuerySkuPriceListRequest) SetCommodityCode(v string) *QuerySkuPriceListRequest {
	s.CommodityCode = &v
	return s
}

func (s *QuerySkuPriceListRequest) SetLang(v string) *QuerySkuPriceListRequest {
	s.Lang = &v
	return s
}

func (s *QuerySkuPriceListRequest) SetNextPageToken(v string) *QuerySkuPriceListRequest {
	s.NextPageToken = &v
	return s
}

func (s *QuerySkuPriceListRequest) SetPageSize(v int32) *QuerySkuPriceListRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySkuPriceListRequest) SetPriceEntityCode(v string) *QuerySkuPriceListRequest {
	s.PriceEntityCode = &v
	return s
}

func (s *QuerySkuPriceListRequest) SetPriceFactorConditionMap(v map[string][]*string) *QuerySkuPriceListRequest {
	s.PriceFactorConditionMap = v
	return s
}

func (s *QuerySkuPriceListRequest) Validate() error {
	return dara.Validate(s)
}
