// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAddrAttributeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddrs(v string) *DescribeDnsGtmAddrAttributeInfoRequest
	GetAddrs() *string
	SetLang(v string) *DescribeDnsGtmAddrAttributeInfoRequest
	GetLang() *string
	SetType(v string) *DescribeDnsGtmAddrAttributeInfoRequest
	GetType() *string
}

type DescribeDnsGtmAddrAttributeInfoRequest struct {
	// The list of addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       "1.1.XX.XX"
	//
	// ]
	Addrs *string `json:"Addrs,omitempty" xml:"Addrs,omitempty"`
	// The language of some returned parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The address type:
	//
	// - IPV4: IPv4 address
	//
	// - IPV6: IPv6 address
	//
	// - DOMAIN: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// IPV4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDnsGtmAddrAttributeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddrAttributeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) GetAddrs() *string {
	return s.Addrs
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) SetAddrs(v string) *DescribeDnsGtmAddrAttributeInfoRequest {
	s.Addrs = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) SetLang(v string) *DescribeDnsGtmAddrAttributeInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) SetType(v string) *DescribeDnsGtmAddrAttributeInfoRequest {
	s.Type = &v
	return s
}

func (s *DescribeDnsGtmAddrAttributeInfoRequest) Validate() error {
	return dara.Validate(s)
}
