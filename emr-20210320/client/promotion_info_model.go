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
	SetIsSelected(v string) *PromotionInfo
	GetIsSelected() *string
	SetPromotionDesc(v string) *PromotionInfo
	GetPromotionDesc() *string
	SetPromotionName(v string) *PromotionInfo
	GetPromotionName() *string
	SetPromotionOptionCode(v string) *PromotionInfo
	GetPromotionOptionCode() *string
	SetPromotionOptionNo(v string) *PromotionInfo
	GetPromotionOptionNo() *string
}

type PromotionInfo struct {
	CanPromFee          *string `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	IsSelected          *string `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	PromotionDesc       *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionName       *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
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

func (s *PromotionInfo) GetIsSelected() *string {
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

func (s *PromotionInfo) SetCanPromFee(v string) *PromotionInfo {
	s.CanPromFee = &v
	return s
}

func (s *PromotionInfo) SetIsSelected(v string) *PromotionInfo {
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

func (s *PromotionInfo) Validate() error {
	return dara.Validate(s)
}
