// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v string) *CreateChatConfigurationRequest
	GetConfiguration() *string
	SetName(v string) *CreateChatConfigurationRequest
	GetName() *string
	SetRamRole(v string) *CreateChatConfigurationRequest
	GetRamRole() *string
	SetRegionId(v string) *CreateChatConfigurationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateChatConfigurationRequest
	GetResourceGroupId() *string
	SetType(v string) *CreateChatConfigurationRequest
	GetType() *string
}

type CreateChatConfigurationRequest struct {
	// example:
	//
	// {"BotName":"chatops"}
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// example:
	//
	// chatops
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AliyunOOSDefaultRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// DingTalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateChatConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateChatConfigurationRequest) GetConfiguration() *string {
	return s.Configuration
}

func (s *CreateChatConfigurationRequest) GetName() *string {
	return s.Name
}

func (s *CreateChatConfigurationRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateChatConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateChatConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateChatConfigurationRequest) GetType() *string {
	return s.Type
}

func (s *CreateChatConfigurationRequest) SetConfiguration(v string) *CreateChatConfigurationRequest {
	s.Configuration = &v
	return s
}

func (s *CreateChatConfigurationRequest) SetName(v string) *CreateChatConfigurationRequest {
	s.Name = &v
	return s
}

func (s *CreateChatConfigurationRequest) SetRamRole(v string) *CreateChatConfigurationRequest {
	s.RamRole = &v
	return s
}

func (s *CreateChatConfigurationRequest) SetRegionId(v string) *CreateChatConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateChatConfigurationRequest) SetResourceGroupId(v string) *CreateChatConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateChatConfigurationRequest) SetType(v string) *CreateChatConfigurationRequest {
	s.Type = &v
	return s
}

func (s *CreateChatConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
