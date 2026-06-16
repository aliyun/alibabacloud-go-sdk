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
	// A client token used to ensure the idempotence of the request. The client generates this value to make sure that it is unique among different requests. The value can be up to 64 ASCII characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The details of the custom privacy policy content.
	CustomPrivacyPolicyContents []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents `json:"CustomPrivacyPolicyContents,omitempty" xml:"CustomPrivacyPolicyContents,omitempty" type:"Repeated"`
	// The name of the custom privacy policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom Privacy Policy Name
	CustomPrivacyPolicyName *string `json:"CustomPrivacyPolicyName,omitempty" xml:"CustomPrivacyPolicyName,omitempty"`
	// The default language of the privacy policy.
	//
	// example:
	//
	// zh-Hans-CN
	DefaultLanguageCode *string `json:"DefaultLanguageCode,omitempty" xml:"DefaultLanguageCode,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the custom privacy policy.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The consent type for the privacy policy.
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
	// The items of the custom privacy policy.
	CustomPrivacyPolicyItems []*CreateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems `json:"CustomPrivacyPolicyItems,omitempty" xml:"CustomPrivacyPolicyItems,omitempty" type:"Repeated"`
	// The prompt for the custom privacy policy.
	//
	// example:
	//
	// Please read and agree：
	CustomPrivacyPolicyTip *string `json:"CustomPrivacyPolicyTip,omitempty" xml:"CustomPrivacyPolicyTip,omitempty"`
	// The language of the custom privacy policy. The value is the LanguageCode returned by the ListLanguages operation.
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
	// The name of the custom privacy policy item.
	//
	// example:
	//
	// Custom Privacy Policy Name
	CustomPrivacyPolicyItemName *string `json:"CustomPrivacyPolicyItemName,omitempty" xml:"CustomPrivacyPolicyItemName,omitempty"`
	// The endpoint of the custom privacy policy item.
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
