// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerImageSpecShrink(v string) *AddImageShrinkRequest
	GetContainerImageSpecShrink() *string
	SetDescription(v string) *AddImageShrinkRequest
	GetDescription() *string
	SetImageType(v string) *AddImageShrinkRequest
	GetImageType() *string
	SetImageVersion(v string) *AddImageShrinkRequest
	GetImageVersion() *string
	SetName(v string) *AddImageShrinkRequest
	GetName() *string
	SetVMImageSpecShrink(v string) *AddImageShrinkRequest
	GetVMImageSpecShrink() *string
}

type AddImageShrinkRequest struct {
	// The configurations of the container image.
	ContainerImageSpecShrink *string `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty"`
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
	VMImageSpecShrink *string `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty"`
}

func (s AddImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddImageShrinkRequest) GetContainerImageSpecShrink() *string {
	return s.ContainerImageSpecShrink
}

func (s *AddImageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddImageShrinkRequest) GetImageType() *string {
	return s.ImageType
}

func (s *AddImageShrinkRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *AddImageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddImageShrinkRequest) GetVMImageSpecShrink() *string {
	return s.VMImageSpecShrink
}

func (s *AddImageShrinkRequest) SetContainerImageSpecShrink(v string) *AddImageShrinkRequest {
	s.ContainerImageSpecShrink = &v
	return s
}

func (s *AddImageShrinkRequest) SetDescription(v string) *AddImageShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddImageShrinkRequest) SetImageType(v string) *AddImageShrinkRequest {
	s.ImageType = &v
	return s
}

func (s *AddImageShrinkRequest) SetImageVersion(v string) *AddImageShrinkRequest {
	s.ImageVersion = &v
	return s
}

func (s *AddImageShrinkRequest) SetName(v string) *AddImageShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddImageShrinkRequest) SetVMImageSpecShrink(v string) *AddImageShrinkRequest {
	s.VMImageSpecShrink = &v
	return s
}

func (s *AddImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
