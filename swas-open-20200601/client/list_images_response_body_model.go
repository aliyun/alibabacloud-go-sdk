// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody
	GetImages() []*ListImagesResponseBodyImages
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
}

type ListImagesResponseBody struct {
	// Details of the queried images.
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) GetImages() []*ListImagesResponseBodyImages {
	return s.Images
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImagesResponseBodyImages struct {
	// The description of the image.
	//
	// example:
	//
	// WordPress is a popular and powerful platform that allows you to build blogs and content management websites and implement custom features by using third-party templates and plug-ins. This WordPress image is built based on CentOS 7.9, PHP 7.4, MySQL 5.7, and Apache 2.4.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// 794c230fd3e64ea19f83f4d7a0ad****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// WordPress-4.8.1
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the image. Valid values:
	//
	// - system
	//
	// - app
	//
	// - custom
	//
	// example:
	//
	// app
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// - Linux
	//
	// - Windows
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesResponseBodyImages) GetImageName() *string {
	return s.ImageName
}

func (s *ListImagesResponseBodyImages) GetImageType() *string {
	return s.ImageType
}

func (s *ListImagesResponseBodyImages) GetPlatform() *string {
	return s.Platform
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageName(v string) *ListImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageType(v string) *ListImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetPlatform(v string) *ListImagesResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
