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
	// The ID of the address pool. To obtain the ID, call [DescribeDnsGtmInstanceAddressPools](https://www.alibabacloud.com/help/en/dns/latest/api-alidns-2015-01-09-describednsgtminstanceaddresspools).
	//
	// This parameter is required.
	//
	// example:
	//
	// testp******
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The language for some returned parameters. Default: en. Valid values: en, zh, and ja.
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
