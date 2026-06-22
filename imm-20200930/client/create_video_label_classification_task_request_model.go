// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoLabelClassificationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CreateVideoLabelClassificationTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetNotification(v *Notification) *CreateVideoLabelClassificationTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateVideoLabelClassificationTaskRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateVideoLabelClassificationTaskRequest
	GetSourceURI() *string
	SetTags(v map[string]interface{}) *CreateVideoLabelClassificationTaskRequest
	GetTags() map[string]interface{}
	SetUserData(v string) *CreateVideoLabelClassificationTaskRequest
	GetUserData() *string
}

type CreateVideoLabelClassificationTaskRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The chained authorization configuration. This parameter is not required. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For more information, click Notification. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project name. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The Object Storage Service (OSS) URI of the video.
	//
	// The OSS URI must follow the format oss\\://${Bucket}/${Object}. ${Bucket} is the name of the OSS bucket that is in the same region as the current project. ${Object} is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket1/object.mp4
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// Custom tags that you can use to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"test":"val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Custom information. This information is returned in the asynchronous notification message. You can use this information to associate the notification message with your services. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {
	//
	//       "ID": "testuid",
	//
	//       "Name": "test-user",
	//
	//       "Avatar": "http://test.com/testuid"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateVideoLabelClassificationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateVideoLabelClassificationTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateVideoLabelClassificationTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateVideoLabelClassificationTaskRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateVideoLabelClassificationTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateVideoLabelClassificationTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateVideoLabelClassificationTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateVideoLabelClassificationTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetNotification(v *Notification) *CreateVideoLabelClassificationTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetProjectName(v string) *CreateVideoLabelClassificationTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetSourceURI(v string) *CreateVideoLabelClassificationTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetTags(v map[string]interface{}) *CreateVideoLabelClassificationTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) SetUserData(v string) *CreateVideoLabelClassificationTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskRequest) Validate() error {
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
