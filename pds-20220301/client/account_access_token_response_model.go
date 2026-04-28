// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *AccountAccessTokenResponse
	GetAccessToken() *string
	SetAvatar(v string) *AccountAccessTokenResponse
	GetAvatar() *string
	SetDefaultDriveId(v string) *AccountAccessTokenResponse
	GetDefaultDriveId() *string
	SetDefaultSboxDriveId(v string) *AccountAccessTokenResponse
	GetDefaultSboxDriveId() *string
	SetDeviceId(v string) *AccountAccessTokenResponse
	GetDeviceId() *string
	SetDeviceName(v string) *AccountAccessTokenResponse
	GetDeviceName() *string
	SetDomainId(v string) *AccountAccessTokenResponse
	GetDomainId() *string
	SetExistLink(v []*LinkInfo) *AccountAccessTokenResponse
	GetExistLink() []*LinkInfo
	SetExpireTime(v string) *AccountAccessTokenResponse
	GetExpireTime() *string
	SetExpiresIn(v int64) *AccountAccessTokenResponse
	GetExpiresIn() *int64
	SetIsFirstLogin(v bool) *AccountAccessTokenResponse
	GetIsFirstLogin() *bool
	SetNeedLink(v bool) *AccountAccessTokenResponse
	GetNeedLink() *bool
	SetNeedRpVerify(v bool) *AccountAccessTokenResponse
	GetNeedRpVerify() *bool
	SetNickName(v string) *AccountAccessTokenResponse
	GetNickName() *string
	SetPathStatus(v string) *AccountAccessTokenResponse
	GetPathStatus() *string
	SetPinSetup(v bool) *AccountAccessTokenResponse
	GetPinSetup() *bool
	SetRefreshToken(v string) *AccountAccessTokenResponse
	GetRefreshToken() *string
	SetRole(v string) *AccountAccessTokenResponse
	GetRole() *string
	SetState(v string) *AccountAccessTokenResponse
	GetState() *string
	SetStatus(v string) *AccountAccessTokenResponse
	GetStatus() *string
	SetTokenType(v string) *AccountAccessTokenResponse
	GetTokenType() *string
	SetUserData(v map[string]*string) *AccountAccessTokenResponse
	GetUserData() map[string]*string
	SetUserId(v string) *AccountAccessTokenResponse
	GetUserId() *string
	SetUserName(v string) *AccountAccessTokenResponse
	GetUserName() *string
}

