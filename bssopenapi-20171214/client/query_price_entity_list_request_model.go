// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPriceEntityListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *QueryPriceEntityListRequest
	GetCommodityCode() *string
	SetLang(v string) *QueryPriceEntityListRequest
	GetLang() *string
}

type QueryPriceEntityListRequest struct {
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryPriceEntityListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceEntityListRequest) GoString() string {
	return s.String()
}

func (s *QueryPriceEntityListRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryPriceEntityListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryPriceEntityListRequest) SetCommodityCode(v string) *QueryPriceEntityListRequest {
	s.CommodityCode = &v
	return s
}

func (s *QueryPriceEntityListRequest) SetLang(v string) *QueryPriceEntityListRequest {
	s.Lang = &v
	return s
}

func (s *QueryPriceEntityListRequest) Validate() error {
	return dara.Validate(s)
}
