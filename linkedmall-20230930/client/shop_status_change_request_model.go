// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShopStatusChangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetShopId(v string) *ShopStatusChangeRequest
	GetShopId() *string
	SetShopStatus(v string) *ShopStatusChangeRequest
	GetShopStatus() *string
}

type ShopStatusChangeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WORKING
	ShopStatus *string `json:"shopStatus,omitempty" xml:"shopStatus,omitempty"`
}

func (s ShopStatusChangeRequest) String() string {
	return dara.Prettify(s)
}

func (s ShopStatusChangeRequest) GoString() string {
	return s.String()
}

func (s *ShopStatusChangeRequest) GetShopId() *string {
	return s.ShopId
}

func (s *ShopStatusChangeRequest) GetShopStatus() *string {
	return s.ShopStatus
}

func (s *ShopStatusChangeRequest) SetShopId(v string) *ShopStatusChangeRequest {
	s.ShopId = &v
	return s
}

func (s *ShopStatusChangeRequest) SetShopStatus(v string) *ShopStatusChangeRequest {
	s.ShopStatus = &v
	return s
}

func (s *ShopStatusChangeRequest) Validate() error {
	return dara.Validate(s)
}
