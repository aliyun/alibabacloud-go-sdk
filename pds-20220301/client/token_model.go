// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToken interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *Token
	GetAccessToken() *string
	SetAvatar(v string) *Token
	GetAvatar() *string
	SetDefaultDriveId(v string) *Token
	GetDefaultDriveId() *string
	SetDefaultSboxDriveId(v string) *Token
	GetDefaultSboxDriveId() *string
	SetDeviceName(v string) *Token
	GetDeviceName() *string
	SetDomainId(v string) *Token
	GetDomainId() *string
	SetExistLink(v []*LinkInfo) *Token
	GetExistLink() []*LinkInfo
	SetExpireTime(v string) *Token
	GetExpireTime() *string
	SetExpiresIn(v int64) *Token
	GetExpiresIn() *int64
	SetIsFirstLogin(v bool) *Token
	GetIsFirstLogin() *bool
	SetNeedLink(v bool) *Token
	GetNeedLink() *bool
	SetNeedRpVerify(v bool) *Token
	GetNeedRpVerify() *bool
	SetNickName(v string) *Token
	GetNickName() *string
	SetPinSetup(v bool) *Token
	GetPinSetup() *bool
	SetRefreshToken(v string) *Token
	GetRefreshToken() *string
	SetRole(v string) *Token
	GetRole() *string
	SetState(v string) *Token
	GetState() *string
	SetStatus(v string) *Token
	GetStatus() *string
	SetTokenType(v string) *Token
	GetTokenType() *string
	SetUserData(v map[string]*string) *Token
	GetUserData() map[string]*string
	SetUserId(v string) *Token
	GetUserId() *string
	SetUserName(v string) *Token
	GetUserName() *string
}

