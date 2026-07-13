// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateModelShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateModelShrinkRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateModelShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateModelShrinkRequest
	GetName() *string
	SetProtocolsShrink(v string) *CreateModelShrinkRequest
	GetProtocolsShrink() *string
	SetProvider(v string) *CreateModelShrinkRequest
	GetProvider() *string
	SetProviderId(v string) *CreateModelShrinkRequest
	GetProviderId() *string
	SetProviderName(v string) *CreateModelShrinkRequest
	GetProviderName() *string
}

type CreateModelShrinkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ProtocolsShrink *string `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
	Provider        *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// This parameter is required.
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// This parameter is required.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s CreateModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateModelShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateModelShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelShrinkRequest) GetProtocolsShrink() *string {
	return s.ProtocolsShrink
}

func (s *CreateModelShrinkRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelShrinkRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *CreateModelShrinkRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *CreateModelShrinkRequest) SetClientToken(v string) *CreateModelShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelShrinkRequest) SetDescription(v string) *CreateModelShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateModelShrinkRequest) SetInstanceId(v string) *CreateModelShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateModelShrinkRequest) SetName(v string) *CreateModelShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateModelShrinkRequest) SetProtocolsShrink(v string) *CreateModelShrinkRequest {
	s.ProtocolsShrink = &v
	return s
}

func (s *CreateModelShrinkRequest) SetProvider(v string) *CreateModelShrinkRequest {
	s.Provider = &v
	return s
}

func (s *CreateModelShrinkRequest) SetProviderId(v string) *CreateModelShrinkRequest {
	s.ProviderId = &v
	return s
}

func (s *CreateModelShrinkRequest) SetProviderName(v string) *CreateModelShrinkRequest {
	s.ProviderName = &v
	return s
}

func (s *CreateModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
