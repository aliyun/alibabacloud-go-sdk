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
	// A list of domain names.
	//
	// > Separate multiple domain names with a comma (,). You can specify up to 100 domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.net
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The ID of the Alibaba Cloud DNS instance. You can call the [ListCloudGtmInstances](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-listcloudgtminstances) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default value: zh
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
