// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *CreateAiModelProviderRequest
	GetDisplayName() *string
	SetGatewayId(v string) *CreateAiModelProviderRequest
	GetGatewayId() *string
	SetProvider(v string) *CreateAiModelProviderRequest
	GetProvider() *string
	SetServiceIds(v []*string) *CreateAiModelProviderRequest
	GetServiceIds() []*string
	SetClientToken(v string) *CreateAiModelProviderRequest
	GetClientToken() *string
}

type CreateAiModelProviderRequest struct {
	// The display name of the model provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenAI
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-ucbx3s2m****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The model provider identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// openai
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The list of service IDs to bind to the provider.
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateAiModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAiModelProviderRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateAiModelProviderRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateAiModelProviderRequest) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *CreateAiModelProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAiModelProviderRequest) SetDisplayName(v string) *CreateAiModelProviderRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAiModelProviderRequest) SetGatewayId(v string) *CreateAiModelProviderRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateAiModelProviderRequest) SetProvider(v string) *CreateAiModelProviderRequest {
	s.Provider = &v
	return s
}

func (s *CreateAiModelProviderRequest) SetServiceIds(v []*string) *CreateAiModelProviderRequest {
	s.ServiceIds = v
	return s
}

func (s *CreateAiModelProviderRequest) SetClientToken(v string) *CreateAiModelProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAiModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
