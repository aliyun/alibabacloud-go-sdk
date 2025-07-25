// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstanceDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BindInstanceDomainsRequest
	GetDomainNames() *string
	SetInstanceId(v string) *BindInstanceDomainsRequest
	GetInstanceId() *string
	SetLang(v string) *BindInstanceDomainsRequest
	GetLang() *string
}

type BindInstanceDomainsRequest struct {
	// The domain names.
	//
	// >  Separate multiple domain names with commas (,). Up to 100 domain names can be entered.
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
	// sdfasdf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s BindInstanceDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s BindInstanceDomainsRequest) GoString() string {
	return s.String()
}

func (s *BindInstanceDomainsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BindInstanceDomainsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindInstanceDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *BindInstanceDomainsRequest) SetDomainNames(v string) *BindInstanceDomainsRequest {
	s.DomainNames = &v
	return s
}

func (s *BindInstanceDomainsRequest) SetInstanceId(v string) *BindInstanceDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *BindInstanceDomainsRequest) SetLang(v string) *BindInstanceDomainsRequest {
	s.Lang = &v
	return s
}

func (s *BindInstanceDomainsRequest) Validate() error {
	return dara.Validate(s)
}
