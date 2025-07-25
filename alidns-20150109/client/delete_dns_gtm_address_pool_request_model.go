// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *DeleteDnsGtmAddressPoolRequest
	GetAddrPoolId() *string
	SetLang(v string) *DeleteDnsGtmAddressPoolRequest
	GetLang() *string
}

type DeleteDnsGtmAddressPoolRequest struct {
	// The ID of the address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpool1
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteDnsGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAddressPoolRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DeleteDnsGtmAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDnsGtmAddressPoolRequest) SetAddrPoolId(v string) *DeleteDnsGtmAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *DeleteDnsGtmAddressPoolRequest) SetLang(v string) *DeleteDnsGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDnsGtmAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}
