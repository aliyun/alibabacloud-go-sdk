// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckUrl(v string) *BatchAddDcdnDomainRequest
	GetCheckUrl() *string
	SetDomainName(v string) *BatchAddDcdnDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *BatchAddDcdnDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchAddDcdnDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *BatchAddDcdnDomainRequest
	GetResourceGroupId() *string
	SetScope(v string) *BatchAddDcdnDomainRequest
	GetScope() *string
	SetSecurityToken(v string) *BatchAddDcdnDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *BatchAddDcdnDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *BatchAddDcdnDomainRequest
	GetTopLevelDomain() *string
}

type BatchAddDcdnDomainRequest struct {
	// The URL that is used for health checks.
	//
	// example:
	//
	// www.example.com/test.html
	CheckUrl *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	// You can add up to 20 domain names to DCDN for each of your Alibaba Cloud account. If you specify multiple domain names, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.org
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group. If you do not specify a value for this parameter, the system uses the ID of the default resource group.
	//
	// example:
	//
	// testID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The acceleration region. Default value: domestic. Valid values:
	//
	// 	- domestic: Chinese mainland
	//
	// 	- overseas: global (excluding the Chinese mainland)
	//
	// 	- global: global
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
	// example.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s BatchAddDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchAddDcdnDomainRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *BatchAddDcdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchAddDcdnDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchAddDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchAddDcdnDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BatchAddDcdnDomainRequest) GetScope() *string {
	return s.Scope
}

func (s *BatchAddDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchAddDcdnDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *BatchAddDcdnDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *BatchAddDcdnDomainRequest) SetCheckUrl(v string) *BatchAddDcdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetDomainName(v string) *BatchAddDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetOwnerAccount(v string) *BatchAddDcdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetOwnerId(v int64) *BatchAddDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetResourceGroupId(v string) *BatchAddDcdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetScope(v string) *BatchAddDcdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetSecurityToken(v string) *BatchAddDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetSources(v string) *BatchAddDcdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) SetTopLevelDomain(v string) *BatchAddDcdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *BatchAddDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
