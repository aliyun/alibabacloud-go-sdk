// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyCdnDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *ModifyCdnDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyCdnDomainRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *ModifyCdnDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *ModifyCdnDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *ModifyCdnDomainRequest
	GetTopLevelDomain() *string
}

type ModifyCdnDomainRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the addresses of origin servers.
	//
	// > Do not set **Sources*	- and **TopLevelDomain*	- at the same time. If you set **Sources*	- and **TopLevelDomain*	- at the same time, **TopLevelDomain*	- does not take effect.
	//
	// example:
	//
	// [{"content":"1.1.1.1","type":"ipaddr","priority":"20","port":80,"weight":"15"}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The root domain. To add a root domain name, you must be added to the whitelist specified by the CDN_TOP_LEVEL_DOMAIN_GREY_USER_LIST parameter.
	//
	// > Do not set **Sources*	- and **TopLevelDomain*	- at the same time. If you set **Sources*	- and **TopLevelDomain*	- at the same time, **TopLevelDomain*	- does not take effect.
	//
	// example:
	//
	// aliyundoc.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s ModifyCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCdnDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyCdnDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *ModifyCdnDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *ModifyCdnDomainRequest) SetDomainName(v string) *ModifyCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetOwnerId(v int64) *ModifyCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetResourceGroupId(v string) *ModifyCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetSecurityToken(v string) *ModifyCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetSources(v string) *ModifyCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetTopLevelDomain(v string) *ModifyCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *ModifyCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
