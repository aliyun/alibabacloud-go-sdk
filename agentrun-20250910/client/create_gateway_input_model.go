// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayInput interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityId(v string) *CreateGatewayInput
	GetIdentityId() *string
	SetName(v string) *CreateGatewayInput
	GetName() *string
	SetNetworkConfiguration(v *GatewayNetworkConfiguration) *CreateGatewayInput
	GetNetworkConfiguration() *GatewayNetworkConfiguration
	SetType(v string) *CreateGatewayInput
	GetType() *string
}

type CreateGatewayInput struct {
	IdentityId           *string                      `json:"identityId,omitempty" xml:"identityId,omitempty"`
	Name                 *string                      `json:"name,omitempty" xml:"name,omitempty"`
	NetworkConfiguration *GatewayNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	Type                 *string                      `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateGatewayInput) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayInput) GoString() string {
	return s.String()
}

func (s *CreateGatewayInput) GetIdentityId() *string {
	return s.IdentityId
}

func (s *CreateGatewayInput) GetName() *string {
	return s.Name
}

func (s *CreateGatewayInput) GetNetworkConfiguration() *GatewayNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateGatewayInput) GetType() *string {
	return s.Type
}

func (s *CreateGatewayInput) SetIdentityId(v string) *CreateGatewayInput {
	s.IdentityId = &v
	return s
}

func (s *CreateGatewayInput) SetName(v string) *CreateGatewayInput {
	s.Name = &v
	return s
}

func (s *CreateGatewayInput) SetNetworkConfiguration(v *GatewayNetworkConfiguration) *CreateGatewayInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateGatewayInput) SetType(v string) *CreateGatewayInput {
	s.Type = &v
	return s
}

func (s *CreateGatewayInput) Validate() error {
	return dara.Validate(s)
}
