// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromotionParam interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *PromotionParam
	GetProductCode() *string
	SetPromotionOptionCode(v string) *PromotionParam
	GetPromotionOptionCode() *string
	SetPromotionOptionNo(v string) *PromotionParam
	GetPromotionOptionNo() *string
}

type PromotionParam struct {
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s PromotionParam) String() string {
	return dara.Prettify(s)
}

func (s PromotionParam) GoString() string {
	return s.String()
}

func (s *PromotionParam) GetProductCode() *string {
	return s.ProductCode
}

func (s *PromotionParam) GetPromotionOptionCode() *string {
	return s.PromotionOptionCode
}

func (s *PromotionParam) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *PromotionParam) SetProductCode(v string) *PromotionParam {
	s.ProductCode = &v
	return s
}

func (s *PromotionParam) SetPromotionOptionCode(v string) *PromotionParam {
	s.PromotionOptionCode = &v
	return s
}

func (s *PromotionParam) SetPromotionOptionNo(v string) *PromotionParam {
	s.PromotionOptionNo = &v
	return s
}

func (s *PromotionParam) Validate() error {
	return dara.Validate(s)
}
