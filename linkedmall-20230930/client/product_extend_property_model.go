// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductExtendProperty interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ProductExtendProperty
	GetKey() *string
	SetValue(v string) *ProductExtendProperty
	GetValue() *string
}

type ProductExtendProperty struct {
	// The property key.
	//
	// > Valid values:
	//
	// >
	//
	// > - - `ss_picture_scene` (scene picture)
	//
	// >
	//
	// > - - `ss_picture_white_background` (white background picture)
	//
	// >
	//
	// > - - `extraPeriod` (shelf life)
	//
	// >
	//
	// > - - `itemBoundaryInventoryZeroTag` (Reserved. Ignore this parameter.)
	//
	// >
	//
	// > - - `shoppingShowTitle` (shopping guide title)
	//
	// >
	//
	// > - - `itemCCStatus` (Reserved. Ignore this parameter.)
	//
	// >
	//
	// > - - `brandLogo` (brand logo)
	//
	// >
	//
	// > - - `multipleBuyLimit` (purchase multiple)
	//
	// >
	//
	// > - - `eticket_type` (electronic coupon type)
	//
	// >
	//
	// > - - `eticket_upper_buy_limit` (maximum purchase quantity of electronic coupons per order)
	//
	// >
	//
	// > - - `validity_type` (validity period type of electronic coupon)
	//
	// >
	//
	// > - - `etc_expiry_date` (Validity period of the electronic coupon. Valid only when `validity_type` is `1`.)
	//
	// >
	//
	// > - - `etc_duration_date` (Validity period of the electronic coupon. Valid only when `validity_type` is `2`, `3`, or `5`.)
	//
	// >
	//
	// > - - `f_refund` (Automatic refund ratio for valid electronic coupons)
	//
	// >
	//
	// > - - `refund` (automatic refund ratio for expired electronic coupons)
	//
	// >
	//
	// > - - `writeoff` (Reserved. Ignore this parameter.)
	//
	// example:
	//
	// ss_picture_scene
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The property value.
	//
	// example:
	//
	// 场景图
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProductExtendProperty) String() string {
	return dara.Prettify(s)
}

func (s ProductExtendProperty) GoString() string {
	return s.String()
}

func (s *ProductExtendProperty) GetKey() *string {
	return s.Key
}

func (s *ProductExtendProperty) GetValue() *string {
	return s.Value
}

func (s *ProductExtendProperty) SetKey(v string) *ProductExtendProperty {
	s.Key = &v
	return s
}

func (s *ProductExtendProperty) SetValue(v string) *ProductExtendProperty {
	s.Value = &v
	return s
}

func (s *ProductExtendProperty) Validate() error {
	return dara.Validate(s)
}
