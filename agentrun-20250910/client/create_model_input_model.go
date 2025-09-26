// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateModelInput
	GetAddress() *string
	SetApiKey(v string) *CreateModelInput
	GetApiKey() *string
	SetDesc(v string) *CreateModelInput
	GetDesc() *string
	SetModels(v []*string) *CreateModelInput
	GetModels() []*string
	SetName(v string) *CreateModelInput
	GetName() *string
	SetProvider(v string) *CreateModelInput
	GetProvider() *string
	SetType(v string) *CreateModelInput
	GetType() *string
}

type CreateModelInput struct {
	Address  *string   `json:"address,omitempty" xml:"address,omitempty"`
	ApiKey   *string   `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Desc     *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Models   []*string `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	Name     *string   `json:"name,omitempty" xml:"name,omitempty"`
	Provider *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateModelInput) String() string {
	return dara.Prettify(s)
}

func (s CreateModelInput) GoString() string {
	return s.String()
}

func (s *CreateModelInput) GetAddress() *string {
	return s.Address
}

func (s *CreateModelInput) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateModelInput) GetDesc() *string {
	return s.Desc
}

func (s *CreateModelInput) GetModels() []*string {
	return s.Models
}

func (s *CreateModelInput) GetName() *string {
	return s.Name
}

func (s *CreateModelInput) GetProvider() *string {
	return s.Provider
}

func (s *CreateModelInput) GetType() *string {
	return s.Type
}

func (s *CreateModelInput) SetAddress(v string) *CreateModelInput {
	s.Address = &v
	return s
}

func (s *CreateModelInput) SetApiKey(v string) *CreateModelInput {
	s.ApiKey = &v
	return s
}

func (s *CreateModelInput) SetDesc(v string) *CreateModelInput {
	s.Desc = &v
	return s
}

func (s *CreateModelInput) SetModels(v []*string) *CreateModelInput {
	s.Models = v
	return s
}

func (s *CreateModelInput) SetName(v string) *CreateModelInput {
	s.Name = &v
	return s
}

func (s *CreateModelInput) SetProvider(v string) *CreateModelInput {
	s.Provider = &v
	return s
}

func (s *CreateModelInput) SetType(v string) *CreateModelInput {
	s.Type = &v
	return s
}

func (s *CreateModelInput) Validate() error {
	return dara.Validate(s)
}
