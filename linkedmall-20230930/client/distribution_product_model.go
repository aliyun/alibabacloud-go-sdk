// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributionProduct interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryChain(v string) *DistributionProduct
	GetCategoryChain() *string
	SetCategoryLeafId(v int64) *DistributionProduct
	GetCategoryLeafId() *int64
	SetCategoryLeafName(v string) *DistributionProduct
	GetCategoryLeafName() *string
	SetChannelCode(v string) *DistributionProduct
	GetChannelCode() *string
	SetDistributeStatus(v string) *DistributionProduct
	GetDistributeStatus() *string
	SetPicUrl(v string) *DistributionProduct
	GetPicUrl() *string
	SetProductId(v string) *DistributionProduct
	GetProductId() *string
	SetSellerId(v string) *DistributionProduct
	GetSellerId() *string
	SetShortTitle(v string) *DistributionProduct
	GetShortTitle() *string
	SetSkus(v []*DistributionSku) *DistributionProduct
	GetSkus() []*DistributionSku
	SetTitle(v string) *DistributionProduct
	GetTitle() *string
	SetWhiteBackgroundPicUrl(v string) *DistributionProduct
	GetWhiteBackgroundPicUrl() *string
}

type DistributionProduct struct {
	CategoryChain         *string            `json:"categoryChain,omitempty" xml:"categoryChain,omitempty"`
	CategoryLeafId        *int64             `json:"categoryLeafId,omitempty" xml:"categoryLeafId,omitempty"`
	CategoryLeafName      *string            `json:"categoryLeafName,omitempty" xml:"categoryLeafName,omitempty"`
	ChannelCode           *string            `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	DistributeStatus      *string            `json:"distributeStatus,omitempty" xml:"distributeStatus,omitempty"`
	PicUrl                *string            `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	ProductId             *string            `json:"productId,omitempty" xml:"productId,omitempty"`
	SellerId              *string            `json:"sellerId,omitempty" xml:"sellerId,omitempty"`
	ShortTitle            *string            `json:"shortTitle,omitempty" xml:"shortTitle,omitempty"`
	Skus                  []*DistributionSku `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	Title                 *string            `json:"title,omitempty" xml:"title,omitempty"`
	WhiteBackgroundPicUrl *string            `json:"whiteBackgroundPicUrl,omitempty" xml:"whiteBackgroundPicUrl,omitempty"`
}

func (s DistributionProduct) String() string {
	return dara.Prettify(s)
}

func (s DistributionProduct) GoString() string {
	return s.String()
}

func (s *DistributionProduct) GetCategoryChain() *string {
	return s.CategoryChain
}

func (s *DistributionProduct) GetCategoryLeafId() *int64 {
	return s.CategoryLeafId
}

func (s *DistributionProduct) GetCategoryLeafName() *string {
	return s.CategoryLeafName
}

func (s *DistributionProduct) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *DistributionProduct) GetDistributeStatus() *string {
	return s.DistributeStatus
}

func (s *DistributionProduct) GetPicUrl() *string {
	return s.PicUrl
}

func (s *DistributionProduct) GetProductId() *string {
	return s.ProductId
}

func (s *DistributionProduct) GetSellerId() *string {
	return s.SellerId
}

func (s *DistributionProduct) GetShortTitle() *string {
	return s.ShortTitle
}

func (s *DistributionProduct) GetSkus() []*DistributionSku {
	return s.Skus
}

func (s *DistributionProduct) GetTitle() *string {
	return s.Title
}

func (s *DistributionProduct) GetWhiteBackgroundPicUrl() *string {
	return s.WhiteBackgroundPicUrl
}

func (s *DistributionProduct) SetCategoryChain(v string) *DistributionProduct {
	s.CategoryChain = &v
	return s
}

func (s *DistributionProduct) SetCategoryLeafId(v int64) *DistributionProduct {
	s.CategoryLeafId = &v
	return s
}

func (s *DistributionProduct) SetCategoryLeafName(v string) *DistributionProduct {
	s.CategoryLeafName = &v
	return s
}

func (s *DistributionProduct) SetChannelCode(v string) *DistributionProduct {
	s.ChannelCode = &v
	return s
}

func (s *DistributionProduct) SetDistributeStatus(v string) *DistributionProduct {
	s.DistributeStatus = &v
	return s
}

func (s *DistributionProduct) SetPicUrl(v string) *DistributionProduct {
	s.PicUrl = &v
	return s
}

func (s *DistributionProduct) SetProductId(v string) *DistributionProduct {
	s.ProductId = &v
	return s
}

func (s *DistributionProduct) SetSellerId(v string) *DistributionProduct {
	s.SellerId = &v
	return s
}

func (s *DistributionProduct) SetShortTitle(v string) *DistributionProduct {
	s.ShortTitle = &v
	return s
}

func (s *DistributionProduct) SetSkus(v []*DistributionSku) *DistributionProduct {
	s.Skus = v
	return s
}

func (s *DistributionProduct) SetTitle(v string) *DistributionProduct {
	s.Title = &v
	return s
}

func (s *DistributionProduct) SetWhiteBackgroundPicUrl(v string) *DistributionProduct {
	s.WhiteBackgroundPicUrl = &v
	return s
}

func (s *DistributionProduct) Validate() error {
	if s.Skus != nil {
		for _, item := range s.Skus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
