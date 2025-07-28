// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnsCategory(v string) *UpdateCustomLineRequest
	GetDnsCategory() *string
	SetIpv4s(v []*string) *UpdateCustomLineRequest
	GetIpv4s() []*string
	SetLang(v string) *UpdateCustomLineRequest
	GetLang() *string
	SetLineId(v string) *UpdateCustomLineRequest
	GetLineId() *string
	SetName(v string) *UpdateCustomLineRequest
	GetName() *string
}

type UpdateCustomLineRequest struct {
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
	// The unique ID of the custom line.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100003
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The name of the custom line.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCustomLineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineRequest) GetDnsCategory() *string {
	return s.DnsCategory
}

func (s *UpdateCustomLineRequest) GetIpv4s() []*string {
	return s.Ipv4s
}

func (s *UpdateCustomLineRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateCustomLineRequest) GetLineId() *string {
	return s.LineId
}

func (s *UpdateCustomLineRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCustomLineRequest) SetDnsCategory(v string) *UpdateCustomLineRequest {
	s.DnsCategory = &v
	return s
}

func (s *UpdateCustomLineRequest) SetIpv4s(v []*string) *UpdateCustomLineRequest {
	s.Ipv4s = v
	return s
}

func (s *UpdateCustomLineRequest) SetLang(v string) *UpdateCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCustomLineRequest) SetLineId(v string) *UpdateCustomLineRequest {
	s.LineId = &v
	return s
}

func (s *UpdateCustomLineRequest) SetName(v string) *UpdateCustomLineRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomLineRequest) Validate() error {
	return dara.Validate(s)
}
