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
	// The ID of the destination Container Registry Enterprise instance. This field does not inherit the instance ID of the source image.
	//
	// example:
	//
	// cri-****
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// Specifies whether to enable the image replication action. This parameter is required when this action block is passed.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The destination image address for the push operation. If this parameter is not provided, the service derives a new tag in the repository where the source image is located.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/sandbox-demo/python:3.12-envd
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The credentials, TLS settings, and network configurations required to access the destination image repository.
	RegistryConfig *CreateTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The type of the destination image repository. Valid values: oci (standard OCI repository), acr (Container Registry Personal Edition), and acree (Container Registry Enterprise Edition). By default, the value is derived from the destination image address.
	//
	// example:
	//
	// acree
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
