// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *UpdateCredentialOutput
	GetCreatedAt() *string
	SetCredentialAuthType(v string) *UpdateCredentialOutput
	GetCredentialAuthType() *string
	SetCredentialId(v string) *UpdateCredentialOutput
	GetCredentialId() *string
	SetCredentialName(v string) *UpdateCredentialOutput
	GetCredentialName() *string
	SetCredentialPublicConfig(v map[string]*string) *UpdateCredentialOutput
	GetCredentialPublicConfig() map[string]*string
	SetCredentialSecret(v string) *UpdateCredentialOutput
	GetCredentialSecret() *string
	SetCredentialSourceType(v string) *UpdateCredentialOutput
	GetCredentialSourceType() *string
	SetDescription(v string) *UpdateCredentialOutput
	GetDescription() *string
	SetEnabled(v bool) *UpdateCredentialOutput
	GetEnabled() *bool
	SetRelatedResources(v []*RelatedResource) *UpdateCredentialOutput
	GetRelatedResources() []*RelatedResource
	SetUpdatedAt(v string) *UpdateCredentialOutput
	GetUpdatedAt() *string
}

type UpdateCredentialOutput struct {
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

func (s UpdateCredentialOutput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialOutput) GoString() string {
	return s.String()
}

func (s *UpdateCredentialOutput) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCredentialOutput) GetCredentialAuthType() *string {
	return s.CredentialAuthType
}

func (s *UpdateCredentialOutput) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateCredentialOutput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *UpdateCredentialOutput) GetCredentialPublicConfig() map[string]*string {
	return s.CredentialPublicConfig
}

func (s *UpdateCredentialOutput) GetCredentialSecret() *string {
	return s.CredentialSecret
}

func (s *UpdateCredentialOutput) GetCredentialSourceType() *string {
	return s.CredentialSourceType
}

func (s *UpdateCredentialOutput) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialOutput) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCredentialOutput) GetRelatedResources() []*RelatedResource {
	return s.RelatedResources
}

func (s *UpdateCredentialOutput) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCredentialOutput) SetCreatedAt(v string) *UpdateCredentialOutput {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCredentialOutput) SetCredentialAuthType(v string) *UpdateCredentialOutput {
	s.CredentialAuthType = &v
	return s
}

func (s *UpdateCredentialOutput) SetCredentialId(v string) *UpdateCredentialOutput {
	s.CredentialId = &v
	return s
}

func (s *UpdateCredentialOutput) SetCredentialName(v string) *UpdateCredentialOutput {
	s.CredentialName = &v
	return s
}

func (s *UpdateCredentialOutput) SetCredentialPublicConfig(v map[string]*string) *UpdateCredentialOutput {
	s.CredentialPublicConfig = v
	return s
}

func (s *UpdateCredentialOutput) SetCredentialSecret(v string) *UpdateCredentialOutput {
	s.CredentialSecret = &v
	return s
}

func (s *UpdateCredentialOutput) SetCredentialSourceType(v string) *UpdateCredentialOutput {
	s.CredentialSourceType = &v
	return s
}

func (s *UpdateCredentialOutput) SetDescription(v string) *UpdateCredentialOutput {
	s.Description = &v
	return s
}

func (s *UpdateCredentialOutput) SetEnabled(v bool) *UpdateCredentialOutput {
	s.Enabled = &v
	return s
}

func (s *UpdateCredentialOutput) SetRelatedResources(v []*RelatedResource) *UpdateCredentialOutput {
	s.RelatedResources = v
	return s
}

func (s *UpdateCredentialOutput) SetUpdatedAt(v string) *UpdateCredentialOutput {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCredentialOutput) Validate() error {
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
