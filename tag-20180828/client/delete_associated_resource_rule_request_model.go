// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAssociatedResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteAssociatedResourceRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAssociatedResourceRuleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteAssociatedResourceRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteAssociatedResourceRuleRequest
	GetResourceOwnerAccount() *string
	SetSettingName(v string) *DeleteAssociatedResourceRuleRequest
	GetSettingName() *string
}

type DeleteAssociatedResourceRuleRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The name of the associated resource tagging rule.
	//
	// For more information, see the **Rule Name*	- column in [Resource types that support the Associated Resource Tagging feature](https://help.aliyun.com/document_detail/2586330.html).
	//
	// example:
	//
	// rule:AttachEni-DetachEni-TagInstance:Ecs-Instance:Ecs-Eni
	SettingName *string `json:"SettingName,omitempty" xml:"SettingName,omitempty"`
}

func (s DeleteAssociatedResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAssociatedResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAssociatedResourceRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAssociatedResourceRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAssociatedResourceRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAssociatedResourceRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAssociatedResourceRuleRequest) GetSettingName() *string {
	return s.SettingName
}

func (s *DeleteAssociatedResourceRuleRequest) SetOwnerAccount(v string) *DeleteAssociatedResourceRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAssociatedResourceRuleRequest) SetOwnerId(v int64) *DeleteAssociatedResourceRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAssociatedResourceRuleRequest) SetRegionId(v string) *DeleteAssociatedResourceRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAssociatedResourceRuleRequest) SetResourceOwnerAccount(v string) *DeleteAssociatedResourceRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAssociatedResourceRuleRequest) SetSettingName(v string) *DeleteAssociatedResourceRuleRequest {
	s.SettingName = &v
	return s
}

func (s *DeleteAssociatedResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}
