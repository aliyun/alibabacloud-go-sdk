// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckUrl(v string) *AddVodDomainRequest
	GetCheckUrl() *string
	SetDomainName(v string) *AddVodDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *AddVodDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddVodDomainRequest
	GetOwnerId() *int64
	SetScope(v string) *AddVodDomainRequest
	GetScope() *string
	SetSecurityToken(v string) *AddVodDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *AddVodDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *AddVodDomainRequest
	GetTopLevelDomain() *string
}

type AddVodDomainRequest struct {
	// The URL that is used for health checks.
	//
	// example:
	//
	// www.example.com/test.html
	CheckUrl *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	// The domain name that you want to accelerate. Wildcard domain names that start with periods (.) are supported. Example: .example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is applicable to users of level 3 or higher in the Chinese mainland and users outside the Chinese mainland. Default value: domestic. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// 	- **global**: regions in and outside the Chinese mainland
	//
	// example:
	//
	// domestic
	Scope         *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the addresses of origin servers. For more information, see the **Sources*	- table in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"content":"1.1.1.1","type":"ipaddr","priority":"20","port":80}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The top-level domain.
	//
	// example:
	//
	// example.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddVodDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVodDomainRequest) GoString() string {
	return s.String()
}

func (s *AddVodDomainRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *AddVodDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddVodDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddVodDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddVodDomainRequest) GetScope() *string {
	return s.Scope
}

func (s *AddVodDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddVodDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *AddVodDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *AddVodDomainRequest) SetCheckUrl(v string) *AddVodDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddVodDomainRequest) SetDomainName(v string) *AddVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddVodDomainRequest) SetOwnerAccount(v string) *AddVodDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddVodDomainRequest) SetOwnerId(v int64) *AddVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVodDomainRequest) SetScope(v string) *AddVodDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddVodDomainRequest) SetSecurityToken(v string) *AddVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddVodDomainRequest) SetSources(v string) *AddVodDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddVodDomainRequest) SetTopLevelDomain(v string) *AddVodDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *AddVodDomainRequest) Validate() error {
	return dara.Validate(s)
}
