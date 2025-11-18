// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaSite(v string) *GetLoginTokenRequest
	GetAreaSite() *string
	SetAuthenticationCode(v string) *GetLoginTokenRequest
	GetAuthenticationCode() *string
	SetAvailableFeatures(v map[string]*string) *GetLoginTokenRequest
	GetAvailableFeatures() map[string]*string
	SetChannel(v string) *GetLoginTokenRequest
	GetChannel() *string
	SetClientId(v string) *GetLoginTokenRequest
	GetClientId() *string
	SetClientName(v string) *GetLoginTokenRequest
	GetClientName() *string
	SetClientOS(v string) *GetLoginTokenRequest
	GetClientOS() *string
	SetClientType(v string) *GetLoginTokenRequest
	GetClientType() *string
	SetClientVersion(v string) *GetLoginTokenRequest
	GetClientVersion() *string
	SetCurrentStage(v string) *GetLoginTokenRequest
	GetCurrentStage() *string
	SetDirectoryId(v string) *GetLoginTokenRequest
	GetDirectoryId() *string
	SetEncryptedFingerPrintData(v string) *GetLoginTokenRequest
	GetEncryptedFingerPrintData() *string
	SetEncryptedKey(v string) *GetLoginTokenRequest
	GetEncryptedKey() *string
	SetEncryptedPassword(v string) *GetLoginTokenRequest
	GetEncryptedPassword() *string
	SetEndUserId(v string) *GetLoginTokenRequest
	GetEndUserId() *string
	SetFingerPrintData(v string) *GetLoginTokenRequest
	GetFingerPrintData() *string
	SetIdpId(v string) *GetLoginTokenRequest
	GetIdpId() *string
	SetImageUrl(v string) *GetLoginTokenRequest
	GetImageUrl() *string
	SetKeepAlive(v bool) *GetLoginTokenRequest
	GetKeepAlive() *bool
	SetKeepAliveToken(v string) *GetLoginTokenRequest
	GetKeepAliveToken() *string
	SetLoginIdentifier(v string) *GetLoginTokenRequest
	GetLoginIdentifier() *string
	SetLoginName(v string) *GetLoginTokenRequest
	GetLoginName() *string
	SetMfaType(v string) *GetLoginTokenRequest
	GetMfaType() *string
	SetNetworkType(v string) *GetLoginTokenRequest
	GetNetworkType() *string
	SetNewPassword(v string) *GetLoginTokenRequest
	GetNewPassword() *string
	SetOfficeSiteId(v string) *GetLoginTokenRequest
	GetOfficeSiteId() *string
	SetOldPassword(v string) *GetLoginTokenRequest
	GetOldPassword() *string
	SetPassword(v string) *GetLoginTokenRequest
	GetPassword() *string
	SetPhone(v string) *GetLoginTokenRequest
	GetPhone() *string
	SetPhoneVerifyCode(v string) *GetLoginTokenRequest
	GetPhoneVerifyCode() *string
	SetProfileRegion(v string) *GetLoginTokenRequest
	GetProfileRegion() *string
	SetRegionId(v string) *GetLoginTokenRequest
	GetRegionId() *string
	SetSessionId(v string) *GetLoginTokenRequest
	GetSessionId() *string
	SetSsoExtendsCookies(v string) *GetLoginTokenRequest
	GetSsoExtendsCookies() *string
	SetSsoSessionToken(v string) *GetLoginTokenRequest
	GetSsoSessionToken() *string
	SetTokenCode(v string) *GetLoginTokenRequest
	GetTokenCode() *string
	SetUmidToken(v string) *GetLoginTokenRequest
	GetUmidToken() *string
	SetUuid(v string) *GetLoginTokenRequest
	GetUuid() *string
}

