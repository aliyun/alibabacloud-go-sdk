// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseAcceleratePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationType(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetAccelerationType() *string
	SetDescription(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetDescription() *string
	SetName(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetName() *string
	SetPriority(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetPriority() *string
	SetShowInClient(v int32) *CreateEnterpriseAcceleratePolicyRequest
	GetShowInClient() *int32
	SetUpstreamHost(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetUpstreamHost() *string
	SetUpstreamPort(v int32) *CreateEnterpriseAcceleratePolicyRequest
	GetUpstreamPort() *int32
	SetUpstreamType(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetUpstreamType() *string
	SetUserAttributeGroup(v string) *CreateEnterpriseAcceleratePolicyRequest
	GetUserAttributeGroup() *string
}

type CreateEnterpriseAcceleratePolicyRequest struct {
	// Acceleration pattern:
	//
	// - **whitelist**: Whitelist acceleration
	//
	// - **global**: Global acceleration
	//
	// - **build-in-list**: Built-in application acceleration
	//
	// example:
	//
	// whitelist
	AccelerationType *string `json:"AccelerationType,omitempty" xml:"AccelerationType,omitempty"`
	// Policy description. Length: 1 to 512 characters.
	//
	// example:
	//
	// 用于全局网络访问的加速策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Policy Name.
	//
	// example:
	//
	// 全局加速策略
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Priority.
	//
	// example:
	//
	// 99
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Whether to display this policy in the client:
	//
	// - **0**: Do not display
	//
	// - **1**: Display
	//
	// example:
	//
	// 1
	ShowInClient *int32 `json:"ShowInClient,omitempty" xml:"ShowInClient,omitempty"`
	// The IP address or domain name of the acceleration instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.34.56.XX
	UpstreamHost *string `json:"UpstreamHost,omitempty" xml:"UpstreamHost,omitempty"`
	// Port for the accelerated instance. The port must be between 1000 and 60000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	UpstreamPort *int32 `json:"UpstreamPort,omitempty" xml:"UpstreamPort,omitempty"`
	// Accelerated instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	// User group for acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试用户组
	UserAttributeGroup *string `json:"UserAttributeGroup,omitempty" xml:"UserAttributeGroup,omitempty"`
}

func (s CreateEnterpriseAcceleratePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseAcceleratePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetShowInClient() *int32 {
	return s.ShowInClient
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUpstreamHost() *string {
	return s.UpstreamHost
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUpstreamPort() *int32 {
	return s.UpstreamPort
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *CreateEnterpriseAcceleratePolicyRequest) GetUserAttributeGroup() *string {
	return s.UserAttributeGroup
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetAccelerationType(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.AccelerationType = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetDescription(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetName(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetPriority(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetShowInClient(v int32) *CreateEnterpriseAcceleratePolicyRequest {
	s.ShowInClient = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUpstreamHost(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.UpstreamHost = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUpstreamPort(v int32) *CreateEnterpriseAcceleratePolicyRequest {
	s.UpstreamPort = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUpstreamType(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.UpstreamType = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) SetUserAttributeGroup(v string) *CreateEnterpriseAcceleratePolicyRequest {
	s.UserAttributeGroup = &v
	return s
}

func (s *CreateEnterpriseAcceleratePolicyRequest) Validate() error {
	return dara.Validate(s)
}
