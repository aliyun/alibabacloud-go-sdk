// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileUncompressionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CreateFileUncompressionTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetNotification(v *Notification) *CreateFileUncompressionTaskRequest
	GetNotification() *Notification
	SetPassword(v string) *CreateFileUncompressionTaskRequest
	GetPassword() *string
	SetProjectName(v string) *CreateFileUncompressionTaskRequest
	GetProjectName() *string
	SetSelectedFiles(v []*string) *CreateFileUncompressionTaskRequest
	GetSelectedFiles() []*string
	SetSourceURI(v string) *CreateFileUncompressionTaskRequest
	GetSourceURI() *string
	SetTargetURI(v string) *CreateFileUncompressionTaskRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateFileUncompressionTaskRequest
	GetUserData() *string
}

type CreateFileUncompressionTaskRequest struct {
	// **Leave this parameter empty unless you have specific requirements.**
	//
	// The chained authorization configuration. This parameter is optional. For more information, see [Use chained authorization to access other entity resources](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification configuration. For details, click Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > Intelligent Media Management API callbacks do not support custom webhook addresses. Use MNS instead.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	SelectedFiles []*string `json:"SelectedFiles,omitempty" xml:"SelectedFiles,omitempty" type:"Repeated"`
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

func (s CreateFileUncompressionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileUncompressionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateFileUncompressionTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateFileUncompressionTaskRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateFileUncompressionTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFileUncompressionTaskRequest) GetSelectedFiles() []*string {
	return s.SelectedFiles
}

func (s *CreateFileUncompressionTaskRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateFileUncompressionTaskRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateFileUncompressionTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFileUncompressionTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateFileUncompressionTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetNotification(v *Notification) *CreateFileUncompressionTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetPassword(v string) *CreateFileUncompressionTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetProjectName(v string) *CreateFileUncompressionTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetSelectedFiles(v []*string) *CreateFileUncompressionTaskRequest {
	s.SelectedFiles = v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetSourceURI(v string) *CreateFileUncompressionTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetTargetURI(v string) *CreateFileUncompressionTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) SetUserData(v string) *CreateFileUncompressionTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateFileUncompressionTaskRequest) Validate() error {
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
