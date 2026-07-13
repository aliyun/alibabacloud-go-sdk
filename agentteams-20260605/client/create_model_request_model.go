// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateModelRequest
	GetClientToken() *string
	SetDescription(v string) *CreateModelRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateModelRequest
	GetInstanceId() *string
	SetName(v string) *CreateModelRequest
	GetName() *string
	SetProtocols(v []*string) *CreateModelRequest
	GetProtocols() []*string
	SetProvider(v string) *CreateModelRequest
	GetProvider() *string
	SetProviderId(v string) *CreateModelRequest
	GetProviderId() *string
	SetProviderName(v string) *CreateModelRequest
	GetProviderName() *string
}

type CreateModelRequest struct {
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
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Provider  *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// This parameter is required.
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// This parameter is required.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s CreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateModelRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelRequest) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateModelRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *CreateModelRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *CreateModelRequest) SetClientToken(v string) *CreateModelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelRequest) SetDescription(v string) *CreateModelRequest {
	s.Description = &v
	return s
}

func (s *CreateModelRequest) SetInstanceId(v string) *CreateModelRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateModelRequest) SetName(v string) *CreateModelRequest {
	s.Name = &v
	return s
}

func (s *CreateModelRequest) SetProtocols(v []*string) *CreateModelRequest {
	s.Protocols = v
	return s
}

func (s *CreateModelRequest) SetProvider(v string) *CreateModelRequest {
	s.Provider = &v
	return s
}

func (s *CreateModelRequest) SetProviderId(v string) *CreateModelRequest {
	s.ProviderId = &v
	return s
}

func (s *CreateModelRequest) SetProviderName(v string) *CreateModelRequest {
	s.ProviderName = &v
	return s
}

func (s *CreateModelRequest) Validate() error {
	return dara.Validate(s)
}
