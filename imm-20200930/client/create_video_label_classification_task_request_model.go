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
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom data, which is returned in an asynchronous notification and facilitates notification management. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
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
	return dara.Validate(s)
}
