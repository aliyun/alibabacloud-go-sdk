// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageSplicingTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlign(v int64) *CreateImageSplicingTaskRequest
	GetAlign() *int64
	SetBackgroundColor(v string) *CreateImageSplicingTaskRequest
	GetBackgroundColor() *string
	SetCredentialConfig(v *CredentialConfig) *CreateImageSplicingTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetDirection(v string) *CreateImageSplicingTaskRequest
	GetDirection() *string
	SetImageFormat(v string) *CreateImageSplicingTaskRequest
	GetImageFormat() *string
	SetMargin(v int64) *CreateImageSplicingTaskRequest
	GetMargin() *int64
	SetNotification(v *Notification) *CreateImageSplicingTaskRequest
	GetNotification() *Notification
	SetPadding(v int64) *CreateImageSplicingTaskRequest
	GetPadding() *int64
	SetProjectName(v string) *CreateImageSplicingTaskRequest
	GetProjectName() *string
	SetQuality(v int64) *CreateImageSplicingTaskRequest
	GetQuality() *int64
	SetScaleType(v string) *CreateImageSplicingTaskRequest
	GetScaleType() *string
	SetSources(v []*CreateImageSplicingTaskRequestSources) *CreateImageSplicingTaskRequest
	GetSources() []*CreateImageSplicingTaskRequestSources
	SetTags(v map[string]interface{}) *CreateImageSplicingTaskRequest
	GetTags() map[string]interface{}
	SetTargetURI(v string) *CreateImageSplicingTaskRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateImageSplicingTaskRequest
	GetUserData() *string
}

