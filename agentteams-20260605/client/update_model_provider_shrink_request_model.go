// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateModelProviderShrinkRequest
	GetAddress() *string
	SetApiKeysShrink(v string) *UpdateModelProviderShrinkRequest
	GetApiKeysShrink() *string
	SetClientToken(v string) *UpdateModelProviderShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateModelProviderShrinkRequest
	GetDescription() *string
	SetId(v string) *UpdateModelProviderShrinkRequest
	GetId() *string
	SetInstanceId(v string) *UpdateModelProviderShrinkRequest
	GetInstanceId() *string
	SetProtocolsShrink(v string) *UpdateModelProviderShrinkRequest
	GetProtocolsShrink() *string
}

type UpdateModelProviderShrinkRequest struct {
	// This parameter is required.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// This parameter is required.
	ApiKeysShrink *string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ProtocolsShrink *string `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
}

func (s UpdateModelProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateModelProviderShrinkRequest) GetApiKeysShrink() *string {
	return s.ApiKeysShrink
}

func (s *UpdateModelProviderShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelProviderShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelProviderShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateModelProviderShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateModelProviderShrinkRequest) GetProtocolsShrink() *string {
	return s.ProtocolsShrink
}

func (s *UpdateModelProviderShrinkRequest) SetAddress(v string) *UpdateModelProviderShrinkRequest {
	s.Address = &v
	return s
}

func (s *UpdateModelProviderShrinkRequest) SetApiKeysShrink(v string) *UpdateModelProviderShrinkRequest {
	s.ApiKeysShrink = &v
	return s
}

func (s *UpdateModelProviderShrinkRequest) SetClientToken(v string) *UpdateModelProviderShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelProviderShrinkRequest) SetDescription(v string) *UpdateModelProviderShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelProviderShrinkRequest) SetId(v string) *UpdateModelProviderShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateModelProviderShrinkRequest) SetInstanceId(v string) *UpdateModelProviderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateModelProviderShrinkRequest) SetProtocolsShrink(v string) *UpdateModelProviderShrinkRequest {
	s.ProtocolsShrink = &v
	return s
}

func (s *UpdateModelProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
