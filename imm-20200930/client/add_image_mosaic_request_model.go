// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageMosaicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *AddImageMosaicRequest
	GetCredentialConfig() *CredentialConfig
	SetImageFormat(v string) *AddImageMosaicRequest
	GetImageFormat() *string
	SetProjectName(v string) *AddImageMosaicRequest
	GetProjectName() *string
	SetQuality(v int32) *AddImageMosaicRequest
	GetQuality() *int32
	SetSourceURI(v string) *AddImageMosaicRequest
	GetSourceURI() *string
	SetTargetURI(v string) *AddImageMosaicRequest
	GetTargetURI() *string
	SetTargets(v []*AddImageMosaicRequestTargets) *AddImageMosaicRequest
	GetTargets() []*AddImageMosaicRequestTargets
}

type AddImageMosaicRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	Targets []*AddImageMosaicRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s AddImageMosaicRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageMosaicRequest) GoString() string {
	return s.String()
}

func (s *AddImageMosaicRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *AddImageMosaicRequest) GetImageFormat() *string {
	return s.ImageFormat
}

func (s *AddImageMosaicRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *AddImageMosaicRequest) GetQuality() *int32 {
	return s.Quality
}

func (s *AddImageMosaicRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *AddImageMosaicRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *AddImageMosaicRequest) GetTargets() []*AddImageMosaicRequestTargets {
	return s.Targets
}

func (s *AddImageMosaicRequest) SetCredentialConfig(v *CredentialConfig) *AddImageMosaicRequest {
	s.CredentialConfig = v
	return s
}

func (s *AddImageMosaicRequest) SetImageFormat(v string) *AddImageMosaicRequest {
	s.ImageFormat = &v
	return s
}

func (s *AddImageMosaicRequest) SetProjectName(v string) *AddImageMosaicRequest {
	s.ProjectName = &v
	return s
}

func (s *AddImageMosaicRequest) SetQuality(v int32) *AddImageMosaicRequest {
	s.Quality = &v
	return s
}

func (s *AddImageMosaicRequest) SetSourceURI(v string) *AddImageMosaicRequest {
	s.SourceURI = &v
	return s
}

func (s *AddImageMosaicRequest) SetTargetURI(v string) *AddImageMosaicRequest {
	s.TargetURI = &v
	return s
}

func (s *AddImageMosaicRequest) SetTargets(v []*AddImageMosaicRequestTargets) *AddImageMosaicRequest {
	s.Targets = v
	return s
}

