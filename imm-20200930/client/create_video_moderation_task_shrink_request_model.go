// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoModerationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateVideoModerationTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetInterval(v int64) *CreateVideoModerationTaskShrinkRequest
	GetInterval() *int64
	SetMaxFrames(v int64) *CreateVideoModerationTaskShrinkRequest
	GetMaxFrames() *int64
	SetNotificationShrink(v string) *CreateVideoModerationTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateVideoModerationTaskShrinkRequest
	GetProjectName() *string
	SetScenesShrink(v string) *CreateVideoModerationTaskShrinkRequest
	GetScenesShrink() *string
	SetSourceURI(v string) *CreateVideoModerationTaskShrinkRequest
	GetSourceURI() *string
	SetTagsShrink(v string) *CreateVideoModerationTaskShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateVideoModerationTaskShrinkRequest
	GetUserData() *string
}

type CreateVideoModerationTaskShrinkRequest struct {
	// The chained authorization configuration. This parameter is optional. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The interval for video snapshots, in seconds. The value can be an integer from 1 to 600. The default value is 1.
	//
	// example:
	//
	// 1
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The maximum number of frames that can be captured for this detection task. The value can be an integer from 5 to 3,600. The default value is 200.
	//
	// example:
	//
	// 200
	MaxFrames *int64 `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	// The notification configuration. For the format of asynchronous notification messages, see the metadata index section in [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information about how to get the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The video detection scenarios.
	ScenesShrink *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	// The OSS URI of the video.
	//
	// The OSS URI must follow the format oss\\://${Bucket}/${Object}. `${Bucket}` is the name of the OSS bucket in the same region as the project. `${Object}` is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// Custom tags used to search for tasks.
	//
	// example:
	//
	// {"test": "val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Custom information that is returned in the asynchronous notification message. Use this information to associate the notification message with your internal system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateVideoModerationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoModerationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateVideoModerationTaskShrinkRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateVideoModerationTaskShrinkRequest) GetMaxFrames() *int64 {
	return s.MaxFrames
}

func (s *CreateVideoModerationTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateVideoModerationTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateVideoModerationTaskShrinkRequest) GetScenesShrink() *string {
	return s.ScenesShrink
}

func (s *CreateVideoModerationTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateVideoModerationTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateVideoModerationTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateVideoModerationTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetInterval(v int64) *CreateVideoModerationTaskShrinkRequest {
	s.Interval = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetMaxFrames(v int64) *CreateVideoModerationTaskShrinkRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetNotificationShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetProjectName(v string) *CreateVideoModerationTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetScenesShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.ScenesShrink = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetSourceURI(v string) *CreateVideoModerationTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetTagsShrink(v string) *CreateVideoModerationTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) SetUserData(v string) *CreateVideoModerationTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateVideoModerationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
