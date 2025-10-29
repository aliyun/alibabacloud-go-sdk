// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApigLLMModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateApigLLMModelInput
	GetAddress() *string
	SetApiKey(v string) *UpdateApigLLMModelInput
	GetApiKey() *string
	SetDesc(v string) *UpdateApigLLMModelInput
	GetDesc() *string
	SetModels(v []*string) *UpdateApigLLMModelInput
	GetModels() []*string
	SetName(v string) *UpdateApigLLMModelInput
	GetName() *string
	SetProvider(v string) *UpdateApigLLMModelInput
	GetProvider() *string
	SetType(v string) *UpdateApigLLMModelInput
	GetType() *string
}

type UpdateApigLLMModelInput struct {
	Address  *string   `json:"address,omitempty" xml:"address,omitempty"`
	ApiKey   *string   `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Desc     *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Models   []*string `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	Name     *string   `json:"name,omitempty" xml:"name,omitempty"`
	Provider *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateApigLLMModelInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateApigLLMModelInput) GoString() string {
	return s.String()
}

func (s *UpdateApigLLMModelInput) GetAddress() *string {
	return s.Address
}

func (s *UpdateApigLLMModelInput) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateApigLLMModelInput) GetDesc() *string {
	return s.Desc
}

func (s *UpdateApigLLMModelInput) GetModels() []*string {
	return s.Models
}

func (s *UpdateApigLLMModelInput) GetName() *string {
	return s.Name
}

func (s *UpdateApigLLMModelInput) GetProvider() *string {
	return s.Provider
}

func (s *UpdateApigLLMModelInput) GetType() *string {
	return s.Type
}

func (s *UpdateApigLLMModelInput) SetAddress(v string) *UpdateApigLLMModelInput {
	s.Address = &v
	return s
}

func (s *UpdateApigLLMModelInput) SetApiKey(v string) *UpdateApigLLMModelInput {
	s.ApiKey = &v
	return s
}

func (s *UpdateApigLLMModelInput) SetDesc(v string) *UpdateApigLLMModelInput {
	s.Desc = &v
	return s
}

func (s *UpdateApigLLMModelInput) SetModels(v []*string) *UpdateApigLLMModelInput {
	s.Models = v
	return s
}

func (s *UpdateApigLLMModelInput) SetName(v string) *UpdateApigLLMModelInput {
	s.Name = &v
	return s
}

func (s *UpdateApigLLMModelInput) SetProvider(v string) *UpdateApigLLMModelInput {
	s.Provider = &v
	return s
}

func (s *UpdateApigLLMModelInput) SetType(v string) *UpdateApigLLMModelInput {
	s.Type = &v
	return s
}

func (s *UpdateApigLLMModelInput) Validate() error {
	return dara.Validate(s)
}
