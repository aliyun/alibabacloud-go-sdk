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
}

type CreateAiModelProviderRequest struct {
	// This parameter is required.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// This parameter is required.
	Provider   *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
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

func (s *CreateAiModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
