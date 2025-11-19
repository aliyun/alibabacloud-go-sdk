// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *CreateCredentialOutput
	GetCreatedAt() *string
	SetCredentialAuthType(v string) *CreateCredentialOutput
	GetCredentialAuthType() *string
	SetCredentialId(v string) *CreateCredentialOutput
	GetCredentialId() *string
	SetCredentialName(v string) *CreateCredentialOutput
	GetCredentialName() *string
	SetCredentialPublicConfig(v map[string]*string) *CreateCredentialOutput
	GetCredentialPublicConfig() map[string]*string
	SetCredentialSecret(v string) *CreateCredentialOutput
	GetCredentialSecret() *string
	SetCredentialSourceType(v string) *CreateCredentialOutput
	GetCredentialSourceType() *string
	SetDescription(v string) *CreateCredentialOutput
	GetDescription() *string
	SetEnabled(v bool) *CreateCredentialOutput
	GetEnabled() *bool
	SetRelatedResources(v []*RelatedResource) *CreateCredentialOutput
	GetRelatedResources() []*RelatedResource
	SetUpdatedAt(v string) *CreateCredentialOutput
	GetUpdatedAt() *string
}

type CreateCredentialOutput struct {
	CreatedAt              *string            `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialAuthType     *string            `json:"credentialAuthType,omitempty" xml:"credentialAuthType,omitempty"`
	CredentialId           *string            `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	CredentialName         *string            `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	CredentialPublicConfig map[string]*string `json:"credentialPublicConfig,omitempty" xml:"credentialPublicConfig,omitempty"`
	CredentialSecret       *string            `json:"credentialSecret,omitempty" xml:"credentialSecret,omitempty"`
	CredentialSourceType   *string            `json:"credentialSourceType,omitempty" xml:"credentialSourceType,omitempty"`
	Description            *string            `json:"description,omitempty" xml:"description,omitempty"`
	Enabled                *bool              `json:"enabled,omitempty" xml:"enabled,omitempty"`
	RelatedResources       []*RelatedResource `json:"relatedResources,omitempty" xml:"relatedResources,omitempty" type:"Repeated"`
	UpdatedAt              *string            `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateCredentialOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialOutput) GoString() string {
	return s.String()
}

func (s *CreateCredentialOutput) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCredentialOutput) GetCredentialAuthType() *string {
	return s.CredentialAuthType
}

func (s *CreateCredentialOutput) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateCredentialOutput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateCredentialOutput) GetCredentialPublicConfig() map[string]*string {
	return s.CredentialPublicConfig
}

func (s *CreateCredentialOutput) GetCredentialSecret() *string {
	return s.CredentialSecret
}

func (s *CreateCredentialOutput) GetCredentialSourceType() *string {
	return s.CredentialSourceType
}

func (s *CreateCredentialOutput) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialOutput) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateCredentialOutput) GetRelatedResources() []*RelatedResource {
	return s.RelatedResources
}

func (s *CreateCredentialOutput) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCredentialOutput) SetCreatedAt(v string) *CreateCredentialOutput {
	s.CreatedAt = &v
	return s
}

func (s *CreateCredentialOutput) SetCredentialAuthType(v string) *CreateCredentialOutput {
	s.CredentialAuthType = &v
	return s
}

func (s *CreateCredentialOutput) SetCredentialId(v string) *CreateCredentialOutput {
	s.CredentialId = &v
	return s
}

func (s *CreateCredentialOutput) SetCredentialName(v string) *CreateCredentialOutput {
	s.CredentialName = &v
	return s
}

func (s *CreateCredentialOutput) SetCredentialPublicConfig(v map[string]*string) *CreateCredentialOutput {
	s.CredentialPublicConfig = v
	return s
}

func (s *CreateCredentialOutput) SetCredentialSecret(v string) *CreateCredentialOutput {
	s.CredentialSecret = &v
	return s
}

func (s *CreateCredentialOutput) SetCredentialSourceType(v string) *CreateCredentialOutput {
	s.CredentialSourceType = &v
	return s
}

func (s *CreateCredentialOutput) SetDescription(v string) *CreateCredentialOutput {
	s.Description = &v
	return s
}

func (s *CreateCredentialOutput) SetEnabled(v bool) *CreateCredentialOutput {
	s.Enabled = &v
	return s
}

func (s *CreateCredentialOutput) SetRelatedResources(v []*RelatedResource) *CreateCredentialOutput {
	s.RelatedResources = v
	return s
}

func (s *CreateCredentialOutput) SetUpdatedAt(v string) *CreateCredentialOutput {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCredentialOutput) Validate() error {
	if s.RelatedResources != nil {
		for _, item := range s.RelatedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
