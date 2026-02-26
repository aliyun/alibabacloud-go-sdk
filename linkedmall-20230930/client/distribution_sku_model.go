// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributionSku interface {
	dara.Model
	String() string
	GoString() string
	SetAliasTitle(v string) *DistributionSku
	GetAliasTitle() *string
	SetBarCode(v string) *DistributionSku
	GetBarCode() *string
	SetCreditPeriod(v int32) *DistributionSku
	GetCreditPeriod() *int32
	SetDxPrice(v int64) *DistributionSku
	GetDxPrice() *int64
	SetHasCredit(v bool) *DistributionSku
	GetHasCredit() *bool
	SetHasInvoice(v bool) *DistributionSku
	GetHasInvoice() *bool
	SetJxPrice(v int64) *DistributionSku
	GetJxPrice() *int64
	SetPicUrl(v string) *DistributionSku
	GetPicUrl() *string
	SetQuantity(v int32) *DistributionSku
	GetQuantity() *int32
	SetSkuId(v string) *DistributionSku
	GetSkuId() *string
	SetSkuStatus(v string) *DistributionSku
	GetSkuStatus() *string
	SetTaxCode(v string) *DistributionSku
	GetTaxCode() *string
	SetTaxRate(v int32) *DistributionSku
	GetTaxRate() *int32
	SetTitle(v string) *DistributionSku
	GetTitle() *string
}

type DistributionSku struct {
	AliasTitle   *string `json:"aliasTitle,omitempty" xml:"aliasTitle,omitempty"`
	BarCode      *string `json:"barCode,omitempty" xml:"barCode,omitempty"`
	CreditPeriod *int32  `json:"creditPeriod,omitempty" xml:"creditPeriod,omitempty"`
	DxPrice      *int64  `json:"dxPrice,omitempty" xml:"dxPrice,omitempty"`
	HasCredit    *bool   `json:"hasCredit,omitempty" xml:"hasCredit,omitempty"`
	HasInvoice   *bool   `json:"hasInvoice,omitempty" xml:"hasInvoice,omitempty"`
	JxPrice      *int64  `json:"jxPrice,omitempty" xml:"jxPrice,omitempty"`
	PicUrl       *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	Quantity     *int32  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	SkuId        *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	SkuStatus    *string `json:"skuStatus,omitempty" xml:"skuStatus,omitempty"`
	TaxCode      *string `json:"taxCode,omitempty" xml:"taxCode,omitempty"`
	TaxRate      *int32  `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DistributionSku) String() string {
	return dara.Prettify(s)
}

func (s DistributionSku) GoString() string {
	return s.String()
}

func (s *DistributionSku) GetAliasTitle() *string {
	return s.AliasTitle
}

func (s *DistributionSku) GetBarCode() *string {
	return s.BarCode
}

func (s *DistributionSku) GetCreditPeriod() *int32 {
	return s.CreditPeriod
}

func (s *DistributionSku) GetDxPrice() *int64 {
	return s.DxPrice
}

func (s *DistributionSku) GetHasCredit() *bool {
	return s.HasCredit
}

func (s *DistributionSku) GetHasInvoice() *bool {
	return s.HasInvoice
}

func (s *DistributionSku) GetJxPrice() *int64 {
	return s.JxPrice
}

func (s *DistributionSku) GetPicUrl() *string {
	return s.PicUrl
}

func (s *DistributionSku) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DistributionSku) GetSkuId() *string {
	return s.SkuId
}

func (s *DistributionSku) GetSkuStatus() *string {
	return s.SkuStatus
}

func (s *DistributionSku) GetTaxCode() *string {
	return s.TaxCode
}

func (s *DistributionSku) GetTaxRate() *int32 {
	return s.TaxRate
}

func (s *DistributionSku) GetTitle() *string {
	return s.Title
}

func (s *DistributionSku) SetAliasTitle(v string) *DistributionSku {
	s.AliasTitle = &v
	return s
}

func (s *DistributionSku) SetBarCode(v string) *DistributionSku {
	s.BarCode = &v
	return s
}

func (s *DistributionSku) SetCreditPeriod(v int32) *DistributionSku {
	s.CreditPeriod = &v
	return s
}

func (s *DistributionSku) SetDxPrice(v int64) *DistributionSku {
	s.DxPrice = &v
	return s
}

func (s *DistributionSku) SetHasCredit(v bool) *DistributionSku {
	s.HasCredit = &v
	return s
}

func (s *DistributionSku) SetHasInvoice(v bool) *DistributionSku {
	s.HasInvoice = &v
	return s
}

func (s *DistributionSku) SetJxPrice(v int64) *DistributionSku {
	s.JxPrice = &v
	return s
}

func (s *DistributionSku) SetPicUrl(v string) *DistributionSku {
	s.PicUrl = &v
	return s
}

func (s *DistributionSku) SetQuantity(v int32) *DistributionSku {
	s.Quantity = &v
	return s
}

func (s *DistributionSku) SetSkuId(v string) *DistributionSku {
	s.SkuId = &v
	return s
}

func (s *DistributionSku) SetSkuStatus(v string) *DistributionSku {
	s.SkuStatus = &v
	return s
}

func (s *DistributionSku) SetTaxCode(v string) *DistributionSku {
	s.TaxCode = &v
	return s
}

func (s *DistributionSku) SetTaxRate(v int32) *DistributionSku {
	s.TaxRate = &v
	return s
}

func (s *DistributionSku) SetTitle(v string) *DistributionSku {
	s.Title = &v
	return s
}

func (s *DistributionSku) Validate() error {
	return dara.Validate(s)
}