func (s *AddImageMosaicRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddImageMosaicRequestTargets struct {
	// The radius of the Gaussian blur. Valid values: 1 to 50. Default value: 3. Unit: pixels.
	//
	// >  This parameter takes effect only for a Gaussian blur.
	//
	// example:
	//
	// 3
	BlurRadius *int32 `json:"BlurRadius,omitempty" xml:"BlurRadius,omitempty"`
	// The position of the bounding box.
	//
	// This parameter is required.
	Boundary *AddImageMosaicRequestTargetsBoundary `json:"Boundary,omitempty" xml:"Boundary,omitempty" type:"Struct"`
	// The color of the color shape. You can specify a color by using a color code such as`#RRGGBB` or preset color names such as `red` and `white`. The default value is #FFFFFF, which is white.
	//
	// >  This parameter takes effect only for solid color shapes.
	//
	// example:
	//
	// #FFFFFF
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The radius of the mosaic. Default value: 5. Unit: pixels.
	//
	// >  This parameter does not take effect for Gaussian blurs and solid color shapes.
	//
	// example:
	//
	// 5
	MosaicRadius *int32 `json:"MosaicRadius,omitempty" xml:"MosaicRadius,omitempty"`
	// The standard deviation of the Gaussian blur. The value must be greater than 0. Default value: 5.
	//
	// >  This parameter takes effect only for a Gaussian blur.
	//
	// example:
	//
	// 5
	Sigma *int32 `json:"Sigma,omitempty" xml:"Sigma,omitempty"`
	// The type of the mosaic effect. Valid values:
	//
	// 	- square: squares.
	//
	// 	- diamond: diamonds.
	//
	// 	- hexagon: hexagons.
	//
	// 	- blur: Gaussian blurs.
	//
	// 	- pure: solid color shapes.
	//
	// This parameter is required.
	//
	// example:
	//
	// square
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddImageMosaicRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s AddImageMosaicRequestTargets) GoString() string {
	return s.String()
}

func (s *AddImageMosaicRequestTargets) GetBlurRadius() *int32 {
	return s.BlurRadius
}

func (s *AddImageMosaicRequestTargets) GetBoundary() *AddImageMosaicRequestTargetsBoundary {
	return s.Boundary
}

func (s *AddImageMosaicRequestTargets) GetColor() *string {
	return s.Color
}

func (s *AddImageMosaicRequestTargets) GetMosaicRadius() *int32 {
	return s.MosaicRadius
}

func (s *AddImageMosaicRequestTargets) GetSigma() *int32 {
	return s.Sigma
}

func (s *AddImageMosaicRequestTargets) GetType() *string {
	return s.Type
}

func (s *AddImageMosaicRequestTargets) SetBlurRadius(v int32) *AddImageMosaicRequestTargets {
	s.BlurRadius = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetBoundary(v *AddImageMosaicRequestTargetsBoundary) *AddImageMosaicRequestTargets {
	s.Boundary = v
	return s
}

func (s *AddImageMosaicRequestTargets) SetColor(v string) *AddImageMosaicRequestTargets {
	s.Color = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetMosaicRadius(v int32) *AddImageMosaicRequestTargets {
	s.MosaicRadius = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetSigma(v int32) *AddImageMosaicRequestTargets {
	s.Sigma = &v
	return s
}

func (s *AddImageMosaicRequestTargets) SetType(v string) *AddImageMosaicRequestTargets {
	s.Type = &v
	return s
}

func (s *AddImageMosaicRequestTargets) Validate() error {
	if s.Boundary != nil {
		if err := s.Boundary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddImageMosaicRequestTargetsBoundary struct {
	// The height of the bounding box. The value can be an integer greater than or equal to 0 or a decimal within the range of [0,1):
	//
	// 	- An integer value greater than or equal to 0 indicates the height of the bounding box in pixels.
	//
	// 	- A decimal value within the range of [0,1) indicates the height of the bounding box as a ratio of its height to the image height.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The reference position of the bounding box on the image. Valid values:
	//
	// 	- topright: the upper-right corner.
	//
	// 	- topleft: the upper-left corner. This is the default value.
	//
	// 	- bottomright: the lower-right corner.
	//
	// 	- bottomleft: the lower-left corner.
	//
	// example:
	//
	// topleft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The width of the bounding box. The value can be an integer greater than or equal to 0 or a decimal within the range of [0,1):
	//
	// 	- An integer value greater than or equal to 0 indicates the width of the bounding box in pixels.
	//
	// 	- A decimal value within the range of [0,1) indicates the width of the bounding box as a ratio of its width to the image width.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The horizontal offset relative to the reference position. The value can be an integer greater than or equal to 0 or a decimal within the range of [0,1):
	//
	// 	- An integer value greater than or equal to 0 indicates the horizontal offset in pixels.
	//
	// 	- A decimal value within the range of [0,1) indicates the horizontal offset as a ratio of the offset to the image width.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The vertical offset relative to the reference position. The value can be an integer greater than or equal to 0 or a decimal within the range of [0,1):
	//
	// 	- An integer value greater than or equal to 0 indicates the vertical offset in pixels.
	//
	// 	- A decimal value within the range of [0,1) indicates the vertical offset as a ratio of the offset to the image height.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s AddImageMosaicRequestTargetsBoundary) String() string {
	return dara.Prettify(s)
}

func (s AddImageMosaicRequestTargetsBoundary) GoString() string {
	return s.String()
}

func (s *AddImageMosaicRequestTargetsBoundary) GetHeight() *float32 {
	return s.Height
}

func (s *AddImageMosaicRequestTargetsBoundary) GetReferPos() *string {
	return s.ReferPos
}

func (s *AddImageMosaicRequestTargetsBoundary) GetWidth() *float32 {
	return s.Width
}

func (s *AddImageMosaicRequestTargetsBoundary) GetX() *float32 {
	return s.X
}

func (s *AddImageMosaicRequestTargetsBoundary) GetY() *float32 {
	return s.Y
}

func (s *AddImageMosaicRequestTargetsBoundary) SetHeight(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.Height = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetReferPos(v string) *AddImageMosaicRequestTargetsBoundary {
	s.ReferPos = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetWidth(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.Width = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetX(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.X = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) SetY(v float32) *AddImageMosaicRequestTargetsBoundary {
	s.Y = &v
	return s
}

func (s *AddImageMosaicRequestTargetsBoundary) Validate() error {
	return dara.Validate(s)
}
