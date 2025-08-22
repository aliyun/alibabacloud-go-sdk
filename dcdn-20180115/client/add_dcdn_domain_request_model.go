// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckUrl(v string) *AddDcdnDomainRequest
	GetCheckUrl() *string
	SetDomainName(v string) *AddDcdnDomainRequest
	GetDomainName() *string
	SetFunctionType(v string) *AddDcdnDomainRequest
	GetFunctionType() *string
	SetOwnerAccount(v string) *AddDcdnDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddDcdnDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *AddDcdnDomainRequest
	GetResourceGroupId() *string
	SetScene(v string) *AddDcdnDomainRequest
	GetScene() *string
	SetScope(v string) *AddDcdnDomainRequest
	GetScope() *string
	SetSecurityToken(v string) *AddDcdnDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *AddDcdnDomainRequest
	GetSources() *string
	SetTag(v []*AddDcdnDomainRequestTag) *AddDcdnDomainRequest
	GetTag() []*AddDcdnDomainRequestTag
	SetTopLevelDomain(v string) *AddDcdnDomainRequest
	GetTopLevelDomain() *string
}

type AddDcdnDomainRequest struct {
	// The URL that is used for health checks.
	//
	// example:
	//
	// example.com
	CheckUrl *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	// The domain name that you want to add. You can specify only one domain name in each request.
	//
	// Wildcard domain names are supported. A wildcard domain name must start with a period (.), such as .example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Computing service type. Valid values:
	//
	// 	- **routine**
	//
	// 	- **image**
	//
	// 	- **cloudFunction**
	//
	// example:
	//
	// routine
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Acceleration scen. Supported:
	//
	// 	- apiscene:API acceleration.
	//
	// 	- webservicescene: accelerate website business.
	//
	// 	- staticscene: video and graphic acceleration.
	//
	// 	- (Empty): no scene.
	//
	// example:
	//
	// apiscene
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The acceleration region. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// 	- **global**: global
	//
	// Default value: **domestic**.
	//
	// example:
	//
	// domestic
	Scope         *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the addresses of origin servers.
	//
	// example:
	//
	// [{"content":"10.10.10.10","type":"ipaddr","priority":"20","port":80}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The information about the tags.
	Tag []*AddDcdnDomainRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The top-level domain.
	//
	// example:
	//
	// yourTopLevelDomain
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDcdnDomainRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *AddDcdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDcdnDomainRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *AddDcdnDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddDcdnDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddDcdnDomainRequest) GetScene() *string {
	return s.Scene
}

func (s *AddDcdnDomainRequest) GetScope() *string {
	return s.Scope
}

func (s *AddDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddDcdnDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *AddDcdnDomainRequest) GetTag() []*AddDcdnDomainRequestTag {
	return s.Tag
}

func (s *AddDcdnDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *AddDcdnDomainRequest) SetCheckUrl(v string) *AddDcdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddDcdnDomainRequest) SetDomainName(v string) *AddDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDcdnDomainRequest) SetFunctionType(v string) *AddDcdnDomainRequest {
	s.FunctionType = &v
	return s
}

func (s *AddDcdnDomainRequest) SetOwnerAccount(v string) *AddDcdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddDcdnDomainRequest) SetOwnerId(v int64) *AddDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddDcdnDomainRequest) SetResourceGroupId(v string) *AddDcdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddDcdnDomainRequest) SetScene(v string) *AddDcdnDomainRequest {
	s.Scene = &v
	return s
}

func (s *AddDcdnDomainRequest) SetScope(v string) *AddDcdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddDcdnDomainRequest) SetSecurityToken(v string) *AddDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddDcdnDomainRequest) SetSources(v string) *AddDcdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddDcdnDomainRequest) SetTag(v []*AddDcdnDomainRequestTag) *AddDcdnDomainRequest {
	s.Tag = v
	return s
}

func (s *AddDcdnDomainRequest) SetTopLevelDomain(v string) *AddDcdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *AddDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}

type AddDcdnDomainRequestTag struct {
	// The key of a tag. Valid values of N: **1 to 20**.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag. Valid values of N: **1 to 20**.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddDcdnDomainRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddDcdnDomainRequestTag) GoString() string {
	return s.String()
}

func (s *AddDcdnDomainRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddDcdnDomainRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddDcdnDomainRequestTag) SetKey(v string) *AddDcdnDomainRequestTag {
	s.Key = &v
	return s
}

func (s *AddDcdnDomainRequestTag) SetValue(v string) *AddDcdnDomainRequestTag {
	s.Value = &v
	return s
}

func (s *AddDcdnDomainRequestTag) Validate() error {
	return dara.Validate(s)
}
