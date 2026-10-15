// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommodityListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *QueryCommodityListRequest
	GetLang() *string
	SetProductCode(v string) *QueryCommodityListRequest
	GetProductCode() *string
}

type QueryCommodityListRequest struct {
	// The language. If this parameter is not specified, the default value zh (Chinese) is used.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s QueryCommodityListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCommodityListRequest) GoString() string {
	return s.String()
}

func (s *QueryCommodityListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryCommodityListRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryCommodityListRequest) SetLang(v string) *QueryCommodityListRequest {
	s.Lang = &v
	return s
}

func (s *QueryCommodityListRequest) SetProductCode(v string) *QueryCommodityListRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryCommodityListRequest) Validate() error {
	return dara.Validate(s)
}
