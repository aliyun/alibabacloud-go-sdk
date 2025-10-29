// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApigLLMModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateApigLLMModelInput
	GetAddress() *string
	SetApiKey(v string) *CreateApigLLMModelInput
	GetApiKey() *string
	SetDesc(v string) *CreateApigLLMModelInput
	GetDesc() *string
	SetModels(v []*string) *CreateApigLLMModelInput
	GetModels() []*string
	SetName(v string) *CreateApigLLMModelInput
	GetName() *string
	SetProvider(v string) *CreateApigLLMModelInput
	GetProvider() *string
	SetType(v string) *CreateApigLLMModelInput
	GetType() *string
}

type CreateApigLLMModelInput struct {
	Address  *string   `json:"address,omitempty" xml:"address,omitempty"`
	ApiKey   *string   `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Desc     *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Models   []*string `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	Name     *string   `json:"name,omitempty" xml:"name,omitempty"`
	Provider *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateApigLLMModelInput) String() string {
	return dara.Prettify(s)
}

func (s CreateApigLLMModelInput) GoString() string {
	return s.String()
}

func (s *CreateApigLLMModelInput) GetAddress() *string {
	return s.Address
}

func (s *CreateApigLLMModelInput) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateApigLLMModelInput) GetDesc() *string {
	return s.Desc
}

func (s *CreateApigLLMModelInput) GetModels() []*string {
	return s.Models
}

func (s *CreateApigLLMModelInput) GetName() *string {
	return s.Name
}

func (s *CreateApigLLMModelInput) GetProvider() *string {
	return s.Provider
}

func (s *CreateApigLLMModelInput) GetType() *string {
	return s.Type
}

func (s *CreateApigLLMModelInput) SetAddress(v string) *CreateApigLLMModelInput {
	s.Address = &v
	return s
}

func (s *CreateApigLLMModelInput) SetApiKey(v string) *CreateApigLLMModelInput {
	s.ApiKey = &v
	return s
}

func (s *CreateApigLLMModelInput) SetDesc(v string) *CreateApigLLMModelInput {
	s.Desc = &v
	return s
}

func (s *CreateApigLLMModelInput) SetModels(v []*string) *CreateApigLLMModelInput {
	s.Models = v
	return s
}

func (s *CreateApigLLMModelInput) SetName(v string) *CreateApigLLMModelInput {
	s.Name = &v
	return s
}

func (s *CreateApigLLMModelInput) SetProvider(v string) *CreateApigLLMModelInput {
	s.Provider = &v
	return s
}

func (s *CreateApigLLMModelInput) SetType(v string) *CreateApigLLMModelInput {
	s.Type = &v
	return s
}

func (s *CreateApigLLMModelInput) Validate() error {
	return dara.Validate(s)
}
