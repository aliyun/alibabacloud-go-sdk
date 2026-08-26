// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckUrl(v string) *AddLiveDomainRequest
	GetCheckUrl() *string
	SetDomainName(v string) *AddLiveDomainRequest
	GetDomainName() *string
	SetLiveDomainType(v string) *AddLiveDomainRequest
	GetLiveDomainType() *string
	SetOwnerAccount(v string) *AddLiveDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddLiveDomainRequest
	GetOwnerId() *int64
	SetRegion(v string) *AddLiveDomainRequest
	GetRegion() *string
	SetResourceGroupId(v string) *AddLiveDomainRequest
	GetResourceGroupId() *string
	SetScope(v string) *AddLiveDomainRequest
	GetScope() *string
	SetSecurityToken(v string) *AddLiveDomainRequest
	GetSecurityToken() *string
	SetTag(v []*AddLiveDomainRequestTag) *AddLiveDomainRequest
	GetTag() []*AddLiveDomainRequestTag
	SetTopLevelDomain(v string) *AddLiveDomainRequest
	GetTopLevelDomain() *string
}

type AddLiveDomainRequest struct {
	// The health check URL.
	//
	// example:
	//
	// http://demo.aliyundoc.com/status.html
	CheckUrl *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	// The ingest domain or streaming domain to be connected to ApsaraVideo Live. Wildcard domain names are supported and must start with a period (.).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. Valid values:
	//
	// - **liveVideo**: streaming domain. If you set DomainName (the domain name to be connected to ApsaraVideo Live) to a streaming domain, you must set this parameter to liveVideo.
	//
	// - **liveEdge**: edge ingest domain. If you set DomainName (the domain name to be connected to ApsaraVideo Live) to an ingest domain, you must set this parameter to liveEdge.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveVideo
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit information of the live streaming domain name. Valid values:
	//
	// - **cn-beijing**: Beijing.
	//
	// - **cn-shanghai**: Shanghai.
	//
	// - **cn-shenzhen**: Shenzhen.
	//
	// - **cn-qingdao**: Qingdao.
	//
	// - **ap-southeast-1**: Singapore.
	//
	// - **eu-central-1**: Germany.
	//
	// - **ap-northeast-1**: Tokyo.
	//
	// - **ap-southeast-5**: Jakarta.
	//
	// >Region (unit information of the live streaming domain name) and Scope (acceleration region) do not restrict each other.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource group ID. For more information about resource groups, see [What is a resource group](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-aekzw******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The acceleration region. This parameter takes effect for international users and China site users at L3 or above. Valid values:
	//
	// - **domestic*	- (default): the Chinese mainland.
	//
	// - **overseas**: outside the Chinese mainland, including Hong Kong (China), Macao (China), and Taiwan (China).
	//
	// - **global**: global acceleration.
	//
	// example:
	//
	// domestic
	Scope         *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The list of tags.
	Tag []*AddLiveDomainRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The top-level domain name for access.
	//
	// example:
	//
	// learn.aliyundoc.com
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddLiveDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDomainRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *AddLiveDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveDomainRequest) GetLiveDomainType() *string {
	return s.LiveDomainType
}

func (s *AddLiveDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddLiveDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveDomainRequest) GetRegion() *string {
	return s.Region
}

func (s *AddLiveDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddLiveDomainRequest) GetScope() *string {
	return s.Scope
}

func (s *AddLiveDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddLiveDomainRequest) GetTag() []*AddLiveDomainRequestTag {
	return s.Tag
}

func (s *AddLiveDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *AddLiveDomainRequest) SetCheckUrl(v string) *AddLiveDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddLiveDomainRequest) SetDomainName(v string) *AddLiveDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveDomainRequest) SetLiveDomainType(v string) *AddLiveDomainRequest {
	s.LiveDomainType = &v
	return s
}

func (s *AddLiveDomainRequest) SetOwnerAccount(v string) *AddLiveDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddLiveDomainRequest) SetOwnerId(v int64) *AddLiveDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveDomainRequest) SetRegion(v string) *AddLiveDomainRequest {
	s.Region = &v
	return s
}

func (s *AddLiveDomainRequest) SetResourceGroupId(v string) *AddLiveDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddLiveDomainRequest) SetScope(v string) *AddLiveDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddLiveDomainRequest) SetSecurityToken(v string) *AddLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveDomainRequest) SetTag(v []*AddLiveDomainRequestTag) *AddLiveDomainRequest {
	s.Tag = v
	return s
}

func (s *AddLiveDomainRequest) SetTopLevelDomain(v string) *AddLiveDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *AddLiveDomainRequest) Validate() error {
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

type AddLiveDomainRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddLiveDomainRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainRequestTag) GoString() string {
	return s.String()
}

func (s *AddLiveDomainRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddLiveDomainRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddLiveDomainRequestTag) SetKey(v string) *AddLiveDomainRequestTag {
	s.Key = &v
	return s
}

func (s *AddLiveDomainRequestTag) SetValue(v string) *AddLiveDomainRequestTag {
	s.Value = &v
	return s
}

func (s *AddLiveDomainRequestTag) Validate() error {
	return dara.Validate(s)
}
