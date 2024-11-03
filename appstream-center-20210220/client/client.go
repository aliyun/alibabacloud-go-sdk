// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type FindIdpListByLoginIdentifierRequest struct {
	AvailableFeatures map[string]*string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
	// example:
	//
	// pc
	ClientChannel *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	// example:
	//
	// 370b56f8-2812-4b6c-bfa6-2560791c****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Enterprise\\" 10.0 (Build 14393)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Aliyun123***
	LoginIdentifier *string   `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	SupportTypes    []*string `json:"SupportTypes,omitempty" xml:"SupportTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s FindIdpListByLoginIdentifierRequest) String() string {
	return tea.Prettify(s)
}

func (s FindIdpListByLoginIdentifierRequest) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierRequest) SetAvailableFeatures(v map[string]*string) *FindIdpListByLoginIdentifierRequest {
	s.AvailableFeatures = v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientChannel(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientChannel = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientId(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientIp(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientIp = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientOS(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientOS = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetClientVersion(v string) *FindIdpListByLoginIdentifierRequest {
	s.ClientVersion = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetLoginIdentifier(v string) *FindIdpListByLoginIdentifierRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetSupportTypes(v []*string) *FindIdpListByLoginIdentifierRequest {
	s.SupportTypes = v
	return s
}

func (s *FindIdpListByLoginIdentifierRequest) SetUuid(v string) *FindIdpListByLoginIdentifierRequest {
	s.Uuid = &v
	return s
}

type FindIdpListByLoginIdentifierShrinkRequest struct {
	AvailableFeaturesShrink *string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
	// example:
	//
	// pc
	ClientChannel *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	// example:
	//
	// 370b56f8-2812-4b6c-bfa6-2560791c****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Enterprise\\" 10.0 (Build 14393)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Aliyun123***
	LoginIdentifier *string   `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	SupportTypes    []*string `json:"SupportTypes,omitempty" xml:"SupportTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s FindIdpListByLoginIdentifierShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FindIdpListByLoginIdentifierShrinkRequest) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetAvailableFeaturesShrink(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.AvailableFeaturesShrink = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientChannel(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientChannel = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientId(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientIp(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientIp = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientOS(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientOS = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetClientVersion(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetLoginIdentifier(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetSupportTypes(v []*string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.SupportTypes = v
	return s
}

func (s *FindIdpListByLoginIdentifierShrinkRequest) SetUuid(v string) *FindIdpListByLoginIdentifierShrinkRequest {
	s.Uuid = &v
	return s
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
	return tea.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBodyIdpInfos) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBodyOfficeSiteInfo) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) SetAccessType(v string) *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo {
	s.AccessType = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo) SetTenantAlias(v string) *FindIdpListByLoginIdentifierResponseBodyTenantAliasInfo {
	s.TenantAlias = &v
	return s
}

type FindIdpListByLoginIdentifierResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindIdpListByLoginIdentifierResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindIdpListByLoginIdentifierResponse) String() string {
	return tea.Prettify(s)
}

func (s FindIdpListByLoginIdentifierResponse) GoString() string {
	return s.String()
}

func (s *FindIdpListByLoginIdentifierResponse) SetHeaders(v map[string]*string) *FindIdpListByLoginIdentifierResponse {
	s.Headers = v
	return s
}

func (s *FindIdpListByLoginIdentifierResponse) SetStatusCode(v int32) *FindIdpListByLoginIdentifierResponse {
	s.StatusCode = &v
	return s
}

func (s *FindIdpListByLoginIdentifierResponse) SetBody(v *FindIdpListByLoginIdentifierResponseBody) *FindIdpListByLoginIdentifierResponse {
	s.Body = v
	return s
}

type GetLoginTokenRequest struct {
	// example:
	//
	// 182901
	AuthenticationCode *string            `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	AvailableFeatures  map[string]*string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 05967f80-6f51-46cb-a27c-****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Pro\\" 10.0 (Build 22631)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 7.3.0-20240619.143924
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// SsoTokenLogin
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// example:
	//
	// cn-beijing+dir-j9dd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// null
	EncryptedFingerPrintData *string `json:"EncryptedFingerPrintData,omitempty" xml:"EncryptedFingerPrintData,omitempty"`
	// example:
	//
	// 4d7****8e90bb0484fc
	EncryptedKey *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	// example:
	//
	// 04d7****8e90bb0484fc;gJ1GLca1vQRRqQbRvByU0A==;5kOWZE7AtbQhki+4LAo69A==
	EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// null
	FingerPrintData *string `json:"FingerPrintData,omitempty" xml:"FingerPrintData,omitempty"`
	// example:
	//
	// idp-iwntrlbb98q7v****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// null
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// true
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// example:
	//
	// 00ugamMAoO1+u1MlhmgCeZJ75q094x3H/4kg7ZaTI3f/joVTVPIpJgfv9JFkPLNxnQjblrvsByNas08mc6FtVWvQPOE68fqmt6QMM4UbRtahm8luxEXvicF58qSPXW1hxOtV/Ev6d92VBz2Bck/N4CYyjD0iLocfN8jkBnt231****
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// example:
	//
	// Fe04****
	LoginIdentifier *string `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	// example:
	//
	// null
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	MfaType   *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
	// example:
	//
	// INTERNET
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// Admin@1234****
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// example:
	//
	// cn-beijing+dir-j9dd****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// Admin@1234****
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
	// example:
	//
	// Admin@1234****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 1822727****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// example:
	//
	// 321123
	PhoneVerifyCode *string `json:"PhoneVerifyCode,omitempty" xml:"PhoneVerifyCode,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// c6f3cd91-65fc-4c7b-b189-2a73da0****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// null
	SsoExtendsCookies *string `json:"SsoExtendsCookies,omitempty" xml:"SsoExtendsCookies,omitempty"`
	// example:
	//
	// 04d707a6-fb23-44a7-aae7-8e90****
	SsoSessionToken *string `json:"SsoSessionToken,omitempty" xml:"SsoSessionToken,omitempty"`
	// example:
	//
	// 1234***
	TokenCode *string `json:"TokenCode,omitempty" xml:"TokenCode,omitempty"`
	// example:
	//
	// 04d707a6-fb23-44a7-aae7-8e90bb04****
	UmidToken *string `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// example:
	//
	// C50973691A6D2BE23F2CDD73B85B****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequest) SetAuthenticationCode(v string) *GetLoginTokenRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetAvailableFeatures(v map[string]*string) *GetLoginTokenRequest {
	s.AvailableFeatures = v
	return s
}

func (s *GetLoginTokenRequest) SetClientId(v string) *GetLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientOS(v string) *GetLoginTokenRequest {
	s.ClientOS = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientType(v string) *GetLoginTokenRequest {
	s.ClientType = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientVersion(v string) *GetLoginTokenRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetLoginTokenRequest) SetCurrentStage(v string) *GetLoginTokenRequest {
	s.CurrentStage = &v
	return s
}

func (s *GetLoginTokenRequest) SetDirectoryId(v string) *GetLoginTokenRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetLoginTokenRequest) SetEncryptedFingerPrintData(v string) *GetLoginTokenRequest {
	s.EncryptedFingerPrintData = &v
	return s
}

func (s *GetLoginTokenRequest) SetEncryptedKey(v string) *GetLoginTokenRequest {
	s.EncryptedKey = &v
	return s
}

func (s *GetLoginTokenRequest) SetEncryptedPassword(v string) *GetLoginTokenRequest {
	s.EncryptedPassword = &v
	return s
}

func (s *GetLoginTokenRequest) SetEndUserId(v string) *GetLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenRequest) SetFingerPrintData(v string) *GetLoginTokenRequest {
	s.FingerPrintData = &v
	return s
}

func (s *GetLoginTokenRequest) SetIdpId(v string) *GetLoginTokenRequest {
	s.IdpId = &v
	return s
}

func (s *GetLoginTokenRequest) SetImageUrl(v string) *GetLoginTokenRequest {
	s.ImageUrl = &v
	return s
}

func (s *GetLoginTokenRequest) SetKeepAlive(v bool) *GetLoginTokenRequest {
	s.KeepAlive = &v
	return s
}

func (s *GetLoginTokenRequest) SetKeepAliveToken(v string) *GetLoginTokenRequest {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenRequest) SetLoginIdentifier(v string) *GetLoginTokenRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *GetLoginTokenRequest) SetLoginName(v string) *GetLoginTokenRequest {
	s.LoginName = &v
	return s
}

func (s *GetLoginTokenRequest) SetMfaType(v string) *GetLoginTokenRequest {
	s.MfaType = &v
	return s
}

func (s *GetLoginTokenRequest) SetNetworkType(v string) *GetLoginTokenRequest {
	s.NetworkType = &v
	return s
}

func (s *GetLoginTokenRequest) SetNewPassword(v string) *GetLoginTokenRequest {
	s.NewPassword = &v
	return s
}

func (s *GetLoginTokenRequest) SetOfficeSiteId(v string) *GetLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetLoginTokenRequest) SetOldPassword(v string) *GetLoginTokenRequest {
	s.OldPassword = &v
	return s
}

