// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageSplicingTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlign(v int64) *CreateImageSplicingTaskShrinkRequest
	GetAlign() *int64
	SetBackgroundColor(v string) *CreateImageSplicingTaskShrinkRequest
	GetBackgroundColor() *string
	SetCredentialConfigShrink(v string) *CreateImageSplicingTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetDirection(v string) *CreateImageSplicingTaskShrinkRequest
	GetDirection() *string
	SetImageFormat(v string) *CreateImageSplicingTaskShrinkRequest
	GetImageFormat() *string
	SetMargin(v int64) *CreateImageSplicingTaskShrinkRequest
	GetMargin() *int64
	SetNotificationShrink(v string) *CreateImageSplicingTaskShrinkRequest
	GetNotificationShrink() *string
	SetPadding(v int64) *CreateImageSplicingTaskShrinkRequest
	GetPadding() *int64
	SetProjectName(v string) *CreateImageSplicingTaskShrinkRequest
	GetProjectName() *string
	SetQuality(v int64) *CreateImageSplicingTaskShrinkRequest
	GetQuality() *int64
	SetScaleType(v string) *CreateImageSplicingTaskShrinkRequest
	GetScaleType() *string
	SetSourcesShrink(v string) *CreateImageSplicingTaskShrinkRequest
	GetSourcesShrink() *string
	SetTagsShrink(v string) *CreateImageSplicingTaskShrinkRequest
	GetTagsShrink() *string
	SetTargetURI(v string) *CreateImageSplicingTaskShrinkRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateImageSplicingTaskShrinkRequest
	GetUserData() *string
}

type CreateImageSplicingTaskShrinkRequest struct {
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
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// Custom tags used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateImageSplicingTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageSplicingTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskShrinkRequest) GetAlign() *int64 {
	return s.Align
}

func (s *CreateImageSplicingTaskShrinkRequest) GetBackgroundColor() *string {
	return s.BackgroundColor
}

func (s *CreateImageSplicingTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateImageSplicingTaskShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *CreateImageSplicingTaskShrinkRequest) GetImageFormat() *string {
	return s.ImageFormat
}

func (s *CreateImageSplicingTaskShrinkRequest) GetMargin() *int64 {
	return s.Margin
}

func (s *CreateImageSplicingTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateImageSplicingTaskShrinkRequest) GetPadding() *int64 {
	return s.Padding
}

func (s *CreateImageSplicingTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateImageSplicingTaskShrinkRequest) GetQuality() *int64 {
	return s.Quality
}

func (s *CreateImageSplicingTaskShrinkRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *CreateImageSplicingTaskShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreateImageSplicingTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateImageSplicingTaskShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateImageSplicingTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateImageSplicingTaskShrinkRequest) SetAlign(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Align = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetBackgroundColor(v string) *CreateImageSplicingTaskShrinkRequest {
	s.BackgroundColor = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetDirection(v string) *CreateImageSplicingTaskShrinkRequest {
	s.Direction = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetImageFormat(v string) *CreateImageSplicingTaskShrinkRequest {
	s.ImageFormat = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetMargin(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Margin = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetNotificationShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetPadding(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Padding = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetProjectName(v string) *CreateImageSplicingTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetQuality(v int64) *CreateImageSplicingTaskShrinkRequest {
	s.Quality = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetScaleType(v string) *CreateImageSplicingTaskShrinkRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetSourcesShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetTagsShrink(v string) *CreateImageSplicingTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetTargetURI(v string) *CreateImageSplicingTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) SetUserData(v string) *CreateImageSplicingTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateImageSplicingTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
