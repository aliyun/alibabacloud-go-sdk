// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginTokenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaSite(v string) *GetLoginTokenShrinkRequest
	GetAreaSite() *string
	SetAuthenticationCode(v string) *GetLoginTokenShrinkRequest
	GetAuthenticationCode() *string
	SetAvailableFeaturesShrink(v string) *GetLoginTokenShrinkRequest
	GetAvailableFeaturesShrink() *string
	SetChannel(v string) *GetLoginTokenShrinkRequest
	GetChannel() *string
	SetClientId(v string) *GetLoginTokenShrinkRequest
	GetClientId() *string
	SetClientName(v string) *GetLoginTokenShrinkRequest
	GetClientName() *string
	SetClientOS(v string) *GetLoginTokenShrinkRequest
	GetClientOS() *string
	SetClientType(v string) *GetLoginTokenShrinkRequest
	GetClientType() *string
	SetClientVersion(v string) *GetLoginTokenShrinkRequest
	GetClientVersion() *string
	SetCurrentStage(v string) *GetLoginTokenShrinkRequest
	GetCurrentStage() *string
	SetDirectoryId(v string) *GetLoginTokenShrinkRequest
	GetDirectoryId() *string
	SetEncryptedFingerPrintData(v string) *GetLoginTokenShrinkRequest
	GetEncryptedFingerPrintData() *string
	SetEncryptedKey(v string) *GetLoginTokenShrinkRequest
	GetEncryptedKey() *string
	SetEncryptedPassword(v string) *GetLoginTokenShrinkRequest
	GetEncryptedPassword() *string
	SetEndUserId(v string) *GetLoginTokenShrinkRequest
	GetEndUserId() *string
	SetFingerPrintData(v string) *GetLoginTokenShrinkRequest
	GetFingerPrintData() *string
	SetIdpId(v string) *GetLoginTokenShrinkRequest
	GetIdpId() *string
	SetImageUrl(v string) *GetLoginTokenShrinkRequest
	GetImageUrl() *string
	SetKeepAlive(v bool) *GetLoginTokenShrinkRequest
	GetKeepAlive() *bool
	SetKeepAliveToken(v string) *GetLoginTokenShrinkRequest
	GetKeepAliveToken() *string
	SetLoginIdentifier(v string) *GetLoginTokenShrinkRequest
	GetLoginIdentifier() *string
	SetLoginName(v string) *GetLoginTokenShrinkRequest
	GetLoginName() *string
	SetMfaType(v string) *GetLoginTokenShrinkRequest
	GetMfaType() *string
	SetNetworkType(v string) *GetLoginTokenShrinkRequest
	GetNetworkType() *string
	SetNewPassword(v string) *GetLoginTokenShrinkRequest
	GetNewPassword() *string
	SetOfficeSiteId(v string) *GetLoginTokenShrinkRequest
	GetOfficeSiteId() *string
	SetOldPassword(v string) *GetLoginTokenShrinkRequest
	GetOldPassword() *string
	SetPassword(v string) *GetLoginTokenShrinkRequest
	GetPassword() *string
	SetPhone(v string) *GetLoginTokenShrinkRequest
	GetPhone() *string
	SetPhoneVerifyCode(v string) *GetLoginTokenShrinkRequest
	GetPhoneVerifyCode() *string
	SetProfileRegion(v string) *GetLoginTokenShrinkRequest
	GetProfileRegion() *string
	SetRegionId(v string) *GetLoginTokenShrinkRequest
	GetRegionId() *string
	SetSessionId(v string) *GetLoginTokenShrinkRequest
	GetSessionId() *string
	SetSsoExtendsCookies(v string) *GetLoginTokenShrinkRequest
	GetSsoExtendsCookies() *string
	SetSsoSessionToken(v string) *GetLoginTokenShrinkRequest
	GetSsoSessionToken() *string
	SetTokenCode(v string) *GetLoginTokenShrinkRequest
	GetTokenCode() *string
	SetUmidToken(v string) *GetLoginTokenShrinkRequest
	GetUmidToken() *string
	SetUuid(v string) *GetLoginTokenShrinkRequest
	GetUuid() *string
}

