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
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URI of the video file.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket1/object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The custom tags, which can be used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"test":"val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom data, which is returned in an asynchronous notification and facilitates notification management. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
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
