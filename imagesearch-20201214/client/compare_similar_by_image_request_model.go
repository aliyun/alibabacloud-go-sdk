// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareSimilarByImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CompareSimilarByImageRequest
	GetInstanceName() *string
	SetPrimaryPicContent(v string) *CompareSimilarByImageRequest
	GetPrimaryPicContent() *string
	SetSecondaryPicContent(v string) *CompareSimilarByImageRequest
	GetSecondaryPicContent() *string
}

type CompareSimilarByImageRequest struct {
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the <props="intl">[Image Search console](https://imagesearch.console.alibabacloud.com)<props="china">[Image Search console](https://imagesearch.console.aliyun.com) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Make sure you distinguish between them.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The image content.
	//
	// - The image size must not exceed 4 MB.
	//
	// - Image formats: PNG, JPG, JPEG, BMP, GIF, WEBP, TIFF, and PPM.
	//
	// - The transmission wait time must not exceed 5 seconds.
	//
	// <props="china">
	//
	// - If the service type is product image search, generic image search, furniture image search, or industrial hardware image search, the image width and height must be at least 100 px and at most 4096 px.
	//
	// <props="china">
	//
	// - If the service type is trademark image search, the image width and height must be at least 200 px and less than 4096 px.
	//
	// <props="china">
	//
	// - If the service type is fabric image search, the image width and height must be at least 448 px and at most 4096 px.
	//
	// <props="intl">
	//
	// - If the service type is product image search or generic image search, the image width and height must be at least 100 px and at most 4096 px.
	//
	// - The image must not contain rotation information.
	//
	// > **When calling by using an SDK:**- Only V3 SDKs are supported. You do not need to set the PrimaryPicContent field. The SDK encapsulates this field as PrimaryPicContentObject and automatically converts it to Base64 encoding. For examples, refer to [JAVA SDK](https://help.aliyun.com/document_detail/179188.html).- The SDK does not support passing image URLs directly. V3 SDKs provide an alternative way to upload images by URL. For examples, refer to [JAVA SDK](https://help.aliyun.com/document_detail/179188.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PrimaryPicContent *string `json:"PrimaryPicContent,omitempty" xml:"PrimaryPicContent,omitempty"`
	// The image content.
	//
	// - The image size must not exceed 4 MB.
	//
	// - Image formats: PNG, JPG, JPEG, BMP, GIF, WEBP, TIFF, and PPM.
	//
	// - The transmission wait time must not exceed 5 seconds.
	//
	// <props="china">
	//
	// - If the service type is product image search, generic image search, furniture image search, or industrial hardware image search, the image width and height must be at least 100 px and at most 4096 px.
	//
	// <props="china">
	//
	// - If the service type is trademark image search, the image width and height must be at least 200 px and less than 4096 px.
	//
	// <props="china">
	//
	// - If the service type is fabric image search, the image width and height must be at least 448 px and at most 4096 px.
	//
	// <props="intl">
	//
	// - If the service type is product image search or generic image search, the image width and height must be at least 100 px and at most 4096 px.
	//
	// - The image must not contain rotation information.
	//
	// > **When calling by using an SDK:**- Only V3 SDKs are supported. You do not need to set the PrimaryPicContent field. The SDK encapsulates this field as PrimaryPicContentObject and automatically converts it to Base64 encoding. For examples, refer to [JAVA SDK](https://help.aliyun.com/document_detail/179188.html).- The SDK does not support passing image URLs directly. V3 SDKs provide an alternative way to upload images by URL. For examples, refer to [JAVA SDK](https://help.aliyun.com/document_detail/179188.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	SecondaryPicContent *string `json:"SecondaryPicContent,omitempty" xml:"SecondaryPicContent,omitempty"`
}

func (s CompareSimilarByImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareSimilarByImageRequest) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CompareSimilarByImageRequest) GetPrimaryPicContent() *string {
	return s.PrimaryPicContent
}

func (s *CompareSimilarByImageRequest) GetSecondaryPicContent() *string {
	return s.SecondaryPicContent
}

func (s *CompareSimilarByImageRequest) SetInstanceName(v string) *CompareSimilarByImageRequest {
	s.InstanceName = &v
	return s
}

func (s *CompareSimilarByImageRequest) SetPrimaryPicContent(v string) *CompareSimilarByImageRequest {
	s.PrimaryPicContent = &v
	return s
}

func (s *CompareSimilarByImageRequest) SetSecondaryPicContent(v string) *CompareSimilarByImageRequest {
	s.SecondaryPicContent = &v
	return s
}

func (s *CompareSimilarByImageRequest) Validate() error {
	return dara.Validate(s)
}
