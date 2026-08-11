// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlashSmsAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfile(v *UpdateFlashSmsAccessProfileRequestAccessProfile) *UpdateFlashSmsAccessProfileRequest
	GetAccessProfile() *UpdateFlashSmsAccessProfileRequestAccessProfile
	SetAccessProfileId(v string) *UpdateFlashSmsAccessProfileRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *UpdateFlashSmsAccessProfileRequest
	GetInstanceId() *string
	SetProviderId(v string) *UpdateFlashSmsAccessProfileRequest
	GetProviderId() *string
}

type UpdateFlashSmsAccessProfileRequest struct {
	// The access configuration.
	AccessProfile *UpdateFlashSmsAccessProfileRequestAccessProfile `json:"AccessProfile,omitempty" xml:"AccessProfile,omitempty" type:"Struct"`
	// The access configuration ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The provider ID. Valid values:
	//
	// - Uincall: Beijing Youyin Communication Co., Ltd.
	//
	// - ChuangLan: Beijing Chuanglan Cloud Intelligence Information Co., Ltd.
	//
	// - ChinaMobile: China Mobile.
	//
	// - ShangHaiTianNan: Shanghai Tiannan.
	//
	// - HeDao: Galaxis.
	//
	// - DySms: Alibaba Communication.
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s UpdateFlashSmsAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlashSmsAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlashSmsAccessProfileRequest) GetAccessProfile() *UpdateFlashSmsAccessProfileRequestAccessProfile {
	return s.AccessProfile
}

func (s *UpdateFlashSmsAccessProfileRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateFlashSmsAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFlashSmsAccessProfileRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *UpdateFlashSmsAccessProfileRequest) SetAccessProfile(v *UpdateFlashSmsAccessProfileRequestAccessProfile) *UpdateFlashSmsAccessProfileRequest {
	s.AccessProfile = v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequest) SetAccessProfileId(v string) *UpdateFlashSmsAccessProfileRequest {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequest) SetInstanceId(v string) *UpdateFlashSmsAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequest) SetProviderId(v string) *UpdateFlashSmsAccessProfileRequest {
	s.ProviderId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequest) Validate() error {
	if s.AccessProfile != nil {
		if err := s.AccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFlashSmsAccessProfileRequestAccessProfile struct {
	// Required when ProviderId is set to ShangHaiTianNan or Uincall.
	//
	// example:
	//
	// 6004200267
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// Required when ProviderId is set to ChinaMobile.
	//
	// example:
	//
	// TQChVEAabhaNp2AB
	AesKey *string `json:"AesKey,omitempty" xml:"AesKey,omitempty"`
	// Required when ProviderId is set to ChuangLan.
	//
	// example:
	//
	// N92685567
	ApiAccount *string `json:"ApiAccount,omitempty" xml:"ApiAccount,omitempty"`
	// Required when ProviderId is set to ChinaMobile.
	//
	// example:
	//
	// 100235
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// Required when ProviderId is set to ChinaMobile.
	//
	// example:
	//
	// 3aRsPrTsDG3OPNq5
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// Required when ProviderId is set to ChuangLan.
	//
	// example:
	//
	// Rp7hyUbtXMef23
	ApiPassword *string `json:"ApiPassword,omitempty" xml:"ApiPassword,omitempty"`
	// Required when ProviderId is set to ChinaMobile.
	//
	// example:
	//
	// 300012117547
	CapAppId *string `json:"CapAppId,omitempty" xml:"CapAppId,omitempty"`
	// The list of Alibaba Communication configurations. Required when ProviderId is set to DySms.
	DySmsAccessProfiles []*UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles `json:"DySmsAccessProfiles,omitempty" xml:"DySmsAccessProfiles,omitempty" type:"Repeated"`
	// Required when ProviderId is set to ShangHaiTianNan.
	//
	// example:
	//
	// 10690101220
	Extno *string `json:"Extno,omitempty" xml:"Extno,omitempty"`
	// Required when ProviderId is set to ChuangLan.
	//
	// example:
	//
	// Rp7hyUbtXMef23
	ManagementPassword *string `json:"ManagementPassword,omitempty" xml:"ManagementPassword,omitempty"`
	// Required when ProviderId is set to ChuangLan.
	//
	// example:
	//
	// chuanglanrobot2
	ManagementSubUserId *string `json:"ManagementSubUserId,omitempty" xml:"ManagementSubUserId,omitempty"`
	// Required when ProviderId is set to ChuangLan.
	//
	// example:
	//
	// chuanglanrobot
	ManagementUsername *string `json:"ManagementUsername,omitempty" xml:"ManagementUsername,omitempty"`
	// Required when ProviderId is set to ShangHaiTianNan or HeDao.
	//
	// example:
	//
	// nu2DxxfZtY46
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Required when ProviderId is set to Uincall.
	//
	// example:
	//
	// 828ee92ebc8241d3b37d0238dde6345e
	Pwd *string `json:"Pwd,omitempty" xml:"Pwd,omitempty"`
	// Required when ProviderId is set to Uincall.
	//
	// example:
	//
	// 6004200267_dev
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// Required when ProviderId is set to HeDao.
	//
	// example:
	//
	// TEST10
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateFlashSmsAccessProfileRequestAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlashSmsAccessProfileRequestAccessProfile) GoString() string {
	return s.String()
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetAccount() *string {
	return s.Account
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetAesKey() *string {
	return s.AesKey
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetApiAccount() *string {
	return s.ApiAccount
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetApiId() *string {
	return s.ApiId
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetApiPassword() *string {
	return s.ApiPassword
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetCapAppId() *string {
	return s.CapAppId
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetDySmsAccessProfiles() []*UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	return s.DySmsAccessProfiles
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetExtno() *string {
	return s.Extno
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetManagementPassword() *string {
	return s.ManagementPassword
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetManagementSubUserId() *string {
	return s.ManagementSubUserId
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetManagementUsername() *string {
	return s.ManagementUsername
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetPassword() *string {
	return s.Password
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetPwd() *string {
	return s.Pwd
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetUser() *string {
	return s.User
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) GetUserName() *string {
	return s.UserName
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetAccount(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.Account = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetAesKey(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.AesKey = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetApiAccount(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiAccount = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetApiId(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetApiKey(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiKey = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetApiPassword(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiPassword = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetCapAppId(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.CapAppId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetDySmsAccessProfiles(v []*UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.DySmsAccessProfiles = v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetExtno(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.Extno = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetManagementPassword(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.ManagementPassword = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetManagementSubUserId(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.ManagementSubUserId = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetManagementUsername(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.ManagementUsername = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetPassword(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.Password = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetPwd(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.Pwd = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetUser(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.User = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) SetUserName(v string) *UpdateFlashSmsAccessProfileRequestAccessProfile {
	s.UserName = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfile) Validate() error {
	if s.DySmsAccessProfiles != nil {
		for _, item := range s.DySmsAccessProfiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles struct {
	// The template content.
	//
	// example:
	//
	// We tried to reach you but you were unavailable. Our staff will contact you again shortly. We apologize for any inconvenience
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template name.
	//
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The signature name.
	//
	// example:
	//
	// Cloud Call Center
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The template code.
	//
	// example:
	//
	// SMS_469075249
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GoString() string {
	return s.String()
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetName() *string {
	return s.Name
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetSignName() *string {
	return s.SignName
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetDescription(v string) *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.Description = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetName(v string) *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.Name = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetSignName(v string) *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.SignName = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetTemplateCode(v string) *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.TemplateCode = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) Validate() error {
	return dara.Validate(s)
}
