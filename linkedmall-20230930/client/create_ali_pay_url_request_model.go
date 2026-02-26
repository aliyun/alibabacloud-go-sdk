// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliPayUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetShopId(v string) *CreateAliPayUrlRequest
	GetShopId() *string
}

type CreateAliPayUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s CreateAliPayUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAliPayUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateAliPayUrlRequest) GetShopId() *string {
	return s.ShopId
}

func (s *CreateAliPayUrlRequest) SetShopId(v string) *CreateAliPayUrlRequest {
	s.ShopId = &v
	return s
}

func (s *CreateAliPayUrlRequest) Validate() error {
	return dara.Validate(s)
}
