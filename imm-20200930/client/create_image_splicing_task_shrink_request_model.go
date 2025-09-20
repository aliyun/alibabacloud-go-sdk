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
	// The width or height with which the input images must align. Valid values: 1 to 4096. Unit: px.
	//
	// 	- If you set **Direction*	- to `vertical`, this parameter specifies the width with which the input images must align.
	//
	// 	- If you set **Direction*	- to `horizontal`, this parameter specifies the height with which the input images must align.
	//
	// >  If you do not specify this parameter, the width or height of the first input image is used.
	//
	// example:
	//
	// 192
	Align *int64 `json:"Align,omitempty" xml:"Align,omitempty"`
	// The padding color of the spaces specified by `Padding` and `Margin`. Colors encoded in the `#FFFFFF` format and colors that are related to preset keywords such as `red` and `alpha` are supported.
	//
	// example:
	//
	// red
	BackgroundColor *string `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The splicing method. Valid values:
	//
	// 	- vertical (default): All input images are vertically aligned and have the same width.
	//
	// 	- horizontal: All input images are horizontally aligned and have the same height.
	//
	// example:
	//
	// vertical
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The compression format of the output image. Valid values:
	//
	// 	- jpg (default)
	//
	// 	- png
	//
	// 	- webp
	//
	// example:
	//
	// jpg
	ImageFormat *string `json:"ImageFormat,omitempty" xml:"ImageFormat,omitempty"`
	// The empty space or border around the edges of the output image. Default value: 0. Unit: px.
	//
	// example:
	//
	// 2
	Margin *int64 `json:"Margin,omitempty" xml:"Margin,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The space between component images in the output image. Default value: 0. Unit: px.
	//
	// example:
	//
	// 2
	Padding *int64 `json:"Padding,omitempty" xml:"Padding,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The compression quality of the output image. This parameter takes effect only for JPG and WebP images. Valid values: 0 to 100. Default value: 80.
	//
	// example:
	//
	// 80
	Quality *int64 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The scaling mode of the input images that are vertically or horizontally aligned. Valid values:
	//
	// 	- fit (default): Input images are scaled proportionally, and black edges are not retained.
	//
	// 	- stretch: Input images are stretched to fill the space.
	//
	// 	- horizon: Input images are horizontally stretched.
	//
	// 	- vertical: Input images are vertically stretched.
	//
	// example:
	//
	// stretch
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	// The input images. The images are sliced in the order of the input image URIs.
	//
	// This parameter is required.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The OSS bucket in which you want to store the output image.
	//
	// Specify the value in the oss://${bucketname}/${objectname} format. ${bucketname} specifies the name of the OSS bucket that resides in the same region as the current project. ${objectname} specifies the path to the output image.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://examplebucket/outputImage.jpg
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The user data, which is returned as asynchronous notifications to help manage notifications within your system. The maximum length of the user data is 2,048 bytes.
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
