// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *GetCredentialOutput
	GetCreatedAt() *string
	SetCredentialAuthType(v string) *GetCredentialOutput
	GetCredentialAuthType() *string
	SetCredentialId(v string) *GetCredentialOutput
	GetCredentialId() *string
	SetCredentialName(v string) *GetCredentialOutput
	GetCredentialName() *string
	SetCredentialPublicConfig(v map[string]*string) *GetCredentialOutput
	GetCredentialPublicConfig() map[string]*string
	SetCredentialSecret(v string) *GetCredentialOutput
	GetCredentialSecret() *string
	SetCredentialSourceType(v string) *GetCredentialOutput
	GetCredentialSourceType() *string
	SetDescription(v string) *GetCredentialOutput
	GetDescription() *string
	SetEnabled(v bool) *GetCredentialOutput
	GetEnabled() *bool
	SetRelatedResources(v []*RelatedResource) *GetCredentialOutput
	GetRelatedResources() []*RelatedResource
	SetUpdatedAt(v string) *GetCredentialOutput
	GetUpdatedAt() *string
}

type GetCredentialOutput struct {
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

func (s GetCredentialOutput) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialOutput) GoString() string {
	return s.String()
}

func (s *GetCredentialOutput) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCredentialOutput) GetCredentialAuthType() *string {
	return s.CredentialAuthType
}

func (s *GetCredentialOutput) GetCredentialId() *string {
	return s.CredentialId
}

func (s *GetCredentialOutput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *GetCredentialOutput) GetCredentialPublicConfig() map[string]*string {
	return s.CredentialPublicConfig
}

func (s *GetCredentialOutput) GetCredentialSecret() *string {
	return s.CredentialSecret
}

func (s *GetCredentialOutput) GetCredentialSourceType() *string {
	return s.CredentialSourceType
}

func (s *GetCredentialOutput) GetDescription() *string {
	return s.Description
}

func (s *GetCredentialOutput) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetCredentialOutput) GetRelatedResources() []*RelatedResource {
	return s.RelatedResources
}

func (s *GetCredentialOutput) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCredentialOutput) SetCreatedAt(v string) *GetCredentialOutput {
	s.CreatedAt = &v
	return s
}

func (s *GetCredentialOutput) SetCredentialAuthType(v string) *GetCredentialOutput {
	s.CredentialAuthType = &v
	return s
}

func (s *GetCredentialOutput) SetCredentialId(v string) *GetCredentialOutput {
	s.CredentialId = &v
	return s
}

func (s *GetCredentialOutput) SetCredentialName(v string) *GetCredentialOutput {
	s.CredentialName = &v
	return s
}

func (s *GetCredentialOutput) SetCredentialPublicConfig(v map[string]*string) *GetCredentialOutput {
	s.CredentialPublicConfig = v
	return s
}

func (s *GetCredentialOutput) SetCredentialSecret(v string) *GetCredentialOutput {
	s.CredentialSecret = &v
	return s
}

func (s *GetCredentialOutput) SetCredentialSourceType(v string) *GetCredentialOutput {
	s.CredentialSourceType = &v
	return s
}

func (s *GetCredentialOutput) SetDescription(v string) *GetCredentialOutput {
	s.Description = &v
	return s
}

func (s *GetCredentialOutput) SetEnabled(v bool) *GetCredentialOutput {
	s.Enabled = &v
	return s
}

func (s *GetCredentialOutput) SetRelatedResources(v []*RelatedResource) *GetCredentialOutput {
	s.RelatedResources = v
	return s
}

func (s *GetCredentialOutput) SetUpdatedAt(v string) *GetCredentialOutput {
	s.UpdatedAt = &v
	return s
}

func (s *GetCredentialOutput) Validate() error {
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
