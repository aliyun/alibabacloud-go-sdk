// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolRequest
	GetAddrPoolId() *string
	SetLang(v string) *DescribeGtmInstanceAddressPoolRequest
	GetLang() *string
}

type DescribeGtmInstanceAddressPoolRequest struct {
	// The ID of the address pool that you want to query.
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

func (s DescribeGtmInstanceAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeGtmInstanceAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmInstanceAddressPoolRequest) SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolRequest {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolRequest) SetLang(v string) *DescribeGtmInstanceAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}
