// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssociatedResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExistingStatus(v string) *UpdateAssociatedResourceRuleRequest
	GetExistingStatus() *string
	SetOwnerAccount(v string) *UpdateAssociatedResourceRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateAssociatedResourceRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateAssociatedResourceRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateAssociatedResourceRuleRequest
	GetResourceOwnerAccount() *string
	SetSettingName(v string) *UpdateAssociatedResourceRuleRequest
	GetSettingName() *string
	SetStatus(v string) *UpdateAssociatedResourceRuleRequest
	GetStatus() *string
	SetTagKeys(v []*string) *UpdateAssociatedResourceRuleRequest
	GetTagKeys() []*string
}

type UpdateAssociatedResourceRuleRequest struct {
	ExistingStatus *string `json:"ExistingStatus,omitempty" xml:"ExistingStatus,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The setting name of the associated resource tag rule.
	//
	// For specific values, see the **Rule Setting Name*	- column in [Resources that Support Associated Resource Tag Settings](https://help.aliyun.com/document_detail/2586330.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rule:AssociateEip-UnassociateEip-TagInstance:Ecs-Instance:Vpc-Eip
	SettingName *string `json:"SettingName,omitempty" xml:"SettingName,omitempty"`
	// Indicates whether to enable the associated resource tag rule. Values:
	//
	// - Enable: Enabled.
	//
	// - Disable: Disabled.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of tag keys affected by the associated resource tag rule.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s UpdateAssociatedResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssociatedResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssociatedResourceRuleRequest) GetExistingStatus() *string {
	return s.ExistingStatus
}

func (s *UpdateAssociatedResourceRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateAssociatedResourceRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAssociatedResourceRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAssociatedResourceRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAssociatedResourceRuleRequest) GetSettingName() *string {
	return s.SettingName
}

func (s *UpdateAssociatedResourceRuleRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAssociatedResourceRuleRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *UpdateAssociatedResourceRuleRequest) SetExistingStatus(v string) *UpdateAssociatedResourceRuleRequest {
	s.ExistingStatus = &v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) SetOwnerAccount(v string) *UpdateAssociatedResourceRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) SetOwnerId(v int64) *UpdateAssociatedResourceRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) SetRegionId(v string) *UpdateAssociatedResourceRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) SetResourceOwnerAccount(v string) *UpdateAssociatedResourceRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) SetSettingName(v string) *UpdateAssociatedResourceRuleRequest {
	s.SettingName = &v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) SetStatus(v string) *UpdateAssociatedResourceRuleRequest {
	s.Status = &v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) SetTagKeys(v []*string) *UpdateAssociatedResourceRuleRequest {
	s.TagKeys = v
	return s
}

func (s *UpdateAssociatedResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}
