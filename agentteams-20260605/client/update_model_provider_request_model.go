// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateModelProviderRequest
	GetAddress() *string
	SetApiKeys(v []*string) *UpdateModelProviderRequest
	GetApiKeys() []*string
	SetClientToken(v string) *UpdateModelProviderRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateModelProviderRequest
	GetDescription() *string
	SetId(v string) *UpdateModelProviderRequest
	GetId() *string
	SetInstanceId(v string) *UpdateModelProviderRequest
	GetInstanceId() *string
	SetProtocols(v []*string) *UpdateModelProviderRequest
	GetProtocols() []*string
}

type UpdateModelProviderRequest struct {
	// This parameter is required.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// This parameter is required.
	ApiKeys     []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
}

func (s UpdateModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateModelProviderRequest) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *UpdateModelProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelProviderRequest) GetId() *string {
	return s.Id
}

func (s *UpdateModelProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateModelProviderRequest) GetProtocols() []*string {
	return s.Protocols
}

func (s *UpdateModelProviderRequest) SetAddress(v string) *UpdateModelProviderRequest {
	s.Address = &v
	return s
}

func (s *UpdateModelProviderRequest) SetApiKeys(v []*string) *UpdateModelProviderRequest {
	s.ApiKeys = v
	return s
}

func (s *UpdateModelProviderRequest) SetClientToken(v string) *UpdateModelProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelProviderRequest) SetDescription(v string) *UpdateModelProviderRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelProviderRequest) SetId(v string) *UpdateModelProviderRequest {
	s.Id = &v
	return s
}

func (s *UpdateModelProviderRequest) SetInstanceId(v string) *UpdateModelProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateModelProviderRequest) SetProtocols(v []*string) *UpdateModelProviderRequest {
	s.Protocols = v
	return s
}

func (s *UpdateModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
