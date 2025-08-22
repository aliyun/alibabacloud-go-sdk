// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDcdnIpaDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckUrl(v string) *AddDcdnIpaDomainRequest
	GetCheckUrl() *string
	SetDomainName(v string) *AddDcdnIpaDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *AddDcdnIpaDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddDcdnIpaDomainRequest
	GetOwnerId() *int64
	SetProtocol(v string) *AddDcdnIpaDomainRequest
	GetProtocol() *string
	SetResourceGroupId(v string) *AddDcdnIpaDomainRequest
	GetResourceGroupId() *string
	SetScope(v string) *AddDcdnIpaDomainRequest
	GetScope() *string
	SetSecurityToken(v string) *AddDcdnIpaDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *AddDcdnIpaDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *AddDcdnIpaDomainRequest
	GetTopLevelDomain() *string
}

type AddDcdnIpaDomainRequest struct {
	// The URL that is used for health checks.
	//
	// example:
	//
	// example.com/image_01.png
	CheckUrl *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	// The domain name to be added to IPA.
	//
	// A wildcard domain that starts with a period (.) is supported, such as .example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **udp**
	//
	// 	- **tcp**
	//
	// **
	//
	// **Description*	- For example: `{"protocol":"udp"}`.
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
	//
	// example:
	//
	// rg
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The acceleration region. Default value: domestic. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// 	- **global**: global
	//
	// example:
	//
	// domestic
	Scope         *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the addresses of origin servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"content":"10.10.10.10","type":"ipaddr","priority":"20","port":80,"weight":"15"}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The top-level domain.
	//
	// example:
	//
	// *.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddDcdnIpaDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDcdnIpaDomainRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *AddDcdnIpaDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDcdnIpaDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddDcdnIpaDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddDcdnIpaDomainRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *AddDcdnIpaDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddDcdnIpaDomainRequest) GetScope() *string {
	return s.Scope
}

func (s *AddDcdnIpaDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddDcdnIpaDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *AddDcdnIpaDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *AddDcdnIpaDomainRequest) SetCheckUrl(v string) *AddDcdnIpaDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetDomainName(v string) *AddDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetOwnerAccount(v string) *AddDcdnIpaDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetOwnerId(v int64) *AddDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetProtocol(v string) *AddDcdnIpaDomainRequest {
	s.Protocol = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetResourceGroupId(v string) *AddDcdnIpaDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetScope(v string) *AddDcdnIpaDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetSecurityToken(v string) *AddDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetSources(v string) *AddDcdnIpaDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) SetTopLevelDomain(v string) *AddDcdnIpaDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *AddDcdnIpaDomainRequest) Validate() error {
	return dara.Validate(s)
}
