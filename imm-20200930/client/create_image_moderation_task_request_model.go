// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageModerationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CreateImageModerationTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetInterval(v int64) *CreateImageModerationTaskRequest
	GetInterval() *int64
	SetMaxFrames(v int64) *CreateImageModerationTaskRequest
	GetMaxFrames() *int64
	SetNotification(v *Notification) *CreateImageModerationTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateImageModerationTaskRequest
	GetProjectName() *string
	SetScenes(v []*string) *CreateImageModerationTaskRequest
	GetScenes() []*string
	SetSourceURI(v string) *CreateImageModerationTaskRequest
	GetSourceURI() *string
	SetTags(v map[string]interface{}) *CreateImageModerationTaskRequest
	GetTags() map[string]interface{}
	SetUserData(v string) *CreateImageModerationTaskRequest
	GetUserData() *string
}

type CreateImageModerationTaskRequest struct {
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The scenarios in which you want to apply the image moderation task.
	Scenes []*string `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The user data, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the user data is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateImageModerationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageModerationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateImageModerationTaskRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateImageModerationTaskRequest) GetMaxFrames() *int64 {
	return s.MaxFrames
}

func (s *CreateImageModerationTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateImageModerationTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateImageModerationTaskRequest) GetScenes() []*string {
	return s.Scenes
}

func (s *CreateImageModerationTaskRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateImageModerationTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateImageModerationTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateImageModerationTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateImageModerationTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateImageModerationTaskRequest) SetInterval(v int64) *CreateImageModerationTaskRequest {
	s.Interval = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetMaxFrames(v int64) *CreateImageModerationTaskRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetNotification(v *Notification) *CreateImageModerationTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateImageModerationTaskRequest) SetProjectName(v string) *CreateImageModerationTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetScenes(v []*string) *CreateImageModerationTaskRequest {
	s.Scenes = v
	return s
}

func (s *CreateImageModerationTaskRequest) SetSourceURI(v string) *CreateImageModerationTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateImageModerationTaskRequest) SetTags(v map[string]interface{}) *CreateImageModerationTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateImageModerationTaskRequest) SetUserData(v string) *CreateImageModerationTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateImageModerationTaskRequest) Validate() error {
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
	return nil
}
