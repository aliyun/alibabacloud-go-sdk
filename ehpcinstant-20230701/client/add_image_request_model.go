// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerImageSpec(v *AddImageRequestContainerImageSpec) *AddImageRequest
	GetContainerImageSpec() *AddImageRequestContainerImageSpec
	SetDescription(v string) *AddImageRequest
	GetDescription() *string
	SetImageType(v string) *AddImageRequest
	GetImageType() *string
	SetImageVersion(v string) *AddImageRequest
	GetImageVersion() *string
	SetName(v string) *AddImageRequest
	GetName() *string
	SetVMImageSpec(v *AddImageRequestVMImageSpec) *AddImageRequest
	GetVMImageSpec() *AddImageRequestVMImageSpec
}

type AddImageRequest struct {
	// The configurations of the container image.
	ContainerImageSpec *AddImageRequestContainerImageSpec `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty" type:"Struct"`
	// The description of the image.
	//
	// example:
	//
	// Test image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the images. Valid values:
	//
	// 	- VM: virtual machine image.
	//
	// 	- Container: the container image.
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// V1.0
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The name of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The image configuration of the virtual machine.
	VMImageSpec *AddImageRequestVMImageSpec `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty" type:"Struct"`
}

func (s AddImageRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) GetContainerImageSpec() *AddImageRequestContainerImageSpec {
	return s.ContainerImageSpec
}

func (s *AddImageRequest) GetDescription() *string {
	return s.Description
}

func (s *AddImageRequest) GetImageType() *string {
	return s.ImageType
}

func (s *AddImageRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *AddImageRequest) GetName() *string {
	return s.Name
}

func (s *AddImageRequest) GetVMImageSpec() *AddImageRequestVMImageSpec {
	return s.VMImageSpec
}

func (s *AddImageRequest) SetContainerImageSpec(v *AddImageRequestContainerImageSpec) *AddImageRequest {
	s.ContainerImageSpec = v
	return s
}

func (s *AddImageRequest) SetDescription(v string) *AddImageRequest {
	s.Description = &v
	return s
}

func (s *AddImageRequest) SetImageType(v string) *AddImageRequest {
	s.ImageType = &v
	return s
}

func (s *AddImageRequest) SetImageVersion(v string) *AddImageRequest {
	s.ImageVersion = &v
	return s
}

func (s *AddImageRequest) SetName(v string) *AddImageRequest {
	s.Name = &v
	return s
}

func (s *AddImageRequest) SetVMImageSpec(v *AddImageRequestVMImageSpec) *AddImageRequest {
	s.VMImageSpec = v
	return s
}

func (s *AddImageRequest) Validate() error {
	if s.ContainerImageSpec != nil {
		if err := s.ContainerImageSpec.Validate(); err != nil {
			return err
		}
	}
	if s.VMImageSpec != nil {
		if err := s.VMImageSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddImageRequestContainerImageSpec struct {
	// Whether the instance is an Alibaba Cloud image repository Enterprise Edition.
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	IsACREnterprise *bool `json:"IsACREnterprise,omitempty" xml:"IsACREnterprise,omitempty"`
	// Whether it is an Alibaba Cloud image repository.
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	IsACRRegistry *bool `json:"IsACRRegistry,omitempty" xml:"IsACRRegistry,omitempty"`
	// The authentication of the private image repository.
	RegistryCredential *AddImageRequestContainerImageSpecRegistryCredential `json:"RegistryCredential,omitempty" xml:"RegistryCredential,omitempty" type:"Struct"`
	// The ID of the Container Registry Enterprise Edition image repository.
	//
	// example:
	//
	// cri-xyz795ygf8k9****
	RegistryCriId *string `json:"RegistryCriId,omitempty" xml:"RegistryCriId,omitempty"`
	// The endpoint of the container image.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc_open/nginx:latest
	RegistryUrl *string `json:"RegistryUrl,omitempty" xml:"RegistryUrl,omitempty"`
}

func (s AddImageRequestContainerImageSpec) String() string {
	return dara.Prettify(s)
}

func (s AddImageRequestContainerImageSpec) GoString() string {
	return s.String()
}

func (s *AddImageRequestContainerImageSpec) GetIsACREnterprise() *bool {
	return s.IsACREnterprise
}

func (s *AddImageRequestContainerImageSpec) GetIsACRRegistry() *bool {
	return s.IsACRRegistry
}

func (s *AddImageRequestContainerImageSpec) GetRegistryCredential() *AddImageRequestContainerImageSpecRegistryCredential {
	return s.RegistryCredential
}

func (s *AddImageRequestContainerImageSpec) GetRegistryCriId() *string {
	return s.RegistryCriId
}

func (s *AddImageRequestContainerImageSpec) GetRegistryUrl() *string {
	return s.RegistryUrl
}

func (s *AddImageRequestContainerImageSpec) SetIsACREnterprise(v bool) *AddImageRequestContainerImageSpec {
	s.IsACREnterprise = &v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetIsACRRegistry(v bool) *AddImageRequestContainerImageSpec {
	s.IsACRRegistry = &v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetRegistryCredential(v *AddImageRequestContainerImageSpecRegistryCredential) *AddImageRequestContainerImageSpec {
	s.RegistryCredential = v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetRegistryCriId(v string) *AddImageRequestContainerImageSpec {
	s.RegistryCriId = &v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetRegistryUrl(v string) *AddImageRequestContainerImageSpec {
	s.RegistryUrl = &v
	return s
}

func (s *AddImageRequestContainerImageSpec) Validate() error {
	if s.RegistryCredential != nil {
		if err := s.RegistryCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddImageRequestContainerImageSpecRegistryCredential struct {
	// The password of the logon user.
	//
	// example:
	//
	// userpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The registered address of the image repository.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username of the logon user.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddImageRequestContainerImageSpecRegistryCredential) String() string {
	return dara.Prettify(s)
}

func (s AddImageRequestContainerImageSpecRegistryCredential) GoString() string {
	return s.String()
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) GetPassword() *string {
	return s.Password
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) GetServer() *string {
	return s.Server
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) GetUserName() *string {
	return s.UserName
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) SetPassword(v string) *AddImageRequestContainerImageSpecRegistryCredential {
	s.Password = &v
	return s
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) SetServer(v string) *AddImageRequestContainerImageSpecRegistryCredential {
	s.Server = &v
	return s
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) SetUserName(v string) *AddImageRequestContainerImageSpecRegistryCredential {
	s.UserName = &v
	return s
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) Validate() error {
	return dara.Validate(s)
}

type AddImageRequestVMImageSpec struct {
	// The image ID.
	//
	// example:
	//
	// m-bp1akkkr1rkxtb******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s AddImageRequestVMImageSpec) String() string {
	return dara.Prettify(s)
}

func (s AddImageRequestVMImageSpec) GoString() string {
	return s.String()
}

func (s *AddImageRequestVMImageSpec) GetImageId() *string {
	return s.ImageId
}

func (s *AddImageRequestVMImageSpec) SetImageId(v string) *AddImageRequestVMImageSpec {
	s.ImageId = &v
	return s
}

func (s *AddImageRequestVMImageSpec) Validate() error {
	return dara.Validate(s)
}
