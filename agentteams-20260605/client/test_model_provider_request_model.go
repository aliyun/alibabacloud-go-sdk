// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *TestModelProviderRequest
	GetInstanceId() *string
	SetModelName(v string) *TestModelProviderRequest
	GetModelName() *string
	SetPrompt(v string) *TestModelProviderRequest
	GetPrompt() *string
	SetProviderId(v string) *TestModelProviderRequest
	GetProviderId() *string
	SetProviderName(v string) *TestModelProviderRequest
	GetProviderName() *string
}

type TestModelProviderRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RUNNING
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// This parameter is required.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s TestModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s TestModelProviderRequest) GoString() string {
	return s.String()
}

func (s *TestModelProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TestModelProviderRequest) GetModelName() *string {
	return s.ModelName
}

func (s *TestModelProviderRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *TestModelProviderRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *TestModelProviderRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *TestModelProviderRequest) SetInstanceId(v string) *TestModelProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *TestModelProviderRequest) SetModelName(v string) *TestModelProviderRequest {
	s.ModelName = &v
	return s
}

func (s *TestModelProviderRequest) SetPrompt(v string) *TestModelProviderRequest {
	s.Prompt = &v
	return s
}

func (s *TestModelProviderRequest) SetProviderId(v string) *TestModelProviderRequest {
	s.ProviderId = &v
	return s
}

func (s *TestModelProviderRequest) SetProviderName(v string) *TestModelProviderRequest {
	s.ProviderName = &v
	return s
}

func (s *TestModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
