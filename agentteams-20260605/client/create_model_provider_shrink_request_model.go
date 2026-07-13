// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateModelProviderShrinkRequest
	GetAddress() *string
	SetApiKeysShrink(v string) *CreateModelProviderShrinkRequest
	GetApiKeysShrink() *string
	SetClientToken(v string) *CreateModelProviderShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateModelProviderShrinkRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateModelProviderShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateModelProviderShrinkRequest
	GetName() *string
	SetProtocolsShrink(v string) *CreateModelProviderShrinkRequest
	GetProtocolsShrink() *string
	SetProvider(v string) *CreateModelProviderShrinkRequest
	GetProvider() *string
}

type CreateModelProviderShrinkRequest struct {
	// This parameter is required.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// This parameter is required.
	ApiKeysShrink *string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// RUNNING
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
}

func (s CreateModelProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateModelProviderShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateModelProviderShrinkRequest) GetApiKeysShrink() *string {
	return s.ApiKeysShrink
}

func (s *CreateModelProviderShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelProviderShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelProviderShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateModelProviderShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelProviderShrinkRequest) GetProtocolsShrink() *string {
	return s.ProtocolsShrink
}

func (s *CreateModelProviderShrinkRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelProviderShrinkRequest) SetAddress(v string) *CreateModelProviderShrinkRequest {
	s.Address = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) SetApiKeysShrink(v string) *CreateModelProviderShrinkRequest {
	s.ApiKeysShrink = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) SetClientToken(v string) *CreateModelProviderShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) SetDescription(v string) *CreateModelProviderShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) SetInstanceId(v string) *CreateModelProviderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) SetName(v string) *CreateModelProviderShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) SetProtocolsShrink(v string) *CreateModelProviderShrinkRequest {
	s.ProtocolsShrink = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) SetProvider(v string) *CreateModelProviderShrinkRequest {
	s.Provider = &v
	return s
}

func (s *CreateModelProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
