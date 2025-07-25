// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *DeleteGtmAddressPoolRequest
	GetAddrPoolId() *string
	SetLang(v string) *DeleteGtmAddressPoolRequest
	GetLang() *string
}

type DeleteGtmAddressPoolRequest struct {
	// The ID of the address pool that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteGtmAddressPoolRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DeleteGtmAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteGtmAddressPoolRequest) SetAddrPoolId(v string) *DeleteGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *DeleteGtmAddressPoolRequest) SetLang(v string) *DeleteGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DeleteGtmAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}
