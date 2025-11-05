// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdnType(v string) *AddCdnDomainRequest
	GetCdnType() *string
	SetCheckUrl(v string) *AddCdnDomainRequest
	GetCheckUrl() *string
	SetDomainName(v string) *AddCdnDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *AddCdnDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddCdnDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *AddCdnDomainRequest
	GetResourceGroupId() *string
	SetScope(v string) *AddCdnDomainRequest
	GetScope() *string
	SetSecurityToken(v string) *AddCdnDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *AddCdnDomainRequest
	GetSources() *string
	SetTag(v []*AddCdnDomainRequestTag) *AddCdnDomainRequest
	GetTag() []*AddCdnDomainRequestTag
	SetTopLevelDomain(v string) *AddCdnDomainRequest
	GetTopLevelDomain() *string
}

type AddCdnDomainRequest struct {
	// The workload type of the accelerated domain name. Valid values:
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
	// The URL that is used to check the accessibility of the origin server.
	//
	// example:
	//
	// www.example.com/test.html
	CheckUrl *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	// The domain name that you want to add to Alibaba Cloud CDN.
	//
	// A wildcard domain that starts with a period (.) is supported, such as .example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// .example.com
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// If you do not set this parameter, the system uses the ID of the default resource group.
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
	// Details about the tags. You can specify up to 20 tags.
	Tag []*AddCdnDomainRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The top-level domain.
	//
	// example:
	//
	// example.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *AddCdnDomainRequest) GetCdnType() *string {
	return s.CdnType
}

func (s *AddCdnDomainRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *AddCdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddCdnDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCdnDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddCdnDomainRequest) GetScope() *string {
	return s.Scope
}

func (s *AddCdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddCdnDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *AddCdnDomainRequest) GetTag() []*AddCdnDomainRequestTag {
	return s.Tag
}

func (s *AddCdnDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *AddCdnDomainRequest) SetCdnType(v string) *AddCdnDomainRequest {
	s.CdnType = &v
	return s
}

func (s *AddCdnDomainRequest) SetCheckUrl(v string) *AddCdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddCdnDomainRequest) SetDomainName(v string) *AddCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddCdnDomainRequest) SetOwnerAccount(v string) *AddCdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddCdnDomainRequest) SetOwnerId(v int64) *AddCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCdnDomainRequest) SetResourceGroupId(v string) *AddCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddCdnDomainRequest) SetScope(v string) *AddCdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddCdnDomainRequest) SetSecurityToken(v string) *AddCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddCdnDomainRequest) SetSources(v string) *AddCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddCdnDomainRequest) SetTag(v []*AddCdnDomainRequestTag) *AddCdnDomainRequest {
	s.Tag = v
	return s
}

func (s *AddCdnDomainRequest) SetTopLevelDomain(v string) *AddCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *AddCdnDomainRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddCdnDomainRequestTag struct {
	// The key of the tag. Valid values of N: **1 to 20**.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. Valid values of N: **1 to 20**.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddCdnDomainRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddCdnDomainRequestTag) GoString() string {
	return s.String()
}

func (s *AddCdnDomainRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddCdnDomainRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddCdnDomainRequestTag) SetKey(v string) *AddCdnDomainRequestTag {
	s.Key = &v
	return s
}

func (s *AddCdnDomainRequestTag) SetValue(v string) *AddCdnDomainRequestTag {
	s.Value = &v
	return s
}

func (s *AddCdnDomainRequestTag) Validate() error {
	return dara.Validate(s)
}
