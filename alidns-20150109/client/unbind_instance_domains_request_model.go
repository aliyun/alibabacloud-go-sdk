// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInstanceDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *UnbindInstanceDomainsRequest
	GetDomainNames() *string
	SetInstanceId(v string) *UnbindInstanceDomainsRequest
	GetInstanceId() *string
	SetLang(v string) *UnbindInstanceDomainsRequest
	GetLang() *string
}

type UnbindInstanceDomainsRequest struct {
	// The domain names.
	//
	// Separate multiple domain names with commas (,). Up to 100 domain names can be entered.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.net
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s UnbindInstanceDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindInstanceDomainsRequest) GoString() string {
	return s.String()
}

func (s *UnbindInstanceDomainsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *UnbindInstanceDomainsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnbindInstanceDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *UnbindInstanceDomainsRequest) SetDomainNames(v string) *UnbindInstanceDomainsRequest {
	s.DomainNames = &v
	return s
}

func (s *UnbindInstanceDomainsRequest) SetInstanceId(v string) *UnbindInstanceDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindInstanceDomainsRequest) SetLang(v string) *UnbindInstanceDomainsRequest {
	s.Lang = &v
	return s
}

func (s *UnbindInstanceDomainsRequest) Validate() error {
	return dara.Validate(s)
}
