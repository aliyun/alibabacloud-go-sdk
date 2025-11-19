// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredential interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Credential
	GetCreatedAt() *string
	SetCredentialAuthType(v string) *Credential
	GetCredentialAuthType() *string
	SetCredentialId(v string) *Credential
	GetCredentialId() *string
	SetCredentialName(v string) *Credential
	GetCredentialName() *string
	SetCredentialPublicConfig(v map[string]*string) *Credential
	GetCredentialPublicConfig() map[string]*string
	SetCredentialSecret(v string) *Credential
	GetCredentialSecret() *string
	SetCredentialSourceType(v string) *Credential
	GetCredentialSourceType() *string
	SetDescription(v string) *Credential
	GetDescription() *string
	SetEnabled(v bool) *Credential
	GetEnabled() *bool
	SetRelatedResources(v []*RelatedResource) *Credential
	GetRelatedResources() []*RelatedResource
	SetUpdatedAt(v string) *Credential
	GetUpdatedAt() *string
}

type Credential struct {
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

func (s Credential) String() string {
	return dara.Prettify(s)
}

func (s Credential) GoString() string {
	return s.String()
}

func (s *Credential) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Credential) GetCredentialAuthType() *string {
	return s.CredentialAuthType
}

func (s *Credential) GetCredentialId() *string {
	return s.CredentialId
}

func (s *Credential) GetCredentialName() *string {
	return s.CredentialName
}

func (s *Credential) GetCredentialPublicConfig() map[string]*string {
	return s.CredentialPublicConfig
}

func (s *Credential) GetCredentialSecret() *string {
	return s.CredentialSecret
}

func (s *Credential) GetCredentialSourceType() *string {
	return s.CredentialSourceType
}

func (s *Credential) GetDescription() *string {
	return s.Description
}

func (s *Credential) GetEnabled() *bool {
	return s.Enabled
}

func (s *Credential) GetRelatedResources() []*RelatedResource {
	return s.RelatedResources
}

func (s *Credential) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Credential) SetCreatedAt(v string) *Credential {
	s.CreatedAt = &v
	return s
}

func (s *Credential) SetCredentialAuthType(v string) *Credential {
	s.CredentialAuthType = &v
	return s
}

func (s *Credential) SetCredentialId(v string) *Credential {
	s.CredentialId = &v
	return s
}

func (s *Credential) SetCredentialName(v string) *Credential {
	s.CredentialName = &v
	return s
}

func (s *Credential) SetCredentialPublicConfig(v map[string]*string) *Credential {
	s.CredentialPublicConfig = v
	return s
}

func (s *Credential) SetCredentialSecret(v string) *Credential {
	s.CredentialSecret = &v
	return s
}

func (s *Credential) SetCredentialSourceType(v string) *Credential {
	s.CredentialSourceType = &v
	return s
}

func (s *Credential) SetDescription(v string) *Credential {
	s.Description = &v
	return s
}

func (s *Credential) SetEnabled(v bool) *Credential {
	s.Enabled = &v
	return s
}

func (s *Credential) SetRelatedResources(v []*RelatedResource) *Credential {
	s.RelatedResources = v
	return s
}

func (s *Credential) SetUpdatedAt(v string) *Credential {
	s.UpdatedAt = &v
	return s
}

func (s *Credential) Validate() error {
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
