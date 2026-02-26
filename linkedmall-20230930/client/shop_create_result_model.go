// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShopCreateResult interface {
	dara.Model
	String() string
	GoString() string
	SetShopId(v string) *ShopCreateResult
	GetShopId() *string
	SetShopStatus(v string) *ShopCreateResult
	GetShopStatus() *string
}

type ShopCreateResult struct {
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// example:
	//
	// WORKING
	ShopStatus *string `json:"shopStatus,omitempty" xml:"shopStatus,omitempty"`
}

func (s ShopCreateResult) String() string {
	return dara.Prettify(s)
}

func (s ShopCreateResult) GoString() string {
	return s.String()
}

func (s *ShopCreateResult) GetShopId() *string {
	return s.ShopId
}

func (s *ShopCreateResult) GetShopStatus() *string {
	return s.ShopStatus
}

func (s *ShopCreateResult) SetShopId(v string) *ShopCreateResult {
	s.ShopId = &v
	return s
}

func (s *ShopCreateResult) SetShopStatus(v string) *ShopCreateResult {
	s.ShopStatus = &v
	return s
}

func (s *ShopCreateResult) Validate() error {
	return dara.Validate(s)
}
