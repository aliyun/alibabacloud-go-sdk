// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *ContainerConfiguration
	GetAcrInstanceId() *string
	SetImage(v string) *ContainerConfiguration
	GetImage() *string
	SetRegistryCredential(v *ContainerConfigurationRegistryCredential) *ContainerConfiguration
	GetRegistryCredential() *ContainerConfigurationRegistryCredential
}

type ContainerConfiguration struct {
	// The ID of the Container Registry Enterprise instance. This parameter is required when registryType is set to acree.
	//
	// example:
	//
	// cri-****
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The address of the sandbox base image.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/sandbox-demo/python:3.12
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The repository access credentials used to pull images. You do not need to configure this parameter when you use public images.
	RegistryCredential *ContainerConfigurationRegistryCredential `json:"registryCredential,omitempty" xml:"registryCredential,omitempty" type:"Struct"`
}

func (s ContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ContainerConfiguration) GoString() string {
	return s.String()
}

func (s *ContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *ContainerConfiguration) GetRegistryCredential() *ContainerConfigurationRegistryCredential {
	return s.RegistryCredential
}

func (s *ContainerConfiguration) SetAcrInstanceId(v string) *ContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *ContainerConfiguration) SetImage(v string) *ContainerConfiguration {
	s.Image = &v
	return s
}

func (s *ContainerConfiguration) SetRegistryCredential(v *ContainerConfigurationRegistryCredential) *ContainerConfiguration {
	s.RegistryCredential = v
	return s
}

func (s *ContainerConfiguration) Validate() error {
	if s.RegistryCredential != nil {
		if err := s.RegistryCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ContainerConfigurationRegistryCredential struct {
	// The logon password or access token for the image repository. The value is encrypted and stored on the server, and plaintext is not returned in query responses.
	//
	// example:
	//
	// Passw0rd****
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The logon username for the image repository.
	//
	// example:
	//
	// sandbox-puller
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ContainerConfigurationRegistryCredential) String() string {
	return dara.Prettify(s)
}

func (s ContainerConfigurationRegistryCredential) GoString() string {
	return s.String()
}

func (s *ContainerConfigurationRegistryCredential) GetPassword() *string {
	return s.Password
}

func (s *ContainerConfigurationRegistryCredential) GetUsername() *string {
	return s.Username
}

func (s *ContainerConfigurationRegistryCredential) SetPassword(v string) *ContainerConfigurationRegistryCredential {
	s.Password = &v
	return s
}

func (s *ContainerConfigurationRegistryCredential) SetUsername(v string) *ContainerConfigurationRegistryCredential {
	s.Username = &v
	return s
}

func (s *ContainerConfigurationRegistryCredential) Validate() error {
	return dara.Validate(s)
}
