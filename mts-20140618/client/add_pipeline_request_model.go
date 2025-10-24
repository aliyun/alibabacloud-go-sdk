// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddPipelineRequest
	GetName() *string
	SetNotifyConfig(v string) *AddPipelineRequest
	GetNotifyConfig() *string
	SetOwnerAccount(v string) *AddPipelineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddPipelineRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddPipelineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddPipelineRequest
	GetResourceOwnerId() *int64
	SetRole(v string) *AddPipelineRequest
	GetRole() *string
	SetSpeed(v string) *AddPipelineRequest
	GetSpeed() *string
	SetSpeedLevel(v int64) *AddPipelineRequest
	GetSpeedLevel() *int64
}

type AddPipelineRequest struct {
	// The name of the MPS queue. The name can be up to 128 bytes in size.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-pipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Message Service (MNS) configuration.
	//
	// example:
	//
	// {"Topic":"mts-topic-1"}
	NotifyConfig         *string `json:"NotifyConfig,omitempty" xml:"NotifyConfig,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role.
	//
	// example:
	//
	// AliyunMTSDefaultRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The type of the MPS queue. Valid values:
	//
	// 	- **Boost**: MPS queue with transcoding speed boosted.
	//
	// 	- **Standard**: standard MPS queue.
	//
	// 	- **NarrowBandHDV2**: MPS queue that supports Narrowband HD 2.0.
	//
	// 	- **AIVideoCover**: MPS queue for intelligent snapshot capture.
	//
	// 	- **AIVideoTag**: MPS queue for video tagging. The supported regions are China (Shanghai), China (Beijing), and China (Hangzhou).
	//
	// Default value: **Standard**.
	//
	// example:
	//
	// Standard
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The level of the MPS queue. Valid values: **1 to 3**.
	//
	// example:
	//
	// 1
	SpeedLevel *int64 `json:"SpeedLevel,omitempty" xml:"SpeedLevel,omitempty"`
}

func (s AddPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineRequest) GoString() string {
	return s.String()
}

func (s *AddPipelineRequest) GetName() *string {
	return s.Name
}

func (s *AddPipelineRequest) GetNotifyConfig() *string {
	return s.NotifyConfig
}

func (s *AddPipelineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddPipelineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddPipelineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddPipelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddPipelineRequest) GetRole() *string {
	return s.Role
}

func (s *AddPipelineRequest) GetSpeed() *string {
	return s.Speed
}

func (s *AddPipelineRequest) GetSpeedLevel() *int64 {
	return s.SpeedLevel
}

func (s *AddPipelineRequest) SetName(v string) *AddPipelineRequest {
	s.Name = &v
	return s
}

func (s *AddPipelineRequest) SetNotifyConfig(v string) *AddPipelineRequest {
	s.NotifyConfig = &v
	return s
}

func (s *AddPipelineRequest) SetOwnerAccount(v string) *AddPipelineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddPipelineRequest) SetOwnerId(v int64) *AddPipelineRequest {
	s.OwnerId = &v
	return s
}

func (s *AddPipelineRequest) SetResourceOwnerAccount(v string) *AddPipelineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddPipelineRequest) SetResourceOwnerId(v int64) *AddPipelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddPipelineRequest) SetRole(v string) *AddPipelineRequest {
	s.Role = &v
	return s
}

func (s *AddPipelineRequest) SetSpeed(v string) *AddPipelineRequest {
	s.Speed = &v
	return s
}

func (s *AddPipelineRequest) SetSpeedLevel(v int64) *AddPipelineRequest {
	s.SpeedLevel = &v
	return s
}

func (s *AddPipelineRequest) Validate() error {
	return dara.Validate(s)
}