type AccountAccessTokenResponse struct {
	AccessToken        *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Avatar             *string            `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DefaultDriveId     *string            `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DefaultSboxDriveId *string            `json:"default_sbox_drive_id,omitempty" xml:"default_sbox_drive_id,omitempty"`
	DeviceId           *string            `json:"device_id,omitempty" xml:"device_id,omitempty"`
	DeviceName         *string            `json:"device_name,omitempty" xml:"device_name,omitempty"`
	DomainId           *string            `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	ExistLink          []*LinkInfo        `json:"exist_link,omitempty" xml:"exist_link,omitempty" type:"Repeated"`
	ExpireTime         *string            `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	ExpiresIn          *int64             `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	IsFirstLogin       *bool              `json:"is_first_login,omitempty" xml:"is_first_login,omitempty"`
	NeedLink           *bool              `json:"need_link,omitempty" xml:"need_link,omitempty"`
	NeedRpVerify       *bool              `json:"need_rp_verify,omitempty" xml:"need_rp_verify,omitempty"`
	NickName           *string            `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	PathStatus         *string            `json:"path_status,omitempty" xml:"path_status,omitempty"`
	PinSetup           *bool              `json:"pin_setup,omitempty" xml:"pin_setup,omitempty"`
	RefreshToken       *string            `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Role               *string            `json:"role,omitempty" xml:"role,omitempty"`
	State              *string            `json:"state,omitempty" xml:"state,omitempty"`
	Status             *string            `json:"status,omitempty" xml:"status,omitempty"`
	TokenType          *string            `json:"token_type,omitempty" xml:"token_type,omitempty"`
	UserData           map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	UserId             *string            `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string            `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s AccountAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *AccountAccessTokenResponse) GetAccessToken() *string {
	return s.AccessToken
}

func (s *AccountAccessTokenResponse) GetAvatar() *string {
	return s.Avatar
}

func (s *AccountAccessTokenResponse) GetDefaultDriveId() *string {
	return s.DefaultDriveId
}

func (s *AccountAccessTokenResponse) GetDefaultSboxDriveId() *string {
	return s.DefaultSboxDriveId
}

func (s *AccountAccessTokenResponse) GetDeviceId() *string {
	return s.DeviceId
}

func (s *AccountAccessTokenResponse) GetDeviceName() *string {
	return s.DeviceName
}

func (s *AccountAccessTokenResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *AccountAccessTokenResponse) GetExistLink() []*LinkInfo {
	return s.ExistLink
}

func (s *AccountAccessTokenResponse) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *AccountAccessTokenResponse) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *AccountAccessTokenResponse) GetIsFirstLogin() *bool {
	return s.IsFirstLogin
}

func (s *AccountAccessTokenResponse) GetNeedLink() *bool {
	return s.NeedLink
}

func (s *AccountAccessTokenResponse) GetNeedRpVerify() *bool {
	return s.NeedRpVerify
}

func (s *AccountAccessTokenResponse) GetNickName() *string {
	return s.NickName
}

func (s *AccountAccessTokenResponse) GetPathStatus() *string {
	return s.PathStatus
}

func (s *AccountAccessTokenResponse) GetPinSetup() *bool {
	return s.PinSetup
}

func (s *AccountAccessTokenResponse) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *AccountAccessTokenResponse) GetRole() *string {
	return s.Role
}

func (s *AccountAccessTokenResponse) GetState() *string {
	return s.State
}

func (s *AccountAccessTokenResponse) GetStatus() *string {
	return s.Status
}

func (s *AccountAccessTokenResponse) GetTokenType() *string {
	return s.TokenType
}

func (s *AccountAccessTokenResponse) GetUserData() map[string]*string {
	return s.UserData
}

func (s *AccountAccessTokenResponse) GetUserId() *string {
	return s.UserId
}

func (s *AccountAccessTokenResponse) GetUserName() *string {
	return s.UserName
}

func (s *AccountAccessTokenResponse) SetAccessToken(v string) *AccountAccessTokenResponse {
	s.AccessToken = &v
	return s
}

func (s *AccountAccessTokenResponse) SetAvatar(v string) *AccountAccessTokenResponse {
	s.Avatar = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDefaultDriveId(v string) *AccountAccessTokenResponse {
	s.DefaultDriveId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDefaultSboxDriveId(v string) *AccountAccessTokenResponse {
	s.DefaultSboxDriveId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDeviceId(v string) *AccountAccessTokenResponse {
	s.DeviceId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDeviceName(v string) *AccountAccessTokenResponse {
	s.DeviceName = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDomainId(v string) *AccountAccessTokenResponse {
	s.DomainId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetExistLink(v []*LinkInfo) *AccountAccessTokenResponse {
	s.ExistLink = v
	return s
}

func (s *AccountAccessTokenResponse) SetExpireTime(v string) *AccountAccessTokenResponse {
	s.ExpireTime = &v
	return s
}

func (s *AccountAccessTokenResponse) SetExpiresIn(v int64) *AccountAccessTokenResponse {
	s.ExpiresIn = &v
	return s
}

func (s *AccountAccessTokenResponse) SetIsFirstLogin(v bool) *AccountAccessTokenResponse {
	s.IsFirstLogin = &v
	return s
}

func (s *AccountAccessTokenResponse) SetNeedLink(v bool) *AccountAccessTokenResponse {
	s.NeedLink = &v
	return s
}

func (s *AccountAccessTokenResponse) SetNeedRpVerify(v bool) *AccountAccessTokenResponse {
	s.NeedRpVerify = &v
	return s
}

func (s *AccountAccessTokenResponse) SetNickName(v string) *AccountAccessTokenResponse {
	s.NickName = &v
	return s
}

func (s *AccountAccessTokenResponse) SetPathStatus(v string) *AccountAccessTokenResponse {
	s.PathStatus = &v
	return s
}

func (s *AccountAccessTokenResponse) SetPinSetup(v bool) *AccountAccessTokenResponse {
	s.PinSetup = &v
	return s
}

func (s *AccountAccessTokenResponse) SetRefreshToken(v string) *AccountAccessTokenResponse {
	s.RefreshToken = &v
	return s
}

func (s *AccountAccessTokenResponse) SetRole(v string) *AccountAccessTokenResponse {
	s.Role = &v
	return s
}

func (s *AccountAccessTokenResponse) SetState(v string) *AccountAccessTokenResponse {
	s.State = &v
	return s
}

func (s *AccountAccessTokenResponse) SetStatus(v string) *AccountAccessTokenResponse {
	s.Status = &v
	return s
}

func (s *AccountAccessTokenResponse) SetTokenType(v string) *AccountAccessTokenResponse {
	s.TokenType = &v
	return s
}

func (s *AccountAccessTokenResponse) SetUserData(v map[string]*string) *AccountAccessTokenResponse {
	s.UserData = v
	return s
}

func (s *AccountAccessTokenResponse) SetUserId(v string) *AccountAccessTokenResponse {
	s.UserId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetUserName(v string) *AccountAccessTokenResponse {
	s.UserName = &v
	return s
}

func (s *AccountAccessTokenResponse) Validate() error {
	if s.ExistLink != nil {
		for _, item := range s.ExistLink {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
