// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssociatedTransferSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedTransferSetting(v *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) *ListAssociatedTransferSettingResponseBody
	GetAssociatedTransferSetting() *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting
	SetRequestId(v string) *ListAssociatedTransferSettingResponseBody
	GetRequestId() *string
}

type ListAssociatedTransferSettingResponseBody struct {
	// The settings of the Transfer Associated Resources feature.
	AssociatedTransferSetting *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting `json:"AssociatedTransferSetting,omitempty" xml:"AssociatedTransferSetting,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7556FD65-45D2-5C45-9FC9-A7DE831C775C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssociatedTransferSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedTransferSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponseBody) GetAssociatedTransferSetting() *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	return s.AssociatedTransferSetting
}

func (s *ListAssociatedTransferSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssociatedTransferSettingResponseBody) SetAssociatedTransferSetting(v *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) *ListAssociatedTransferSettingResponseBody {
	s.AssociatedTransferSetting = v
	return s
}

func (s *ListAssociatedTransferSettingResponseBody) SetRequestId(v string) *ListAssociatedTransferSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBody) Validate() error {
	if s.AssociatedTransferSetting != nil {
		if err := s.AssociatedTransferSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 121998723923****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether the Transfer Existing Associated Resources feature is enabled. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	EnableExistingResourcesTransfer *string `json:"EnableExistingResourcesTransfer,omitempty" xml:"EnableExistingResourcesTransfer,omitempty"`
	// The settings of transfer rules.
	RuleSettings []*ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings `json:"RuleSettings,omitempty" xml:"RuleSettings,omitempty" type:"Repeated"`
	// The status of the Transfer Associated Resources feature. Valid values:
	//
	// 	- Enable: enabled
	//
	// 	- Disable: disabled
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) GetEnableExistingResourcesTransfer() *string {
	return s.EnableExistingResourcesTransfer
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) GetRuleSettings() []*ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	return s.RuleSettings
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) GetStatus() *string {
	return s.Status
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetAccountId(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.AccountId = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetEnableExistingResourcesTransfer(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.EnableExistingResourcesTransfer = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetRuleSettings(v []*ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.RuleSettings = v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) SetStatus(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting {
	s.Status = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSetting) Validate() error {
	if s.RuleSettings != nil {
		for _, item := range s.RuleSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings struct {
	// The type of the associated resource.
	//
	// example:
	//
	// disk
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
	// The service code of the associated resource.
	//
	// example:
	//
	// ecs
	AssociatedService *string `json:"AssociatedService,omitempty" xml:"AssociatedService,omitempty"`
	// The type of the primary resource.
	//
	// example:
	//
	// instance
	MasterResourceType *string `json:"MasterResourceType,omitempty" xml:"MasterResourceType,omitempty"`
	// The service code of the primary resource.
	//
	// example:
	//
	// ecs
	MasterService *string `json:"MasterService,omitempty" xml:"MasterService,omitempty"`
	// The status of the Transfer Associated Resources feature. Valid values:
	//
	// 	- Enable: enabled
	//
	// 	- Disable: disabled
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) GoString() string {
	return s.String()
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) GetAssociatedResourceType() *string {
	return s.AssociatedResourceType
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) GetAssociatedService() *string {
	return s.AssociatedService
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) GetMasterResourceType() *string {
	return s.MasterResourceType
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) GetMasterService() *string {
	return s.MasterService
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) GetStatus() *string {
	return s.Status
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetAssociatedResourceType(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.AssociatedResourceType = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetAssociatedService(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.AssociatedService = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetMasterResourceType(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.MasterResourceType = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetMasterService(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.MasterService = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) SetStatus(v string) *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings {
	s.Status = &v
	return s
}

func (s *ListAssociatedTransferSettingResponseBodyAssociatedTransferSettingRuleSettings) Validate() error {
	return dara.Validate(s)
}
