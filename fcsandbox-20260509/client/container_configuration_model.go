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
	// The ID of the ACR Enterprise instance to which the container image belongs.
	//
	// example:
	//
	// cri-abcd1234efgh
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The container image address.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/demo/nginx:latest
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The repository credentials used to pull private images.
	//
	// example:
	//
	// {"username":"registry-user","password":"******"}
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
	// The password of the image repository.
	//
	// example:
	//
	// ******
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username of the image repository.
	//
	// example:
	//
	// registry-user
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
