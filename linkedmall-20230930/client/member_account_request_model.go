// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMemberAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetShopId(v string) *MemberAccountRequest
	GetShopId() *string
}

type MemberAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s MemberAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s MemberAccountRequest) GoString() string {
	return s.String()
}

func (s *MemberAccountRequest) GetShopId() *string {
	return s.ShopId
}

func (s *MemberAccountRequest) SetShopId(v string) *MemberAccountRequest {
	s.ShopId = &v
	return s
}

func (s *MemberAccountRequest) Validate() error {
	return dara.Validate(s)
}
