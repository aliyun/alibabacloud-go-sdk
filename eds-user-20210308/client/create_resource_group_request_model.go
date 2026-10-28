// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *CreateResourceGroupRequest
	GetAgentType() *string
	SetBusinessChannel(v string) *CreateResourceGroupRequest
	GetBusinessChannel() *string
	SetEnableAliyunResourceGroup(v bool) *CreateResourceGroupRequest
	GetEnableAliyunResourceGroup() *bool
	SetIsResourceGroupWithOfficeSite(v int64) *CreateResourceGroupRequest
	GetIsResourceGroupWithOfficeSite() *int64
	SetPlatform(v string) *CreateResourceGroupRequest
	GetPlatform() *string
	SetResourceClassification(v string) *CreateResourceGroupRequest
	GetResourceClassification() *string
	SetResourceGroupDescription(v string) *CreateResourceGroupRequest
	GetResourceGroupDescription() *string
	SetResourceGroupName(v string) *CreateResourceGroupRequest
	GetResourceGroupName() *string
}

type CreateResourceGroupRequest struct {
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel           *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	EnableAliyunResourceGroup *bool   `json:"EnableAliyunResourceGroup,omitempty" xml:"EnableAliyunResourceGroup,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	IsResourceGroupWithOfficeSite *int64 `json:"IsResourceGroupWithOfficeSite,omitempty" xml:"IsResourceGroupWithOfficeSite,omitempty"`
	// For Elastic Desktop Service (EDS) Enterprise Edition, set the value to `AliyunConsole`. Other platforms are not supported.
	//
	// example:
	//
	// AliyunConsole
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The resource category.
	//
	// example:
	//
	// Default
	ResourceClassification *string `json:"ResourceClassification,omitempty" xml:"ResourceClassification,omitempty"`
	// The description of the resource group. This parameter is optional. The description can be up to 500 characters in length and is saved as-is.
	//
	// example:
	//
	// Department A resource group
	ResourceGroupDescription *string `json:"ResourceGroupDescription,omitempty" xml:"ResourceGroupDescription,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// Department A resource group
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateResourceGroupRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *CreateResourceGroupRequest) GetEnableAliyunResourceGroup() *bool {
	return s.EnableAliyunResourceGroup
}

func (s *CreateResourceGroupRequest) GetIsResourceGroupWithOfficeSite() *int64 {
	return s.IsResourceGroupWithOfficeSite
}

func (s *CreateResourceGroupRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateResourceGroupRequest) GetResourceClassification() *string {
	return s.ResourceClassification
}

func (s *CreateResourceGroupRequest) GetResourceGroupDescription() *string {
	return s.ResourceGroupDescription
}

func (s *CreateResourceGroupRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CreateResourceGroupRequest) SetAgentType(v string) *CreateResourceGroupRequest {
	s.AgentType = &v
	return s
}

func (s *CreateResourceGroupRequest) SetBusinessChannel(v string) *CreateResourceGroupRequest {
	s.BusinessChannel = &v
	return s
}

func (s *CreateResourceGroupRequest) SetEnableAliyunResourceGroup(v bool) *CreateResourceGroupRequest {
	s.EnableAliyunResourceGroup = &v
	return s
}

func (s *CreateResourceGroupRequest) SetIsResourceGroupWithOfficeSite(v int64) *CreateResourceGroupRequest {
	s.IsResourceGroupWithOfficeSite = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPlatform(v string) *CreateResourceGroupRequest {
	s.Platform = &v
	return s
}

func (s *CreateResourceGroupRequest) SetResourceClassification(v string) *CreateResourceGroupRequest {
	s.ResourceClassification = &v
	return s
}

func (s *CreateResourceGroupRequest) SetResourceGroupDescription(v string) *CreateResourceGroupRequest {
	s.ResourceGroupDescription = &v
	return s
}

func (s *CreateResourceGroupRequest) SetResourceGroupName(v string) *CreateResourceGroupRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
