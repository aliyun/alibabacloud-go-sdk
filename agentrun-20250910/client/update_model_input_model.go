// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateModelInput
	GetAddress() *string
	SetApiKey(v string) *UpdateModelInput
	GetApiKey() *string
	SetDesc(v string) *UpdateModelInput
	GetDesc() *string
	SetModels(v []*string) *UpdateModelInput
	GetModels() []*string
	SetName(v string) *UpdateModelInput
	GetName() *string
	SetProvider(v string) *UpdateModelInput
	GetProvider() *string
	SetType(v string) *UpdateModelInput
	GetType() *string
}

type UpdateModelInput struct {
	Address  *string   `json:"address,omitempty" xml:"address,omitempty"`
	ApiKey   *string   `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Desc     *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Models   []*string `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	Name     *string   `json:"name,omitempty" xml:"name,omitempty"`
	Provider *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateModelInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelInput) GoString() string {
	return s.String()
}

func (s *UpdateModelInput) GetAddress() *string {
	return s.Address
}

func (s *UpdateModelInput) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateModelInput) GetDesc() *string {
	return s.Desc
}

func (s *UpdateModelInput) GetModels() []*string {
	return s.Models
}

func (s *UpdateModelInput) GetName() *string {
	return s.Name
}

func (s *UpdateModelInput) GetProvider() *string {
	return s.Provider
}

func (s *UpdateModelInput) GetType() *string {
	return s.Type
}

func (s *UpdateModelInput) SetAddress(v string) *UpdateModelInput {
	s.Address = &v
	return s
}

func (s *UpdateModelInput) SetApiKey(v string) *UpdateModelInput {
	s.ApiKey = &v
	return s
}

func (s *UpdateModelInput) SetDesc(v string) *UpdateModelInput {
	s.Desc = &v
	return s
}

func (s *UpdateModelInput) SetModels(v []*string) *UpdateModelInput {
	s.Models = v
	return s
}

func (s *UpdateModelInput) SetName(v string) *UpdateModelInput {
	s.Name = &v
	return s
}

func (s *UpdateModelInput) SetProvider(v string) *UpdateModelInput {
	s.Provider = &v
	return s
}

func (s *UpdateModelInput) SetType(v string) *UpdateModelInput {
	s.Type = &v
	return s
}

func (s *UpdateModelInput) Validate() error {
	return dara.Validate(s)
}
