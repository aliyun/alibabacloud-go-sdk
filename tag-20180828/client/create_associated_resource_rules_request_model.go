// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssociatedResourceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateRulesList(v []*CreateAssociatedResourceRulesRequestCreateRulesList) *CreateAssociatedResourceRulesRequest
	GetCreateRulesList() []*CreateAssociatedResourceRulesRequestCreateRulesList
	SetOwnerAccount(v string) *CreateAssociatedResourceRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAssociatedResourceRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateAssociatedResourceRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateAssociatedResourceRulesRequest
	GetResourceOwnerAccount() *string
}

type CreateAssociatedResourceRulesRequest struct {
	// The associated resource tagging rules that you want to create.
	CreateRulesList []*CreateAssociatedResourceRulesRequestCreateRulesList `json:"CreateRulesList,omitempty" xml:"CreateRulesList,omitempty" type:"Repeated"`
	OwnerAccount    *string                                                `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64                                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s CreateAssociatedResourceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAssociatedResourceRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateAssociatedResourceRulesRequest) GetCreateRulesList() []*CreateAssociatedResourceRulesRequestCreateRulesList {
	return s.CreateRulesList
}

func (s *CreateAssociatedResourceRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAssociatedResourceRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAssociatedResourceRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAssociatedResourceRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAssociatedResourceRulesRequest) SetCreateRulesList(v []*CreateAssociatedResourceRulesRequestCreateRulesList) *CreateAssociatedResourceRulesRequest {
	s.CreateRulesList = v
	return s
}

func (s *CreateAssociatedResourceRulesRequest) SetOwnerAccount(v string) *CreateAssociatedResourceRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAssociatedResourceRulesRequest) SetOwnerId(v int64) *CreateAssociatedResourceRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAssociatedResourceRulesRequest) SetRegionId(v string) *CreateAssociatedResourceRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAssociatedResourceRulesRequest) SetResourceOwnerAccount(v string) *CreateAssociatedResourceRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAssociatedResourceRulesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAssociatedResourceRulesRequestCreateRulesList struct {
	ExistingStatus *string `json:"ExistingStatus,omitempty" xml:"ExistingStatus,omitempty"`
	// The name of the associated resource tagging rule.
	//
	// For more information, see the **Rule Name*	- column in [Resource types that support the Associated Resource Tagging feature](https://help.aliyun.com/document_detail/2586330.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rule:AttachEni-DetachEni-TagInstance:Ecs-Instance:Ecs-Eni
	SettingName *string `json:"SettingName,omitempty" xml:"SettingName,omitempty"`
	// Specifies whether to enable the associated resource tagging rule. Valid values:
	//
	// 	- Enable (default)
	//
	// 	- Disable
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag keys to which the associated resource tagging rule applies.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s CreateAssociatedResourceRulesRequestCreateRulesList) String() string {
	return dara.Prettify(s)
}

func (s CreateAssociatedResourceRulesRequestCreateRulesList) GoString() string {
	return s.String()
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) GetExistingStatus() *string {
	return s.ExistingStatus
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) GetSettingName() *string {
	return s.SettingName
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) GetStatus() *string {
	return s.Status
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) SetExistingStatus(v string) *CreateAssociatedResourceRulesRequestCreateRulesList {
	s.ExistingStatus = &v
	return s
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) SetSettingName(v string) *CreateAssociatedResourceRulesRequestCreateRulesList {
	s.SettingName = &v
	return s
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) SetStatus(v string) *CreateAssociatedResourceRulesRequestCreateRulesList {
	s.Status = &v
	return s
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) SetTagKeys(v []*string) *CreateAssociatedResourceRulesRequestCreateRulesList {
	s.TagKeys = v
	return s
}

func (s *CreateAssociatedResourceRulesRequestCreateRulesList) Validate() error {
	return dara.Validate(s)
}
