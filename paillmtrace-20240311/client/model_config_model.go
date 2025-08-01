// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ModelConfig
	GetApiKey() *string
	SetEndpoint(v string) *ModelConfig
	GetEndpoint() *string
	SetIsSelfHost(v bool) *ModelConfig
	GetIsSelfHost() *bool
	SetName(v string) *ModelConfig
	GetName() *string
	SetTemperature(v float32) *ModelConfig
	GetTemperature() *float32
	SetTopP(v float32) *ModelConfig
	GetTopP() *float32
	SetUseFunctionCall(v bool) *ModelConfig
	GetUseFunctionCall() *bool
}

type ModelConfig struct {
	ApiKey          *string  `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	Endpoint        *string  `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	IsSelfHost      *bool    `json:"IsSelfHost,omitempty" xml:"IsSelfHost,omitempty"`
	Name            *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Temperature     *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	TopP            *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	UseFunctionCall *bool    `json:"UseFunctionCall,omitempty" xml:"UseFunctionCall,omitempty"`
}

func (s ModelConfig) String() string {
	return dara.Prettify(s)
}

func (s ModelConfig) GoString() string {
	return s.String()
}

func (s *ModelConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModelConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModelConfig) GetIsSelfHost() *bool {
	return s.IsSelfHost
}

func (s *ModelConfig) GetName() *string {
	return s.Name
}

func (s *ModelConfig) GetTemperature() *float32 {
	return s.Temperature
}

func (s *ModelConfig) GetTopP() *float32 {
	return s.TopP
}

func (s *ModelConfig) GetUseFunctionCall() *bool {
	return s.UseFunctionCall
}

func (s *ModelConfig) SetApiKey(v string) *ModelConfig {
	s.ApiKey = &v
	return s
}

func (s *ModelConfig) SetEndpoint(v string) *ModelConfig {
	s.Endpoint = &v
	return s
}

func (s *ModelConfig) SetIsSelfHost(v bool) *ModelConfig {
	s.IsSelfHost = &v
	return s
}

func (s *ModelConfig) SetName(v string) *ModelConfig {
	s.Name = &v
	return s
}

func (s *ModelConfig) SetTemperature(v float32) *ModelConfig {
	s.Temperature = &v
	return s
}

func (s *ModelConfig) SetTopP(v float32) *ModelConfig {
	s.TopP = &v
	return s
}

func (s *ModelConfig) SetUseFunctionCall(v bool) *ModelConfig {
	s.UseFunctionCall = &v
	return s
}

func (s *ModelConfig) Validate() error {
	return dara.Validate(s)
}
