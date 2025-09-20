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
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The configurations of authorization chains. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	SelectedFiles []*string `json:"SelectedFiles,omitempty" xml:"SelectedFiles,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}
