// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssociatedTransferSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableExistingResourcesTransfer(v string) *UpdateAssociatedTransferSettingRequest
	GetEnableExistingResourcesTransfer() *string
	SetRuleSettings(v []*UpdateAssociatedTransferSettingRequestRuleSettings) *UpdateAssociatedTransferSettingRequest
	GetRuleSettings() []*UpdateAssociatedTransferSettingRequestRuleSettings
}

type UpdateAssociatedTransferSettingRequest struct {
	// Specifies whether to enable the Transfer Existing Associated Resources feature. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	EnableExistingResourcesTransfer *string `json:"EnableExistingResourcesTransfer,omitempty" xml:"EnableExistingResourcesTransfer,omitempty"`
	// The settings of transfer rules.
	RuleSettings []*UpdateAssociatedTransferSettingRequestRuleSettings `json:"RuleSettings,omitempty" xml:"RuleSettings,omitempty" type:"Repeated"`
}

func (s UpdateAssociatedTransferSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssociatedTransferSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingRequest) GetEnableExistingResourcesTransfer() *string {
	return s.EnableExistingResourcesTransfer
}

func (s *UpdateAssociatedTransferSettingRequest) GetRuleSettings() []*UpdateAssociatedTransferSettingRequestRuleSettings {
	return s.RuleSettings
}

func (s *UpdateAssociatedTransferSettingRequest) SetEnableExistingResourcesTransfer(v string) *UpdateAssociatedTransferSettingRequest {
	s.EnableExistingResourcesTransfer = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequest) SetRuleSettings(v []*UpdateAssociatedTransferSettingRequestRuleSettings) *UpdateAssociatedTransferSettingRequest {
	s.RuleSettings = v
	return s
}

func (s *UpdateAssociatedTransferSettingRequest) Validate() error {
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

type UpdateAssociatedTransferSettingRequestRuleSettings struct {
	// The type of the associated resource.
	//
	// You can obtain the resource type from the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// disk
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
	// The service code of the associated resource.
	//
	// You can obtain the service code from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	AssociatedService *string `json:"AssociatedService,omitempty" xml:"AssociatedService,omitempty"`
	// The type of the primary resource.
	//
	// You can obtain the resource type from the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	MasterResourceType *string `json:"MasterResourceType,omitempty" xml:"MasterResourceType,omitempty"`
	// The service code of the primary resource.
	//
	// You can obtain the service code from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
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
	// This parameter is required.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAssociatedTransferSettingRequestRuleSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssociatedTransferSettingRequestRuleSettings) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) GetAssociatedResourceType() *string {
	return s.AssociatedResourceType
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) GetAssociatedService() *string {
	return s.AssociatedService
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) GetMasterResourceType() *string {
	return s.MasterResourceType
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) GetMasterService() *string {
	return s.MasterService
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) GetStatus() *string {
	return s.Status
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetAssociatedResourceType(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.AssociatedResourceType = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetAssociatedService(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.AssociatedService = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetMasterResourceType(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.MasterResourceType = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetMasterService(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.MasterService = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) SetStatus(v string) *UpdateAssociatedTransferSettingRequestRuleSettings {
	s.Status = &v
	return s
}

func (s *UpdateAssociatedTransferSettingRequestRuleSettings) Validate() error {
	return dara.Validate(s)
}
