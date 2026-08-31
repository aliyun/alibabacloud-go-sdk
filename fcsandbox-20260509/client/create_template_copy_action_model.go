// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateCopyAction interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *CreateTemplateCopyAction
	GetAcrInstanceId() *string
	SetEnabled(v bool) *CreateTemplateCopyAction
	GetEnabled() *bool
	SetImage(v string) *CreateTemplateCopyAction
	GetImage() *string
	SetRegistryConfig(v *CreateTemplateRegistryConfig) *CreateTemplateCopyAction
	GetRegistryConfig() *CreateTemplateRegistryConfig
	SetRegistryType(v string) *CreateTemplateCopyAction
	GetRegistryType() *string
}

type CreateTemplateCopyAction struct {
	// The ID of the ACR Enterprise instance where the source image resides.
	//
	// example:
	//
	// cri-abcd1234efgh
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// Specifies whether to enable image replication.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The source image address.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/demo/nginx:latest
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The source image repository configuration.
	RegistryConfig *CreateTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The source image repository type.
	//
	// example:
	//
	// acr
	RegistryType *string `json:"registryType,omitempty" xml:"registryType,omitempty"`
}

func (s CreateTemplateCopyAction) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateCopyAction) GoString() string {
	return s.String()
}

func (s *CreateTemplateCopyAction) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateTemplateCopyAction) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateTemplateCopyAction) GetImage() *string {
	return s.Image
}

func (s *CreateTemplateCopyAction) GetRegistryConfig() *CreateTemplateRegistryConfig {
	return s.RegistryConfig
}

func (s *CreateTemplateCopyAction) GetRegistryType() *string {
	return s.RegistryType
}

func (s *CreateTemplateCopyAction) SetAcrInstanceId(v string) *CreateTemplateCopyAction {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateTemplateCopyAction) SetEnabled(v bool) *CreateTemplateCopyAction {
	s.Enabled = &v
	return s
}

func (s *CreateTemplateCopyAction) SetImage(v string) *CreateTemplateCopyAction {
	s.Image = &v
	return s
}

func (s *CreateTemplateCopyAction) SetRegistryConfig(v *CreateTemplateRegistryConfig) *CreateTemplateCopyAction {
	s.RegistryConfig = v
	return s
}

func (s *CreateTemplateCopyAction) SetRegistryType(v string) *CreateTemplateCopyAction {
	s.RegistryType = &v
	return s
}

func (s *CreateTemplateCopyAction) Validate() error {
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
