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
	// **Leave this parameter empty unless you have specific requirements.**
	//
	// The chained authorization configuration. This parameter is optional. For more information, see [Use chained authorization to access other entity resources](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For details, click Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > Intelligent Media Management API callbacks do not support custom webhook addresses. Use MNS instead.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The password for the encrypted compressed package.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The list of files to decompress. If this parameter is empty, the entire compressed package is decompressed. The default value is empty.
	SelectedFilesShrink *string `json:"SelectedFiles,omitempty" xml:"SelectedFiles,omitempty"`
	// The Object Storage Service (OSS) address where the compressed file is stored.
	//
	// The OSS address must be in the \\`oss\\://${Bucket}/${Object}\\` format. \\`${Bucket}\\` is the name of the OSS bucket in the same region as the current project. \\`${Object}\\` is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object.zip
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The destination OSS address for the decompressed files. The selected files or the entire compressed package are decompressed to this address.
	//
	// The OSS address must be in the \\`oss\\://${Bucket}/${Object}\\` format. \\`${Bucket}\\` is the name of the OSS bucket in the same region as the current project. \\`${Object}\\` is the full path of the file, including the file name extension.
	//
	// example:
	//
	// oss://test-bucket/test-dir/
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// Custom user information. It is returned in the asynchronous notification message to help you associate the notification with your system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// test-data
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
