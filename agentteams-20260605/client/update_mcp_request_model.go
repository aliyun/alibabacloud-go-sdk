// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *UpdateMcpRequest
	GetAddresses() []*string
	SetAuthConfig(v string) *UpdateMcpRequest
	GetAuthConfig() *string
	SetAuthEnabled(v bool) *UpdateMcpRequest
	GetAuthEnabled() *bool
	SetClientToken(v string) *UpdateMcpRequest
	GetClientToken() *string
	SetCreateType(v string) *UpdateMcpRequest
	GetCreateType() *string
	SetDescription(v string) *UpdateMcpRequest
	GetDescription() *string
	SetId(v string) *UpdateMcpRequest
	GetId() *string
	SetInstanceId(v string) *UpdateMcpRequest
	GetInstanceId() *string
	SetSwaggerConfig(v string) *UpdateMcpRequest
	GetSwaggerConfig() *string
}

type UpdateMcpRequest struct {
	// This parameter is required.
	Addresses   []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	AuthConfig  *string   `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	AuthEnabled *bool     `json:"AuthEnabled,omitempty" xml:"AuthEnabled,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreateType  *string   `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s UpdateMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdateMcpRequest) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *UpdateMcpRequest) GetAuthEnabled() *bool {
	return s.AuthEnabled
}

func (s *UpdateMcpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMcpRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *UpdateMcpRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMcpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateMcpRequest) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *UpdateMcpRequest) SetAddresses(v []*string) *UpdateMcpRequest {
	s.Addresses = v
	return s
}

func (s *UpdateMcpRequest) SetAuthConfig(v string) *UpdateMcpRequest {
	s.AuthConfig = &v
	return s
}

func (s *UpdateMcpRequest) SetAuthEnabled(v bool) *UpdateMcpRequest {
	s.AuthEnabled = &v
	return s
}

func (s *UpdateMcpRequest) SetClientToken(v string) *UpdateMcpRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMcpRequest) SetCreateType(v string) *UpdateMcpRequest {
	s.CreateType = &v
	return s
}

func (s *UpdateMcpRequest) SetDescription(v string) *UpdateMcpRequest {
	s.Description = &v
	return s
}

func (s *UpdateMcpRequest) SetId(v string) *UpdateMcpRequest {
	s.Id = &v
	return s
}

func (s *UpdateMcpRequest) SetInstanceId(v string) *UpdateMcpRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMcpRequest) SetSwaggerConfig(v string) *UpdateMcpRequest {
	s.SwaggerConfig = &v
	return s
}

func (s *UpdateMcpRequest) Validate() error {
	return dara.Validate(s)
}