func (s *GetLoginTokenRequest) SetPassword(v string) *GetLoginTokenRequest {
	s.Password = &v
	return s
}

func (s *GetLoginTokenRequest) SetPhone(v string) *GetLoginTokenRequest {
	s.Phone = &v
	return s
}

func (s *GetLoginTokenRequest) SetPhoneVerifyCode(v string) *GetLoginTokenRequest {
	s.PhoneVerifyCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetRegionId(v string) *GetLoginTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetLoginTokenRequest) SetSessionId(v string) *GetLoginTokenRequest {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenRequest) SetSsoExtendsCookies(v string) *GetLoginTokenRequest {
	s.SsoExtendsCookies = &v
	return s
}

func (s *GetLoginTokenRequest) SetSsoSessionToken(v string) *GetLoginTokenRequest {
	s.SsoSessionToken = &v
	return s
}

func (s *GetLoginTokenRequest) SetTokenCode(v string) *GetLoginTokenRequest {
	s.TokenCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetUmidToken(v string) *GetLoginTokenRequest {
	s.UmidToken = &v
	return s
}

func (s *GetLoginTokenRequest) SetUuid(v string) *GetLoginTokenRequest {
	s.Uuid = &v
	return s
}

type GetLoginTokenShrinkRequest struct {
	// example:
	//
	// 182901
	AuthenticationCode      *string `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	AvailableFeaturesShrink *string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 05967f80-6f51-46cb-a27c-****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Pro\\" 10.0 (Build 22631)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 7.3.0-20240619.143924
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// SsoTokenLogin
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// example:
	//
	// cn-beijing+dir-j9dd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// null
	EncryptedFingerPrintData *string `json:"EncryptedFingerPrintData,omitempty" xml:"EncryptedFingerPrintData,omitempty"`
	// example:
	//
	// 4d7****8e90bb0484fc
	EncryptedKey *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	// example:
	//
	// 04d7****8e90bb0484fc;gJ1GLca1vQRRqQbRvByU0A==;5kOWZE7AtbQhki+4LAo69A==
	EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// null
	FingerPrintData *string `json:"FingerPrintData,omitempty" xml:"FingerPrintData,omitempty"`
	// example:
	//
	// idp-iwntrlbb98q7v****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// null
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// true
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// example:
	//
	// 00ugamMAoO1+u1MlhmgCeZJ75q094x3H/4kg7ZaTI3f/joVTVPIpJgfv9JFkPLNxnQjblrvsByNas08mc6FtVWvQPOE68fqmt6QMM4UbRtahm8luxEXvicF58qSPXW1hxOtV/Ev6d92VBz2Bck/N4CYyjD0iLocfN8jkBnt231****
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// example:
	//
	// Fe04****
	LoginIdentifier *string `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	// example:
	//
	// null
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	MfaType   *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
	// example:
	//
	// INTERNET
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// Admin@1234****
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// example:
	//
	// cn-beijing+dir-j9dd****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// Admin@1234****
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
	// example:
	//
	// Admin@1234****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 1822727****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// example:
	//
	// 321123
	PhoneVerifyCode *string `json:"PhoneVerifyCode,omitempty" xml:"PhoneVerifyCode,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// c6f3cd91-65fc-4c7b-b189-2a73da0****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// null
	SsoExtendsCookies *string `json:"SsoExtendsCookies,omitempty" xml:"SsoExtendsCookies,omitempty"`
	// example:
	//
	// 04d707a6-fb23-44a7-aae7-8e90****
	SsoSessionToken *string `json:"SsoSessionToken,omitempty" xml:"SsoSessionToken,omitempty"`
	// example:
	//
	// 1234***
	TokenCode *string `json:"TokenCode,omitempty" xml:"TokenCode,omitempty"`
	// example:
	//
	// 04d707a6-fb23-44a7-aae7-8e90bb04****
	UmidToken *string `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// example:
	//
	// C50973691A6D2BE23F2CDD73B85B****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetLoginTokenShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenShrinkRequest) SetAuthenticationCode(v string) *GetLoginTokenShrinkRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetAvailableFeaturesShrink(v string) *GetLoginTokenShrinkRequest {
	s.AvailableFeaturesShrink = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetClientId(v string) *GetLoginTokenShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetClientOS(v string) *GetLoginTokenShrinkRequest {
	s.ClientOS = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetClientType(v string) *GetLoginTokenShrinkRequest {
	s.ClientType = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetClientVersion(v string) *GetLoginTokenShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetCurrentStage(v string) *GetLoginTokenShrinkRequest {
	s.CurrentStage = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetDirectoryId(v string) *GetLoginTokenShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetEncryptedFingerPrintData(v string) *GetLoginTokenShrinkRequest {
	s.EncryptedFingerPrintData = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetEncryptedKey(v string) *GetLoginTokenShrinkRequest {
	s.EncryptedKey = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetEncryptedPassword(v string) *GetLoginTokenShrinkRequest {
	s.EncryptedPassword = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetEndUserId(v string) *GetLoginTokenShrinkRequest {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetFingerPrintData(v string) *GetLoginTokenShrinkRequest {
	s.FingerPrintData = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetIdpId(v string) *GetLoginTokenShrinkRequest {
	s.IdpId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetImageUrl(v string) *GetLoginTokenShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetKeepAlive(v bool) *GetLoginTokenShrinkRequest {
	s.KeepAlive = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetKeepAliveToken(v string) *GetLoginTokenShrinkRequest {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetLoginIdentifier(v string) *GetLoginTokenShrinkRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetLoginName(v string) *GetLoginTokenShrinkRequest {
	s.LoginName = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetMfaType(v string) *GetLoginTokenShrinkRequest {
	s.MfaType = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetNetworkType(v string) *GetLoginTokenShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetNewPassword(v string) *GetLoginTokenShrinkRequest {
	s.NewPassword = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetOfficeSiteId(v string) *GetLoginTokenShrinkRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetOldPassword(v string) *GetLoginTokenShrinkRequest {
	s.OldPassword = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetPassword(v string) *GetLoginTokenShrinkRequest {
	s.Password = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetPhone(v string) *GetLoginTokenShrinkRequest {
	s.Phone = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetPhoneVerifyCode(v string) *GetLoginTokenShrinkRequest {
	s.PhoneVerifyCode = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetRegionId(v string) *GetLoginTokenShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetSessionId(v string) *GetLoginTokenShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetSsoExtendsCookies(v string) *GetLoginTokenShrinkRequest {
	s.SsoExtendsCookies = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetSsoSessionToken(v string) *GetLoginTokenShrinkRequest {
	s.SsoSessionToken = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetTokenCode(v string) *GetLoginTokenShrinkRequest {
	s.TokenCode = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetUmidToken(v string) *GetLoginTokenShrinkRequest {
	s.UmidToken = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetUuid(v string) *GetLoginTokenShrinkRequest {
	s.Uuid = &v
	return s
}

type GetLoginTokenResponseBody struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// ad
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// easthp***.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// example:
	//
	// alice***@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// idp-7ttvs4ove8bo5***
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// edu
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// 006YwvYMsesWWsDBZnVB+Wq9AvJDVIqOY3YCktvtb7+KxMb3ClnNlV8+l/knhZYrXUmeP06IzkjF+IgcZ3vZKOyMprDyFHjCy1r27FRE/U7+geWCl8iQ+yF8GaCRHfJEkC2+ROs93HkT4tfHxyY1J8W7O7ZQGUC/cdCvm+cCP6FIy73IUuPuVR6PcKYXIp***
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// example:
	//
	// test:wuying
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// v185fdd7f6d39fa7861981639366085772e150a390a5bb7b43c4e62440d94fc392b945770e1596cebe90085ce0af4d****
	LoginToken  *string                                 `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	MfaTypeList []*GetLoginTokenResponseBodyMfaTypeList `json:"MfaTypeList,omitempty" xml:"MfaTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// MFABind
	NextStage        *string                                    `json:"NextStage,omitempty" xml:"NextStage,omitempty"`
	OfficeSites      []*string                                  `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	PasswordStrategy *GetLoginTokenResponseBodyPasswordStrategy `json:"PasswordStrategy,omitempty" xml:"PasswordStrategy,omitempty" type:"Struct"`
	// example:
	//
	// 1826717****
	Phone *string            `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Props map[string]*string `json:"Props,omitempty" xml:"Props,omitempty"`
	// example:
	//
	// 5OCLLKKOJU5HPBX66H3QCTWY******
	QrCodePng *string `json:"QrCodePng,omitempty" xml:"QrCodePng,omitempty"`
	// example:
	//
	// PasswordExpired
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// EEA72491-B731-53D6-83ED-209769D6****
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskVerifyInfo *GetLoginTokenResponseBodyRiskVerifyInfo `json:"RiskVerifyInfo,omitempty" xml:"RiskVerifyInfo,omitempty" type:"Struct"`
	// example:
	//
	// 4JZNSDHDM3T6AZ4G2O5OWXBLLE4P****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// example:
	//
	// cc15c91c-821b-4edd-9af2-6df66cc****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// Rq201***
	TenantAlias *string `json:"TenantAlias,omitempty" xml:"TenantAlias,omitempty"`
	// example:
	//
	// 13747924304****
	TenantId    *int64                                  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantInfos []*GetLoginTokenResponseBodyTenantInfos `json:"TenantInfos,omitempty" xml:"TenantInfos,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// example:
	//
	// Windowed
	WindowDisplayMode *string `json:"WindowDisplayMode,omitempty" xml:"WindowDisplayMode,omitempty"`
	// example:
	//
	// 0aba1403b337a***
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s GetLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBody) SetAccessType(v string) *GetLoginTokenResponseBody {
	s.AccessType = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetAccountType(v string) *GetLoginTokenResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetAdDomain(v string) *GetLoginTokenResponseBody {
	s.AdDomain = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetEmail(v string) *GetLoginTokenResponseBody {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetEndUserId(v string) *GetLoginTokenResponseBody {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetIdpId(v string) *GetLoginTokenResponseBody {
	s.IdpId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetIndustry(v string) *GetLoginTokenResponseBody {
	s.Industry = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetKeepAliveToken(v string) *GetLoginTokenResponseBody {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLabel(v string) *GetLoginTokenResponseBody {
	s.Label = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLoginToken(v string) *GetLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetMfaTypeList(v []*GetLoginTokenResponseBodyMfaTypeList) *GetLoginTokenResponseBody {
	s.MfaTypeList = v
	return s
}

func (s *GetLoginTokenResponseBody) SetNextStage(v string) *GetLoginTokenResponseBody {
	s.NextStage = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetOfficeSites(v []*string) *GetLoginTokenResponseBody {
	s.OfficeSites = v
	return s
}

func (s *GetLoginTokenResponseBody) SetPasswordStrategy(v *GetLoginTokenResponseBodyPasswordStrategy) *GetLoginTokenResponseBody {
	s.PasswordStrategy = v
	return s
}

func (s *GetLoginTokenResponseBody) SetPhone(v string) *GetLoginTokenResponseBody {
	s.Phone = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetProps(v map[string]*string) *GetLoginTokenResponseBody {
	s.Props = v
	return s
}

func (s *GetLoginTokenResponseBody) SetQrCodePng(v string) *GetLoginTokenResponseBody {
	s.QrCodePng = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetReason(v string) *GetLoginTokenResponseBody {
	s.Reason = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRequestId(v string) *GetLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRiskVerifyInfo(v *GetLoginTokenResponseBodyRiskVerifyInfo) *GetLoginTokenResponseBody {
	s.RiskVerifyInfo = v
	return s
}

func (s *GetLoginTokenResponseBody) SetSecret(v string) *GetLoginTokenResponseBody {
	s.Secret = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetSessionId(v string) *GetLoginTokenResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantAlias(v string) *GetLoginTokenResponseBody {
	s.TenantAlias = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantId(v int64) *GetLoginTokenResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantInfos(v []*GetLoginTokenResponseBodyTenantInfos) *GetLoginTokenResponseBody {
	s.TenantInfos = v
	return s
}

func (s *GetLoginTokenResponseBody) SetVpcRegionId(v string) *GetLoginTokenResponseBody {
	s.VpcRegionId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetWindowDisplayMode(v string) *GetLoginTokenResponseBody {
	s.WindowDisplayMode = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetWyId(v string) *GetLoginTokenResponseBody {
	s.WyId = &v
	return s
}

type GetLoginTokenResponseBodyMfaTypeList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
}

func (s GetLoginTokenResponseBodyMfaTypeList) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyMfaTypeList) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyMfaTypeList) SetName(v string) *GetLoginTokenResponseBodyMfaTypeList {
	s.Name = &v
	return s
}

func (s *GetLoginTokenResponseBodyMfaTypeList) SetStage(v string) *GetLoginTokenResponseBodyMfaTypeList {
	s.Stage = &v
	return s
}

type GetLoginTokenResponseBodyPasswordStrategy struct {
	TenantAlternativeChars []*string `json:"TenantAlternativeChars,omitempty" xml:"TenantAlternativeChars,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TenantPasswordLength *int32 `json:"TenantPasswordLength,omitempty" xml:"TenantPasswordLength,omitempty"`
}

func (s GetLoginTokenResponseBodyPasswordStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyPasswordStrategy) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantAlternativeChars(v []*string) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantAlternativeChars = v
	return s
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantPasswordLength(v int32) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantPasswordLength = &v
	return s
}

type GetLoginTokenResponseBodyRiskVerifyInfo struct {
	// example:
	//
	// ppas***@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 12000
	LastLockDuration *int64 `json:"LastLockDuration,omitempty" xml:"LastLockDuration,omitempty"`
	// example:
	//
	// true
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// example:
	//
	// 1826717****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetEmail(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLastLockDuration(v int64) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.LastLockDuration = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLocked(v bool) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Locked = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetPhone(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Phone = &v
	return s
}

type GetLoginTokenResponseBodyTenantInfos struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// Up830***
	TenantAlias *string `json:"TenantAlias,omitempty" xml:"TenantAlias,omitempty"`
}

func (s GetLoginTokenResponseBodyTenantInfos) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyTenantInfos) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyTenantInfos) SetAccessType(v string) *GetLoginTokenResponseBodyTenantInfos {
	s.AccessType = &v
	return s
}

func (s *GetLoginTokenResponseBodyTenantInfos) SetTenantAlias(v string) *GetLoginTokenResponseBodyTenantInfos {
	s.TenantAlias = &v
	return s
}

type GetLoginTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponse) SetHeaders(v map[string]*string) *GetLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLoginTokenResponse) SetStatusCode(v int32) *GetLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginTokenResponse) SetBody(v *GetLoginTokenResponseBody) *GetLoginTokenResponse {
	s.Body = v
	return s
}

type RefreshLoginTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// Alibaba****
	LoginIdentifier *string `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1046341d8d4e2f05c4aa168196009613594aaf451499bfc75e54699efa7230bc968e1debb1fa4063b01e5d327b467****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// cn-shenzhen+dir-436909****
	OfficeSiteId  *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	ProfileRegion *string `json:"ProfileRegion,omitempty" xml:"ProfileRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6442b2fd-ed3e-423a-8e6e-352d26a4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RefreshLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenRequest) SetClientId(v string) *RefreshLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetClientType(v string) *RefreshLoginTokenRequest {
	s.ClientType = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetEndUserId(v string) *RefreshLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginIdentifier(v string) *RefreshLoginTokenRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginToken(v string) *RefreshLoginTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetOfficeSiteId(v string) *RefreshLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetProfileRegion(v string) *RefreshLoginTokenRequest {
	s.ProfileRegion = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetSessionId(v string) *RefreshLoginTokenRequest {
	s.SessionId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetUuid(v string) *RefreshLoginTokenRequest {
	s.Uuid = &v
	return s
}

type RefreshLoginTokenResponseBody struct {
	// example:
	//
	// v12369636c721ba6b3ddb1683341016775c3f63e4d0e78f120f9a0544ed826b7af7daf747c402f0d0730b52f451b70****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 419F31B9-1FDF-5644-ABA3-D00026FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponseBody) SetLoginToken(v string) *RefreshLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenResponseBody) SetRequestId(v string) *RefreshLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshLoginTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponse) SetHeaders(v map[string]*string) *RefreshLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshLoginTokenResponse) SetStatusCode(v int32) *RefreshLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLoginTokenResponse) SetBody(v *RefreshLoginTokenResponseBody) *RefreshLoginTokenResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 身份认证查询接口
//
// @param tmpReq - FindIdpListByLoginIdentifierRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindIdpListByLoginIdentifierResponse
func (client *Client) FindIdpListByLoginIdentifierWithOptions(tmpReq *FindIdpListByLoginIdentifierRequest, runtime *util.RuntimeOptions) (_result *FindIdpListByLoginIdentifierResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FindIdpListByLoginIdentifierShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AvailableFeatures)) {
		request.AvailableFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvailableFeatures, tea.String("AvailableFeatures"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvailableFeaturesShrink)) {
		query["AvailableFeatures"] = request.AvailableFeaturesShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientChannel)) {
		body["ClientChannel"] = request.ClientChannel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.LoginIdentifier)) {
		body["LoginIdentifier"] = request.LoginIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SupportTypes)) {
		body["SupportTypes"] = request.SupportTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FindIdpListByLoginIdentifier"),
		Version:     tea.String("2021-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FindIdpListByLoginIdentifierResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份认证查询接口
//
// @param request - FindIdpListByLoginIdentifierRequest
//
// @return FindIdpListByLoginIdentifierResponse
func (client *Client) FindIdpListByLoginIdentifier(request *FindIdpListByLoginIdentifierRequest) (_result *FindIdpListByLoginIdentifierResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FindIdpListByLoginIdentifierResponse{}
	_body, _err := client.FindIdpListByLoginIdentifierWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetLoginToken
//
// @param tmpReq - GetLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginTokenResponse
func (client *Client) GetLoginTokenWithOptions(tmpReq *GetLoginTokenRequest, runtime *util.RuntimeOptions) (_result *GetLoginTokenResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetLoginTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AvailableFeatures)) {
		request.AvailableFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvailableFeatures, tea.String("AvailableFeatures"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationCode)) {
		query["AuthenticationCode"] = request.AuthenticationCode
	}

	if !tea.BoolValue(util.IsUnset(request.AvailableFeaturesShrink)) {
		query["AvailableFeatures"] = request.AvailableFeaturesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentStage)) {
		query["CurrentStage"] = request.CurrentStage
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedFingerPrintData)) {
		query["EncryptedFingerPrintData"] = request.EncryptedFingerPrintData
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedKey)) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedPassword)) {
		query["EncryptedPassword"] = request.EncryptedPassword
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FingerPrintData)) {
		query["FingerPrintData"] = request.FingerPrintData
	}

	if !tea.BoolValue(util.IsUnset(request.IdpId)) {
		query["IdpId"] = request.IdpId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAlive)) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAliveToken)) {
		query["KeepAliveToken"] = request.KeepAliveToken
	}

	if !tea.BoolValue(util.IsUnset(request.LoginIdentifier)) {
		query["LoginIdentifier"] = request.LoginIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.LoginName)) {
		query["LoginName"] = request.LoginName
	}

	if !tea.BoolValue(util.IsUnset(request.MfaType)) {
		query["MfaType"] = request.MfaType
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OldPassword)) {
		query["OldPassword"] = request.OldPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneVerifyCode)) {
		query["PhoneVerifyCode"] = request.PhoneVerifyCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SsoExtendsCookies)) {
		query["SsoExtendsCookies"] = request.SsoExtendsCookies
	}

	if !tea.BoolValue(util.IsUnset(request.SsoSessionToken)) {
		query["SsoSessionToken"] = request.SsoSessionToken
	}

	if !tea.BoolValue(util.IsUnset(request.TokenCode)) {
		query["TokenCode"] = request.TokenCode
	}

	if !tea.BoolValue(util.IsUnset(request.UmidToken)) {
		query["UmidToken"] = request.UmidToken
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoginToken"),
		Version:     tea.String("2021-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetLoginToken
//
// @param request - GetLoginTokenRequest
//
// @return GetLoginTokenResponse
func (client *Client) GetLoginToken(request *GetLoginTokenRequest) (_result *GetLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.GetLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RefreshLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshLoginTokenResponse
func (client *Client) RefreshLoginTokenWithOptions(request *RefreshLoginTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshLoginTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginIdentifier)) {
		query["LoginIdentifier"] = request.LoginIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.ProfileRegion)) {
		query["ProfileRegion"] = request.ProfileRegion
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshLoginToken"),
		Version:     tea.String("2021-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshLoginTokenRequest
//
// @return RefreshLoginTokenResponse
func (client *Client) RefreshLoginToken(request *RefreshLoginTokenRequest) (_result *RefreshLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.RefreshLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
