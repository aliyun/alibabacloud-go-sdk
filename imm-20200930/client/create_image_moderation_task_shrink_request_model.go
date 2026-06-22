// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageModerationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateImageModerationTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetInterval(v int64) *CreateImageModerationTaskShrinkRequest
	GetInterval() *int64
	SetMaxFrames(v int64) *CreateImageModerationTaskShrinkRequest
	GetMaxFrames() *int64
	SetNotificationShrink(v string) *CreateImageModerationTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateImageModerationTaskShrinkRequest
	GetProjectName() *string
	SetScenesShrink(v string) *CreateImageModerationTaskShrinkRequest
	GetScenesShrink() *string
	SetSourceURI(v string) *CreateImageModerationTaskShrinkRequest
	GetSourceURI() *string
	SetTagsShrink(v string) *CreateImageModerationTaskShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateImageModerationTaskShrinkRequest
	GetUserData() *string
}

type CreateImageModerationTaskShrinkRequest struct {
	// The chained authorization configuration. This parameter is optional. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The frame capture frequency. This parameter is used for GIF and long image detection. The default value is 1.
	//
	// example:
	//
	// 2
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The maximum number of frames to capture. This parameter is used for GIF and long image detection. The default value is 1.
	//
	// example:
	//
	// 10
	MaxFrames *int64 `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	// The notification configuration. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project name. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The image detection scenarios.
	ScenesShrink *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	// The OSS URI of the image.
	//
	// The URI must follow the `oss://<Bucket>/<Object>` format. `<Bucket>` is the name of the OSS bucket that is in the same region as the project. `<Object>` is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The custom tags. You can use tags to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information. This information is returned in the asynchronous notification message to help you associate the message with your system. The value can be up to 2,048 bytes long.
	//
	// example:
	//
	// test-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageModerationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageModerationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateImageModerationTaskShrinkRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateImageModerationTaskShrinkRequest) GetMaxFrames() *int64 {
	return s.MaxFrames
}

func (s *CreateImageModerationTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateImageModerationTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateImageModerationTaskShrinkRequest) GetScenesShrink() *string {
	return s.ScenesShrink
}

func (s *CreateImageModerationTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateImageModerationTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateImageModerationTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateImageModerationTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetInterval(v int64) *CreateImageModerationTaskShrinkRequest {
	s.Interval = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetMaxFrames(v int64) *CreateImageModerationTaskShrinkRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetNotificationShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetProjectName(v string) *CreateImageModerationTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetScenesShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.ScenesShrink = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetSourceURI(v string) *CreateImageModerationTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetTagsShrink(v string) *CreateImageModerationTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) SetUserData(v string) *CreateImageModerationTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateImageModerationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
