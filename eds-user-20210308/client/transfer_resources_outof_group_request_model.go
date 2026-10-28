// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferResourcesOutofGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppRuleId(v string) *TransferResourcesOutofGroupRequest
	GetAppRuleId() *string
	SetBusinessChannel(v string) *TransferResourcesOutofGroupRequest
	GetBusinessChannel() *string
	SetPlatform(v string) *TransferResourcesOutofGroupRequest
	GetPlatform() *string
	SetPolicyId(v string) *TransferResourcesOutofGroupRequest
	GetPolicyId() *string
	SetResourceGroupId(v string) *TransferResourcesOutofGroupRequest
	GetResourceGroupId() *string
	SetResources(v []*TransferResourcesOutofGroupRequestResources) *TransferResourcesOutofGroupRequest
	GetResources() []*TransferResourcesOutofGroupRequestResources
	SetTargetResourceGroupId(v string) *TransferResourcesOutofGroupRequest
	GetTargetResourceGroupId() *string
	SetTimerId(v string) *TransferResourcesOutofGroupRequest
	GetTimerId() *string
}

type TransferResourcesOutofGroupRequest struct {
	// The software blacklists and whitelists rule ID to associate with the resources after they are moved out of the resource group. This parameter is used only for pure move-out operations and cannot be set together with TargetResourceGroupId. If this parameter is left empty, no rule is associated.
	//
	// example:
	//
	// bwr-47wdwfe5x3ql7qwf4
	AppRuleId *string `json:"AppRuleId,omitempty" xml:"AppRuleId,omitempty"`
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// > Fixed value: AliyunConsole.
	//
	// > - For Wuying Enterprise Edition, set this parameter to AliyunConsole. Other platforms are not publicly available.
	//
	// example:
	//
	// AliyunConsole
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The policy ID to associate with the resources after they are moved out of the resource group. This parameter is used only for pure move-out operations and cannot be specified together with TargetResourceGroupId. If this parameter is left empty, the system default behavior is used.
	//
	// example:
	//
	// system-all-enabled-policy
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-66konqltcpx43****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resources.
	Resources []*TransferResourcesOutofGroupRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The ID of the target resource group to which you want to move the resources.
	//
	// example:
	//
	// rg-hcjo6dugpo1****
	TargetResourceGroupId *string `json:"TargetResourceGroupId,omitempty" xml:"TargetResourceGroupId,omitempty"`
	// The scheduled task ID to associate with the resources after they are moved out of the resource group. This parameter is used only for pure move-out operations and cannot be specified together with TargetResourceGroupId. If this parameter is left empty, no scheduled task is associated.
	//
	// example:
	//
	// ccg-0bf5d87dtipve4aqt
	TimerId *string `json:"TimerId,omitempty" xml:"TimerId,omitempty"`
}

func (s TransferResourcesOutofGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesOutofGroupRequest) GoString() string {
	return s.String()
}

func (s *TransferResourcesOutofGroupRequest) GetAppRuleId() *string {
	return s.AppRuleId
}

func (s *TransferResourcesOutofGroupRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *TransferResourcesOutofGroupRequest) GetPlatform() *string {
	return s.Platform
}

func (s *TransferResourcesOutofGroupRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *TransferResourcesOutofGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransferResourcesOutofGroupRequest) GetResources() []*TransferResourcesOutofGroupRequestResources {
	return s.Resources
}

func (s *TransferResourcesOutofGroupRequest) GetTargetResourceGroupId() *string {
	return s.TargetResourceGroupId
}

func (s *TransferResourcesOutofGroupRequest) GetTimerId() *string {
	return s.TimerId
}

func (s *TransferResourcesOutofGroupRequest) SetAppRuleId(v string) *TransferResourcesOutofGroupRequest {
	s.AppRuleId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetBusinessChannel(v string) *TransferResourcesOutofGroupRequest {
	s.BusinessChannel = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetPlatform(v string) *TransferResourcesOutofGroupRequest {
	s.Platform = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetPolicyId(v string) *TransferResourcesOutofGroupRequest {
	s.PolicyId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetResourceGroupId(v string) *TransferResourcesOutofGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetResources(v []*TransferResourcesOutofGroupRequestResources) *TransferResourcesOutofGroupRequest {
	s.Resources = v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetTargetResourceGroupId(v string) *TransferResourcesOutofGroupRequest {
	s.TargetResourceGroupId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) SetTimerId(v string) *TransferResourcesOutofGroupRequest {
	s.TimerId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequest) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TransferResourcesOutofGroupRequestResources struct {
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ecd-1shc5sk0vo****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s TransferResourcesOutofGroupRequestResources) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesOutofGroupRequestResources) GoString() string {
	return s.String()
}

func (s *TransferResourcesOutofGroupRequestResources) GetRegionId() *string {
	return s.RegionId
}

func (s *TransferResourcesOutofGroupRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *TransferResourcesOutofGroupRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *TransferResourcesOutofGroupRequestResources) SetRegionId(v string) *TransferResourcesOutofGroupRequestResources {
	s.RegionId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequestResources) SetResourceId(v string) *TransferResourcesOutofGroupRequestResources {
	s.ResourceId = &v
	return s
}

func (s *TransferResourcesOutofGroupRequestResources) SetResourceType(v string) *TransferResourcesOutofGroupRequestResources {
	s.ResourceType = &v
	return s
}

func (s *TransferResourcesOutofGroupRequestResources) Validate() error {
	return dara.Validate(s)
}
