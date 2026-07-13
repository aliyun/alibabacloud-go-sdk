// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressesShrink(v string) *CreateMcpShrinkRequest
	GetAddressesShrink() *string
	SetAuthConfig(v string) *CreateMcpShrinkRequest
	GetAuthConfig() *string
	SetAuthEnabled(v bool) *CreateMcpShrinkRequest
	GetAuthEnabled() *bool
	SetClientToken(v string) *CreateMcpShrinkRequest
	GetClientToken() *string
	SetCreateType(v string) *CreateMcpShrinkRequest
	GetCreateType() *string
	SetDescription(v string) *CreateMcpShrinkRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateMcpShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateMcpShrinkRequest
	GetName() *string
	SetProtocol(v string) *CreateMcpShrinkRequest
	GetProtocol() *string
	SetSwaggerConfig(v string) *CreateMcpShrinkRequest
	GetSwaggerConfig() *string
}

type CreateMcpShrinkRequest struct {
	// This parameter is required.
	AddressesShrink *string `json:"Addresses,omitempty" xml:"Addresses,omitempty"`
	AuthConfig      *string `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	AuthEnabled     *bool   `json:"AuthEnabled,omitempty" xml:"AuthEnabled,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol      *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SwaggerConfig *string `json:"SwaggerConfig,omitempty" xml:"SwaggerConfig,omitempty"`
}

func (s CreateMcpShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpShrinkRequest) GetAddressesShrink() *string {
	return s.AddressesShrink
}

func (s *CreateMcpShrinkRequest) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *CreateMcpShrinkRequest) GetAuthEnabled() *bool {
	return s.AuthEnabled
}

func (s *CreateMcpShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMcpShrinkRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CreateMcpShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMcpShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMcpShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpShrinkRequest) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *CreateMcpShrinkRequest) SetAddressesShrink(v string) *CreateMcpShrinkRequest {
	s.AddressesShrink = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetAuthConfig(v string) *CreateMcpShrinkRequest {
	s.AuthConfig = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetAuthEnabled(v bool) *CreateMcpShrinkRequest {
	s.AuthEnabled = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetClientToken(v string) *CreateMcpShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetCreateType(v string) *CreateMcpShrinkRequest {
	s.CreateType = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetDescription(v string) *CreateMcpShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetInstanceId(v string) *CreateMcpShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetName(v string) *CreateMcpShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetProtocol(v string) *CreateMcpShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetSwaggerConfig(v string) *CreateMcpShrinkRequest {
	s.SwaggerConfig = &v
	return s
}

func (s *CreateMcpShrinkRequest) Validate() error {
	return dara.Validate(s)
}
