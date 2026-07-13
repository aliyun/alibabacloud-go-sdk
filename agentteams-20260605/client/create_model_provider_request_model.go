// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateModelProviderRequest
	GetAddress() *string
	SetApiKeys(v []*string) *CreateModelProviderRequest
	GetApiKeys() []*string
	SetClientToken(v string) *CreateModelProviderRequest
	GetClientToken() *string
	SetDescription(v string) *CreateModelProviderRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateModelProviderRequest
	GetInstanceId() *string
	SetName(v string) *CreateModelProviderRequest
	GetName() *string
	SetProtocols(v []*string) *CreateModelProviderRequest
	GetProtocols() []*string
	SetProvider(v string) *CreateModelProviderRequest
	GetProvider() *string
}

type CreateModelProviderRequest struct {
	// This parameter is required.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// This parameter is required.
	ApiKeys     []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// RUNNING
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
}

func (s CreateModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateModelProviderRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateModelProviderRequest) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *CreateModelProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateModelProviderRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelProviderRequest) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateModelProviderRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelProviderRequest) SetAddress(v string) *CreateModelProviderRequest {
	s.Address = &v
	return s
}

func (s *CreateModelProviderRequest) SetApiKeys(v []*string) *CreateModelProviderRequest {
	s.ApiKeys = v
	return s
}

func (s *CreateModelProviderRequest) SetClientToken(v string) *CreateModelProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelProviderRequest) SetDescription(v string) *CreateModelProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateModelProviderRequest) SetInstanceId(v string) *CreateModelProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateModelProviderRequest) SetName(v string) *CreateModelProviderRequest {
	s.Name = &v
	return s
}

func (s *CreateModelProviderRequest) SetProtocols(v []*string) *CreateModelProviderRequest {
	s.Protocols = v
	return s
}

func (s *CreateModelProviderRequest) SetProvider(v string) *CreateModelProviderRequest {
	s.Provider = &v
	return s
}

func (s *CreateModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
