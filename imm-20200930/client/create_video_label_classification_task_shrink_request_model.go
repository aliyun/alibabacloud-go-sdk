// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoLabelClassificationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetNotificationShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateVideoLabelClassificationTaskShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateVideoLabelClassificationTaskShrinkRequest
	GetSourceURI() *string
	SetTagsShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateVideoLabelClassificationTaskShrinkRequest
	GetUserData() *string
}

type CreateVideoLabelClassificationTaskShrinkRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The chained authorization configuration. This parameter is not required. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For more information, click Notification. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateVideoLabelClassificationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetNotificationShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetProjectName(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetSourceURI(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetTagsShrink(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) SetUserData(v string) *CreateVideoLabelClassificationTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
