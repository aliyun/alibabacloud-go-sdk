// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnsCategory(v string) *AddCustomLineRequest
	GetDnsCategory() *string
	SetIpv4s(v []*string) *AddCustomLineRequest
	GetIpv4s() []*string
	SetLang(v string) *AddCustomLineRequest
	GetLang() *string
	SetName(v string) *AddCustomLineRequest
	GetName() *string
	SetShareScope(v string) *AddCustomLineRequest
	GetShareScope() *string
}

type AddCustomLineRequest struct {
	// This parameter is not available. You can ignore it.
	//
	// example:
	//
	// INTRANET
	DnsCategory *string `json:"DnsCategory,omitempty" xml:"DnsCategory,omitempty"`
	// The IPv4 CIDR blocks.
	//
	// This parameter is required.
	Ipv4s []*string `json:"Ipv4s,omitempty" xml:"Ipv4s,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the custom line.
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is not available. You can ignore it.
	//
	// example:
	//
	// GLOBAL
	ShareScope *string `json:"ShareScope,omitempty" xml:"ShareScope,omitempty"`
}

func (s AddCustomLineRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLineRequest) GoString() string {
	return s.String()
}

func (s *AddCustomLineRequest) GetDnsCategory() *string {
	return s.DnsCategory
}

func (s *AddCustomLineRequest) GetIpv4s() []*string {
	return s.Ipv4s
}

func (s *AddCustomLineRequest) GetLang() *string {
	return s.Lang
}

func (s *AddCustomLineRequest) GetName() *string {
	return s.Name
}

func (s *AddCustomLineRequest) GetShareScope() *string {
	return s.ShareScope
}

func (s *AddCustomLineRequest) SetDnsCategory(v string) *AddCustomLineRequest {
	s.DnsCategory = &v
	return s
}

func (s *AddCustomLineRequest) SetIpv4s(v []*string) *AddCustomLineRequest {
	s.Ipv4s = v
	return s
}

func (s *AddCustomLineRequest) SetLang(v string) *AddCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *AddCustomLineRequest) SetName(v string) *AddCustomLineRequest {
	s.Name = &v
	return s
}

func (s *AddCustomLineRequest) SetShareScope(v string) *AddCustomLineRequest {
	s.ShareScope = &v
	return s
}

func (s *AddCustomLineRequest) Validate() error {
	return dara.Validate(s)
}
