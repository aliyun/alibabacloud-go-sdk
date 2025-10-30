// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomPrivacyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicy(v *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) *GetCustomPrivacyPolicyResponseBody
	GetCustomPrivacyPolicy() *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy
	SetRequestId(v string) *GetCustomPrivacyPolicyResponseBody
	GetRequestId() *string
}

type GetCustomPrivacyPolicyResponseBody struct {
	CustomPrivacyPolicy *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy `json:"CustomPrivacyPolicy,omitempty" xml:"CustomPrivacyPolicy,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCustomPrivacyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomPrivacyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomPrivacyPolicyResponseBody) GetCustomPrivacyPolicy() *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	return s.CustomPrivacyPolicy
}

func (s *GetCustomPrivacyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomPrivacyPolicyResponseBody) SetCustomPrivacyPolicy(v *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) *GetCustomPrivacyPolicyResponseBody {
	s.CustomPrivacyPolicy = v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBody) SetRequestId(v string) *GetCustomPrivacyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBody) Validate() error {
	if s.CustomPrivacyPolicy != nil {
		if err := s.CustomPrivacyPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy struct {
	CustomPrivacyPolicyContents []*GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents `json:"CustomPrivacyPolicyContents,omitempty" xml:"CustomPrivacyPolicyContents,omitempty" type:"Repeated"`
	// example:
	//
	// pp_xxxxx
	CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
	// example:
	//
	// Custom Privacy Policy Name
	CustomPrivacyPolicyName *string `json:"CustomPrivacyPolicyName,omitempty" xml:"CustomPrivacyPolicyName,omitempty"`
	// example:
	//
	// zh-Hans-CN
	DefaultLanguageCode *string `json:"DefaultLanguageCode,omitempty" xml:"DefaultLanguageCode,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// implied_consent
	UserConsentType *string `json:"UserConsentType,omitempty" xml:"UserConsentType,omitempty"`
}

func (s GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GoString() string {
	return s.String()
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GetCustomPrivacyPolicyContents() []*GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents {
	return s.CustomPrivacyPolicyContents
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GetCustomPrivacyPolicyName() *string {
	return s.CustomPrivacyPolicyName
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GetDefaultLanguageCode() *string {
	return s.DefaultLanguageCode
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GetStatus() *string {
	return s.Status
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) GetUserConsentType() *string {
	return s.UserConsentType
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) SetCustomPrivacyPolicyContents(v []*GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	s.CustomPrivacyPolicyContents = v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) SetCustomPrivacyPolicyId(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) SetCustomPrivacyPolicyName(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	s.CustomPrivacyPolicyName = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) SetDefaultLanguageCode(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	s.DefaultLanguageCode = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) SetInstanceId(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	s.InstanceId = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) SetStatus(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	s.Status = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) SetUserConsentType(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy {
	s.UserConsentType = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicy) Validate() error {
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

type GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents struct {
	CustomPrivacyPolicyItems []*GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems `json:"CustomPrivacyPolicyItems,omitempty" xml:"CustomPrivacyPolicyItems,omitempty" type:"Repeated"`
	// example:
	//
	// Please read and agree：
	CustomPrivacyPolicyTip *string `json:"CustomPrivacyPolicyTip,omitempty" xml:"CustomPrivacyPolicyTip,omitempty"`
	// example:
	//
	// zh-Hans-CN
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
}

func (s GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) String() string {
	return dara.Prettify(s)
}

func (s GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) GoString() string {
	return s.String()
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) GetCustomPrivacyPolicyItems() []*GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	return s.CustomPrivacyPolicyItems
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) GetCustomPrivacyPolicyTip() *string {
	return s.CustomPrivacyPolicyTip
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) SetCustomPrivacyPolicyItems(v []*GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents {
	s.CustomPrivacyPolicyItems = v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) SetCustomPrivacyPolicyTip(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents {
	s.CustomPrivacyPolicyTip = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) SetLanguageCode(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents {
	s.LanguageCode = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContents) Validate() error {
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

type GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems struct {
	// example:
	//
	// item name
	CustomPrivacyPolicyItemName *string `json:"CustomPrivacyPolicyItemName,omitempty" xml:"CustomPrivacyPolicyItemName,omitempty"`
	// example:
	//
	// https://example.com
	CustomPrivacyPolicyItemUrl *string `json:"CustomPrivacyPolicyItemUrl,omitempty" xml:"CustomPrivacyPolicyItemUrl,omitempty"`
}

func (s GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) String() string {
	return dara.Prettify(s)
}

func (s GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GoString() string {
	return s.String()
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GetCustomPrivacyPolicyItemName() *string {
	return s.CustomPrivacyPolicyItemName
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) GetCustomPrivacyPolicyItemUrl() *string {
	return s.CustomPrivacyPolicyItemUrl
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) SetCustomPrivacyPolicyItemName(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	s.CustomPrivacyPolicyItemName = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) SetCustomPrivacyPolicyItemUrl(v string) *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems {
	s.CustomPrivacyPolicyItemUrl = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponseBodyCustomPrivacyPolicyCustomPrivacyPolicyContentsCustomPrivacyPolicyItems) Validate() error {
	return dara.Validate(s)
}
