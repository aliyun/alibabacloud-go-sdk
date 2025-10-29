// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialOutput interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]*string) *GetCredentialOutput
	GetConfig() map[string]*string
	SetCreatedAt(v string) *GetCredentialOutput
	GetCreatedAt() *string
	SetDescription(v string) *GetCredentialOutput
	GetDescription() *string
	SetId(v string) *GetCredentialOutput
	GetId() *string
	SetName(v string) *GetCredentialOutput
	GetName() *string
	SetSecret(v string) *GetCredentialOutput
	GetSecret() *string
	SetType(v string) *GetCredentialOutput
	GetType() *string
	SetUpdatedAt(v string) *GetCredentialOutput
	GetUpdatedAt() *string
}

type GetCredentialOutput struct {
	// 凭证的配置参数，以键值对形式存储
	//
	// example:
	//
	// api_endpoint=https://api.example.com,timeout=30
	Config      map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	CreatedAt   *string            `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string            `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string            `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string            `json:"name,omitempty" xml:"name,omitempty"`
	Secret      *string            `json:"secret,omitempty" xml:"secret,omitempty"`
	Type        *string            `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt   *string            `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetCredentialOutput) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialOutput) GoString() string {
	return s.String()
}

func (s *GetCredentialOutput) GetConfig() map[string]*string {
	return s.Config
}

func (s *GetCredentialOutput) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCredentialOutput) GetDescription() *string {
	return s.Description
}

func (s *GetCredentialOutput) GetId() *string {
	return s.Id
}

func (s *GetCredentialOutput) GetName() *string {
	return s.Name
}

func (s *GetCredentialOutput) GetSecret() *string {
	return s.Secret
}

func (s *GetCredentialOutput) GetType() *string {
	return s.Type
}

func (s *GetCredentialOutput) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCredentialOutput) SetConfig(v map[string]*string) *GetCredentialOutput {
	s.Config = v
	return s
}

func (s *GetCredentialOutput) SetCreatedAt(v string) *GetCredentialOutput {
	s.CreatedAt = &v
	return s
}

func (s *GetCredentialOutput) SetDescription(v string) *GetCredentialOutput {
	s.Description = &v
	return s
}

func (s *GetCredentialOutput) SetId(v string) *GetCredentialOutput {
	s.Id = &v
	return s
}

func (s *GetCredentialOutput) SetName(v string) *GetCredentialOutput {
	s.Name = &v
	return s
}

func (s *GetCredentialOutput) SetSecret(v string) *GetCredentialOutput {
	s.Secret = &v
	return s
}

func (s *GetCredentialOutput) SetType(v string) *GetCredentialOutput {
	s.Type = &v
	return s
}

func (s *GetCredentialOutput) SetUpdatedAt(v string) *GetCredentialOutput {
	s.UpdatedAt = &v
	return s
}

func (s *GetCredentialOutput) Validate() error {
	return dara.Validate(s)
}
