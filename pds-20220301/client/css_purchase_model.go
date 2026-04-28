// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCssPurchase interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CssPurchase
	GetChargeType() *string
	SetCommodityCode(v string) *CssPurchase
	GetCommodityCode() *string
	SetEndDate(v int64) *CssPurchase
	GetEndDate() *int64
	SetGmtCreate(v int64) *CssPurchase
	GetGmtCreate() *int64
	SetInstanceComponents(v []*CssInstanceComponent) *CssPurchase
	GetInstanceComponents() []*CssInstanceComponent
	SetInstanceId(v string) *CssPurchase
	GetInstanceId() *string
	SetOrderType(v string) *CssPurchase
	GetOrderType() *string
	SetPurchaseParams(v map[string]*string) *CssPurchase
	GetPurchaseParams() map[string]*string
	SetStartDate(v int64) *CssPurchase
	GetStartDate() *int64
}

type CssPurchase struct {
	ChargeType         *string                 `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	CommodityCode      *string                 `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	EndDate            *int64                  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	GmtCreate          *int64                  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	InstanceComponents []*CssInstanceComponent `json:"instanceComponents,omitempty" xml:"instanceComponents,omitempty" type:"Repeated"`
	InstanceId         *string                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrderType          *string                 `json:"orderType,omitempty" xml:"orderType,omitempty"`
	PurchaseParams     map[string]*string      `json:"purchaseParams,omitempty" xml:"purchaseParams,omitempty"`
	StartDate          *int64                  `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s CssPurchase) String() string {
	return dara.Prettify(s)
}

func (s CssPurchase) GoString() string {
	return s.String()
}

func (s *CssPurchase) GetChargeType() *string {
	return s.ChargeType
}

func (s *CssPurchase) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CssPurchase) GetEndDate() *int64 {
	return s.EndDate
}

func (s *CssPurchase) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CssPurchase) GetInstanceComponents() []*CssInstanceComponent {
	return s.InstanceComponents
}

func (s *CssPurchase) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CssPurchase) GetOrderType() *string {
	return s.OrderType
}

func (s *CssPurchase) GetPurchaseParams() map[string]*string {
	return s.PurchaseParams
}

func (s *CssPurchase) GetStartDate() *int64 {
	return s.StartDate
}

func (s *CssPurchase) SetChargeType(v string) *CssPurchase {
	s.ChargeType = &v
	return s
}

func (s *CssPurchase) SetCommodityCode(v string) *CssPurchase {
	s.CommodityCode = &v
	return s
}

func (s *CssPurchase) SetEndDate(v int64) *CssPurchase {
	s.EndDate = &v
	return s
}

func (s *CssPurchase) SetGmtCreate(v int64) *CssPurchase {
	s.GmtCreate = &v
	return s
}

func (s *CssPurchase) SetInstanceComponents(v []*CssInstanceComponent) *CssPurchase {
	s.InstanceComponents = v
	return s
}

func (s *CssPurchase) SetInstanceId(v string) *CssPurchase {
	s.InstanceId = &v
	return s
}

func (s *CssPurchase) SetOrderType(v string) *CssPurchase {
	s.OrderType = &v
	return s
}

func (s *CssPurchase) SetPurchaseParams(v map[string]*string) *CssPurchase {
	s.PurchaseParams = v
	return s
}

func (s *CssPurchase) SetStartDate(v int64) *CssPurchase {
	s.StartDate = &v
	return s
}

func (s *CssPurchase) Validate() error {
	if s.InstanceComponents != nil {
		for _, item := range s.InstanceComponents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
