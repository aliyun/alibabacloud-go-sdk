// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *BatchUpdateCdnDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *BatchUpdateCdnDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *BatchUpdateCdnDomainRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *BatchUpdateCdnDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *BatchUpdateCdnDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *BatchUpdateCdnDomainRequest
	GetTopLevelDomain() *string
}

type BatchUpdateCdnDomainRequest struct {
	// The accelerated domain names. You can specify one or more accelerated domain names. Separate domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.org
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
	// [{"content":"10.10.10.10","type":"ipaddr","priority":"20","port":80,"weight":"15"}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The root domain.
	//
	// > Do not set **Sources*	- and **TopLevelDomain*	- at the same time. If you set **Sources*	- and **TopLevelDomain*	- at the same time, **TopLevelDomain*	- does not take effect.
	//
	// example:
	//
	// example.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s BatchUpdateCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateCdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchUpdateCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchUpdateCdnDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BatchUpdateCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchUpdateCdnDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *BatchUpdateCdnDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *BatchUpdateCdnDomainRequest) SetDomainName(v string) *BatchUpdateCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetOwnerId(v int64) *BatchUpdateCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetResourceGroupId(v string) *BatchUpdateCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetSecurityToken(v string) *BatchUpdateCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetSources(v string) *BatchUpdateCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetTopLevelDomain(v string) *BatchUpdateCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