type Token struct {
	// The access token.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjOWI3YTVhYTA0ZDE0YWUzODY3ZmRjODg2ZmEwMWRhNCIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwiMjVkelgzdmJZcWt0Vnh5WFwiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJTSEFSRS5BTExcIixcIkZJTEUuQUxMXCIsXCJVU0VSLkFMTFwiLFwiVklFVy5BTExcIixcIlNUT1JBR0UuQUxMXCIsXCJTVE9SQUdFRklMRS5MSVNUXCIsXCJCQVRDSFwiLFwiT0FVVEguQUxMXCIsXCJJTUFHRS5BTExcIixcIklOVklURS5BTExcIixcIkFDQ09VTlQuQUxMXCJdLFwicm9sZVwiOlwidXNlclwiLFwicmVmXCI6XCJodHRwczovL3d3dy5hbGl5dW5kcml2ZS5jb20vXCIsXCJkZXZpY2VfaWRcIjpcImIyODIwNWU1YzU5NzRjY2JiODI3MDNiNjhkYjhjNDUxXCJ9IiwiZXhwIjoxNjQ4NjE0NDkzLCJpYXQiOjE2NDg2MDcyMzN9.d3HVLvv_LFw2QhPrhvjH_kICWQJX9sKKt7NjQEqI_xE2JO_b7D8rPsFTZz93PLvZ7MhCmudTjGImUpd-ehFnI4Go-1S7BGaKaHFILvP-sWy18Wpikowjxx9mSbzBM_cO6D1LI-kyYhXKWHgVdADfVIPniTDA7-ffhUpi7cAebEs
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// The profile picture of the user.
	//
	// example:
	//
	// aliyunpds.com/a.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The ID of the default space of the user.
	//
	// example:
	//
	// 1
	DefaultDriveId     *string `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DefaultSboxDriveId *string `json:"default_sbox_drive_id,omitempty" xml:"default_sbox_drive_id,omitempty"`
	// The name of the device that is bound to OAuth 2.0 Device Authorization Grant.
	//
	// example:
	//
	// 4683C25F
	DeviceName *string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId  *string     `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	ExistLink []*LinkInfo `json:"exist_link,omitempty" xml:"exist_link,omitempty" type:"Repeated"`
	// The time when the credential expires.
	//
	// example:
	//
	// 2019-09-01T06:57:48.813Z
	ExpireTime *string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// The validity period of the token.
	//
	// example:
	//
	// 3600
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// Indicates whether this is the first logon of the user.
	//
	// example:
	//
	// false
	IsFirstLogin *bool `json:"is_first_login,omitempty" xml:"is_first_login,omitempty"`
	NeedLink     *bool `json:"need_link,omitempty" xml:"need_link,omitempty"`
	NeedRpVerify *bool `json:"need_rp_verify,omitempty" xml:"need_rp_verify,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// pdsuser
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	PinSetup *bool   `json:"pin_setup,omitempty" xml:"pin_setup,omitempty"`
	// The refresh token.
	//
	// example:
	//
	// 060e78d36afb4879b51e4264e9541c16
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// The role of the user.
	//
	// example:
	//
	// admin
	Role  *string `json:"role,omitempty" xml:"role,omitempty"`
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The status of the user.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the token.
	//
	// 	- Only Bearer is supported.
	//
	// example:
	//
	// Bearer
	TokenType *string            `json:"token_type,omitempty" xml:"token_type,omitempty"`
	UserData  map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID.
	//
	// example:
	//
	// DING-xxxxx
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The name of the user.
	//
	// example:
	//
	// pdsuser
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s Token) String() string {
	return dara.Prettify(s)
}

func (s Token) GoString() string {
	return s.String()
}

func (s *Token) GetAccessToken() *string {
	return s.AccessToken
}

func (s *Token) GetAvatar() *string {
	return s.Avatar
}

func (s *Token) GetDefaultDriveId() *string {
	return s.DefaultDriveId
}

func (s *Token) GetDefaultSboxDriveId() *string {
	return s.DefaultSboxDriveId
}

func (s *Token) GetDeviceName() *string {
	return s.DeviceName
}

func (s *Token) GetDomainId() *string {
	return s.DomainId
}

func (s *Token) GetExistLink() []*LinkInfo {
	return s.ExistLink
}

func (s *Token) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *Token) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *Token) GetIsFirstLogin() *bool {
	return s.IsFirstLogin
}

func (s *Token) GetNeedLink() *bool {
	return s.NeedLink
}

func (s *Token) GetNeedRpVerify() *bool {
	return s.NeedRpVerify
}

func (s *Token) GetNickName() *string {
	return s.NickName
}

func (s *Token) GetPinSetup() *bool {
	return s.PinSetup
}

func (s *Token) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *Token) GetRole() *string {
	return s.Role
}

func (s *Token) GetState() *string {
	return s.State
}

func (s *Token) GetStatus() *string {
	return s.Status
}

func (s *Token) GetTokenType() *string {
	return s.TokenType
}

func (s *Token) GetUserData() map[string]*string {
	return s.UserData
}

func (s *Token) GetUserId() *string {
	return s.UserId
}

func (s *Token) GetUserName() *string {
	return s.UserName
}

func (s *Token) SetAccessToken(v string) *Token {
	s.AccessToken = &v
	return s
}

func (s *Token) SetAvatar(v string) *Token {
	s.Avatar = &v
	return s
}

func (s *Token) SetDefaultDriveId(v string) *Token {
	s.DefaultDriveId = &v
	return s
}

func (s *Token) SetDefaultSboxDriveId(v string) *Token {
	s.DefaultSboxDriveId = &v
	return s
}

func (s *Token) SetDeviceName(v string) *Token {
	s.DeviceName = &v
	return s
}

func (s *Token) SetDomainId(v string) *Token {
	s.DomainId = &v
	return s
}

func (s *Token) SetExistLink(v []*LinkInfo) *Token {
	s.ExistLink = v
	return s
}

func (s *Token) SetExpireTime(v string) *Token {
	s.ExpireTime = &v
	return s
}

func (s *Token) SetExpiresIn(v int64) *Token {
	s.ExpiresIn = &v
	return s
}

func (s *Token) SetIsFirstLogin(v bool) *Token {
	s.IsFirstLogin = &v
	return s
}

func (s *Token) SetNeedLink(v bool) *Token {
	s.NeedLink = &v
	return s
}

func (s *Token) SetNeedRpVerify(v bool) *Token {
	s.NeedRpVerify = &v
	return s
}

func (s *Token) SetNickName(v string) *Token {
	s.NickName = &v
	return s
}

func (s *Token) SetPinSetup(v bool) *Token {
	s.PinSetup = &v
	return s
}

func (s *Token) SetRefreshToken(v string) *Token {
	s.RefreshToken = &v
	return s
}

func (s *Token) SetRole(v string) *Token {
	s.Role = &v
	return s
}

func (s *Token) SetState(v string) *Token {
	s.State = &v
	return s
}

func (s *Token) SetStatus(v string) *Token {
	s.Status = &v
	return s
}

func (s *Token) SetTokenType(v string) *Token {
	s.TokenType = &v
	return s
}

func (s *Token) SetUserData(v map[string]*string) *Token {
	s.UserData = v
	return s
}

func (s *Token) SetUserId(v string) *Token {
	s.UserId = &v
	return s
}

func (s *Token) SetUserName(v string) *Token {
	s.UserName = &v
	return s
}

func (s *Token) Validate() error {
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