type CreateImageSplicingTaskRequest struct {
	// The alignment value, in pixels, for the width or height of the images to be stitched. The value can range from 1 to 4096.
	//
	// - If you set **Direction*	- to `vertical`, this parameter specifies the width alignment.
	//
	// - If you set **Direction*	- to `horizontal`, this parameter specifies the height alignment.
	//
	// > If you do not specify this parameter, the width or height of the first image is used for alignment by default.
	//
	// example:
	//
	// 192
	Align *int64 `json:"Align,omitempty" xml:"Align,omitempty"`
	// The fill color for the areas specified by `Padding` and `Margin`. The value can be in the `#FFFFFF` format or a keyword such as `red` or `alpha`.
	//
	// example:
	//
	// red
	BackgroundColor *string `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	// The chained authorization configuration. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The image stitching method. Valid values:
	//
	// - vertical (default): Stitches images vertically. The widths of all images must be the same.
	//
	// - horizontal: Stitches images horizontally. The heights of all images must be the same.
	//
	// example:
	//
	// vertical
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The compression format of the output image. Valid values:
	//
	// - jpg (default)
	//
	// - png
	//
	// - webp
	//
	// example:
	//
	// jpg
	ImageFormat *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	// The blank margin, in pixels, of the stitched image. The default value is 0.
	//
	// example:
	//
	// 2
	Margin *int64 `json:"Margin,omitempty" xml:"Margin,omitempty"`
	// The message notification configuration. For information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The space, in pixels, between sub-images in the stitched image. The default value is 0.
	//
	// example:
	//
	// 2
	Padding *int64 `json:"Padding,omitempty" xml:"Padding,omitempty"`
	// The project name. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The compression quality of the output image. This parameter is valid only for JPG and WebP images. The value range is 0 to 100. The default value is 80.
	//
	// example:
	//
	// 80
	Quality *int64 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The scaling method used when the width or height of an image is aligned. Valid values:
	//
	// - fit (default): Scales the image without adding black bars. Only proportional scaling is supported.
	//
	// - stretch: Stretches the image to fill the space.
	//
	// example:
	//
	// stretch
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	// The list of input images. The images are stitched in the order of their URIs in the list.
	//
	// This parameter is required.
	Sources []*CreateImageSplicingTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// Custom tags used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The OSS URI where the output image is stored.
	//
	// The URI must be in the oss\\://${bucketname}/${objectname} format. ${bucketname} is the name of the OSS bucket that is in the same region as the project. ${objectname} is the path of the file, including the file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/outputImage.jpg
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The custom information. This information is returned in the asynchronous notification message. Use this information to associate the notification message with your system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// test-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageSplicingTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageSplicingTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskRequest) GetAlign() *int64 {
	return s.Align
}

func (s *CreateImageSplicingTaskRequest) GetBackgroundColor() *string {
	return s.BackgroundColor
}

func (s *CreateImageSplicingTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateImageSplicingTaskRequest) GetDirection() *string {
	return s.Direction
}

func (s *CreateImageSplicingTaskRequest) GetImageFormat() *string {
	return s.ImageFormat
}

func (s *CreateImageSplicingTaskRequest) GetMargin() *int64 {
	return s.Margin
}

func (s *CreateImageSplicingTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateImageSplicingTaskRequest) GetPadding() *int64 {
	return s.Padding
}

func (s *CreateImageSplicingTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateImageSplicingTaskRequest) GetQuality() *int64 {
	return s.Quality
}

func (s *CreateImageSplicingTaskRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *CreateImageSplicingTaskRequest) GetSources() []*CreateImageSplicingTaskRequestSources {
	return s.Sources
}

func (s *CreateImageSplicingTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateImageSplicingTaskRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateImageSplicingTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateImageSplicingTaskRequest) SetAlign(v int64) *CreateImageSplicingTaskRequest {
	s.Align = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetBackgroundColor(v string) *CreateImageSplicingTaskRequest {
	s.BackgroundColor = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateImageSplicingTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetDirection(v string) *CreateImageSplicingTaskRequest {
	s.Direction = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetImageFormat(v string) *CreateImageSplicingTaskRequest {
	s.ImageFormat = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetMargin(v int64) *CreateImageSplicingTaskRequest {
	s.Margin = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetNotification(v *Notification) *CreateImageSplicingTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetPadding(v int64) *CreateImageSplicingTaskRequest {
	s.Padding = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetProjectName(v string) *CreateImageSplicingTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetQuality(v int64) *CreateImageSplicingTaskRequest {
	s.Quality = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetScaleType(v string) *CreateImageSplicingTaskRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetSources(v []*CreateImageSplicingTaskRequestSources) *CreateImageSplicingTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetTags(v map[string]interface{}) *CreateImageSplicingTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetTargetURI(v string) *CreateImageSplicingTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) SetUserData(v string) *CreateImageSplicingTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateImageSplicingTaskRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImageSplicingTaskRequestSources struct {
	// The rotation angle of the image. Valid values:
	//
	// - 0 (default)
	//
	// - 90
	//
	// - 180
	//
	// - 270
	//
	// example:
	//
	// 90
	Rotate *int64 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The OSS URI of the source image.
	//
	// The URI must be in the oss\\://${Bucket}/${Object} format. `${Bucket}` is the name of the OSS bucket that is in the same region as the project. `${Object}` is the full path of the file, including the file name extension.
	//
	// Supported image formats: JPG and PNG.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/sampleobject.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateImageSplicingTaskRequestSources) String() string {
	return dara.Prettify(s)
}

func (s CreateImageSplicingTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskRequestSources) GetRotate() *int64 {
	return s.Rotate
}

func (s *CreateImageSplicingTaskRequestSources) GetURI() *string {
	return s.URI
}

func (s *CreateImageSplicingTaskRequestSources) SetRotate(v int64) *CreateImageSplicingTaskRequestSources {
	s.Rotate = &v
	return s
}

func (s *CreateImageSplicingTaskRequestSources) SetURI(v string) *CreateImageSplicingTaskRequestSources {
	s.URI = &v
	return s
}

func (s *CreateImageSplicingTaskRequestSources) Validate() error {
	return dara.Validate(s)
}
