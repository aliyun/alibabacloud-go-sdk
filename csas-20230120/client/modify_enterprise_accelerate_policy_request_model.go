// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnterpriseAcceleratePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationType(v string) *ModifyEnterpriseAcceleratePolicyRequest
	GetAccelerationType() *string
	SetDescription(v string) *ModifyEnterpriseAcceleratePolicyRequest
	GetDescription() *string
	SetEapId(v string) *ModifyEnterpriseAcceleratePolicyRequest
	GetEapId() *string
	SetName(v string) *ModifyEnterpriseAcceleratePolicyRequest
	GetName() *string
	SetOnTls(v int32) *ModifyEnterpriseAcceleratePolicyRequest
	GetOnTls() *int32
	SetPriority(v int32) *ModifyEnterpriseAcceleratePolicyRequest
	GetPriority() *int32
	SetShowInClient(v int32) *ModifyEnterpriseAcceleratePolicyRequest
	GetShowInClient() *int32
	SetUpstreamHost(v string) *ModifyEnterpriseAcceleratePolicyRequest
	GetUpstreamHost() *string
	SetUpstreamPort(v int32) *ModifyEnterpriseAcceleratePolicyRequest
	GetUpstreamPort() *int32
	SetUpstreamType(v string) *ModifyEnterpriseAcceleratePolicyRequest
	GetUpstreamType() *string
	SetUserAttributeGroup(v string) *ModifyEnterpriseAcceleratePolicyRequest
	GetUserAttributeGroup() *string
}

type ModifyEnterpriseAcceleratePolicyRequest struct {
	// The acceleration mode:
	//
	// - **whiltelist**: whitelist acceleration
	//
	// - **global**: global acceleration
	//
	// - **build-in-list**: built-in system application acceleration
	//
	// example:
	//
	// whitelist
	AccelerationType *string `json:"AccelerationType,omitempty" xml:"AccelerationType,omitempty"`
	// The description of the enterprise acceleration policy.
	//
	// example:
	//
	// 这是一条内网访问策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the enterprise acceleration policy.
	//
	// example:
	//
	// eap-ce153a7165c8feea
	EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
	// The name of the enterprise acceleration policy.
	//
	// example:
	//
	// 白名单策略001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable TLS mode:
	//
	// - **0**: disable
	//
	// - **1**: enable
	//
	// example:
	//
	// 0
	OnTls *int32 `json:"OnTls,omitempty" xml:"OnTls,omitempty"`
	// The priority.
	//
	// example:
	//
	// 49
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Specifies whether to display on the client:
	//
	// - **0**: do not display
	//
	// - **1**: display
	//
	// example:
	//
	// 0
	ShowInClient *int32 `json:"ShowInClient,omitempty" xml:"ShowInClient,omitempty"`
	// The address (IP or domain name) of the acceleration instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.34.56.XX
	UpstreamHost *string `json:"UpstreamHost,omitempty" xml:"UpstreamHost,omitempty"`
	// The port of the acceleration instance (between 1000 and 60000).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	UpstreamPort *int32 `json:"UpstreamPort,omitempty" xml:"UpstreamPort,omitempty"`
	// The acceleration instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	// The acceleration user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试用户组
	UserAttributeGroup *string `json:"UserAttributeGroup,omitempty" xml:"UserAttributeGroup,omitempty"`
}

func (s ModifyEnterpriseAcceleratePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnterpriseAcceleratePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetEapId() *string {
	return s.EapId
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetOnTls() *int32 {
	return s.OnTls
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetShowInClient() *int32 {
	return s.ShowInClient
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetUpstreamHost() *string {
	return s.UpstreamHost
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetUpstreamPort() *int32 {
	return s.UpstreamPort
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) GetUserAttributeGroup() *string {
	return s.UserAttributeGroup
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetAccelerationType(v string) *ModifyEnterpriseAcceleratePolicyRequest {
	s.AccelerationType = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetDescription(v string) *ModifyEnterpriseAcceleratePolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetEapId(v string) *ModifyEnterpriseAcceleratePolicyRequest {
	s.EapId = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetName(v string) *ModifyEnterpriseAcceleratePolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetOnTls(v int32) *ModifyEnterpriseAcceleratePolicyRequest {
	s.OnTls = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetPriority(v int32) *ModifyEnterpriseAcceleratePolicyRequest {
	s.Priority = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetShowInClient(v int32) *ModifyEnterpriseAcceleratePolicyRequest {
	s.ShowInClient = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetUpstreamHost(v string) *ModifyEnterpriseAcceleratePolicyRequest {
	s.UpstreamHost = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetUpstreamPort(v int32) *ModifyEnterpriseAcceleratePolicyRequest {
	s.UpstreamPort = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetUpstreamType(v string) *ModifyEnterpriseAcceleratePolicyRequest {
	s.UpstreamType = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) SetUserAttributeGroup(v string) *ModifyEnterpriseAcceleratePolicyRequest {
	s.UserAttributeGroup = &v
	return s
}

func (s *ModifyEnterpriseAcceleratePolicyRequest) Validate() error {
	return dara.Validate(s)
}
