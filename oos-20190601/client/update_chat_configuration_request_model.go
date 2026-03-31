// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v string) *UpdateChatConfigurationRequest
	GetConfiguration() *string
	SetName(v string) *UpdateChatConfigurationRequest
	GetName() *string
	SetRamRole(v string) *UpdateChatConfigurationRequest
	GetRamRole() *string
	SetRegionId(v string) *UpdateChatConfigurationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateChatConfigurationRequest
	GetResourceGroupId() *string
	SetType(v string) *UpdateChatConfigurationRequest
	GetType() *string
}

type UpdateChatConfigurationRequest struct {
	// example:
	//
	// {"BotName":"chatops"}
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// example:
	//
	// Ginlong-AIops
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AliyunOOSLifecycleHook4CSRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-2cd3****9dc2
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// DingTalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateChatConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatConfigurationRequest) GetConfiguration() *string {
	return s.Configuration
}

func (s *UpdateChatConfigurationRequest) GetName() *string {
	return s.Name
}

func (s *UpdateChatConfigurationRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *UpdateChatConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateChatConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateChatConfigurationRequest) GetType() *string {
	return s.Type
}

func (s *UpdateChatConfigurationRequest) SetConfiguration(v string) *UpdateChatConfigurationRequest {
	s.Configuration = &v
	return s
}

func (s *UpdateChatConfigurationRequest) SetName(v string) *UpdateChatConfigurationRequest {
	s.Name = &v
	return s
}

func (s *UpdateChatConfigurationRequest) SetRamRole(v string) *UpdateChatConfigurationRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateChatConfigurationRequest) SetRegionId(v string) *UpdateChatConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateChatConfigurationRequest) SetResourceGroupId(v string) *UpdateChatConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateChatConfigurationRequest) SetType(v string) *UpdateChatConfigurationRequest {
	s.Type = &v
	return s
}

func (s *UpdateChatConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
