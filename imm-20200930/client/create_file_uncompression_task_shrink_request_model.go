// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileUncompressionTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateFileUncompressionTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetNotificationShrink(v string) *CreateFileUncompressionTaskShrinkRequest
	GetNotificationShrink() *string
	SetPassword(v string) *CreateFileUncompressionTaskShrinkRequest
	GetPassword() *string
	SetProjectName(v string) *CreateFileUncompressionTaskShrinkRequest
	GetProjectName() *string
	SetSelectedFilesShrink(v string) *CreateFileUncompressionTaskShrinkRequest
	GetSelectedFilesShrink() *string
	SetSourceURI(v string) *CreateFileUncompressionTaskShrinkRequest
	GetSourceURI() *string
	SetTargetURI(v string) *CreateFileUncompressionTaskShrinkRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateFileUncompressionTaskShrinkRequest
	GetUserData() *string
}

type CreateFileUncompressionTaskShrinkRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The configurations of authorization chains. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The password that protects the package.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The files to extract. If you do not specify this parameter, the entire package is decompressed.
	SelectedFilesShrink *string `json:"SelectedFiles,omitempty" xml:"SelectedFiles,omitempty"`
	// The OSS URI of the package.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://imm-apitest-fxf2/name.zip
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The OSS URI to which you want to extract files from the package or decompress the entire package.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-dir/
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFileUncompressionTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileUncompressionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetSelectedFilesShrink() *string {
	return s.SelectedFilesShrink
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateFileUncompressionTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetNotificationShrink(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetPassword(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetProjectName(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetSelectedFilesShrink(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.SelectedFilesShrink = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetSourceURI(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetTargetURI(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) SetUserData(v string) *CreateFileUncompressionTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateFileUncompressionTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
