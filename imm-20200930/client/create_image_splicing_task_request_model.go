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
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	Sources []*CreateImageSplicingTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	return dara.Validate(s)
}

type CreateImageSplicingTaskRequestSources struct {
	// The rotation angle. Valid values:
	//
	// 	- 0 (default)
	//
	// 	- 90
	//
	// 	- 180
	//
	// 	- 270
	//
	// example:
	//
	// 90
	Rotate *int64 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The Object Storage Service (OSS) bucket in which you store the input images.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the input images that have an extension.
	//
	// The following image formats are supported: jpg and png.
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
