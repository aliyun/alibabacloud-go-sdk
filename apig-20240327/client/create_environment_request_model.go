// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *CreateEnvironmentRequest
	GetAlias() *string
	SetDescription(v string) *CreateEnvironmentRequest
	GetDescription() *string
	SetGatewayId(v string) *CreateEnvironmentRequest
	GetGatewayId() *string
	SetName(v string) *CreateEnvironmentRequest
	GetName() *string
	SetResourceGroupId(v string) *CreateEnvironmentRequest
	GetResourceGroupId() *string
}

type CreateEnvironmentRequest struct {
	// The request body.
	//
	// This parameter is required.
	//
	// example:
	//
	// The environment name.
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The environment alias.
	//
	// example:
	//
	// Test environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment description, such as its purpose and owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qasrdc0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Create environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Testing environment for xx project of xxx
	//
	// example:
	//
	// rg-ahr5uil8raz0rq3b
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) GetAlias() *string {
	return s.Alias
}

func (s *CreateEnvironmentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEnvironmentRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateEnvironmentRequest) GetName() *string {
	return s.Name
}

func (s *CreateEnvironmentRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEnvironmentRequest) SetAlias(v string) *CreateEnvironmentRequest {
	s.Alias = &v
	return s
}

func (s *CreateEnvironmentRequest) SetDescription(v string) *CreateEnvironmentRequest {
	s.Description = &v
	return s
}

func (s *CreateEnvironmentRequest) SetGatewayId(v string) *CreateEnvironmentRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateEnvironmentRequest) SetName(v string) *CreateEnvironmentRequest {
	s.Name = &v
	return s
}

func (s *CreateEnvironmentRequest) SetResourceGroupId(v string) *CreateEnvironmentRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
