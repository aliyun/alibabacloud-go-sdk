// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncodeBlindWatermarkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *EncodeBlindWatermarkRequest
  GetContent() *string 
  SetImageQuality(v int32) *EncodeBlindWatermarkRequest
  GetImageQuality() *int32 
  SetProjectName(v string) *EncodeBlindWatermarkRequest
  GetProjectName() *string 
  SetSourceURI(v string) *EncodeBlindWatermarkRequest
  GetSourceURI() *string 
  SetStrengthLevel(v string) *EncodeBlindWatermarkRequest
  GetStrengthLevel() *string 
  SetTargetURI(v string) *EncodeBlindWatermarkRequest
  GetTargetURI() *string 
}

type EncodeBlindWatermarkRequest struct {
  // The text content of watermarks. It can be up to 256 characters in length.
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // This parameter takes effect only if the input image format is JPG.
  // 
  // The storage quality of the output image that carries the watermarks. Default value: 90. Valid values: 70 to 100. The higher the quality, the larger the image size and the higher the watermark resolution quality.
  // 
  // example:
  // 
  // 90
  ImageQuality *int32 `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
  // The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test-project
  ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
  // The Object Storage Service (OSS) URI of the image.
  // 
  // Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region with the current project. `${Object}` specifies the path of the object with the extension included.
  // 
  // Supported image formats: JPG, PNG, BMP, TIFF, and WebP.
  // 
  // Image size limit: 10,000 px maximum and 80 px x 80 px minimum.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // oss://test-bucket/test-object.jpg
  SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
  // The watermark strength level. The higher the strength, the more resistant the watermarked image is to attacks, but the more the image is distorted. Default value: low. Valid values: [low, medium, high].
  // 
  // example:
  // 
  // low
  StrengthLevel *string `json:"StrengthLevel,omitempty" xml:"StrengthLevel,omitempty"`
  // The URI of the output image in OSS.
  // 
  // Specify the URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
  // 
  // > 
  // 
  // 	- The format of the output image is the same as that of the input image.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // oss://test-bucket/target-object.jpg
  TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
}

func (s EncodeBlindWatermarkRequest) String() string {
  return dara.Prettify(s)
}

func (s EncodeBlindWatermarkRequest) GoString() string {
  return s.String()
}

func (s *EncodeBlindWatermarkRequest) GetContent() *string  {
  return s.Content
}

func (s *EncodeBlindWatermarkRequest) GetImageQuality() *int32  {
  return s.ImageQuality
}

func (s *EncodeBlindWatermarkRequest) GetProjectName() *string  {
  return s.ProjectName
}

func (s *EncodeBlindWatermarkRequest) GetSourceURI() *string  {
  return s.SourceURI
}

func (s *EncodeBlindWatermarkRequest) GetStrengthLevel() *string  {
  return s.StrengthLevel
}

func (s *EncodeBlindWatermarkRequest) GetTargetURI() *string  {
  return s.TargetURI
}

func (s *EncodeBlindWatermarkRequest) SetContent(v string) *EncodeBlindWatermarkRequest {
  s.Content = &v
  return s
}

func (s *EncodeBlindWatermarkRequest) SetImageQuality(v int32) *EncodeBlindWatermarkRequest {
  s.ImageQuality = &v
  return s
}

func (s *EncodeBlindWatermarkRequest) SetProjectName(v string) *EncodeBlindWatermarkRequest {
  s.ProjectName = &v
  return s
}

func (s *EncodeBlindWatermarkRequest) SetSourceURI(v string) *EncodeBlindWatermarkRequest {
  s.SourceURI = &v
  return s
}

func (s *EncodeBlindWatermarkRequest) SetStrengthLevel(v string) *EncodeBlindWatermarkRequest {
  s.StrengthLevel = &v
  return s
}

func (s *EncodeBlindWatermarkRequest) SetTargetURI(v string) *EncodeBlindWatermarkRequest {
  s.TargetURI = &v
  return s
}

func (s *EncodeBlindWatermarkRequest) Validate() error {
  return dara.Validate(s)
}

