// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeliveredSupplyItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *QueryDeliveredSupplyItemsRequest
	GetAccountId() *string
	SetCommodityTypeCode(v string) *QueryDeliveredSupplyItemsRequest
	GetCommodityTypeCode() *string
}

type QueryDeliveredSupplyItemsRequest struct {
	AccountId         *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
}

func (s QueryDeliveredSupplyItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeliveredSupplyItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryDeliveredSupplyItemsRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryDeliveredSupplyItemsRequest) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *QueryDeliveredSupplyItemsRequest) SetAccountId(v string) *QueryDeliveredSupplyItemsRequest {
	s.AccountId = &v
	return s
}

func (s *QueryDeliveredSupplyItemsRequest) SetCommodityTypeCode(v string) *QueryDeliveredSupplyItemsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *QueryDeliveredSupplyItemsRequest) Validate() error {
	return dara.Validate(s)
}
