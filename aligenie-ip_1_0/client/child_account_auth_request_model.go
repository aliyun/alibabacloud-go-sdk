// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChildAccountAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ChildAccountAuthRequest
	GetAccount() *string
	SetAppKey(v string) *ChildAccountAuthRequest
	GetAppKey() *string
	SetHotelId(v string) *ChildAccountAuthRequest
	GetHotelId() *string
	SetTbOpenId(v string) *ChildAccountAuthRequest
	GetTbOpenId() *string
}

type ChildAccountAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// lee
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30471753
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEV***E3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s ChildAccountAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s ChildAccountAuthRequest) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthRequest) GetAccount() *string {
	return s.Account
}

func (s *ChildAccountAuthRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *ChildAccountAuthRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ChildAccountAuthRequest) GetTbOpenId() *string {
	return s.TbOpenId
}

func (s *ChildAccountAuthRequest) SetAccount(v string) *ChildAccountAuthRequest {
	s.Account = &v
	return s
}

func (s *ChildAccountAuthRequest) SetAppKey(v string) *ChildAccountAuthRequest {
	s.AppKey = &v
	return s
}

func (s *ChildAccountAuthRequest) SetHotelId(v string) *ChildAccountAuthRequest {
	s.HotelId = &v
	return s
}

func (s *ChildAccountAuthRequest) SetTbOpenId(v string) *ChildAccountAuthRequest {
	s.TbOpenId = &v
	return s
}

func (s *ChildAccountAuthRequest) Validate() error {
	return dara.Validate(s)
}
