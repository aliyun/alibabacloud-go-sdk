// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationType(v string) *ImageConfig
	GetAccelerationType() *string
	SetImage(v string) *ImageConfig
	GetImage() *string
	SetInstanceID(v string) *ImageConfig
	GetInstanceID() *string
	SetRegistryConfig(v *RegistryConfig) *ImageConfig
	GetRegistryConfig() *RegistryConfig
}

type ImageConfig struct {
	AccelerationType *string         `json:"accelerationType,omitempty" xml:"accelerationType,omitempty"`
	Image            *string         `json:"image,omitempty" xml:"image,omitempty"`
	InstanceID       *string         `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	RegistryConfig   *RegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
}

func (s ImageConfig) String() string {
	return dara.Prettify(s)
}

func (s ImageConfig) GoString() string {
	return s.String()
}

func (s *ImageConfig) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *ImageConfig) GetImage() *string {
	return s.Image
}

func (s *ImageConfig) GetInstanceID() *string {
	return s.InstanceID
}

func (s *ImageConfig) GetRegistryConfig() *RegistryConfig {
	return s.RegistryConfig
}

func (s *ImageConfig) SetAccelerationType(v string) *ImageConfig {
	s.AccelerationType = &v
	return s
}

func (s *ImageConfig) SetImage(v string) *ImageConfig {
	s.Image = &v
	return s
}

func (s *ImageConfig) SetInstanceID(v string) *ImageConfig {
	s.InstanceID = &v
	return s
}

func (s *ImageConfig) SetRegistryConfig(v *RegistryConfig) *ImageConfig {
	s.RegistryConfig = v
	return s
}

func (s *ImageConfig) Validate() error {
	return dara.Validate(s)
}