type GetLoginTokenRequest struct {
	AreaSite *string `json:"AreaSite,omitempty" xml:"AreaSite,omitempty"`
	// example:
	//
	// 182901
	AuthenticationCode *string            `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	AvailableFeatures  map[string]*string `json:"AvailableFeatures,omitempty" xml:"AvailableFeatures,omitempty"`
	Channel            *string            `json:"Channel,omitempty" xml:"Channel,omitempty"`
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

func (s GetLoginTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequest) GetAreaSite() *string {
	return s.AreaSite
}

func (s *GetLoginTokenRequest) GetAuthenticationCode() *string {
	return s.AuthenticationCode
}

func (s *GetLoginTokenRequest) GetAvailableFeatures() map[string]*string {
	return s.AvailableFeatures
}

func (s *GetLoginTokenRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetLoginTokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetLoginTokenRequest) GetClientName() *string {
	return s.ClientName
}

func (s *GetLoginTokenRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GetLoginTokenRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GetLoginTokenRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetLoginTokenRequest) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *GetLoginTokenRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetLoginTokenRequest) GetEncryptedFingerPrintData() *string {
	return s.EncryptedFingerPrintData
}

func (s *GetLoginTokenRequest) GetEncryptedKey() *string {
	return s.EncryptedKey
}

func (s *GetLoginTokenRequest) GetEncryptedPassword() *string {
	return s.EncryptedPassword
}

func (s *GetLoginTokenRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetLoginTokenRequest) GetFingerPrintData() *string {
	return s.FingerPrintData
}

func (s *GetLoginTokenRequest) GetIdpId() *string {
	return s.IdpId
}

func (s *GetLoginTokenRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetLoginTokenRequest) GetKeepAlive() *bool {
	return s.KeepAlive
}

func (s *GetLoginTokenRequest) GetKeepAliveToken() *string {
	return s.KeepAliveToken
}

func (s *GetLoginTokenRequest) GetLoginIdentifier() *string {
	return s.LoginIdentifier
}

func (s *GetLoginTokenRequest) GetLoginName() *string {
	return s.LoginName
}

func (s *GetLoginTokenRequest) GetMfaType() *string {
	return s.MfaType
}

func (s *GetLoginTokenRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLoginTokenRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *GetLoginTokenRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetLoginTokenRequest) GetOldPassword() *string {
	return s.OldPassword
}

func (s *GetLoginTokenRequest) GetPassword() *string {
	return s.Password
}

func (s *GetLoginTokenRequest) GetPhone() *string {
	return s.Phone
}

func (s *GetLoginTokenRequest) GetPhoneVerifyCode() *string {
	return s.PhoneVerifyCode
}

func (s *GetLoginTokenRequest) GetProfileRegion() *string {
	return s.ProfileRegion
}

func (s *GetLoginTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLoginTokenRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetLoginTokenRequest) GetSsoExtendsCookies() *string {
	return s.SsoExtendsCookies
}

func (s *GetLoginTokenRequest) GetSsoSessionToken() *string {
	return s.SsoSessionToken
}

func (s *GetLoginTokenRequest) GetTokenCode() *string {
	return s.TokenCode
}

func (s *GetLoginTokenRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *GetLoginTokenRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetLoginTokenRequest) SetAreaSite(v string) *GetLoginTokenRequest {
	s.AreaSite = &v
	return s
}

func (s *GetLoginTokenRequest) SetAuthenticationCode(v string) *GetLoginTokenRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetAvailableFeatures(v map[string]*string) *GetLoginTokenRequest {
	s.AvailableFeatures = v
	return s
}

func (s *GetLoginTokenRequest) SetChannel(v string) *GetLoginTokenRequest {
	s.Channel = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientId(v string) *GetLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientName(v string) *GetLoginTokenRequest {
	s.ClientName = &v
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

func (s *GetLoginTokenRequest) SetProfileRegion(v string) *GetLoginTokenRequest {
	s.ProfileRegion = &v
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

func (s *GetLoginTokenRequest) Validate() error {
	return dara.Validate(s)
}
