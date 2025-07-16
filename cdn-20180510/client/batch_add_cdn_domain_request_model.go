// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdnType(v string) *BatchAddCdnDomainRequest
	GetCdnType() *string
	SetCheckUrl(v string) *BatchAddCdnDomainRequest
	GetCheckUrl() *string
	SetDomainName(v string) *BatchAddCdnDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *BatchAddCdnDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchAddCdnDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *BatchAddCdnDomainRequest
	GetResourceGroupId() *string
	SetScope(v string) *BatchAddCdnDomainRequest
	GetScope() *string
	SetSecurityToken(v string) *BatchAddCdnDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *BatchAddCdnDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *BatchAddCdnDomainRequest
	GetTopLevelDomain() *string
}

type BatchAddCdnDomainRequest struct {
	// The workload type of the domain name to accelerate. Valid values:
	//
	// 	- **web**: images and small files
	//
	// 	- **download**: large files
	//
	// 	- **video**: on-demand video and audio streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// web
	CdnType *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	// The URL that is used for health checks.
	//
	// example:
	//
	// url
	CheckUrl *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	// The domain names that you want to add to Alibaba Cloud CDN. Separate domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,aliyundoc.com
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group. If you do not specify a value for this parameter, the system uses the ID of the default resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The acceleration region. Default value: domestic. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: global (excluding the Chinese mainland)
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
	// [
	//
	//       {
	//
	//             "content": "192.0.2.0",
	//
	//             "type": "ipaddr",
	//
	//             "priority": "20",
	//
	//             "port": 80,
	//
	//             "weight": "15"
	//
	//       }
	//
	// ]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The top-level domain.
	//
	// example:
	//
	// example.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s BatchAddCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchAddCdnDomainRequest) GetCdnType() *string {
	return s.CdnType
}

func (s *BatchAddCdnDomainRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *BatchAddCdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchAddCdnDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchAddCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchAddCdnDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BatchAddCdnDomainRequest) GetScope() *string {
	return s.Scope
}

func (s *BatchAddCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchAddCdnDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *BatchAddCdnDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *BatchAddCdnDomainRequest) SetCdnType(v string) *BatchAddCdnDomainRequest {
	s.CdnType = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetCheckUrl(v string) *BatchAddCdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetDomainName(v string) *BatchAddCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetOwnerAccount(v string) *BatchAddCdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetOwnerId(v int64) *BatchAddCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetResourceGroupId(v string) *BatchAddCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetScope(v string) *BatchAddCdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetSecurityToken(v string) *BatchAddCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetSources(v string) *BatchAddCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetTopLevelDomain(v string) *BatchAddCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *BatchAddCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
