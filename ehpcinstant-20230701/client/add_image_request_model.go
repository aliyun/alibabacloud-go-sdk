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
	ContainerImageSpec *AddImageRequestContainerImageSpec `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty" type:"Struct"`
	Description        *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageType          *string                            `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// V1.0
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// app-image
	Name        *string                     `json:"Name,omitempty" xml:"Name,omitempty"`
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
	return dara.Validate(s)
}

type AddImageRequestContainerImageSpec struct {
	// example:
	//
	// True
	IsACREnterprise *bool `json:"IsACREnterprise,omitempty" xml:"IsACREnterprise,omitempty"`
	// example:
	//
	// True
	IsACRRegistry      *bool                                                `json:"IsACRRegistry,omitempty" xml:"IsACRRegistry,omitempty"`
	RegistryCredential *AddImageRequestContainerImageSpecRegistryCredential `json:"RegistryCredential,omitempty" xml:"RegistryCredential,omitempty" type:"Struct"`
	// example:
	//
	// cri-xyz795ygf8k9****
	RegistryCriId *string `json:"RegistryCriId,omitempty" xml:"RegistryCriId,omitempty"`
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
	return dara.Validate(s)
}

type AddImageRequestContainerImageSpecRegistryCredential struct {
	// example:
	//
	// userpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
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
