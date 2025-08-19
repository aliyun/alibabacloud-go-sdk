// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDNSOption interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DNSOption
	GetName() *string
	SetValue(v string) *DNSOption
	GetValue() *string
}

type DNSOption struct {
	// example:
	//
	// ndots
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DNSOption) String() string {
	return dara.Prettify(s)
}

func (s DNSOption) GoString() string {
	return s.String()
}

func (s *DNSOption) GetName() *string {
	return s.Name
}

func (s *DNSOption) GetValue() *string {
	return s.Value
}

func (s *DNSOption) SetName(v string) *DNSOption {
	s.Name = &v
	return s
}

func (s *DNSOption) SetValue(v string) *DNSOption {
	s.Value = &v
	return s
}

func (s *DNSOption) Validate() error {
	return dara.Validate(s)
}
