// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoModerationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CreateVideoModerationTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetInterval(v int64) *CreateVideoModerationTaskRequest
	GetInterval() *int64
	SetMaxFrames(v int64) *CreateVideoModerationTaskRequest
	GetMaxFrames() *int64
	SetNotification(v *Notification) *CreateVideoModerationTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateVideoModerationTaskRequest
	GetProjectName() *string
	SetScenes(v []*string) *CreateVideoModerationTaskRequest
	GetScenes() []*string
	SetSourceURI(v string) *CreateVideoModerationTaskRequest
	GetSourceURI() *string
	SetTags(v map[string]interface{}) *CreateVideoModerationTaskRequest
	GetTags() map[string]interface{}
	SetUserData(v string) *CreateVideoModerationTaskRequest
	GetUserData() *string
}

type CreateVideoModerationTaskRequest struct {
	// The chained authorization configuration. This parameter is optional. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information about how to get the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The video detection scenarios.
	Scenes []*string `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Custom information that is returned in the asynchronous notification message. Use this information to associate the notification message with your internal system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateVideoModerationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoModerationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateVideoModerationTaskRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateVideoModerationTaskRequest) GetMaxFrames() *int64 {
	return s.MaxFrames
}

func (s *CreateVideoModerationTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateVideoModerationTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateVideoModerationTaskRequest) GetScenes() []*string {
	return s.Scenes
}

func (s *CreateVideoModerationTaskRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateVideoModerationTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateVideoModerationTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateVideoModerationTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateVideoModerationTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetInterval(v int64) *CreateVideoModerationTaskRequest {
	s.Interval = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetMaxFrames(v int64) *CreateVideoModerationTaskRequest {
	s.MaxFrames = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetNotification(v *Notification) *CreateVideoModerationTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetProjectName(v string) *CreateVideoModerationTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetScenes(v []*string) *CreateVideoModerationTaskRequest {
	s.Scenes = v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetSourceURI(v string) *CreateVideoModerationTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetTags(v map[string]interface{}) *CreateVideoModerationTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateVideoModerationTaskRequest) SetUserData(v string) *CreateVideoModerationTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateVideoModerationTaskRequest) Validate() error {
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
