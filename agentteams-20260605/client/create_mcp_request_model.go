// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *CreateMcpRequest
	GetAddresses() []*string
	SetAuthConfig(v string) *CreateMcpRequest
	GetAuthConfig() *string
	SetAuthEnabled(v bool) *CreateMcpRequest
	GetAuthEnabled() *bool
	SetClientToken(v string) *CreateMcpRequest
	GetClientToken() *string
	SetCreateType(v string) *CreateMcpRequest
	GetCreateType() *string
	SetDescription(v string) *CreateMcpRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateMcpRequest
	GetInstanceId() *string
	SetName(v string) *CreateMcpRequest
	GetName() *string
	SetProtocol(v string) *CreateMcpRequest
	GetProtocol() *string
	SetSwaggerConfig(v string) *CreateMcpRequest
	GetSwaggerConfig() *string
}

type CreateMcpRequest struct {
	// This parameter is required.
	Addresses   []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	AuthConfig  *string   `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	AuthEnabled *bool     `json:"AuthEnabled,omitempty" xml:"AuthEnabled,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreateType  *string   `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s CreateMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreateMcpRequest) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *CreateMcpRequest) GetAuthEnabled() *bool {
	return s.AuthEnabled
}

func (s *CreateMcpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMcpRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CreateMcpRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMcpRequest) GetName() *string {
	return s.Name
}

func (s *CreateMcpRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpRequest) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *CreateMcpRequest) SetAddresses(v []*string) *CreateMcpRequest {
	s.Addresses = v
	return s
}

func (s *CreateMcpRequest) SetAuthConfig(v string) *CreateMcpRequest {
	s.AuthConfig = &v
	return s
}

func (s *CreateMcpRequest) SetAuthEnabled(v bool) *CreateMcpRequest {
	s.AuthEnabled = &v
	return s
}

func (s *CreateMcpRequest) SetClientToken(v string) *CreateMcpRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMcpRequest) SetCreateType(v string) *CreateMcpRequest {
	s.CreateType = &v
	return s
}

func (s *CreateMcpRequest) SetDescription(v string) *CreateMcpRequest {
	s.Description = &v
	return s
}

func (s *CreateMcpRequest) SetInstanceId(v string) *CreateMcpRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMcpRequest) SetName(v string) *CreateMcpRequest {
	s.Name = &v
	return s
}

func (s *CreateMcpRequest) SetProtocol(v string) *CreateMcpRequest {
	s.Protocol = &v
	return s
}

func (s *CreateMcpRequest) SetSwaggerConfig(v string) *CreateMcpRequest {
	s.SwaggerConfig = &v
	return s
}

func (s *CreateMcpRequest) Validate() error {
	return dara.Validate(s)
}
