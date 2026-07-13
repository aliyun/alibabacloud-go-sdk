// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressesShrink(v string) *UpdateMcpShrinkRequest
	GetAddressesShrink() *string
	SetAuthConfig(v string) *UpdateMcpShrinkRequest
	GetAuthConfig() *string
	SetAuthEnabled(v bool) *UpdateMcpShrinkRequest
	GetAuthEnabled() *bool
	SetClientToken(v string) *UpdateMcpShrinkRequest
	GetClientToken() *string
	SetCreateType(v string) *UpdateMcpShrinkRequest
	GetCreateType() *string
	SetDescription(v string) *UpdateMcpShrinkRequest
	GetDescription() *string
	SetId(v string) *UpdateMcpShrinkRequest
	GetId() *string
	SetInstanceId(v string) *UpdateMcpShrinkRequest
	GetInstanceId() *string
	SetSwaggerConfig(v string) *UpdateMcpShrinkRequest
	GetSwaggerConfig() *string
}

type UpdateMcpShrinkRequest struct {
	// This parameter is required.
	AddressesShrink *string `json:"Addresses,omitempty" xml:"Addresses,omitempty"`
	AuthConfig      *string `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	AuthEnabled     *bool   `json:"AuthEnabled,omitempty" xml:"AuthEnabled,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SwaggerConfig *string `json:"SwaggerConfig,omitempty" xml:"SwaggerConfig,omitempty"`
}

func (s UpdateMcpShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpShrinkRequest) GetAddressesShrink() *string {
	return s.AddressesShrink
}

func (s *UpdateMcpShrinkRequest) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *UpdateMcpShrinkRequest) GetAuthEnabled() *bool {
	return s.AuthEnabled
}

func (s *UpdateMcpShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMcpShrinkRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *UpdateMcpShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMcpShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateMcpShrinkRequest) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *UpdateMcpShrinkRequest) SetAddressesShrink(v string) *UpdateMcpShrinkRequest {
	s.AddressesShrink = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetAuthConfig(v string) *UpdateMcpShrinkRequest {
	s.AuthConfig = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetAuthEnabled(v bool) *UpdateMcpShrinkRequest {
	s.AuthEnabled = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetClientToken(v string) *UpdateMcpShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetCreateType(v string) *UpdateMcpShrinkRequest {
	s.CreateType = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetDescription(v string) *UpdateMcpShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetId(v string) *UpdateMcpShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetInstanceId(v string) *UpdateMcpShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetSwaggerConfig(v string) *UpdateMcpShrinkRequest {
	s.SwaggerConfig = &v
	return s
}

func (s *UpdateMcpShrinkRequest) Validate() error {
	return dara.Validate(s)
}
