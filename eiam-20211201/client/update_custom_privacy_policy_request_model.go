// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomPrivacyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicyContents(v []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) *UpdateCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyContents() []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents
	SetCustomPrivacyPolicyId(v string) *UpdateCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyId() *string
	SetCustomPrivacyPolicyName(v string) *UpdateCustomPrivacyPolicyRequest
	GetCustomPrivacyPolicyName() *string
	SetDefaultLanguageCode(v string) *UpdateCustomPrivacyPolicyRequest
	GetDefaultLanguageCode() *string
	SetInstanceId(v string) *UpdateCustomPrivacyPolicyRequest
	GetInstanceId() *string
	SetUserConsentType(v string) *UpdateCustomPrivacyPolicyRequest
	GetUserConsentType() *string
}

type UpdateCustomPrivacyPolicyRequest struct {
	CustomPrivacyPolicyContents []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents `json:"CustomPrivacyPolicyContents,omitempty" xml:"CustomPrivacyPolicyContents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pp_neagxpoznsjdtxxxxx
	CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
	// example:
	//
	// Custom Privacy Policy Name
	CustomPrivacyPolicyName *string `json:"CustomPrivacyPolicyName,omitempty" xml:"CustomPrivacyPolicyName,omitempty"`
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
	// example:
	//
	// implied_consent
	UserConsentType *string `json:"UserConsentType,omitempty" xml:"UserConsentType,omitempty"`
}

func (s UpdateCustomPrivacyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomPrivacyPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyContents() []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	return s.CustomPrivacyPolicyContents
}

func (s *UpdateCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *UpdateCustomPrivacyPolicyRequest) GetCustomPrivacyPolicyName() *string {
	return s.CustomPrivacyPolicyName
}

func (s *UpdateCustomPrivacyPolicyRequest) GetDefaultLanguageCode() *string {
	return s.DefaultLanguageCode
}

func (s *UpdateCustomPrivacyPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCustomPrivacyPolicyRequest) GetUserConsentType() *string {
	return s.UserConsentType
}

func (s *UpdateCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyContents(v []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) *UpdateCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyContents = v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyId(v string) *UpdateCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequest) SetCustomPrivacyPolicyName(v string) *UpdateCustomPrivacyPolicyRequest {
	s.CustomPrivacyPolicyName = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequest) SetDefaultLanguageCode(v string) *UpdateCustomPrivacyPolicyRequest {
	s.DefaultLanguageCode = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequest) SetInstanceId(v string) *UpdateCustomPrivacyPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequest) SetUserConsentType(v string) *UpdateCustomPrivacyPolicyRequest {
	s.UserConsentType = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequest) Validate() error {
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

type UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents struct {
	CustomPrivacyPolicyItems []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems `json:"CustomPrivacyPolicyItems,omitempty" xml:"CustomPrivacyPolicyItems,omitempty" type:"Repeated"`
	// example:
	//
	// Please read and agree：
	CustomPrivacyPolicyTip *string `json:"CustomPrivacyPolicyTip,omitempty" xml:"CustomPrivacyPolicyTip,omitempty"`
	// example:
	//
	// zh-Hans-CN
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
}

func (s UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GoString() string {
	return s.String()
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GetCustomPrivacyPolicyItems() []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	return s.CustomPrivacyPolicyItems
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GetCustomPrivacyPolicyTip() *string {
	return s.CustomPrivacyPolicyTip
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) SetCustomPrivacyPolicyItems(v []*UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	s.CustomPrivacyPolicyItems = v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) SetCustomPrivacyPolicyTip(v string) *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	s.CustomPrivacyPolicyTip = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) SetLanguageCode(v string) *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents {
	s.LanguageCode = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContents) Validate() error {
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

type UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems struct {
	// example:
	//
	// Item Name
	CustomPrivacyPolicyItemName *string `json:"CustomPrivacyPolicyItemName,omitempty" xml:"CustomPrivacyPolicyItemName,omitempty"`
	// example:
	//
	// https://example.com
	CustomPrivacyPolicyItemUrl *string `json:"CustomPrivacyPolicyItemUrl,omitempty" xml:"CustomPrivacyPolicyItemUrl,omitempty"`
}

func (s UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GoString() string {
	return s.String()
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GetCustomPrivacyPolicyItemName() *string {
	return s.CustomPrivacyPolicyItemName
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GetCustomPrivacyPolicyItemUrl() *string {
	return s.CustomPrivacyPolicyItemUrl
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) SetCustomPrivacyPolicyItemName(v string) *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	s.CustomPrivacyPolicyItemName = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) SetCustomPrivacyPolicyItemUrl(v string) *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	s.CustomPrivacyPolicyItemUrl = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyRequestCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) Validate() error {
	return dara.Validate(s)
}
