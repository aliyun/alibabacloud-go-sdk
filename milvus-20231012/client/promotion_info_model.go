// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromotionInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCanPromFee(v string) *PromotionInfo
	GetCanPromFee() *string
	SetIsSelected(v bool) *PromotionInfo
	GetIsSelected() *bool
	SetPromotionDesc(v string) *PromotionInfo
	GetPromotionDesc() *string
	SetPromotionName(v string) *PromotionInfo
	GetPromotionName() *string
	SetPromotionOptionCode(v string) *PromotionInfo
	GetPromotionOptionCode() *string
	SetPromotionOptionNo(v string) *PromotionInfo
	GetPromotionOptionNo() *string
	SetSelected(v bool) *PromotionInfo
	GetSelected() *bool
}

type PromotionInfo struct {
	CanPromFee          *string `json:"canPromFee,omitempty" xml:"canPromFee,omitempty"`
	IsSelected          *bool   `json:"isSelected,omitempty" xml:"isSelected,omitempty"`
	PromotionDesc       *string `json:"promotionDesc,omitempty" xml:"promotionDesc,omitempty"`
	PromotionName       *string `json:"promotionName,omitempty" xml:"promotionName,omitempty"`
	PromotionOptionCode *string `json:"promotionOptionCode,omitempty" xml:"promotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"promotionOptionNo,omitempty" xml:"promotionOptionNo,omitempty"`
	Selected            *bool   `json:"selected,omitempty" xml:"selected,omitempty"`
}

func (s PromotionInfo) String() string {
	return dara.Prettify(s)
}

func (s PromotionInfo) GoString() string {
	return s.String()
}

func (s *PromotionInfo) GetCanPromFee() *string {
	return s.CanPromFee
}

func (s *PromotionInfo) GetIsSelected() *bool {
	return s.IsSelected
}

func (s *PromotionInfo) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *PromotionInfo) GetPromotionName() *string {
	return s.PromotionName
}

func (s *PromotionInfo) GetPromotionOptionCode() *string {
	return s.PromotionOptionCode
}

func (s *PromotionInfo) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *PromotionInfo) GetSelected() *bool {
	return s.Selected
}

func (s *PromotionInfo) SetCanPromFee(v string) *PromotionInfo {
	s.CanPromFee = &v
	return s
}

func (s *PromotionInfo) SetIsSelected(v bool) *PromotionInfo {
	s.IsSelected = &v
	return s
}

func (s *PromotionInfo) SetPromotionDesc(v string) *PromotionInfo {
	s.PromotionDesc = &v
	return s
}

func (s *PromotionInfo) SetPromotionName(v string) *PromotionInfo {
	s.PromotionName = &v
	return s
}

func (s *PromotionInfo) SetPromotionOptionCode(v string) *PromotionInfo {
	s.PromotionOptionCode = &v
	return s
}

func (s *PromotionInfo) SetPromotionOptionNo(v string) *PromotionInfo {
	s.PromotionOptionNo = &v
	return s
}

func (s *PromotionInfo) SetSelected(v bool) *PromotionInfo {
	s.Selected = &v
	return s
}

func (s *PromotionInfo) Validate() error {
	return dara.Validate(s)
}
