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
	// The authorization chain. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The time interval between two consecutive frames in a GIF or long image. Default value: 1.
	//
	// example:
	//
	// 2
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The maximum number of frames that can be captured in a GIF or long image. Default value: 1.
	//
	// example:
	//
	// 10
	MaxFrames *int64 `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	// The notification settings. For more information, click Notification. For information about the asynchronous notification format, see [Asynchronous notification format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The scenarios in which you want to apply the image moderation task.
	ScenesShrink *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket in which you store the image.
	//
	// Specify the value in the `oss://<Bucket>/<Object>` format. `<Bucket>` specifies the name of the OSS bucket that resides in the same region as the current project. `<Object>` specifies the complete path to the image file that has an extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {"test": "val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The user data, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the user data is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
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