type GetLoginTokenShrinkRequest struct {
	AreaSite *string `json:"AreaSite,omitempty" xml:"AreaSite,omitempty"`
	// example:
	//
	// 182901
	AuthenticationCode      *string `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	AvailableFeaturesShrink *string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
	Channel                 *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 05967f80-6f51-46cb-a27c-****
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
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
	ProfileRegion   *string `json:"ProfileRegion,omitempty" xml:"ProfileRegion,omitempty"`
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
	return dara.Prettify(s)
}

func (s GetLoginTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenShrinkRequest) GetAreaSite() *string {
	return s.AreaSite
}

func (s *GetLoginTokenShrinkRequest) GetAuthenticationCode() *string {
	return s.AuthenticationCode
}

func (s *GetLoginTokenShrinkRequest) GetAvailableFeaturesShrink() *string {
	return s.AvailableFeaturesShrink
}

func (s *GetLoginTokenShrinkRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetLoginTokenShrinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetLoginTokenShrinkRequest) GetClientName() *string {
	return s.ClientName
}

func (s *GetLoginTokenShrinkRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GetLoginTokenShrinkRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GetLoginTokenShrinkRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetLoginTokenShrinkRequest) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *GetLoginTokenShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetLoginTokenShrinkRequest) GetEncryptedFingerPrintData() *string {
	return s.EncryptedFingerPrintData
}

func (s *GetLoginTokenShrinkRequest) GetEncryptedKey() *string {
	return s.EncryptedKey
}

func (s *GetLoginTokenShrinkRequest) GetEncryptedPassword() *string {
	return s.EncryptedPassword
}

func (s *GetLoginTokenShrinkRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetLoginTokenShrinkRequest) GetFingerPrintData() *string {
	return s.FingerPrintData
}

func (s *GetLoginTokenShrinkRequest) GetIdpId() *string {
	return s.IdpId
}

func (s *GetLoginTokenShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetLoginTokenShrinkRequest) GetKeepAlive() *bool {
	return s.KeepAlive
}

func (s *GetLoginTokenShrinkRequest) GetKeepAliveToken() *string {
	return s.KeepAliveToken
}

func (s *GetLoginTokenShrinkRequest) GetLoginIdentifier() *string {
	return s.LoginIdentifier
}

func (s *GetLoginTokenShrinkRequest) GetLoginName() *string {
	return s.LoginName
}

func (s *GetLoginTokenShrinkRequest) GetMfaType() *string {
	return s.MfaType
}

func (s *GetLoginTokenShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLoginTokenShrinkRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *GetLoginTokenShrinkRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetLoginTokenShrinkRequest) GetOldPassword() *string {
	return s.OldPassword
}

func (s *GetLoginTokenShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *GetLoginTokenShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *GetLoginTokenShrinkRequest) GetPhoneVerifyCode() *string {
	return s.PhoneVerifyCode
}

func (s *GetLoginTokenShrinkRequest) GetProfileRegion() *string {
	return s.ProfileRegion
}

func (s *GetLoginTokenShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLoginTokenShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetLoginTokenShrinkRequest) GetSsoExtendsCookies() *string {
	return s.SsoExtendsCookies
}

func (s *GetLoginTokenShrinkRequest) GetSsoSessionToken() *string {
	return s.SsoSessionToken
}

func (s *GetLoginTokenShrinkRequest) GetTokenCode() *string {
	return s.TokenCode
}

func (s *GetLoginTokenShrinkRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *GetLoginTokenShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetLoginTokenShrinkRequest) SetAreaSite(v string) *GetLoginTokenShrinkRequest {
	s.AreaSite = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetAuthenticationCode(v string) *GetLoginTokenShrinkRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetAvailableFeaturesShrink(v string) *GetLoginTokenShrinkRequest {
	s.AvailableFeaturesShrink = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetChannel(v string) *GetLoginTokenShrinkRequest {
	s.Channel = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetClientId(v string) *GetLoginTokenShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetClientName(v string) *GetLoginTokenShrinkRequest {
	s.ClientName = &v
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

func (s *GetLoginTokenShrinkRequest) SetProfileRegion(v string) *GetLoginTokenShrinkRequest {
	s.ProfileRegion = &v
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

func (s *GetLoginTokenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
