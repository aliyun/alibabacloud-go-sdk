// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlashSmsAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfile(v *CreateFlashSmsAccessProfileRequestAccessProfile) *CreateFlashSmsAccessProfileRequest
	GetAccessProfile() *CreateFlashSmsAccessProfileRequestAccessProfile
	SetInstanceId(v string) *CreateFlashSmsAccessProfileRequest
	GetInstanceId() *string
	SetProviderId(v string) *CreateFlashSmsAccessProfileRequest
	GetProviderId() *string
}

type CreateFlashSmsAccessProfileRequest struct {
	// The access configuration.
	AccessProfile *CreateFlashSmsAccessProfileRequestAccessProfile `json:"AccessProfile,omitempty" xml:"AccessProfile,omitempty" type:"Struct"`
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
	// - ChuangLan: Beijing Chuanglan Yunzhi Information Co., Ltd.
	//
	// - ChinaMobile: China Mobile.
	//
	// - ShangHaiTianNan: Shanghai Tiannan.
	//
	// - HeDao: Galexes.
	//
	// - DySms: Alibaba Communication.
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s CreateFlashSmsAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlashSmsAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateFlashSmsAccessProfileRequest) GetAccessProfile() *CreateFlashSmsAccessProfileRequestAccessProfile {
	return s.AccessProfile
}

func (s *CreateFlashSmsAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFlashSmsAccessProfileRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *CreateFlashSmsAccessProfileRequest) SetAccessProfile(v *CreateFlashSmsAccessProfileRequestAccessProfile) *CreateFlashSmsAccessProfileRequest {
	s.AccessProfile = v
	return s
}

func (s *CreateFlashSmsAccessProfileRequest) SetInstanceId(v string) *CreateFlashSmsAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequest) SetProviderId(v string) *CreateFlashSmsAccessProfileRequest {
	s.ProviderId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequest) Validate() error {
	if s.AccessProfile != nil {
		if err := s.AccessProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFlashSmsAccessProfileRequestAccessProfile struct {
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
	DySmsAccessProfiles []*CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles `json:"DySmsAccessProfiles,omitempty" xml:"DySmsAccessProfiles,omitempty" type:"Repeated"`
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

func (s CreateFlashSmsAccessProfileRequestAccessProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateFlashSmsAccessProfileRequestAccessProfile) GoString() string {
	return s.String()
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetAccount() *string {
	return s.Account
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetAesKey() *string {
	return s.AesKey
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetApiAccount() *string {
	return s.ApiAccount
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetApiId() *string {
	return s.ApiId
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetApiPassword() *string {
	return s.ApiPassword
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetCapAppId() *string {
	return s.CapAppId
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetDySmsAccessProfiles() []*CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	return s.DySmsAccessProfiles
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetExtno() *string {
	return s.Extno
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetManagementPassword() *string {
	return s.ManagementPassword
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetManagementSubUserId() *string {
	return s.ManagementSubUserId
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetManagementUsername() *string {
	return s.ManagementUsername
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetPassword() *string {
	return s.Password
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetPwd() *string {
	return s.Pwd
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetUser() *string {
	return s.User
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) GetUserName() *string {
	return s.UserName
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetAccount(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.Account = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetAesKey(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.AesKey = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetApiAccount(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiAccount = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetApiId(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetApiKey(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiKey = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetApiPassword(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.ApiPassword = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetCapAppId(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.CapAppId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetDySmsAccessProfiles(v []*CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.DySmsAccessProfiles = v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetExtno(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.Extno = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetManagementPassword(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.ManagementPassword = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetManagementSubUserId(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.ManagementSubUserId = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetManagementUsername(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.ManagementUsername = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetPassword(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.Password = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetPwd(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.Pwd = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetUser(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.User = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) SetUserName(v string) *CreateFlashSmsAccessProfileRequestAccessProfile {
	s.UserName = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfile) Validate() error {
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

type CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles struct {
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

func (s CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) String() string {
	return dara.Prettify(s)
}

func (s CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GoString() string {
	return s.String()
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetDescription() *string {
	return s.Description
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetName() *string {
	return s.Name
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetSignName() *string {
	return s.SignName
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetDescription(v string) *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.Description = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetName(v string) *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.Name = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetSignName(v string) *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.SignName = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) SetTemplateCode(v string) *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles {
	s.TemplateCode = &v
	return s
}

func (s *CreateFlashSmsAccessProfileRequestAccessProfileDySmsAccessProfiles) Validate() error {
	return dara.Validate(s)
}
