// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateCopyAction interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *PublicUpdateTemplateCopyAction
	GetAcrInstanceId() *string
	SetEnabled(v bool) *PublicUpdateTemplateCopyAction
	GetEnabled() *bool
	SetImage(v string) *PublicUpdateTemplateCopyAction
	GetImage() *string
	SetRegistryConfig(v *PublicUpdateTemplateRegistryConfig) *PublicUpdateTemplateCopyAction
	GetRegistryConfig() *PublicUpdateTemplateRegistryConfig
	SetRegistryType(v string) *PublicUpdateTemplateCopyAction
	GetRegistryType() *string
}

type PublicUpdateTemplateCopyAction struct {
	// The ID of the destination ACR Enterprise instance.
	//
	// example:
	//
	// cri-abc123
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// Indicates whether image replication is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The destination image address for replication.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/my/ns:v2
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The image repository configuration.
	RegistryConfig *PublicUpdateTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The destination repository type.
	//
	// example:
	//
	// acr
	RegistryType *string `json:"registryType,omitempty" xml:"registryType,omitempty"`
}

func (s PublicUpdateTemplateCopyAction) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateCopyAction) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateCopyAction) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *PublicUpdateTemplateCopyAction) GetEnabled() *bool {
	return s.Enabled
}

func (s *PublicUpdateTemplateCopyAction) GetImage() *string {
	return s.Image
}

func (s *PublicUpdateTemplateCopyAction) GetRegistryConfig() *PublicUpdateTemplateRegistryConfig {
	return s.RegistryConfig
}

func (s *PublicUpdateTemplateCopyAction) GetRegistryType() *string {
	return s.RegistryType
}

func (s *PublicUpdateTemplateCopyAction) SetAcrInstanceId(v string) *PublicUpdateTemplateCopyAction {
	s.AcrInstanceId = &v
	return s
}

func (s *PublicUpdateTemplateCopyAction) SetEnabled(v bool) *PublicUpdateTemplateCopyAction {
	s.Enabled = &v
	return s
}

func (s *PublicUpdateTemplateCopyAction) SetImage(v string) *PublicUpdateTemplateCopyAction {
	s.Image = &v
	return s
}

func (s *PublicUpdateTemplateCopyAction) SetRegistryConfig(v *PublicUpdateTemplateRegistryConfig) *PublicUpdateTemplateCopyAction {
	s.RegistryConfig = v
	return s
}

func (s *PublicUpdateTemplateCopyAction) SetRegistryType(v string) *PublicUpdateTemplateCopyAction {
	s.RegistryType = &v
	return s
}

func (s *PublicUpdateTemplateCopyAction) Validate() error {
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
