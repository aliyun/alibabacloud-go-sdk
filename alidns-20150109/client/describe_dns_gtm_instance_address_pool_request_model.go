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
	// The ID of the address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpool1
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
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
