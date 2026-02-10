// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageItem interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *ImageItem
	GetAcceleratorType() *string
	SetAuthorId(v string) *ImageItem
	GetAuthorId() *string
	SetFramework(v string) *ImageItem
	GetFramework() *string
	SetImageProviderType(v string) *ImageItem
	GetImageProviderType() *string
	SetImageTag(v string) *ImageItem
	GetImageTag() *string
	SetImageUrl(v string) *ImageItem
	GetImageUrl() *string
	SetImageUrlVpc(v string) *ImageItem
	GetImageUrlVpc() *string
}

type ImageItem struct {
	// The type of the image accelerator. Valid values:
	//
	// 	- cpu
	//
	// 	- gpu
	//
	// example:
	//
	// gpu
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The creator of the image.
	//
	// example:
	//
	// ken
	AuthorId *string `json:"AuthorId,omitempty" xml:"AuthorId,omitempty"`
	// The computing framework that is encapsulated by the image. Valid values:
	//
	// 	- TFJob
	//
	// 	- PyTorchJob
	//
	// example:
	//
	// PyTorchJob
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	// The image type. Valid values:
	//
	// 	- Community
	//
	// 	- PAI
	//
	// example:
	//
	// Community
	ImageProviderType *string `json:"ImageProviderType,omitempty" xml:"ImageProviderType,omitempty"`
	// The tag of the docker image.
	//
	// example:
	//
	// tensorflow-training:2.3-cpu-py36-ubuntu18.04
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The image URL address.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/pai-dlc/tensorflow-training:2.3-cpu-py36-ubuntu18.04
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The VPC endpoint of the image URL. This address provides faster access speed.
	//
	// example:
	//
	// registry-vpc.cn-beijing.aliyuncs.com/pai-dlc/tensorflow-training:2.3-cpu-py36-ubuntu18.04
	ImageUrlVpc *string `json:"ImageUrlVpc,omitempty" xml:"ImageUrlVpc,omitempty"`
}

func (s ImageItem) String() string {
	return dara.Prettify(s)
}

func (s ImageItem) GoString() string {
	return s.String()
}

func (s *ImageItem) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ImageItem) GetAuthorId() *string {
	return s.AuthorId
}

func (s *ImageItem) GetFramework() *string {
	return s.Framework
}

func (s *ImageItem) GetImageProviderType() *string {
	return s.ImageProviderType
}

func (s *ImageItem) GetImageTag() *string {
	return s.ImageTag
}

func (s *ImageItem) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageItem) GetImageUrlVpc() *string {
	return s.ImageUrlVpc
}

func (s *ImageItem) SetAcceleratorType(v string) *ImageItem {
	s.AcceleratorType = &v
	return s
}

func (s *ImageItem) SetAuthorId(v string) *ImageItem {
	s.AuthorId = &v
	return s
}

func (s *ImageItem) SetFramework(v string) *ImageItem {
	s.Framework = &v
	return s
}

func (s *ImageItem) SetImageProviderType(v string) *ImageItem {
	s.ImageProviderType = &v
	return s
}

func (s *ImageItem) SetImageTag(v string) *ImageItem {
	s.ImageTag = &v
	return s
}

func (s *ImageItem) SetImageUrl(v string) *ImageItem {
	s.ImageUrl = &v
	return s
}

func (s *ImageItem) SetImageUrlVpc(v string) *ImageItem {
	s.ImageUrlVpc = &v
	return s
}

func (s *ImageItem) Validate() error {
	return dara.Validate(s)
}
