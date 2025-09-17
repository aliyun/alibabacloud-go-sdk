// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindIdpListByLoginIdentifierResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdpInfos(v []*FindIdpListByLoginIdentifierResponseBodyIdpInfos) *FindIdpListByLoginIdentifierResponseBody
	GetIdpInfos() []*FindIdpListByLoginIdentifierResponseBodyIdpInfos
	SetOfficeSiteInfo(v *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) *FindIdpListByLoginIdentifierResponseBody
	GetOfficeSiteInfo() *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo
	SetPopRegionConfig(v map[string]*string) *FindIdpListByLoginIdentifierResponseBody
	GetPopRegionConfig() map[string]*string
	SetProfileRegion(v string) *FindIdpListByLoginIdentifierResponseBody
	GetProfileRegion() *string
	SetRequestId(v string) *FindIdpListByLoginIdentifierResponseBody
	GetRequestId() *string
	SetTenantAliasInfo(v *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) *FindIdpListByLoginIdentifierResponseBody
	GetTenantAliasInfo() *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo
}

type FindIdpListByLoginIdentifierResponseBody struct {
	IdpInfos        []*FindIdpListByLoginIdentifierResponseBodyIdpInfos     `json:"IdpInfos,omitempty" xml:"IdpInfos,omitempty" type:"Repeated"`
	OfficeSiteInfo  *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo `json:"OfficeSiteInfo,omitempty" xml:"OfficeSiteInfo,omitempty" type:"Struct"`
	PopRegionConfig map[string]*string                                      `json:"PopRegionConfig,omitempty" xml:"PopRegionConfig,omitempty"`
	// example:
	//
	// cn_hangzhou
	ProfileRegion *string `json:"ProfileRegion,omitempty" xml:"ProfileRegion,omitempty"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169-D3F8D19C****
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantAliasInfo *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo `json:"TenantAliasInfo,omitempty" xml:"TenantAliasInfo,omitempty" type:"Struct"`
}

func (s FindIdpListByLoginIdentifierResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBody) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierResponseBody) GetIdpInfos() []*FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	return s.IdpInfos
}

func (s *FindIdpListByLoginIdentifierResponseBody) GetOfficeSiteInfo() *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo {
	return s.OfficeSiteInfo
}

func (s *FindIdpListByLoginIdentifierResponseBody) GetPopRegionConfig() map[string]*string {
	return s.PopRegionConfig
}

func (s *FindIdpListByLoginIdentifierResponseBody) GetProfileRegion() *string {
	return s.ProfileRegion
}

func (s *FindIdpListByLoginIdentifierResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindIdpListByLoginIdentifierResponseBody) GetTenantAliasInfo() *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo {
	return s.TenantAliasInfo
}

func (s *FindIdpListByLoginIdentifierResponseBody) SetIdpInfos(v []*FindIdpListByLoginIdentifierResponseBodyIdpInfos) *FindIdpListByLoginIdentifierResponseBody {
	s.IdpInfos = v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBody) SetOfficeSiteInfo(v *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) *FindIdpListByLoginIdentifierResponseBody {
	s.OfficeSiteInfo = v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBody) SetPopRegionConfig(v map[string]*string) *FindIdpListByLoginIdentifierResponseBody {
	s.PopRegionConfig = v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBody) SetProfileRegion(v string) *FindIdpListByLoginIdentifierResponseBody {
	s.ProfileRegion = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBody) SetRequestId(v string) *FindIdpListByLoginIdentifierResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBody) SetTenantAliasInfo(v *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) *FindIdpListByLoginIdentifierResponseBody {
	s.TenantAliasInfo = v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBody) Validate() error {
	return dara.Validate(s)
}

type FindIdpListByLoginIdentifierResponseBodyIdpInfos struct {
	// example:
	//
	// simple
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// null
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// example:
	//
	// idp-hlyexfvwert9m8****
	IdpId       *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	IdpName     *string `json:"IdpName,omitempty" xml:"IdpName,omitempty"`
	IdpNameEN   *string `json:"IdpNameEN,omitempty" xml:"IdpNameEN,omitempty"`
	IdpProvider *string `json:"IdpProvider,omitempty" xml:"IdpProvider,omitempty"`
	// example:
	//
	// true
	JumpSwitch *string `json:"JumpSwitch,omitempty" xml:"JumpSwitch,omitempty"`
	// example:
	//
	// SAML
	SsoProtocol   *string `json:"SsoProtocol,omitempty" xml:"SsoProtocol,omitempty"`
	SsoServiceUrl *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s FindIdpListByLoginIdentifierResponseBodyIdpInfos) String() string {
	return dara.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBodyIdpInfos) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetAccountType() *string {
	return s.AccountType
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetCookies() *string {
	return s.Cookies
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetIdpId() *string {
	return s.IdpId
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetIdpName() *string {
	return s.IdpName
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetIdpNameEN() *string {
	return s.IdpNameEN
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetIdpProvider() *string {
	return s.IdpProvider
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetJumpSwitch() *string {
	return s.JumpSwitch
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetSsoProtocol() *string {
	return s.SsoProtocol
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) GetSsoServiceUrl() *string {
	return s.SsoServiceUrl
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetAccountType(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.AccountType = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetCookies(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.Cookies = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetIdpId(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.IdpId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetIdpName(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.IdpName = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetIdpNameEN(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.IdpNameEN = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetIdpProvider(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.IdpProvider = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetJumpSwitch(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.JumpSwitch = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetSsoProtocol(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.SsoProtocol = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) SetSsoServiceUrl(v string) *FindIdpListByLoginIdentifierResponseBodyIdpInfos {
	s.SsoServiceUrl = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyIdpInfos) Validate() error {
	return dara.Validate(s)
}

type FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// cn-shanghai+dir-448204****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// 26842
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SsoServiceUrl *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) String() string {
	return dara.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) GetAccessType() *string {
	return s.AccessType
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) GetProviderId() *string {
	return s.ProviderId
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) GetSsoServiceUrl() *string {
	return s.SsoServiceUrl
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) SetAccessType(v string) *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo {
	s.AccessType = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) SetOfficeSiteId(v string) *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo {
	s.OfficeSiteId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) SetProviderId(v string) *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo {
	s.ProviderId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) SetRegionId(v string) *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo {
	s.RegionId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) SetSsoServiceUrl(v string) *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo {
	s.SsoServiceUrl = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) Validate() error {
	return dara.Validate(s)
}

type FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// Or09****
	TenantAlias *string `json:"TenantAlias,omitempty" xml:"TenantAlias,omitempty"`
}

func (s FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) String() string {
	return dara.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) GetAccessType() *string {
	return s.AccessType
}

func (s *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) GetTenantAlias() *string {
	return s.TenantAlias
}

func (s *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) SetAccessType(v string) *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo {
	s.AccessType = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) SetTenantAlias(v string) *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo {
	s.TenantAlias = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) Validate() error {
	return dara.Validate(s)
}
