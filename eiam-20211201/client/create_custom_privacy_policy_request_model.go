// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomPrivacyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCustomPrivacyPolicyRequest
	GetClientToken() *string
	SetCustomPrivacyPolicyContents(v []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) *CreateCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyContents() []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents
	SetCustomPrivacyPolicyName(v string) *CreateCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyName() *string
	SetDefaultLanguageCode(v string) *CreateCustomPrivacyPolicyRequest
	GetDefaultLanguageCode() *string
	SetInstanceId(v string) *CreateCustomPrivacyPolicyRequest
	GetInstanceId() *string
	SetStatus(v string) *CreateCustomPrivacyPolicyRequest
	GetStatus() *string
	SetUserConsentType(v string) *CreateCustomPrivacyPolicyRequest
	GetUserConsentType() *string
}

type CreateCustomPrivacyPolicyRequest struct {
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 自定义条款内容详情
	CustomPrivacyPolicyContents []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents `json:"CustomPrivacyPolicyContents,omitempty" xml:"CustomPrivacyPolicyContents,omitempty" type:"Repeated"`
	// 自定义条款名称
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom Privacy Policy Name
	CustomPrivacyPolicyName *string `json:"CustomPrivacyPolicyName,omitempty" xml:"CustomPrivacyPolicyName,omitempty"`
	// 默认条款语言，若其他语言未配置条款，则使用默认的
	//
	// example:
	//
	// zh-Hans-CN
	DefaultLanguageCode *string `json:"DefaultLanguageCode,omitempty" xml:"DefaultLanguageCode,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 自定义条款状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 手动勾选同意，还是默认同意
	//
	// example:
	//
	// implied_consent
	UserConsentType *string `json:"UserConsentType,omitempty" xml:"UserConsentType,omitempty"`
}

func (s CreateCustomPrivacyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomPrivacyPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomPrivacyPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyContents() []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	return s.CustomPrivacyPolicyContents
}

func (s *CreateCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyName() *string {
	return s.CustomPrivacyPolicyName
}

func (s *CreateCustomPrivacyPolicyRequest) GetDefaultLanguageCode() *string {
	return s.DefaultLanguageCode
}

func (s *CreateCustomPrivacyPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomPrivacyPolicyRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateCustomPrivacyPolicyRequest) GetUserConsentType() *string {
	return s.UserConsentType
}

func (s *CreateCustomPrivacyPolicyRequest) SetClientToken(v string) *CreateCustomPrivacyPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyContents(v []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) *CreateCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyContents = v
	return s
}

func (s *CreateCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyName(v string) *CreateCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyName = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequest) SetDefaultLanguageCode(v string) *CreateCustomPrivacyPolicyRequest {
	s.DefaultLanguageCode = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequest) SetInstanceId(v string) *CreateCustomPrivacyPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequest) SetStatus(v string) *CreateCustomPrivacyPolicyRequest {
	s.Status = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequest) SetUserConsentType(v string) *CreateCustomPrivacyPolicyRequest {
	s.UserConsentType = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequest) Validate() error {
	if s.CustomPrivacyPolicyContents != nil {
		for _, item := range s.CustomPrivacyPolicyContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents struct {
	// 自定义条款项
	CustomPrivacyPolicyItems []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems `json:"CustomPrivacyPolicyItems,omitempty" xml:"CustomPrivacyPolicyItems,omitempty" type:"Repeated"`
	// 自定义条款提示
	//
	// example:
	//
	// 登录视为同意此条款
	CustomPrivacyPolicyTip *string `json:"CustomPrivacyPolicyTip,omitempty" xml:"CustomPrivacyPolicyTip,omitempty"`
	// 自定义条款所属语言
	//
	// example:
	//
	// zh-Hans-CN
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
}

func (s CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GoString() string {
	return s.String()
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GetCustomPrivacyPolicyItems() []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	return s.CustomPrivacyPolicyItems
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GetCustomPrivacyPolicyTip() *string {
	return s.CustomPrivacyPolicyTip
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) SetCustomPrivacyPolicyItems(v []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	s.CustomPrivacyPolicyItems = v
	return s
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) SetCustomPrivacyPolicyTip(v string) *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	s.CustomPrivacyPolicyTip = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) SetLanguageCode(v string) *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	s.LanguageCode = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) Validate() error {
	if s.CustomPrivacyPolicyItems != nil {
		for _, item := range s.CustomPrivacyPolicyItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems struct {
	// 自定义条款名称
	//
	// example:
	//
	// xxxx隐私政策条款
	CustomPrivacyPolicyItemName *string `json:"CustomPrivacyPolicyItemName,omitempty" xml:"CustomPrivacyPolicyItemName,omitempty"`
	// 自定义条款访问地址
	//
	// example:
	//
	// http://www.xxxx.com
	CustomPrivacyPolicyItemUrl *string `json:"CustomPrivacyPolicyItemUrl,omitempty" xml:"CustomPrivacyPolicyItemUrl,omitempty"`
}

func (s CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GoString() string {
	return s.String()
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GetCustomPrivacyPolicyItemName() *string {
	return s.CustomPrivacyPolicyItemName
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GetCustomPrivacyPolicyItemUrl() *string {
	return s.CustomPrivacyPolicyItemUrl
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) SetCustomPrivacyPolicyItemName(v string) *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	s.CustomPrivacyPolicyItemName = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) SetCustomPrivacyPolicyItemUrl(v string) *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	s.CustomPrivacyPolicyItemUrl = &v
	return s
}

func (s *CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) Validate() error {
	return dara.Validate(s)
}
