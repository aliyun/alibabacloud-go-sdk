// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolRequest
	GetAddrPoolId() *string
	SetLang(v string) *DescribeDnsGtmInstanceAddressPoolRequest
	GetLang() *string
}

type DescribeDnsGtmInstanceAddressPoolRequest struct {
	// The ID of the address pool.<props="china">You can call the [DescribeDnsGtmInstanceAddressPools](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describednsgtminstanceaddresspools?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_9_7.1cee430dbd1I3y) operation to obtain the ID.
	//
	// <props="intl">You can call the [DescribeDnsGtmInstanceAddressPools](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describednsgtminstanceaddresspools?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test*****
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The language of the response. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) SetLang(v string) *DescribeDnsGtmInstanceAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}
