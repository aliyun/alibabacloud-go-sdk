// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageMosaicShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *AddImageMosaicShrinkRequest
	GetCredentialConfigShrink() *string
	SetImageFormat(v string) *AddImageMosaicShrinkRequest
	GetImageFormat() *string
	SetProjectName(v string) *AddImageMosaicShrinkRequest
	GetProjectName() *string
	SetQuality(v int32) *AddImageMosaicShrinkRequest
	GetQuality() *int32
	SetSourceURI(v string) *AddImageMosaicShrinkRequest
	GetSourceURI() *string
	SetTargetURI(v string) *AddImageMosaicShrinkRequest
	GetTargetURI() *string
	SetTargetsShrink(v string) *AddImageMosaicShrinkRequest
	GetTargetsShrink() *string
}

type AddImageMosaicShrinkRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The encoding of the output image. By default, the output image uses the same encoding as the input image. Valid values: jpg, png, and webp.
	//
	// example:
	//
	// jpg
	ImageFormat *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The quality of the output image. This parameter applies only to JPG and WebP images. Valid values: 0 to 100. Default value: 80.
	//
	// example:
	//
	// 80
	Quality *int32 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The OSS URI of the input image.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// Supported formats of input images include JPG, PNG, TIFF, JP2, and BMP.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/sampleobject.jpg
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The OSS URI of the output image.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/outputImage.jpg
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The bounding boxes and processing parameters.
	//
	// This parameter is required.
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s AddImageMosaicShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageMosaicShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddImageMosaicShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *AddImageMosaicShrinkRequest) GetImageFormat() *string {
	return s.ImageFormat
}

func (s *AddImageMosaicShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *AddImageMosaicShrinkRequest) GetQuality() *int32 {
	return s.Quality
}

func (s *AddImageMosaicShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *AddImageMosaicShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *AddImageMosaicShrinkRequest) GetTargetsShrink() *string {
	return s.TargetsShrink
}

func (s *AddImageMosaicShrinkRequest) SetCredentialConfigShrink(v string) *AddImageMosaicShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetImageFormat(v string) *AddImageMosaicShrinkRequest {
	s.ImageFormat = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetProjectName(v string) *AddImageMosaicShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetQuality(v int32) *AddImageMosaicShrinkRequest {
	s.Quality = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetSourceURI(v string) *AddImageMosaicShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetTargetURI(v string) *AddImageMosaicShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) SetTargetsShrink(v string) *AddImageMosaicShrinkRequest {
	s.TargetsShrink = &v
	return s
}

func (s *AddImageMosaicShrinkRequest) Validate() error {
	return dara.Validate(s)
}
